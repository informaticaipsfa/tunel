package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gorilla/mux"
)

type Persona struct {
}

func (p *Persona) Consultar(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Actualizar Datos Generales
func (p *Persona) Actualizar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
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

//Insertar Persona
func (p *Persona) Insertar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("POST...")
	fmt.Fprintf(w, "Saludos")
}

//Eliminar Persona
func (p *Persona) Eliminar(w http.ResponseWriter, r *http.Request) {

}

//Opciones Persona
func (p *Persona) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
	//fmt.Fprintf(w, "Saludos")

}
