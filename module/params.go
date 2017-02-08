package module

import "github.com/DanielDanteDosSantosViana/microtoolkit/router"

type ParamsM struct {
	Id     string
	Module module
}

type module struct {
	Name     string           `json:"name"`
	HostName string           `json:"hostname"`
	Port     string           `json:"port"`
	Routers  []router.RouterP `json:"routers"`
}
type ParamM func(*ParamsM)

func newParamsM(param ...ParamM) ParamsM {
	params := ParamsM{}
	for _, p := range param {
		p(&params)
	}
	return params
}
