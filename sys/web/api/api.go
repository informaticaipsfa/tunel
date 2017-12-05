package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys/seguridad"
)

var UsuarioConectado seguridad.Usuario

//Militar militares
type Militar struct{}

type Componente struct {
	Componente string
	Grado      string
	Situacion  string
}

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

	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	// fmt.Println("POST...")
	err := json.NewDecoder(r.Body).Decode(&militar)
	M.Tipo = 1
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error al insertar", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	//e := militar.SalvarMGOI("militares", objeto)

	if UsuarioConectado.Login[:3] != "act" {
		e := militar.SalvarMGO()
		if e != nil {
			M.Mensaje = e.Error()
			M.Tipo = 0
		}

		traza.Log = militar.ID
		traza.Documento = "Agregando: " + militar.Grado.Abreviatura + "|" + militar.Situacion +
			"|" + militar.FechaIngresoComponente.String() + "|" + militar.FechaAscenso.String()
		traza.CrearHistoricoConsulta("hmilitar")

	} else {
		M.Mensaje = "Su cuenta no poseé acceso para ingresar nuevos militares"
		M.Tipo = 2
	}

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Eliminar Militar
func (p *Militar) Eliminar(w http.ResponseWriter, r *http.Request) {

}

//EstadisticasPorComponente
func (p *Militar) EstadisticasPorComponente(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	// ip := strings.Split(r.RemoteAddr, ":")
	var militar sssifanb.Militar
	// err := json.NewDecoder(r.Body).Decode(&militar)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println("Estoy en un error ", err.Error())
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }

	j, _ := militar.EstadisticasPorComponente()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Opciones Militar
func (p *Militar) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
	//fmt.Fprintf(w, "Saludos")

}
