package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys/seguridad"
	"github.com/gorilla/mux"
)

var UsuarioConectado seguridad.Usuario

//Militar militares
type Militar struct{}

//Consultar Militares
func (p *Militar) Consultar(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	Cabecera(w, r)
	var dataJSON sssifanb.Militar
	var cedula = mux.Vars(r)
	dataJSON.Persona.DatoBasico.Cedula = cedula["id"]
	j, e := dataJSON.Consultar()
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	//fmt.Println("El usuario ", UsuarioConectado.Nombre, " Esta consultado el documento: ", cedula["id"])
	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = cedula["id"]
	traza.Documento = "Consultando Militar"
	traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Actualizar Datos Generales
func (p *Militar) Actualizar(w http.ResponseWriter, r *http.Request) {

	Cabecera(w, r)
	ip := strings.Split(r.RemoteAddr, ":")
	var dataJSON sssifanb.Militar
	err := json.NewDecoder(r.Body).Decode(&dataJSON)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error ", err.Error())
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	j, _ := dataJSON.Actualizar(UsuarioConectado.Login, ip[0])
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Insertar Militar
func (p *Militar) Insertar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje
	var militar sssifanb.Militar

	// fmt.Println("POST...")
	err := json.NewDecoder(r.Body).Decode(&militar)
	M.Tipo = 1
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error ", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	//e := militar.SalvarMGOI("militares", objeto)
	e := militar.SalvarMGO()
	if e != nil {
		M.Mensaje = e.Error()
		M.Tipo = 0
	}

	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = militar.ID
	traza.Documento = "Agregando: " + militar.Grado.Abreviatura + "|" + militar.Situacion +
		"|" + militar.FechaIngresoComponente.String() + "|" + militar.FechaAscenso.String()
	traza.CrearHistoricoConsulta("hmilitar")

	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Eliminar Militar
func (p *Militar) Eliminar(w http.ResponseWriter, r *http.Request) {

}

//Opciones Militar
func (p *Militar) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
	//fmt.Fprintf(w, "Saludos")

}
