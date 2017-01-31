package server

import (
	"fmt"
	"net/http"
	"reflect"
)

type Server struct {
	Name     string
	Port     string
	HostName string
	Handle   http.Handler
}

func (s *Server) Init() {
	/*
		s.server = &http.Server{
			Addr:           ":" + s.Port,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
	*/

}

func (s *Server) Run() {
	http.ListenAndServe(":"+s.Port, nil)
}

func (s *Server) Handler(handler http.Handler) {
	http.Handle("/", handler)
	m := make(map[string]interface{})
	val := reflect.ValueOf(http.DefaultServeMux).Elem()
	valueField := val.Field(1)
	typeField := val.Type().Field(1)
	if typeField.PkgPath != "" && !typeField.Anonymous {
		fmt.Println(typeField.Name)
		fmt.Println(typeField.Type)
		valor := reflect.ValueOf(valueField)
		m[typeField.Name] = valor.String()
		fmt.Println(valor)

		//fmt.Printf("valueField.FieldByIndex(typeField.Index))
		fmt.Println(valueField)
		dump(m)
		//http.muxEntry

	}
}

func dump(m map[string]interface{}) {
	for k, v := range m {
		fmt.Printf("%s : %v\n", k, v)
	}
}

func (s *Server) Stop() {

}
