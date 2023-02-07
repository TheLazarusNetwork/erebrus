package authenticate

type AuthenticateRequest struct {
	FlowId    string `json:"flowId" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type AuthenticatePayload struct {
	StatusDesc string `json:"statusdesc"`
	Token      string `json:"token"`
}
