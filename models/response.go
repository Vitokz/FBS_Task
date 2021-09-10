package models

type Response struct {
	Numbers string `json:"numbers"`
}

type Err struct{
	Error string `json:"error"`
}