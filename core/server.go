package core

import (
	"errors"
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
		listenPort, err := strconv.Atoi(os.Getenv("WG_ENDPOINT_PORT"))
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
		server.Created = time.Now().UTC()
		server.Updated = server.Created

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
	server.Updated = time.Now().UTC()

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
