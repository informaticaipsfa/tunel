package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gorilla/mux"
)

//WCarnet Familiares
type WCarnet struct{}

//Consultar Militares
func (wca *WCarnet) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	//fmt.Println(r)

	// var recibo sssifanb.Recibo
	// var cedula = mux.Vars(r)
	// dataJSON.Persona.DatoBasico.Cedula = cedula["id"]
	// fmt.Println(dataJSON.Persona.DatoBasico.Cedula)
	// j, e := dataJSON.Consultar()
	// if e != nil {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	w.Write([]byte("Error al consultar los datos"))
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(j)
}

//Insertar Militar
func (wca *WCarnet) Insertar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var Carnet sssifanb.Carnet

	err := json.NewDecoder(r.Body).Decode(&Carnet)
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
	e := Carnet.Salvar()
	if e != nil {
		M.Mensaje = e.Error()
		M.Tipo = 0
		return
	}
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
	//fmt.Fprintf(w, "Saludos")
}

//Listar Militares
func (wca *WCarnet) Listar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var Carnet sssifanb.Carnet
	var estatus = mux.Vars(r)
	nivel, _ := strconv.Atoi(estatus["id"])
	j, e := Carnet.Listar(nivel)
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Listar Militares
func (wca *WCarnet) Aprobar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	Cabecera(w, r)
	var Carnet sssifanb.Carnet
	var nivel = mux.Vars(r)
	estatus, _ := strconv.Atoi(nivel["estatus"])
	serial := nivel["serial"]
	e := Carnet.CambiarEstado(serial, estatus)
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	M.Tipo = 1
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
