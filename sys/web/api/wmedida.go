package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
)

//WMedidaJudicial Medidas
type WMedidaJudicial struct {
}

//Agregar un concepto nuevo
func (WM *WMedidaJudicial) Agregar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var medida sssifanb.MedidaJudicial
	err := json.NewDecoder(r.Body).Decode(&medida)
	M.Tipo = 1
	if err != nil {
		fmt.Println("Estoy en un error al insertar", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	medida.Usuario = UsuarioConectado.Login
	j, _ := medida.Agregar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Actualizar un concepto nuevo
func (WM *WMedidaJudicial) Actualizar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var medida sssifanb.MedidaJudicial
	var id = mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&medida)
	M.Tipo = 1
	if err != nil {
		fmt.Println("Estoy en un error al insertar", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	medida.Usuario = UsuarioConectado.Login
	j, _ := medida.Actualizar(id["id"])
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Consultar Medida Judica
func (WM *WMedidaJudicial) Consultar(w http.ResponseWriter, r *http.Request) {
	var M fanb.Mensaje
	var MedicinaJudicial sssifanb.MedidaJudicial
	Cabecera(w, r)

	MedicinaJudicial.Agregar()
	w.WriteHeader(http.StatusOK)
	M.Tipo = 1
	M.Mensaje = "Impresion"
	j, _ := json.Marshal(M)
	w.Write(j)
}
