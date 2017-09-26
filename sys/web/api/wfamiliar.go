package api

import (
	"encoding/json"
	"fmt"
	"net/http"

<<<<<<< HEAD
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
=======
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
	"github.com/gorilla/mux"
)

//Familiar Familiares
type WFamiliar struct{}

//Consultar Militares
func (f *WFamiliar) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var dataJSON sssifanb.Militar
	var cedula = mux.Vars(r)
	dataJSON.Persona.DatoBasico.Cedula = cedula["id"]
	fmt.Println(dataJSON.Persona.DatoBasico.Cedula)
	j, e := dataJSON.Consultar()
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Actualizar Datos Generales
func (f *WFamiliar) Actualizar(w http.ResponseWriter, r *http.Request) {

	Cabecera(w, r)

	var dataJSON sssifanb.Familiar
	var M sssifanb.Mensaje
	err := json.NewDecoder(r.Body).Decode(&dataJSON)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error ", err.Error())
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	M.Tipo = 1
	dataJSON.Actualizar()
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Insertar Militar
func (f *WFamiliar) Insertar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var familiar sssifanb.Familiar

	fmt.Println("POST...")
	err := json.NewDecoder(r.Body).Decode(&familiar)
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
	e := familiar.IncluirFamiliar()
	if e != nil {
		M.Mensaje = e.Error()
		M.Tipo = 0
		return
	}
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Opciones Militar
func (f *WFamiliar) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
	//fmt.Fprintf(w, "Saludos")

}
