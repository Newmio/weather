package entity

import "net/http"

type Request struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    interface{}
}

type Response struct{
	Body    []byte
	Headers http.Header
	Status  int
}