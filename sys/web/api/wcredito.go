package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/credito"
)

//WCredito API de apoyo a credito
type WCredito struct {
	Fecha   string `json:"fecha" bson:"fecha"`
	Desde   string `json:"desde" bson:"desde"`
	Hasta   string `json:"hasta" bson:"hasta"`
	Estatus int    `json:"estatus" bson:"estatus"`
}

//Guardar Salvando datos del credito a un militar
func (wc *WCredito) Guardar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var wCredito credito.Solicitud
	Cabecera(w, r)

	err := json.NewDecoder(r.Body).Decode(&wCredito)
	if err != nil {
		fmt.Println(err.Error())
		M.Mensaje = "Error de Prestamos " + err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	var codigo = wCredito.NuevoPrestamo(UsuarioConectado.Login)
	M.Mensaje = codigo
	M.Tipo = 1
	w.WriteHeader(http.StatusOK)
	j, _ := json.Marshal(M)
	w.Write(j)
	return

}

//Listar Creditos
func (wc *WCredito) Listar(w http.ResponseWriter, r *http.Request) {
	// var traza fanb.Traza
	var M sssifanb.Mensaje
	Cabecera(w, r)
	var wCred credito.Credito

	err := json.NewDecoder(r.Body).Decode(&wc)
	if err != nil {
		fmt.Println(err.Error())
		M.Mensaje = "Error de Prestamos " + err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	//fmt.Println("Listando")

	j, e := wCred.Listar(wc.Fecha, wc.Desde, wc.Hasta, wc.Estatus)
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	// ip := strings.Split(r.RemoteAddr, ":")
	// traza.IP = ip[0]
	// traza.Time = time.Now()
	// traza.Usuario = UsuarioConectado.Login
	// traza.Log = cedula["id"]
	// traza.Documento = "Consultando Militar"
	// traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Actualizar Creditos
func (wc *WCredito) Actualizar(w http.ResponseWriter, r *http.Request) {

	var M sssifanb.Mensaje
	Cabecera(w, r)
	var wCred credito.WCreditoActualizar
	var CCredito credito.Credito

	err := json.NewDecoder(r.Body).Decode(&wCred)
	if err != nil {
		fmt.Println(err.Error())
		M.Mensaje = "Error de Prestamos " + err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}

	j, e := CCredito.ActualizarLote(wCred, UsuarioConectado.Login)
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	// ip := strings.Split(r.RemoteAddr, ":")
	// traza.IP = ip[0]
	// traza.Time = time.Now()
	// traza.Usuario = UsuarioConectado.Login
	// traza.Log = cedula["id"]
	// traza.Documento = "Consultando Militar"
	// traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
