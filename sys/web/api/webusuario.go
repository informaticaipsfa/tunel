package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/util"
)

//WebUsuario Web de usuario
type WebUsuario struct {
	Cedula     string
	Tipo       string
	Componente string
	Clave      string
	Correo     string
}

//Validar de validación de cuentas
func (Wu *WebUsuario) Validar(w http.ResponseWriter, r *http.Request) {

}

//Identificacion de creación de cuentas
func (Wu *WebUsuario) Identificacion(w http.ResponseWriter, r *http.Request) {
	CabeceraW(w, r)
	// var M util.Mensajes
	var iu sssifanb.IdentificacionUsuario
	var Usr WebUsuario
	var j []byte
	// var usr seguridad.Usuario
	// var datos Clave
	e := json.NewDecoder(r.Body).Decode(&Usr)
	util.Error(e)
	// ok := usr.CambiarClave(datos.Login, datos.Clave, datos.Nueva)
	if Usr.Tipo == "TIT" {
		_, mil := iu.BuscarTitular(Usr.Cedula, Usr.Tipo, Usr.Clave, Usr.Componente, Usr.Correo)
		j, _ = json.Marshal(mil)
	} else if Usr.Tipo == "FAM" {
		_, mil := iu.BuscarSobreviviente(Usr.Cedula, Usr.Tipo)
		j, _ = json.Marshal(mil)

	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Opciones Militar
func (Wu *WebUsuario) Opciones(w http.ResponseWriter, r *http.Request) {
	CabeceraW(w, r)
	fmt.Println("OPTIONS...")
}
