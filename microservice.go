package microtoolkit

import (
	"log"

	"github.com/DanielDanteDosSantosViana/microtoolkit/service"
)

func NewService(params ...service.Param) *service.Service {
	return service.NewService(params...)
}

func ServerName(serverName string) service.Param {
	return func(o *service.Params) {
		log.Print(serverName)
		o.Server.Module.Params.Module.Name = serverName
	}
}

func Port(port string) service.Param {
	return func(o *service.Params) {
		log.Print(port)
		o.Server.Module.Params.Module.Port = port
	}
}

func HostName(hostName string) service.Param {
	return func(o *service.Params) {
		log.Print(hostName)
		o.Server.Module.Params.Module.HostName = hostName
	}
}
