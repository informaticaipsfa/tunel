package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
)

//WRecibo Familiares
type WRecibo struct{}

//Consultar Militares
func (wre *WRecibo) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println(r)

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
func (wre *WRecibo) Insertar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var recibo sssifanb.Recibo
	ip := strings.Split(r.RemoteAddr, ":")

	fmt.Println("Entrando desde: " + ip[0])
	err := json.NewDecoder(r.Body).Decode(&recibo)

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
	recibo.IP = ip[0]
	e := recibo.Salvar()
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
