package status

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
