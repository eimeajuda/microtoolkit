package microtoolkit

type Param func(*Params)

func NewService(params ...Param) *Service {
	return newService(params...)
}

func ServerName(serverName string) Param {
	return func(o *Params) {
		o.Server.Name = serverName
	}
}

func Port(port string) Param {
	return func(o *Params) {
		o.Server.Port = port
	}
}

func HostName(hostName string) Param {
	return func(o *Params) {
		o.Server.HostName = hostName
	}
}
