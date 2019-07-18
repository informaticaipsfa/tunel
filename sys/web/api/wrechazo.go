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
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util/metodobanco"
)

//WRechazos Rechazos en el sistema
type WRechazos struct {
	Codigo string `json:"codigo" bson:"codigo"`
	Banco  string `json:"banco" bson:"banco"`
	Tipo   string `json:"tipo" bson:"tipo"`
	Cuenta string `json:"cuenta" bson:"cuenta"`
}

//Agregar Rechazos al sistema bancario
func (R *WRechazos) Agregar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wRechazos WRechazos //Modulo de WNomina en API
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "rechazosagregar"

	errx := json.NewDecoder(r.Body).Decode(&wRechazos)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
		M.Tipo = 0
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}
	jsonW, ex := json.Marshal(wRechazos)
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

func (R *WRechazos) Listar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "rechazoslistar/" + id["id"]

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
func (R *WRechazos) CrearTxt(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var banfanb metodobanco.Banfanb
	var bicentenario metodobanco.Bicentenario
	var venzuela metodobanco.Venezuela

	var archivos metodobanco.Archivos
	tabla := "pension"
	var id = mux.Vars(r)
	llave := id["id"]
	banfanb.Tabla = tabla
	banfanb.Firma = llave
	banfanb.CodigoEmpresa = "0026"
	banfanb.NumeroEmpresa = "01770006571100173915"
	banfanb.Cantidad, _ = strconv.Atoi(id["cant"])
	banfanb.Generar(sys.PostgreSQLPENSION)

	banfanb.CodigoEmpresa = "0026"
	banfanb.NumeroEmpresa = "01770001421100683232"
	banfanb.Tercero(sys.PostgreSQLPENSION, "0134") //BANESCO

	banfanb.CodigoEmpresa = "0026"
	banfanb.NumeroEmpresa = "01770001441100683245"
	banfanb.Tercero(sys.PostgreSQLPENSION, "0108") //PROVINCIAL

	banfanb.CodigoEmpresa = "0026"
	banfanb.NumeroEmpresa = "01770001411100683238" //TESORO
	banfanb.Tercero(sys.PostgreSQLPENSION, "0163")
	//
	banfanb.CodigoEmpresa = "0026"
	banfanb.NumeroEmpresa = "01770001411100683233"
	banfanb.Tercero(sys.PostgreSQLPENSION, "0105") // MERCANTIL
	// //
	bicentenario.Tabla = tabla
	bicentenario.CodigoEmpresa = "0651"
	bicentenario.NumeroEmpresa = "01750484310076626369"
	bicentenario.Firma = llave
	bicentenario.Cantidad, _ = strconv.Atoi(id["cant"])
	bicentenario.Generar(sys.PostgreSQLPENSION)

	//
	venzuela.Tabla = tabla
	venzuela.CodigoEmpresa = "0"
	venzuela.NumeroEmpresa = "01020488720000002147"
	venzuela.Firma = llave
	venzuela.Cantidad, _ = strconv.Atoi(id["cant"])
	venzuela.Generar(sys.PostgreSQLPENSION, "CA")

	//
	venzuela.CodigoEmpresa = "0"
	venzuela.NumeroEmpresa = "01020488720000002147"
	venzuela.Firma = llave
	venzuela.Cantidad, _ = strconv.Atoi(id["cant"])
	venzuela.Generar(sys.PostgreSQLPENSION, "CC")

	//Comprimir todos los archivos en uno para su descarga
	valor := llave + "-XR"
	M.Mensaje = "Generacion de archivos exitosa "
	if !archivos.ComprimirTxt(valor) {
		M.Mensaje = "La compresion de los archivos presenta problemas"
	}

	M.Tipo = 0
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	return
}
