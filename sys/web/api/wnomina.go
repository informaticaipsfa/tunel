package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util/metodobanco"
)

type WConcepto struct {
	Codigo  string `json:"codigo,omitempty" bson:"codigo"`
	Nombre  string `json:"nombre,omitempty" bson:"nombre"`
	Partida string `json:"partida,omitempty" bson:"partida"`
	Formula string `json:"formula,omitempty" bson:"formula"`
}

type WNomina struct {
	ID          string      `json:"id,omitempty" bson:"id"`
	Nombre      string      `json:"nombre,omitempty" bson:"nombre"`
	Tipo        string      `json:"tipo,omitempty" bson:"tipo"`
	Directiva   string      `json:"directiva,omitempty" bson:"directiva"`
	FechaInicio string      `json:"fechainicio,omitempty" bson:"fechainicio"`
	FechaFin    string      `json:"fechafin,omitempty" bson:"fechafin"`
	Concepto    []WConcepto `json:"Concepto,omitempty" bson:"Concepto"`
}

//Agregar un concepto nuevo
func (N *WNomina) Agregar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var concepto fanb.Concepto
	err := json.NewDecoder(r.Body).Decode(&concepto)
	M.Tipo = 1
	if err != nil {
		fmt.Println("Estoy en un error al insertar", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	fmt.Println(UsuarioConectado.Login)

	j, _ := concepto.Agregar(UsuarioConectado.Login)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Consultar un concepto nuevo
func (N *WNomina) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var concepto fanb.Concepto
	var codigo = mux.Vars(r)
	concepto.Codigo = codigo["id"]
	_, e := concepto.Consultar()
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

//Listar Ver
func (N *WNomina) Listar(w http.ResponseWriter, r *http.Request) {
	var concepto fanb.Concepto
	Cabecera(w, r)
	j, _ := concepto.Listar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Gestionar un concepto nuevo
func (N *WNomina) Gestionar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var codigo = mux.Vars(r)
	oid := codigo["id"]
	estatus := codigo["estatus"]
	url := sys.HostUrlPension + "nominacerrar/" + oid + "/" + estatus
	fmt.Println(url)
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

//WProcesarNomina Nomina
type WProcesarNomina struct {
	ID      int    `json:"id"`
	Estatus int    `json:"estatus"`
	Nombre  string `json:"nombre"`
}

//Procesar un concepto nuevo
func (N *WNomina) Procesar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wProcesar []WProcesarNomina
	url := sys.HostUrlPension + "nominaprocesar"
	errx := json.NewDecoder(r.Body).Decode(&wProcesar)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
		fmt.Println(M.Mensaje)
		M.Tipo = 0
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}
	jsonW, ex := json.Marshal(wProcesar)
	if ex != nil {
		fmt.Println(ex.Error())
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonW))
	if err != nil {

		M.Mensaje = err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	} else {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {

			w.WriteHeader(http.StatusOK)
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

//VerPartidas Militar
func (N *WNomina) VerPartidas(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := sys.HostUrlPension + "verpartida/" + id["id"]
	//fmt.Println(url)
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

//CrearTxt Sistema de generacion bancaria
func (N *WNomina) CrearTxt(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var banfan metodobanco.Banfan
	var archivos metodobanco.Archivos

	var id = mux.Vars(r)
	llave := id["id"]
	banfan.CodigoEmpresa = "0026"
	banfan.NumeroEmpresa = "01770006571100173915"
	banfan.Firma = llave
	banfan.Cantidad, _ = strconv.Atoi(id["cant"])

	banfan.Generar(sys.PostgreSQLPENSION)
	// banfan.Terceros(sys.PostgreSQLPENSION, id["id"])
	M.Mensaje = "Generacion de archivos exitosa "

	if !archivos.ComprimirTxt(llave) {
		M.Mensaje = "La compresion de los archivos presenta problemas"
	}

	M.Tipo = 0
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	return
}

//ListarPendientes Militar
func (N *WNomina) ListarPendientes(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := sys.HostUrlPension + "listartpendientes/" + id["id"]
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

//ListarPagos Militar
func (N *WNomina) ListarPagos(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := sys.HostUrlPension + "listarpagos/" + id["id"]
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

//CuadreBanco Militar
func (N *WNomina) CuadreBanco(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := sys.HostUrlPension + "cuadrebanco/" + id["id"]
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

//Opciones Militar
func (N *WNomina) Opciones(w http.ResponseWriter, r *http.Request) {

	Cabecera(w, r)
	fmt.Println("OPTIONS...")
	//fmt.Println(w, "Saludos")

}
