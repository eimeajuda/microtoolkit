package microtoolkit

import "net/http"

type Service struct {
	params Params
	init   chan bool
}

func newService(param ...Param) *Service {
	params := newParams(param...)
	return &Service{params, make(chan bool)}
}

func (s *Service) Init() {
	s.params.Server.Init()
}

func (s *Service) Run() {
	s.params.Server.Run()
}

func (s *Service) Handler(handler http.Handler) {
	s.params.Server.Handler(handler)
}
