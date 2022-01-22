package server

// swagger:response serverSucessResponse
// Response when the operation suceeds.
type ServerSucessResponse struct {
	// in: body
	Body struct {
		// example: 201
		Status int64
		// example: true
		Sucess bool
		// example: sucess message
		Message string
		Body    Server `json:"server"`
	}
}

// swagger:response serverStatusResponse
// Response for Server Status.
type ServerStatusResponse struct {
	// in: body
	Body Status
}

// swagger:parameters updateServer
type ServerUpdateReqparam struct {
	// Requestbody  used for update server operations.
	// in: body
	Body Server `json:"server"`
}

// swagger:model
// model for server details.
type Server struct {
	//Server address
	// example: ["10.0.0.1/24"]
	Address []string `json:"address"`
	//Port the server listens
	// example: 51280
	ListenPort int64 `json:"listenPort"`
	Mtu        int64 `json:"mtu"`
	//Private key for the server
	// example: UFWsgb/Ax5B8zZGx0YtHBAuQVRrOHrxKz2zS2p1LuUE=
	PrivateKey string `json:"privateKey"`
	//Public key for the server
	// example: T5ZMOnik3YuaRhZgAhcxXrmn2+C0B7qFaqnCypMMcks=
	PublicKey string `json:"publicKey"`
	//Endpoint of the server
	// example: region.example.com
	Endpoint string `json:"endpoint"`
	//Persistent keep alive for server
	// example: 16
	PersistentKeepalive int64 `json:"persistentKeepalive"`
	//DNS of the VPN server
	// example: ["1.1.1.1"]
	DNS []string `json:"dns"`
	//IP addresses allowed to connect
	// example: ["0.0.0.0/0","::/0" ]
	AllowedIPs []string `json:"allowedips"`
	//Pre up command
	// example: echo WireGuard PreUp
	PreUp string `json:"preUp"`
	//Post up command
	// example: iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
	PostUp string `json:"postUp"`
	//Pre down command
	// example: echo WireGuard PreDown
	PreDown string `json:"preDown"`
	//Post down command
	// example: iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
	PostDown string `json:"postDown"`
	// Updater email address
	// example: admin@mail.com
	UpdatedBy string `json:"updatedBy"`
	//Time when server is created
	// example: 26103870
	Created int64 `json:"created"`
	//Time when server is created
	// example: 26103870
	Updated int64 `json:"updated"`
}

// swagger:model
// model for server status.
type Status struct {
	//Server version
	// example: 1.0
	Version string `json:"Version,omitempty"`
	//Server Hostname
	// example: ubuntu
	Hostname string `json:"Hostname,omitempty"`
	// Domain which server is running
	// example: vpn.example.com
	Domain string `json:"Domain,omitempty"`
	// Server's public IP
	// example: 14.10.35.65
	PublicIP string `json:"PublicIP,omitempty"`
	// Port which gRPC service is running
	// example: 5000
	GRPCPort string `json:"gRPCPort,omitempty"`
	// Private IP of server host
	// example: 10.0.1.5
	PrivateIP string `json:"PrivateIP,omitempty"`
	// Port which HTTP service is running
	// example: 4000
	HttpPort string `json:"HttpPort,omitempty"`
	// Region where server running
	// example:India/Banglore
	Region string `json:"Region,omitempty"`
	// VPN port
	// example: 5128
	VPNPort string `json:"VPNPort,omitempty"`
}
