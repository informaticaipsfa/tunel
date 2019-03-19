package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
)

//ConsultarDirectiva Militar
func (p *Militar) ConsultarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	url := "http://localhost/CI-3.1.10/index.php/WServer/directiva"
	//fmt.Println(url);
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

//ConsultarDetalleDirectiva Militar
func (p *Militar) ConsultarDetalleDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var cedula = mux.Vars(r)
	url := "http://localhost/CI-3.1.10/index.php/WServer/dtdirectiva/" + cedula["id"]
	//fmt.Println(url);
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

//ClonarDirectiva Militar
func (p *Militar) ClonarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wNomina WNomina //Modulo de WNomina en API
	url := "http://localhost/CI-3.1.10/index.php/WServer/clonardirectiva"

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

	//fmt.Println("JSON : --> ", wNomina.ID, wNomina.Concepto)

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

//GenerarNomina Militar
func (p *Militar) GenerarNomina(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wNomina WNomina //Modulo de WNomina en API
	url := "http://localhost/CI-3.1.10/index.php/WServer/gnomina"

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

	//fmt.Println("JSON : --> ", wNomina.ID, wNomina.Concepto)

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
