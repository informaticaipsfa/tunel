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
