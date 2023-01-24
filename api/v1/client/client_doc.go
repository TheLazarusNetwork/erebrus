package client

// swagger:response clientSucessResponse
// Response when the operation suceeds.
type ClientSucessResponse struct {
	// in: body
	Body struct {
		// example: 201
		Status int64
		// example: true
		Sucess bool
		// example: sucess message
		Message string
		Body    Client `json:"client"`
	}
}

// swagger:response clientsSucessResponse
// Response for read all clients.
type ClientsSucessResponse struct {
	// in: body
	Body struct {
		// example: 201
		Status int64
		// example: true
		Sucess bool
		// example: sucess message
		Message string
		Body    []Client `json:"clients"`
	}
}

// swagger:response sucessResponse
// Response when the operation suceeds.
type SucessResponse struct {
	// in: body
	Body struct {
		// example: 200
		Status int64
		// example: true
		Sucess bool
		// example: sucess message
		Message string
	}
}

// swagger:response badRequestResponse
// Response when the operation failed with Bad Request.
type BadRequestResponse struct {
	// in:body
	Body struct {
		// example: 400
		Status int64
		// example: false
		Sucess bool
		// example: error message
		Error string
	}
}

// swagger:response unauthorizedResponse
// Response when the operation failed with Bad Request.
type UnauthorizedResponse struct {
	// in:body
	Body struct {
		// example: 401
		Status int64
		// example: false
		Sucess bool
		// example: error message
		Error string
	}
}

// swagger:response serverErrorResponse
// Response when the operation failed with Server Error.
type ServerErrorResponse struct {
	// in:body
	Body struct {
		// example: 500
		Status int64
		// example: false
		Sucess bool
		// example: error message
		Error string
	}
}

// swagger:parameters readClient updateClient deleteClient configClient emailClient
type ClientIDParam struct {
	//The Identifier of the Client
	// in: path
	Id string `json:"id"`
}

// swagger:parameters createClient
type ClientCreateReqparam struct {
	// Requestbody  used for create and update client operations.
	// in: body
	Body ClientReq `json:"client"`
}

// swagger:parameters updateClient
type ClientUpdateReqparam struct {
	// Requestbody  used for create and update client operations.
	// in: body
	Body ClientUpdateReq `json:"client"`
}

// swagger:model
// model for client details.
type Client struct {

	//Client identifier
	// example: 6c8ff96f-ce8a-4c64-a76d-07e9af0b75ab
	UUID string `json:"uuid"`
	//Name of the client
	// example: jon snow
	Name string `json:"name"`
	//Tags for client device
	// example: ["laptop","PC"]
	Tags []string `json:"tags"`
	//Email that the client device belongs
	// example: jonsnow@mail.com
	Email string `json:"email"`
	//Status signal for client
	// example: true
	Enable bool `json:"enable"`
	// example: true
	IgnorePersistentKeepalive bool `json:"ignorePersistentKeepalive"`
	//Preshared key for the client
	// example: twDZk0lehYtst3Zclb+SRniVfoHnug9N6gjxuaipcvc=
	PresharedKey string `json:"presharedKey"`
	//IP addresses allowed to connect
	// example: ["0.0.0.0/0","::/0"]
	AllowedIPs []string `json:"allowedIPs"`
	//Address range client must will assigned
	// example: ["10.0.0.2/32"]
	Address []string `json:"address"`
	//Private key for the client
	// example: KFOyCoR9Eq+LpqT9VzJCilXYmFwhMFw7UDkdRRxoWVg=
	PrivateKey string `json:"privateKey"`
	//Public key for the client
	// example: YeT/lG9L4AeYOHNrkohnmXfljx3/JgThulskllayxi4=
	PublicKey string `json:"publicKey"`
	//Denoting person creates the client
	// example: jonsnow@mail.com
	CreatedBy string `json:"createdBy"`
	// Denoting person updates the client
	// example: jonsnow@mail.com
	UpdatedBy string `json:"updatedBy"`
	//Time the client is created
	// example: 1642409076544
	Created int64 `json:"created"`
	//Time the client is last updated
	// example: 1642409076544
	Updated int64 `json:"updated"`
}

// swagger:model
// model for client details.
type ClientReq struct {

	// required: true
	// example: jon snow
	Name string `json:"name"`
	//Tags for client device
	// required: true
	// example: ["laptop","PC"]
	Tags []string `json:"tags"`
	//Email that the client device belongs
	// required: true
	// example: jonsnow@mail.com
	Email string `json:"email"`
	//Status signal for client
	// required: true
	// example: true
	Enable bool `json:"enable"`
	//IP addresses allowed to connect
	// required: true
	// example: ["0.0.0.0/0","::/0"]
	AllowedIPs []string `json:"allowedIPs"`
	//Address range client must will assigned
	// required: true
	// example: ["10.0.0.0/24"]
	Address []string `json:"address"`
	//Denoting person creates the client
	// required: true
	// example: jonsnow@mail.com
	CreatedBy string `json:"createdBy"`
	// Denoting person updates the client
	// required: true
	// example: jonsnow@mail.com
	UpdatedBy string `json:"updatedBy"`
}

// swagger:model
// model for client details.
type ClientUpdateReq struct {
	//Client identifier
	// required: true
	// example: 6c8ff96f-ce8a-4c64-a76d-07e9af0b75ab
	UUID string `json:"uuid"`
	//Name of the client
	// required: true
	// example: jon snow
	Name string `json:"name"`
	//Tags for client device
	// required: true
	// example: ["laptop","PC"]
	Tags []string `json:"tags"`

	//Email that the client device belongs
	// required: true
	// example: jonsnow@mail.com
	Email string `json:"email"`
	//Status signal for client
	// required: true
	// example: true
	Enable bool `json:"enable"`
	// example: true
	IgnorePersistentKeepalive bool `json:"ignorePersistentKeepalive"`
	//Preshared key for the client
	// example: twDZk0lehYtst3Zclb+SRniVfoHnug9N6gjxuaipcvc=
	PresharedKey string `json:"presharedKey"`
	//IP addresses allowed to connect
	// required: true
	// example: ["0.0.0.0/0","::/0"]
	AllowedIPs []string `json:"allowedIPs"`
	//IP addresses allowed to connect
	// required: true
	// example: ["10.0.0.2/32"]
	Address []string `json:"address"`
	//Private key for the client
	// example: KFOyCoR9Eq+LpqT9VzJCilXYmFwhMFw7UDkdRRxoWVg=
	PrivateKey string `json:"privateKey"`
	//Public key for the client
	// example: YeT/lG9L4AeYOHNrkohnmXfljx3/JgThulskllayxi4=
	PublicKey string `json:"publicKey"`
	//Denoting person creates the client
	// example: jonsnow@mail.com
	CreatedBy string `json:"createdBy"`
	// Denoting person updates the client
	// example: jonsnow@mail.com
	// required: true
	UpdatedBy string `json:"updatedBy"`
	//Time the client is created
	// example: 1642409076544
	Created int64 `json:"created"`
	//Time the client is last updated
	// example: 1642409076544
	Updated int64 `json:"updated"`
}
