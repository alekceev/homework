package server

import (
	"context"
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

func (s *WebServer) Start(ctx context.Context) error {
	log.Printf("Server is listening http://%s ...\n", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *WebServer) Stop(ctx context.Context) error {
	s.server.Shutdown(ctx)
	log.Println("Stoping web server...")
	return nil
}
