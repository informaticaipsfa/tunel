package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gorilla/mux"
)

//Militar militares
type Militar struct{}

//Consultar Militares
func (p *Militar) Consultar(w http.ResponseWriter, r *http.Request) {
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
func (p *Militar) Actualizar(w http.ResponseWriter, r *http.Request) {

	Cabecera(w, r)
	var dataJSON sssifanb.Militar
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&dataJSON)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error ", err.Error())
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	j, _ := dataJSON.Actualizar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Insertar Militar
func (p *Militar) Insertar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var militar sssifanb.Militar

	fmt.Println("POST...")
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
		return
	}
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
	// fmt.Fprintf(w, "Saludos")
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
