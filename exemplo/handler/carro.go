package handler

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/microtoolkit"
	"github.com/DanielDanteDosSantosViana/microtoolkit/registry"
)

type Carro struct {
	Name  string
	Email string
}

func (h *Carro) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	return
}

func (h *Carro) RegisterRouter(service microtoolkit.Service) error {
	registry.RegistryRouter(
		registry.Path("/oi"),
		registry.Method("POST"),
		registry.Description("Descirpiton"),
		registry.Service(service))

	return nil
}
