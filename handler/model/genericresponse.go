package model

type GenericResponse struct {
	Result     string                 `json:"result"`
	StatusCode string                 `json:"statuscode"`
	Metadata   map[string]interface{} `json:"metadata"`
}
