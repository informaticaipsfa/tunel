package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/credito"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util/metodobanco"
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

//EnviarATesoreria Creditos
func (wc *WCredito) EnviarATesoreria(w http.ResponseWriter, r *http.Request) {

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

	j, e := CCredito.EnviarATesoreria(wCred, UsuarioConectado.Login)
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

//Liquidar Salvando datos del credito a un militar
func (wc *WCredito) Liquidar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var wCredito credito.Credito

	var wLiquidar credito.WLiquidar
	Cabecera(w, r)

	err := json.NewDecoder(r.Body).Decode(&wLiquidar)
	if err != nil {
		fmt.Println(err.Error())
		M.Mensaje = "Error de Prestamos " + err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}

	wCredito.Liquidar(wLiquidar, UsuarioConectado.Login)
	M.Mensaje = wLiquidar.Credito
	M.Tipo = 1
	w.WriteHeader(http.StatusOK)
	j, _ := json.Marshal(M)
	w.Write(j)
	return

}

//CrearTxt Sistema de generacion bancaria
func (wc *WCredito) CrearTxt(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var wcob metodobanco.Cobranza
	var rcob []metodobanco.CobranzaDetalle

	var id = mux.Vars(r)
	ano := id["ano"]
	mes := id["mes"]
	desde := ano + "-" + mes + "-01"
	hasta := ano + "-" + mes + "-30"

	lsta := wcob.GenerarCobranza(sys.PostgreSQLPENSIONSIGESP, desde, hasta, "GN")
	rcob = append(rcob, lsta)

	lstb := wcob.GenerarCobranza(sys.PostgreSQLPENSIONSIGESP, desde, hasta, "AV")
	rcob = append(rcob, lstb)

	lstc := wcob.GenerarCobranza(sys.PostgreSQLPENSIONSIGESP, desde, hasta, "AR")
	rcob = append(rcob, lstc)

	lstd := wcob.GenerarCobranza(sys.PostgreSQLPENSIONSIGESP, desde, hasta, "EJ")
	rcob = append(rcob, lstd)

	j, _ := json.Marshal(rcob)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	return
}
