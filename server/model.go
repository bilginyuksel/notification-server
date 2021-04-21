package main

import "net/http"

type baseResponse struct {
	code    int
	message string
}

type Notification struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var (
	success          = baseResponse{code: 0, message: "OK"}
	alreadyConnected = baseResponse{code: 204, message: "Handshake already completed"}
	illegalArgument  = baseResponse{code: 403, message: "Illeagal argument exceptiono"}
	failed           = baseResponse{code: 404, message: "Unknown Error"}
)

type customEndpointHandler func(w http.ResponseWriter, r *http.Request, uniqueKey string) interface{}

type endpointHandler func(w http.ResponseWriter, r *http.Request)
