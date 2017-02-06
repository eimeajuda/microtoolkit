package service

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/microtoolkit/registry"
)

type Service struct {
	Params Params
	init   chan bool
}

func NewService(param ...Param) *Service {
	params := newParams(param...)
	return &Service{params, make(chan bool)}
}

func (s *Service) Init() error {
	return s.Params.Server.Init()
}

func (s *Service) Run() {
	s.Params.Server.Run()
}

func (s *Service) Handler(handler http.Handler) {
	s.Params.Server.Handler(handler)
}

func (s *Service) AddRouter(router registry.Router) {
	s.Params.Server.Routers = append(s.Params.Server.Routers, router)
}
