package api

import (
	"encoding/json"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/estadistica"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
)

type WPanel struct {
	Data string
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

func (wp *WPanel) CrearReduccion(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var Reduccion estadistica.Reduccion
	Cabecera(w, r)
	// e := json.NewDecoder(r.Body).Decode(&wcis)
	go Reduccion.CrearColeccion("reduccion")
	M.Tipo = 1
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
