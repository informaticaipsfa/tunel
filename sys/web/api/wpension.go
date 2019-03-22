package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

//ListarDetalleDirectiva Militar
func (p *Militar) ListarDetalleDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := "http://localhost/CI-3.1.10/index.php/WServer/ldirectiva/" + id["id"]
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

//WDirectivaClonar Militar
type WDirectivaClonar struct {
	ID               int       `json:"id" bson:"id"`
	Nombre           string    `json:"nombre" bson:"nombre"`
	Numero           string    `json:"numero" bson:"numero"`
	Observacion      string    `json:"observacion" bson:"observacion"`
	FechaInicio      time.Time `json:"fechainicio" bson:"fechainicio"`
	FechaVigencia    time.Time `json:"fechavigencia" bson:"fechavigencia"`
	UnidadTributaria float64   `json:"unidadtributaria" bson:"unidadtributaria"`
	Porcentaje       float64   `json:"porcentaje" bson:"porcentaje"`
	SalarioMinimo    float64   `json:"salariominimo" bson:"salariominimo"`
	Usuario          string    `json:"usuario" bson:"usuario"`
}

//ClonarDirectiva Militar
func (p *Militar) ClonarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wDirectiva WDirectivaClonar //Modulo de WNomina en API
	url := "http://localhost/CI-3.1.10/index.php/WServer/clonardirectiva"

	errx := json.NewDecoder(r.Body).Decode(&wDirectiva)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
		M.Tipo = 0
		fmt.Println(M.Mensaje)
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}

	fmt.Println("JSON : --> ", wDirectiva.FechaInicio, UsuarioConectado.Login)
	wDirectiva.Usuario = UsuarioConectado.Login

	jsonW, ex := json.Marshal(wDirectiva)
	if ex != nil {
		fmt.Println(ex.Error())
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonW))
	//fmt.Println(url)

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

//EliminarDirectiva Militar
func (p *Militar) EliminarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := "http://localhost/CI-3.1.10/index.php/WServer/eliminardirectiva/" + id["id"]

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

//WDirectivaPrima Militar
type WDirectivaPrima struct {
	ID          string `json:"id" bson:"id"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Formula     string `json:"formula" bson:"formula"`
	Usuario     string `json:"usuario" bson:"usuario"`
}

//ActualizarPrima Militar
func (p *Militar) ActualizarPrima(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wDirectivaPrima WDirectivaPrima //Modulo de WNomina en API
	url := "http://localhost/CI-3.1.10/index.php/WServer/actualizarprima"

	errx := json.NewDecoder(r.Body).Decode(&wDirectivaPrima)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
		M.Tipo = 0
		fmt.Println(M.Mensaje)
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}

	fmt.Println("JSON : --> ", UsuarioConectado.Login)
	wDirectivaPrima.Usuario = UsuarioConectado.Login

	jsonW, ex := json.Marshal(wDirectivaPrima)
	if ex != nil {
		fmt.Println(ex.Error())
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonW))
	if err != nil {
		M.Mensaje = err.Error()
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			M.Mensaje = err.Error()
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

//WDirectivaActualizar Militar
type WDirectivaActualizar struct {
	ID     string  `json:"id"`
	Factor float64 `json:"factor"`
	Monto  float64 `json:"monto"`
}

type WDActualizar struct {
	ID                   string                 `json:"json"`
	WDirectivaActualizar []WDirectivaActualizar `json:"directivas"`
}

//ActualizarDirectiva Militar
func (p *Militar) ActualizarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wDirectivaActualizar []WDirectivaActualizar //Modulo de WNomina en API
	url := "http://localhost/CI-3.1.10/index.php/WServer/actualizardirectiva"

	errx := json.NewDecoder(r.Body).Decode(&wDirectivaActualizar)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
		M.Tipo = 0
		fmt.Println("Control->> ", M.Mensaje)
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}
	jsonW, ex := json.Marshal(wDirectivaActualizar)
	if ex != nil {
		fmt.Println("Control--->> ", ex.Error())
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
