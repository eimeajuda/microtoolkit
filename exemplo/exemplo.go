package main

import (
	"github.com/DanielDanteDosSantosViana/microtoolkit"
	"github.com/DanielDanteDosSantosViana/microtoolkit/exemplo/handler"
	"github.com/gorilla/mux"
)

func main() {

	service := microtoolkit.NewService(
		microtoolkit.ServerName("vimeo.rest"),
		microtoolkit.Port("3035"),
		microtoolkit.HostName("localhost"))

	carro := new(handler.Carro)
	r := mux.NewRouter()
	r.HandleFunc("/oi", carro.ServeHTTP)

	carro.RegisterRouter(service)
	service.Init()
	service.Handler(r)
	service.Run()
}
