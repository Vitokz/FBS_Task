package models

type Response struct {
	Numbers []Fibonacci `json:"numbers"`
}

type Fibonacci struct {
	Index int  `json:"index"`
	Fibonacci int `json:"value"`
}

type Err struct{
	Error string `json:"error"`
}