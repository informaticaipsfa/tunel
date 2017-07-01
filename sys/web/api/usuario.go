package api

import (
	"net/http"

	"github.com/gesaodin/tunel-ipsfa/mdl/usuario"
	"github.com/gorilla/mux"
)

type WUsuario struct{}

//Consultar
func (u *WUsuario) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var usr usuario.Usuario
	var variable = mux.Vars(r)
	usr.Cedula = variable["id"]
	j, _ := usr.Consultar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
