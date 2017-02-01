package registry

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

var (
	ErrorItIsRegistered = errors.New("The service is registered")
	ErrorNotRegistered  = errors.New("The service is not registered")
)

type RegistryRoute interface {
	RegisterRouter(Service) error
}

type router struct {
	Path        string `json:"path"`
	UrlSrc      string `json:"urlSrc"`
	UrlDest     string `json:"urlDest"`
	Method      string `json:"method"`
	ReturnType  string `json:"returnType"`
	Description string `json:"description"`
}

type Service struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `json:"name"`
	HostName string        `json:"hostname"`
	Port     string        `json:"port"`
	Routers  []router      `json:"routers"`
}
type Param func(*Params)

func RegistryRouter(params ...Param) *Router {
	return newRouter(params...)
}

func register(service Service) error {
	registered, err := isRegistered(service)
	if err != nil {
		return err
	}
	if registered {
		return ErrorItIsRegistered
	}
	return nil
}
func update(service Service) error {
	registered, err := isRegistered(service)
	if err != nil {
		return err
	}

	if !registered {
		return ErrorNotRegistered
	}
	return nil
}

func delete(service Service) {

}

func isRegistered(service Service) (bool, error) {
	return true, nil
}

func Id(idRouter string) Param {
	return func(p *Params) {
		p.Id = idRouter
	}
}

func Method(method string) Param {
	return func(p *Params) {
		p.router.Method = method
	}
}

func Path(path string) Param {
	return func(p *Params) {
		p.router.Path = path
	}
}
func UrlSrc(urlSrc string) Param {
	return func(p *Params) {
		p.router.UrlSrc = urlSrc
	}
}
func UrlDest(urlDest string) Param {
	return func(p *Params) {
		p.router.UrlDest = urlDest
	}
}
func ReturnType(returnType string) Param {
	return func(p *Params) {
		p.router.ReturnType = returnType
	}
}
func Description(description string) Param {
	return func(p *Params) {
		p.router.Description = description
	}
}
func Service(service Service) Param {
	return func(p *Params) {
		service.AddRouter(p.params.Router)
	}
}
