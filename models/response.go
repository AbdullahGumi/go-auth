package models

type Response struct {
	Token string `json:"token"`
	Data  User   `json:"user"`
}
