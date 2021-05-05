package server

import (
	"homework/pkg/interfaces"
	"log"
	"net/http"
)

type WebServer struct {
	server *http.Server
}

var _ interfaces.Service = &WebServer{}

func NewWebServer(server *http.Server) *WebServer {
	return &WebServer{
		server: server,
	}
}

func (s *WebServer) Start() error {
	log.Printf("Server is listening http://%s ...\n", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *WebServer) Stop() error {
	log.Println("Stoping web server...")
	return nil
}
