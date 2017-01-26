package server

import (
	"net/http"
	"time"
)

type Server struct {
	Name     string
	Port     string
	HostName string
	server   *http.Server
}

func (s *Server) Init() {
	s.server = &http.Server{
		Addr:           ":" + s.Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (s *Server) Run() {
	s.server.ListenAndServe()
}

func (s *Server) Handler(handler http.Handler) {

	s.server.Handler = handler

}

func (s *Server) Stop() {

}
