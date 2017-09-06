package api

import (
	"net/http"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gorilla/mux"
)

type APIComponente struct{}

//Consultar Componentess
func (c *APIComponente) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var dataJSON fanb.Componente
	var codigo = mux.Vars(r)

	j, e := dataJSON.Consultar(codigo["id"])
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

type APIEstado struct{}

//Consultar Componentess
func (c *APIEstado) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var dataJSON fanb.Estado
	j, e := dataJSON.ConsultarEstado()
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
