package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/TheLazarusNetwork/erebrus/storage"
	"github.com/TheLazarusNetwork/erebrus/template"
	"github.com/TheLazarusNetwork/erebrus/util"
	log "github.com/sirupsen/logrus"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// ReadServer object, create default one
func ReadServer() (*model.Server, error) {
	if !util.FileExists(filepath.Join(os.Getenv("WG_CONF_DIR"), "server.json")) {
		server := &model.Server{}

		key, err := wgtypes.GeneratePrivateKey()
		if err != nil {
			return nil, err
		}
		server.PrivateKey = key.String()
		server.PublicKey = key.PublicKey().String()
		server.Endpoint = os.Getenv("WG_ENDPOINT_HOST")
		listenPort, _ := strconv.ParseInt(os.Getenv("WG_ENDPOINT_PORT"), 10, 32)

		util.CheckError("Error while reading listen port:", err)
		server.ListenPort = listenPort

		server.Address = make([]string, 0)
		// server.Address = append(server.Address, os.Getenv("WG_IPv6_SUBNET")) //	"fd9f:6666::10:0:0:1/64"
		server.Address = append(server.Address, os.Getenv("WG_IPv4_SUBNET")) //	"10.0.0.1/24"

		server.DNS = make([]string, 0)
		// server.DNS = append(server.DNS, "fd9f::10:0:0:2")
		server.DNS = append(server.DNS, os.Getenv("WG_DNS")) //	"1.1.1.1"

		server.AllowedIPs = make([]string, 0)
		server.AllowedIPs = append(server.AllowedIPs, os.Getenv("WG_ALLOWED_IP_1")) //	"0.0.0.0/0"
		server.AllowedIPs = append(server.AllowedIPs, os.Getenv("WG_ALLOWED_IP_2")) //	"::/0"

		server.PersistentKeepalive = 16
		server.Mtu = 0
		server.PreUp = os.Getenv("WG_PRE_UP")       //	"echo WireGuard PreUp"
		server.PostUp = os.Getenv("WG_POST_UP")     //	"echo WireGuard PostUp"
		server.PreDown = os.Getenv("WG_PRE_DOWN")   //	"echo WireGuard PreDown"
		server.PostDown = os.Getenv("WG_POST_DOWN") //	"echo WireGuard PostDown"
		server.CreatedAt = int64(time.Now().Nanosecond())
		server.UpdatedAt = server.CreatedAt

		err = storage.Serialize("server.json", server)
		if err != nil {
			return nil, err
		}

		// server.json was missing, dump wg config after creation
		err = UpdateServerConfigWg()
		if err != nil {
			return nil, err
		}
	}

	c, err := storage.Deserialize("server.json")
	if err != nil {
		return nil, err
	}

	return c.(*model.Server), nil
}

// UpdateServer keep private values from existing one
func UpdateServer(server *model.Server) (*model.Server, error) {
	current, err := storage.Deserialize("server.json")
	if err != nil {
		return nil, err
	}

	// check if server is valid
	errs := server.IsValid()
	if len(errs) != 0 {
		for _, err := range errs {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("server validation error")
		}
		return nil, errors.New("failed to validate server")
	}

	server.PrivateKey = current.(*model.Server).PrivateKey
	server.PublicKey = current.(*model.Server).PublicKey
	//server.PresharedKey = current.(*model.Server).PresharedKey
	server.UpdatedAt = int64(time.Now().Nanosecond())

	err = storage.Serialize("server.json", server)
	if err != nil {
		return nil, err
	}

	v, err := storage.Deserialize("server.json")
	if err != nil {
		return nil, err
	}
	server = v.(*model.Server)

	return server, UpdateServerConfigWg()
}

// UpdateServerConfigWg in wg format
func UpdateServerConfigWg() error {
	clients, err := ReadClients()
	if err != nil {
		return err
	}

	server, err := ReadServer()
	if err != nil {
		return err
	}

	_, err = template.DumpServerWg(clients, server)
	if err != nil {
		return err
	}

	return nil
}

// GetAllReservedIps the list of all reserved IPs, client and server
func GetAllReservedIps() ([]string, error) {
	clients, err := ReadClients()
	if err != nil {
		return nil, err
	}

	server, err := ReadServer()
	if err != nil {
		return nil, err
	}

	reserverIps := make([]string, 0)

	for _, client := range clients {
		for _, cidr := range client.Address {
			ip, err := util.GetIPFromCidr(cidr)
			if err != nil {
				log.WithFields(log.Fields{
					"err":  err,
					"cidr": cidr,
				}).Error("failed to ip from cidr")
			} else {
				reserverIps = append(reserverIps, ip)
			}
		}
	}

	for _, cidr := range server.Address {
		ip, err := util.GetIPFromCidr(cidr)
		if err != nil {
			log.WithFields(log.Fields{
				"err":  err,
				"cidr": err,
			}).Error("failed to ip from cidr")
		} else {
			reserverIps = append(reserverIps, ip)
		}
	}

	return reserverIps, nil
}

// ReadWgConfigFile return content of wireguard config file
func ReadWgConfigFile() ([]byte, error) {
	return util.ReadFile(filepath.Join(os.Getenv("WG_CONF_DIR"), os.Getenv("WG_INTERFACE_NAME")))
}

// Method to get the server status
func GetServerStatus() (*model.Status, error) {
	var response = &model.Status{}
	resp, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		return nil, err
	}
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response.PublicIP = string(ip)
	hostname, _ := os.Hostname()
	response.Hostname = hostname
	response.Domain = os.Getenv("DOMAIN")
	response.GRPCPort = os.Getenv("GRPC_PORT")
	response.Version = util.Version
	response.HttpPort = os.Getenv("HTTP_PORT")
	response.Region = os.Getenv("REGION")
	response.VPNPort = os.Getenv("WG_ENDPOINT_PORT")

	serverStatus, err := storage.Deserialize("server.json")

	if err != nil {
		log.WithFields(util.StandardFields).Fatal(err)
	} else {
		var server model.Server
		bodybytes, _ := json.Marshal(serverStatus)
		json.Unmarshal(bodybytes, &server)
		response.PublicKey = server.PublicKey
		response.PersistentKeepalive = server.PersistentKeepalive
		response.DNS = server.DNS
	}
	var privateip string
	addrs, _ := net.InterfaceAddrs()

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				privateip = ipnet.IP.String() + ","
			}
		}
	}
	response.PrivateIP = privateip

	return response, nil
}

// success response message
func MakeSucessResponse(status int64, message string, server *model.Server, client *model.Client, clients []*model.Client) *model.Response {
	return &model.Response{
		Status:  status,
		Message: message,
		Server:  server,
		Client:  client,
		Clients: clients,
		Success: true,
		Error:   "",
	}
}

// error response message
func MakeErrorResponse(status int64, err string, server *model.Server, client *model.Client, clients []*model.Client) *model.Response {
	return &model.Response{
		Status:  status,
		Message: "",
		Server:  server,
		Client:  client,
		Clients: clients,
		Success: false,
		Error:   err,
	}

}

func UpdateEndpointDetails() {
	for {
		resp, err := http.Get("https://ipinfo.io/ip")
		if err != nil {
			log.WithFields(util.StandardFields).Fatal(err)
		}
		ip, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithFields(util.StandardFields).Fatal(err)

		}
		var data model.RegionEndpoint
		data.Code = os.Getenv("REGION_CODE")
		data.Name = os.Getenv("REGION_NAME")
		data.ServiceType = "erebrus"
		data.Endpoint = string(ip) + ":" + os.Getenv("GRPC_PORT")
		js, _ := json.Marshal(data)

		client := &http.Client{}
		request, _ := http.NewRequest(http.MethodPatch, os.Getenv("MASTERNODE_URL")+"/api/v1.0/regupdate", bytes.NewBuffer(js))
		_, err = client.Do(request)
		if err != nil {
			log.WithFields(util.StandardFields).Fatal(err)
		}
		log.WithFields(util.StandardFields).Debug("Region Endpoint updation sucess")
		time.Sleep(3 * time.Hour)
	}

}
