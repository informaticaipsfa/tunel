package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/tramitacion"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/util"
	"github.com/gorilla/mux"
)

//WRecibo Familiares
type WCis struct {
	ID        string
	Reembolso tramitacion.Reembolso
	Telefono  tramitacion.Telefono
	Nombre    string
}

type WCisApoyo struct {
	ID     string
	Apoyo  tramitacion.Apoyo
	Nombre string
}

//Consultar Militares
func (wcis *WCis) Registrar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo("semillerocis")

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&wcis)
	wcis.Reembolso.FechaCreacion = time.Now()
	wcis.Reembolso.Usuario = UsuarioConectado.Login
	wcis.Reembolso.Numero = util.CompletarCeros(strconv.Itoa(i), 0, 8)
	util.Error(e)
	cis.CrearReembolso(wcis.ID, wcis.Reembolso, wcis.Telefono, wcis.Nombre)
	M.Tipo = 0
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}

// ActualizarReembolso Militares
func (wcis *WCis) Actualizar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var reemb tramitacion.ActualizarReembolso

	// fmt.Println("Acceso...")
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&reemb)

	for _, v := range reemb.Observaciones {
		var Obs tramitacion.Observacion
		Obs.Contenido = v
		//fmt.Println(Obs.Contenido)
		Obs.FechaCreacion = time.Now()
		Obs.Usuario = UsuarioConectado.Login
		reemb.Reembolso.Seguimiento.Observaciones = append(reemb.Reembolso.Seguimiento.Observaciones, Obs)
	}

	reemb.Reembolso.Usuario = UsuarioConectado.Login
	util.Error(e)
	cis.ActualizarReembolso(reemb)
	M.Tipo = 0
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}

//Consultar Militares
func (wcis *WCis) ListarReembolso(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var variable = mux.Vars(r)
	estatus, _ := strconv.Atoi(variable["id"])
	Cabecera(w, r)
	// fmt.Println("Hola Mundo")
	jSon, _ := cis.ListarReembolso(estatus)
	M.Tipo = 0

	w.WriteHeader(http.StatusOK)
	w.Write(jSon)
}

// ActualizarReembolso Militares
func (wcis *WCis) Estatus(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var Estatus tramitacion.EstatusReembolso

	// fmt.Println("Acceso Estatus...")
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&Estatus)

	util.Error(e)
	cis.EstatusReembolso(Estatus)
	M.Tipo = 0
	M.Mensaje = "Estatus actualizado"
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}

//Consultar Militares
func (wcis *WCis) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
}

//----------------------------------------
// APOYOS

//Registrar Militares por Apoyo
func (wcis *WCisApoyo) Registrar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo("semillerocis")

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&wcis)
	wcis.Apoyo.FechaCreacion = time.Now()
	wcis.Apoyo.Usuario = UsuarioConectado.Login
	wcis.Apoyo.Numero = util.CompletarCeros(strconv.Itoa(i), 0, 8)
	util.Error(e)
	cis.CrearApoyo(wcis.ID, wcis.Apoyo, wcis.Nombre)
	M.Tipo = 0
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}

// ActualizarReembolso Militares
func (wcis *WCisApoyo) Estatus(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var Estatus tramitacion.EstatusApoyo

	// fmt.Println("Acceso Estatus Apoyo...")
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&Estatus)

	util.Error(e)
	cis.EstatusApoyo(Estatus)
	M.Tipo = 0
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}

// ActualizarReembolso Militares
func (wcis *WCisApoyo) Actualizar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var apoyo tramitacion.ActualizarApoyo

	// fmt.Println("Acceso...")
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&apoyo)

	for _, v := range apoyo.Observaciones {
		var Obs tramitacion.Observacion
		Obs.Contenido = v
		// fmt.Println(Obs.Contenido)
		Obs.FechaCreacion = time.Now()
		Obs.Usuario = UsuarioConectado.Login
		apoyo.Apoyo.Seguimiento.Observaciones = append(apoyo.Apoyo.Seguimiento.Observaciones, Obs)
	}

	apoyo.Apoyo.Usuario = UsuarioConectado.Login
	util.Error(e)
	cis.ActualizarApoyo(apoyo)
	M.Tipo = 0
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}

//Consultar Militares
func (wcis *WCisApoyo) ListarApoyo(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var variable = mux.Vars(r)
	estatus, _ := strconv.Atoi(variable["id"])
	Cabecera(w, r)
	// fmt.Println("Hola Mundo")
	jSon, _ := cis.ListarApoyo(estatus)
	M.Tipo = 0

	w.WriteHeader(http.StatusOK)
	w.Write(jSon)
}

//Consultar Militares
func (wcis *WCisApoyo) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
}
