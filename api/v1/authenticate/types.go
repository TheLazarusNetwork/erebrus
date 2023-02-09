package authenticate

type AuthenticateRequest struct {
	ChallengeId string `json:"challengeId" binding:"required"`
	Signature   string `json:"signature" binding:"required"`
}

type AuthenticatePayload struct {
	Status  int64  `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
