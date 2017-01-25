package discovery

import "errors"

var (
	ErrorItIsRegistered = errors.New("The service is registered")
	ErrorNotRegistered  = errors.New("The service is not registered")
)

type RegistryRoute interface {
	Register(Service) error
}

type Router struct {
	id          string
	urlSrc      string
	urlDest     string
	namemodule  string
	method      string
	hostname    string
	port        string
	returnType  string
	description string
}

type Service struct {
	Name    string
	Adress  string
	Routers []Router
}

func newService(name string) {

}

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
