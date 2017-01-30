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

	Home := new(handler.Carro)
	r := mux.NewRouter()
	r.HandleFunc("/oi", Home.ServeHTTP)

	service.Init()
	service.Handler(r)
	service.Run()

}
