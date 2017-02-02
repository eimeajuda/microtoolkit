package microtoolkit

import (
	"errors"
	"log"

	"github.com/DanielDanteDosSantosViana/microtoolkit/registry"
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
func (reg *Registry) RegistryRouter(params ...registry.Param) {
	router := registry.NewRouter(params...)
	reg.service.AddRouter(router)
	log.Println(router)
}

func (reg *Registry) Id(idRouter string) registry.Param {
	return func(p *registry.Params) {
		p.Id = idRouter
	}
}

func (reg *Registry) Method(method string) registry.Param {
	return func(p *registry.Params) {
		p.Router.Method = method
	}
}

func (reg *Registry) Path(path string) registry.Param {
	return func(p *registry.Params) {
		p.Router.Path = path
	}
}
func (reg *Registry) UrlSrc(urlSrc string) registry.Param {
	return func(p *registry.Params) {
		p.Router.UrlSrc = urlSrc
	}
}
func (reg *Registry) UrlDest(urlDest string) registry.Param {
	return func(p *registry.Params) {
		p.Router.UrlDest = urlDest
	}
}
func ReturnType(returnType string) registry.Param {
	return func(p *registry.Params) {
		p.Router.ReturnType = returnType
	}
}
func Description(description string) registry.Param {
	return func(p *registry.Params) {
		p.Router.Description = description
	}
}
