package core

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/TheLazarusNetwork/erebrus/storage"
	"github.com/TheLazarusNetwork/erebrus/template"
	"github.com/TheLazarusNetwork/erebrus/util"
	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateClient client with all necessary data
func CreateClient(client *model.Client) (*model.Client, error) {
	// check if client is valid
	errs := client.IsValid()
	if len(errs) != 0 {
		for _, err := range errs {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("client validation error")
		}
		return nil, errors.New("failed to validate client")
	}

	u, err := uuid.NewRandom()
	client.UUID = u.String()

	key, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}
	// client.PrivateKey = key.String()
	client.PublicKey = key.PublicKey().String()

	presharedKey, err := wgtypes.GenerateKey()
	if err != nil {
		return nil, err
	}
	client.PresharedKey = presharedKey.String()

	reserverIps, err := GetAllReservedIps()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	for _, network := range client.Address {
		ip, err := util.GetAvailableIP(network, reserverIps)
		if err != nil {
			return nil, err
		}
		if util.IsIPv6(ip) {
			ip = ip + "/128"
		} else {
			ip = ip + "/32"
		}
		ips = append(ips, ip)
	}
	client.Address = ips
	client.CreatedAt = timestamppb.Now().AsTime().UnixMilli()

	client.UpdatedAt = client.CreatedAt

	err = storage.Serialize(client.UUID, client)
	if err != nil {
		return nil, err
	}

	v, err := storage.Deserialize(client.UUID)
	if err != nil {
		return nil, err
	}
	client = v.(*model.Client)

	// data modified, dump new config
	return client, UpdateServerConfigWg()
}

// ReadClient client by id
func ReadClient(id string) (*model.Client, error) {
	v, err := storage.Deserialize(id)
	if err != nil {
		return nil, err
	}
	client := v.(*model.Client)

	return client, nil
}

// UpdateClient preserve keys
func UpdateClient(UUID string, client *model.Client) (*model.Client, error) {
	v, err := storage.Deserialize(UUID)
	if err != nil {
		return nil, err
	}
	current := v.(*model.Client)

	if current.UUID != client.UUID {
		return nil, errors.New("records UUID mismatch")
	}

	// check if client is valid
	errs := client.IsValid()
	if len(errs) != 0 {
		for _, err := range errs {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("client validation error")
		}
		return nil, errors.New("failed to validate client")
	}

	// keep keys
	// client.PrivateKey = current.PrivateKey
	client.PublicKey = current.PublicKey
	client.UpdatedAt = timestamppb.Now().AsTime().UnixMilli()

	err = storage.Serialize(client.UUID, client)
	if err != nil {
		return nil, err
	}

	v, err = storage.Deserialize(UUID)
	if err != nil {
		return nil, err
	}
	client = v.(*model.Client)

	// data modified, dump new config
	return client, UpdateServerConfigWg()
}

// DeleteClient from disk
func DeleteClient(id string) error {
	path := filepath.Join(os.Getenv("WG_CLIENTS_DIR"), id)
	err := os.Remove(path)
	if err != nil {
		return err
	}

	// data modified, dump new config
	return UpdateServerConfigWg()
}

// ReadClients all clients
func ReadClients() ([]*model.Client, error) {
	clients := make([]*model.Client, 0)

	files, err := ioutil.ReadDir(filepath.Join(os.Getenv("WG_CLIENTS_DIR")))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		// clients file name is an uuid
		_, err := uuid.Parse(f.Name())
		if err == nil {
			c, err := storage.Deserialize(f.Name())
			if err != nil {
				log.WithFields(log.Fields{
					"err":  err,
					"path": f.Name(),
				}).Error("failed to deserialize client")
			} else {
				clients = append(clients, c.(*model.Client))
			}
		}
	}

	sort.Slice(clients, func(i, j int) bool {
		return clients[i].CreatedAt < (clients[j].CreatedAt)
	})

	return clients, nil
}

// ReadClientConfig in wg format
func ReadClientConfig(id string) ([]byte, error) {
	client, err := ReadClient(id)
	if err != nil {
		return nil, err
	}

	server, err := ReadServer()
	if err != nil {
		return nil, err
	}

	configDataWg, err := template.DumpClientWg(client, server)
	if err != nil {
		return nil, err
	}

	return configDataWg, nil
}
