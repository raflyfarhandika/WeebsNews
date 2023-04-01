package model

type Response struct {
	Error      string
	StatusCode int
	Data       interface{}
}

type LoginResponse struct {
	Response
	Token string `json:"token"`
}