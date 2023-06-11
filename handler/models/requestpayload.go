package models

type RequestPayload struct {
	Payload map[string]interface{} `json:"payload"`
	Entity  string                 `json:"entity"`
}
