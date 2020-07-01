package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	router *mux.Router
	port   string
}

// Listen http server
func (s *Server) Listen() {
	log.Fatal(http.ListenAndServe(":"+s.port, s.router))
}

// NewServer generate
func NewServer(port string) *Server {
	appPort := os.Getenv("PORT")
	// set default port if there's
	if appPort != "" {
		port = appPort
	}

	return &Server{
		router: mux.NewRouter(),
		port:   port,
	}
}
