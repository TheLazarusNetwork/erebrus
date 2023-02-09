package model

import (
	"fmt"

	"github.com/TheLazarusNetwork/erebrus/util"

	"golang.zx2c4.com/wireguard/wgctrl"
	// "golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// Server structure
/*type Server struct {
	Address             []string  `json:"address"`
	ListenPort          int64     `json:"listenPort"`
	Mtu                 int64     `json:"mtu"`
	PrivateKey          string    `json:"privateKey"`
	PublicKey           string    `json:"publicKey"`
	Endpoint            string    `json:"endpoint"`
	PersistentKeepalive int64     `json:"persistentKeepalive"`
	DNS                 []string  `json:"dns"`
	AllowedIPs          []string  `json:"allowedips"`
	PreUp               string    `json:"preUp"`
	PostUp              string    `json:"postUp"`
	PreDown             string    `json:"preDown"`
	PostDown            string    `json:"postDown"`
	UpdatedBy           string    `json:"updatedBy"`
	Created             time.Time `json:"created"`
	Updated             time.Time `json:"updated"`
}*/

/*type Status struct {
	Version   string `json:"version"`
	HostName  string`json:"hostname"`
	Domain    string `json:"domain"`
	PublicIP  string `json:"publicIP"`
	gRPCPort  string`json:"grpcport"`
	PrivateIP string`json:"private"`
	HttpPort  string`json:"httpport"`
	Region    string
	VPNPort   string
}*/

// WireGuardServer supports both Kernel and Userland implementations of WireGuard.
type WireGuardServer struct {
	wg         *wgctrl.Client
	deviceName string
}

// NewServer initializes a Server with a WireGuard client.
func NewServer(wg *wgctrl.Client, deviceName string) (*WireGuardServer, error) {
	return &WireGuardServer{wg: wg, deviceName: deviceName}, nil
}

// ListPeers retrieves information about all Peers known to the current
// WireGuard interface, including allowed IP addresses and usage stats,
// optionally with pagination.
// func (wgServer *WireGuardServer) ListPeers(ctx context.Context, req *Client) (*Client, error) {
// 	if err := validateListPeersRequest(req); err != nil {
// 		return nil, err
// 	}

// 	dev, err := wgServer.wg.Device(s.deviceName)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get WireGuard device: %w", err)
// 	}

// 	var peers []*client.Peer

// 	for _, peer := range dev.Peers {
// 		peers = append(peers, peer2rpc(peer))
// 	}

//  //TODO(jc): pagination

// 	return &Client{
// 		Peers: peers,
// 	}, nil
// }

// func peer2rpc(peer wgtypes.Peer) *client.Peer {
// 	var keepAlive string
// 	if peer.PersistentKeepaliveInterval > 0 {
// 		keepAlive = peer.PersistentKeepaliveInterval.String()
// 	}

// 	var allowedIPs []string
// 	for _, allowedIP := range peer.AllowedIPs {
// 		allowedIPs = append(allowedIPs, allowedIP.String())
// 	}

// 	return &client.Peer{
// 		PublicKey:           peer.PublicKey.String(),
// 		HasPresharedKey:     peer.PresharedKey != wgtypes.Key{},
// 		Endpoint:            peer.Endpoint.String(),
// 		PersistentKeepAlive: keepAlive,
// 		LastHandshake:       peer.LastHandshakeTime,
// 		ReceiveBytes:        peer.ReceiveBytes,
// 		TransmitBytes:       peer.TransmitBytes,
// 		AllowedIPs:          allowedIPs,
// 		ProtocolVersion:     peer.ProtocolVersion,
// 	}
// }

// IsValid check if model is valid
func (a Server) IsValid() []error {
	errs := make([]error, 0)

	// check if the address empty
	if len(a.Address) == 0 {
		errs = append(errs, fmt.Errorf("address is required"))
	}
	// check if the address are valid
	for _, address := range a.Address {
		if !util.IsValidCidr(address) {
			errs = append(errs, fmt.Errorf("address %s is invalid", address))
		}
	}
	// check if the listenPort is valid
	if a.ListenPort < 0 || a.ListenPort > 65535 {
		errs = append(errs, fmt.Errorf("listenPort %d is invalid", a.ListenPort))
	}
	// check if the endpoint empty
	if a.Endpoint == "" {
		errs = append(errs, fmt.Errorf("endpoint is required"))
	}
	// check if the persistentKeepalive is valid
	if a.PersistentKeepalive < 0 {
		errs = append(errs, fmt.Errorf("persistentKeepalive %d is invalid", a.PersistentKeepalive))
	}
	// check if the mtu is valid
	if a.Mtu < 0 {
		errs = append(errs, fmt.Errorf("MTU %d is invalid", a.PersistentKeepalive))
	}
	// check if the address are valid
	for _, dns := range a.DNS {
		if !util.IsValidIP(dns) {
			errs = append(errs, fmt.Errorf("dns %s is invalid", dns))
		}
	}
	// check if the allowedIPs are valid
	for _, allowedIP := range a.AllowedIPs {
		if !util.IsValidCidr(allowedIP) {
			errs = append(errs, fmt.Errorf("allowedIP %s is invalid", allowedIP))
		}
	}

	return errs
}
