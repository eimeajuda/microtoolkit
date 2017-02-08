package handler

import (
	"log"
	"net/http"

	"github.com/DanielDanteDosSantosViana/microtoolkit"
	"github.com/DanielDanteDosSantosViana/microtoolkit/service"
)

type Carro struct {
	Name  string
	Email string
}

func (h *Carro) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Olá")
	return
}

func (h *Carro) RegisterRouter(service *service.Service) error {

	register := microtoolkit.NewRegistry(service)
	register.RegistryRouter(
		register.Method("GET"),
		register.Path("/oi"),
		register.UrlDest("http://www.google.com.br"))

	return nil
}
