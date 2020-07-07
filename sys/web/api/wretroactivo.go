package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/sys"
)

//Retroactivo Datos Formales
type Retroactivo struct {
	Mes      string `json:"mes"`
	Ano      int64  `json:"ano"`
	Tipo     string `json:"tipo"`
	Cedula   string `json:"cedula"`
	Familiar string `json:"familiar"`
}

type XWNomina struct {
	ID          string      `json:"id" bson:"id"`
	Cedula      string      `json:"cedula" bson:"cedula"`
	Nombre      string      `json:"nombre" bson:"nombre"`
	Tipo        string      `json:"tipo" bson:"tipo"`
	Directiva   string      `json:"directiva" bson:"directiva"`
	Fecha       string      `json:"fecha" bson:"fecha"`
	FechaInicio string      `json:"fechainicio" bson:"fechainicio"`
	FechaFin    string      `json:"fechafin" bson:"fechafin"`
	Codigo      string      `json:"codigo" bson:"codigo"`
	Mes         string      `json:"mes" bson:"mes"`
	Concepto    []WConcepto `json:"Concepto,omitempty" bson:"Concepto"`
}

//MesActivo Permite listar los meses activos/generados por nomina
func (Re *Retroactivo) MesActivo(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wProcesar Retroactivo
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "cmesesactivo"
	errx := json.NewDecoder(r.Body).Decode(&wProcesar)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
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

//MesDetalle Permite ver el detalle de un mes de nomina
func (Re *Retroactivo) MesDetalle(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wProcesar Retroactivo
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "cmesdetalle"
	errx := json.NewDecoder(r.Body).Decode(&wProcesar)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
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

//GenerarRetroactivo Militar
func (Re *Retroactivo) GenerarRetroactivo(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wNomina XWNomina //Modulo de WNomina en API
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "gretroactivo"

	errx := json.NewDecoder(r.Body).Decode(&wNomina)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
		M.Tipo = 0
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}
	jsonW, ex := json.Marshal(wNomina)
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
