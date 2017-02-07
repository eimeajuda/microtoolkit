package microtoolkit

import (
	"errors"

	"github.com/DanielDanteDosSantosViana/microtoolkit/router"
	"github.com/DanielDanteDosSantosViana/microtoolkit/service"
)

var (
	ErrorItIsRegistered = errors.New("The service is registered")
	ErrorNotRegistered  = errors.New("The service is not registered")
)

type Registry struct {
	service *service.Service
}

func NewRegistry(service *service.Service) *Registry {
	return &Registry{service}
}
func (reg *Registry) RegistryRouter(params ...router.Param) {
	router := router.NewRouter(params...)
	reg.service.AddRouter(router)
}

func (reg *Registry) Id(idRouter string) router.Param {
	return func(p *router.Params) {
		p.Id = idRouter
	}
}

func (reg *Registry) Method(method string) router.Param {
	return func(p *router.Params) {
		p.Router.Method = method
	}
}

func (reg *Registry) Path(path string) router.Param {
	return func(p *router.Params) {
		p.Router.Path = path
	}
}
func (reg *Registry) UrlSrc(urlSrc string) router.Param {
	return func(p *router.Params) {
		p.Router.UrlSrc = urlSrc
	}
}
func (reg *Registry) UrlDest(urlDest string) router.Param {
	return func(p *router.Params) {
		p.Router.UrlDest = urlDest
	}
}
func ReturnType(returnType string) router.Param {
	return func(p *router.Params) {
		p.Router.ReturnType = returnType
	}
}
func Description(description string) router.Param {
	return func(p *router.Params) {
		p.Router.Description = description
	}
}
