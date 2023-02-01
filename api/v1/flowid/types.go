package flowid

type GetFlowIdPayload struct {
	Eula   string `json:"eula,omitempty"`
	FlowId string `json:"flowId"`
}
