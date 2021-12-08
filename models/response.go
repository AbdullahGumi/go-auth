package models

type Response struct {
	Payload string `json:"payload"`
	Error   string `json:"error"`
}
