package model

type RegionEndpoint struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	ServiceType string `json:"service_type"`
	Endpoint    string `json:"endpoint"`
}
