package main

import (
	"github.com/DanielDanteDosSantosViana/microtoolkit"
	"github.com/DanielDanteDosSantosViana/microtoolkit/exemplo/handler"
	"github.com/gorilla/mux"
)

func main() {

	service := microtoolkit.NewService(
		microtoolkit.ServerName("vimeo.rest"),
		microtoolkit.Port("3035"))

	r := mux.NewRouter()

	Home := new(handler.Carro)
	Home2 := new(handler.Carro)

	r.Handle("/oi", Home)
	r.HandleFunc("/Ola", Home2.ServeHTTP)

	service.Handler(r)
	//service.RegistryHandlers()
	service.Init()
	service.Run()

}
