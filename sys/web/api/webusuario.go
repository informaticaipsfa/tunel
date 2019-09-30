package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/sys/seguridad"
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
	var M util.Mensajes
	var iu sssifanb.IdentificacionUsuario
	var Usr WebUsuario
	var j []byte
	var estatus int
	var existeUsr seguridad.WUsuario
	// var datos Clave
	e := json.NewDecoder(r.Body).Decode(&Usr)
	util.Error(e)
	existeUsr.Existe(Usr.Cedula)
	fmt.Println(existeUsr)
	if existeUsr.Cedula != "" {
		M.Msj = "El usuario ya se encuentra registrado"
		M.Tipo = 0
		j, _ = json.Marshal(M)
		estatus = http.StatusForbidden
	}
	// ok := usr.CambiarClave(datos.Login, datos.Clave, datos.Nueva)
	switch Usr.Tipo {
	case "TIT":
		_, mil := iu.BuscarTitular(Usr.Cedula, Usr.Tipo, Usr.Clave, Usr.Componente, Usr.Correo)
		j, _ = json.Marshal(mil)
		if mil.Cedula != "" {
			estatus = http.StatusOK
		} else {
			M.Msj = "Alguna de sus respuesta no es correctas"
			M.Tipo = 0
			j, _ = json.Marshal(M)
			estatus = http.StatusForbidden
		}
	case "SOB":
		_, mil := iu.BuscarSobreviviente(Usr.Cedula, Usr.Tipo, Usr.Correo, Usr.Clave)
		j, _ = json.Marshal(mil)
		if mil.Cedula != "" {
			estatus = http.StatusOK
		} else {
			M.Msj = "Alguna de sus respuesta no es correctas"
			M.Tipo = 0
			j, _ = json.Marshal(M)
			estatus = http.StatusForbidden
		}
	case "SOB-TIT":

	default:
		M.Msj = "Error de seleccion"
		M.Tipo = 0
		j, _ = json.Marshal(M)
		estatus = http.StatusForbidden
	}

	w.WriteHeader(estatus)
	w.Write(j)

}

//Opciones Militar
func (Wu *WebUsuario) Opciones(w http.ResponseWriter, r *http.Request) {
	CabeceraW(w, r)
	fmt.Println("OPTIONS...")
}
