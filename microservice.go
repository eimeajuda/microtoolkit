package microtoolkit

import "github.com/DanielDanteDosSantosViana/microtoolkit/service"

func RegistryRouter(service *service.Service) *Registry {
	return NewRegistry(service)
}

func NewService(params ...service.Param) *service.Service {
	return service.NewService(params...)
}

func ServerName(serverName string) service.Param {
	return func(o *service.Params) {
		o.Server.Name = serverName
	}
}

func Port(port string) service.Param {
	return func(o *service.Params) {
		o.Server.Port = port
	}
}

func HostName(hostName string) service.Param {
	return func(o *service.Params) {
		o.Server.HostName = hostName
	}
}
