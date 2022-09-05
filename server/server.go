package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(addr string) *http.Server {
	mux := mux.NewRouter().StrictSlash(true)
	mux.NewRoute().HandlerFunc(handleCalRequest).Methods("GET")
	mux.Use(logRequestHandler)
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
