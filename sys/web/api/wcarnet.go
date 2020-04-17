package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
)

//WCarnet Familiares
type WCarnet struct {
	Cedula      string `json:"cedula"`
	Serial 			string `json:"serial"`
	Motivo      string `json:"motivo"`
	Descripcion string `json:"descripcion"`
}

//Consultar Militares
func (wca *WCarnet) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
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
	e := Carnet.Salvar()
	if e != nil {
		M.Mensaje = e.Error()
		M.Tipo = 0
		return
	}
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Listar Militares
func (wca *WCarnet) Listar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var Carnet sssifanb.Carnet
	var estatus = mux.Vars(r)
	nivel, _ := strconv.Atoi(estatus["id"])

	usuario := UsuarioConectado.Login[:3]
	j, e := Carnet.Listar(nivel, strings.ToUpper(usuario))
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Aprobar Militares
func (wca *WCarnet) Aprobar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	Cabecera(w, r)
	var Carnet sssifanb.Carnet
	var nivel = mux.Vars(r)
	serial := nivel["serial"]
	estatus, _ := strconv.Atoi(nivel["estatus"])
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

//Limpiar Militares
func (wca *WCarnet) Limpiar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	Cabecera(w, r)
	var Carnet sssifanb.Carnet
	var nivel = mux.Vars(r)
	estatus, _ := strconv.Atoi(nivel["estatus"])
	sucursal := strings.ToUpper(nivel["sucursal"])
	e := Carnet.Limpiar(estatus, sucursal)
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

//Liberar Militares
func (wca *WCarnet) Liberar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var Carnet sssifanb.Carnet
	var traza fanb.Traza

	Cabecera(w, r)

	errx := json.NewDecoder(r.Body).Decode(&wca)

	if errx != nil {
		M.Mensaje = errx.Error() + " Falla en el formato JSON"
		fmt.Println(M.Mensaje)
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}

	errx = Carnet.CambiarEstadoMilitar(wca.Serial, 0)
	if errx != nil {
		M.Mensaje = errx.Error() + " falla con el serial del carnet"
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}else{
		ip := strings.Split(r.RemoteAddr, ":")
		M.Tipo = 1

		traza.IP = ip[0]
		traza.Time = time.Now()
		traza.Usuario = UsuarioConectado.Login
		traza.Log = wca.Cedula
		traza.Documento = "Liberar Carnet: " + wca.Motivo + "|" + wca.Descripcion
		traza.CrearHistoricoConsulta("historicoconsultas")
		M.Mensaje = "Felicitaciones! El carnet ha sido liberado..."
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	}


}
