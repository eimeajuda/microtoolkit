package main

import "github.com/DanielDanteDosSantosViana/microtoolkit"

func main() {

	service := microtoolkit.NewService(
		microtoolkit.ServerName("vimeo.rest"),
		microtoolkit.Port("3035"))

	service.Init()
	service.Run()
}
