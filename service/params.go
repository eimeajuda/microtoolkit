package service

import "github.com/DanielDanteDosSantosViana/microtoolkit/server"

type Params struct {
	Id     string
	Server server.Server
}
type Param func(*Params)

func newParams(param ...Param) Params {
	params := Params{}
	for _, p := range param {
		p(&params)
	}
	return params
}
