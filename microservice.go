package microtoolkit

import (
	"github.com/DanielDanteDosSantosViana/microtoolkit/service"
	"github.com/fatih/color"
)

func NewService(params ...service.Param) *service.Service {
	return service.NewService(params...)
}

func ServerName(serverName string) service.Param {
	return func(o *service.Params) {
		color.Yellow("Server Name : '%s' \n", serverName)
		o.Server.Module.Params.Module.Name = serverName
	}
}

func Port(port string) service.Param {
	return func(o *service.Params) {
		color.Yellow("Port : '%s' \n", port)
		o.Server.Module.Params.Module.Port = port
	}
}

func HostName(hostName string) service.Param {
	return func(o *service.Params) {
		color.Yellow("Host: '%s' \n", hostName)
		o.Server.Module.Params.Module.HostName = hostName
	}
}
