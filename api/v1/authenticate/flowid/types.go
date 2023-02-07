package flowid

type GetChallengeIdPayload struct {
	Eula        string `json:"eula,omitempty"`
	ChallengeId string `json:"challangeId"`
}
