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
	Routers  []registry.Router
	server   *http.Server
}

func (s *Server) Init() error {
	s.server = createConfigServer(s)
	exist, err := verifyExistinRouters(s)
	if err != nil {
		return err
	}
	if exist {
		registry.Register(s.Routers, s.Name)
	}

	return nil
}
func verifyExistinRouters(s *Server) (bool, error) {

	if len(s.Routers) > 0 {
		return true, nil
	}
	return false, nil
}
func createConfigServer(s *Server) *http.Server {
	return &http.Server{
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
