package server

import (
	"net/http"
	"time"

	"github.com/DanielDanteDosSantosViana/microtoolkit/registry"
)

type Server struct {
	Name     string
	Port     string
	HostName string
	Routers  []*registry.Router
	server   *http.Server
}

func (s *Server) Init() error {
	s.server = &http.Server{
		Addr:           ":" + s.Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return nil
}

func (s *Server) Run() {
	s.server.ListenAndServe()
}

func (s *Server) Handler(handler http.Handler) {
	s.server.Handler = handler
}

func (s *Server) Stop() {

}
