package model

type ResponseTemplate struct {
	Error   bool
	Code    int
	Message string
	Data    any
}

type ResponseError struct {
	StatusCode int
	Error      string
	Message    string
}
