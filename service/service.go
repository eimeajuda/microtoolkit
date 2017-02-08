package service

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
)

type Service struct {
	Params Params
	init   chan bool
}

func NewService(param ...Param) *Service {
	params := newParams(param...)
	return &Service{params, make(chan bool)}
}

func (s *Service) Init() {
	s.Params.Server.Init()
}

func (s *Service) Run() {
	s.Params.Server.Run()
}

func (s *Service) Handler(handler http.Handler) {
	s.Params.Server.Handler(handler)
}

func (s *Service) AddRouter(router router.RouterP) {
	s.Params.Server.Module.Params.Module.Routers = append(s.Params.Server.Module.Params.Module.Routers, router)
}
