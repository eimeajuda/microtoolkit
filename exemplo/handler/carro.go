package handler

import (
	"log"
	"net/http"
)

type Carro struct {
	Name  string
	Email string
}

func (h *Carro) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("passou aqui")
	return
}
