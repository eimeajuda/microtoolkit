package registry

import (
	"errors"

	"github.com/DanielDanteDosSantosViana/microtoolkit"

	"gopkg.in/mgo.v2/bson"
)

var (
	ErrorItIsRegistered = errors.New("The service is registered")
	ErrorNotRegistered  = errors.New("The service is not registered")
)

type RegistryRoute interface {
	RegisterRouter(microtoolkit.Service) error
}

type Service struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `json:"name"`
	HostName string        `json:"hostname"`
	Port     string        `json:"port"`
	Routers  []Router      `json:"routers"`
}
type Param func(*Params)

func Register(service Service) error {
	registered, err := isRegistered(service)
	if err != nil {
		return err
	}
	if registered {
		return ErrorItIsRegistered
	}
	return nil
}
func Update(service Service) error {
	registered, err := isRegistered(service)
	if err != nil {
		return err
	}

	if !registered {
		return ErrorNotRegistered
	}
	return nil
}

func Delete(service Service) {

}

func isRegistered(service Service) (bool, error) {
	return true, nil
}

func update(id string, newValuesService *Service) {
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
