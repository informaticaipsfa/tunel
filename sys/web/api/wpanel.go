package api

import (
	"encoding/json"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/estadistica"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/util"
)

type WPanel struct {
	Data    string `json:"data"`
	Host    string `json:"host"`
	Origen  string `json:"origen"`
	Paquete string `json:"paquete"`
}

//ListarPendientes Ver
func (wp *WPanel) ListarPendientes(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var Reduccion estadistica.Reduccion
	Cabecera(w, r)
	j, _ := Reduccion.ListarPendientes()
	w.WriteHeader(http.StatusOK)
	M.Tipo = 0
	w.Write(j)
}

//ListarColecciones Ver
func (wp *WPanel) ListarColecciones(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var Reduccion estadistica.Reduccion
	Cabecera(w, r)
	j, _ := Reduccion.ListarColecciones()
	w.WriteHeader(http.StatusOK)
	M.Tipo = 0
	w.Write(j)
}

func (wp *WPanel) ValidarReduccion(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var Reduccion estadistica.Reduccion
	Cabecera(w, r)
	// e := json.NewDecoder(r.Body).Decode(&wcis)
	if Reduccion.ValidarColeccion("reduccion") {
		M.Tipo = 1
	} else {
		M.Tipo = 0
	}

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	M.Tipo = 0
	w.Write(j)
}

//ExtraerReduccion Exportar datos
func (wp *WPanel) ExtraerReduccion(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var Reduccion estadistica.Reduccion
	Cabecera(w, r)
	// e := json.NewDecoder(r.Body).Decode(&wcis)
	if Reduccion.ValidarColeccion("reduccion") {
		M.Tipo = 1
		go Reduccion.ExportarCSV("T")
		go Reduccion.ExportarCSV("F")
		M.Mensaje = "Su proceso está en progreso..."
	} else {
		M.Tipo = 0
	}

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	M.Tipo = 0
	w.Write(j)
}

//CrearReduccion Reduciones
func (wp *WPanel) CrearReduccion(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var Reduccion estadistica.Reduccion
	Cabecera(w, r)
	go Reduccion.CrearColeccion("reduccion")
	M.Tipo = 1
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//ExtraerDatosMySQL Exportar datos
func (wp *WPanel) ExtraerDatosMySQL(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	Cabecera(w, r)
	go sssifanb.ExportarMysql()
	M.Tipo = 0
	M.Mensaje = "Solicitud en proceso"
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Compilar Sevicios
func (wp *WPanel) Compilar(w http.ResponseWriter, r *http.Request) {

	Cabecera(w, r)
	go util.EjecutarScript()
	w.WriteHeader(http.StatusOK)
	s := "Este proceso puede durar unos segundos... por favor espere"
	w.Write([]byte(s))
}

//GitAll Actualizacion de paquetes en el sistema
func (wp *WPanel) GitAll(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&wp)
	util.Error(e)
	jSon, _ := util.GitAll(wp.Paquete, wp.Data, wp.Origen)
	w.WriteHeader(http.StatusOK)
	w.Write(jSon)
}
