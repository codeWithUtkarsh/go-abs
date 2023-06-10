package models

type GenericResponse struct {
	Result   string                 `json:"result"`
	Metadata map[string]interface{} `json:"metadata"`
}
