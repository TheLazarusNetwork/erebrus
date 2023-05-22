package challengeid

type GetChallengeIdPayload struct {
	Eula         string `json:"eula,omitempty"`
	ChallengeId  string `json:"challangeId"`
	IsAuthorized bool   `json:"isAuthorized"`
}
