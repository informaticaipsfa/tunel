package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
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

//ListarMedida Militar
func (WM *WMedidaJudicial) ListarMedida(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var codigo = mux.Vars(r)
	oid := codigo["id"]

	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "listarmedidajudicialdt/" + oid
	response, err := http.Get(url)
	if err != nil {
		M.Mensaje = err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			M.Mensaje = err.Error()
			M.Tipo = 0
			j, _ := json.Marshal(M)
			w.Write(j)
			return
		}
		defer response.Body.Close()
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}
}
