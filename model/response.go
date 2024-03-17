package model

type Response struct {
	StatusCode int
	Message    string
	// Data is all type <interface>
	Data interface{}
}
