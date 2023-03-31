package model

type Response struct {
	Error      string
	StatusCode int
	Data       interface{}
}