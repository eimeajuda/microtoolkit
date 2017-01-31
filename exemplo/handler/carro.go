package handler

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/microtoolkit"
)

type Carro struct {
	Name  string
	Email string
}

func (h *Carro) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	return
}

func (h *Carro) RegisterRouter(service microtoolkit.Service) error {
	registry.RegistryRouters(
		registry.Path("/oi"),
		registry.Method("POST"),
		registry.Description("Descirpiton"),
		service)

	return nil
}
