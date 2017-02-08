package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DanielDanteDosSantosViana/microtoolkit/module"
	"github.com/DanielDanteDosSantosViana/microtoolkit/registry"
	"github.com/fatih/color"
)

var hostDiscovery = os.Getenv("MICRO_DISCOVERY")

type Server struct {
	Module module.Module
	server *http.Server
}

func (s *Server) Init() {
	err := verifyVariableEnv()
	if err != nil {
		log.Panic(err)
	}
	s.server = createConfigServer(s)
	exist := verifyExistinRouters(s)
	if exist {
		registry.Register(s.Module)
	} else {
		color.Blue("Could not find route to authenticate...")
	}

}
func verifyExistinRouters(s *Server) bool {
	if len(s.Module.Params.Module.Routers) > 0 {
		return true
	}
	return false
}
func createConfigServer(s *Server) *http.Server {
	return &http.Server{
		Addr:           ":" + s.Module.Params.Module.Port,
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

func verifyVariableEnv() error {
	if hostDiscovery == "" {
		return fmt.Errorf("Not Found 'MICRO_DISCOVERY' variable")
	}
	return nil
}
