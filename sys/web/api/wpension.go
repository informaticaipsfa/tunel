package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
)

//ConsultarDirectiva Militar
func (p *Militar) ConsultarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "directiva"
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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "dtdirectiva/" + cedula["id"]
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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "ldirectiva/" + id["id"]
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

//ConsultarCantidadPensionados Militar
func (p *Militar) ConsultarCantidadPensionados(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "ccpensionados/"
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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "clonardirectiva"

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

	wDirectiva.Usuario = UsuarioConectado.Login

	jsonW, ex := json.Marshal(wDirectiva)
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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "gnomina"

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

//EliminarDirectiva Militar
func (p *Militar) EliminarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "eliminardirectiva/" + id["id"]

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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "actualizarprima"

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

//ActualizarDirectiva Militar
func (p *Militar) ActualizarDirectiva(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wDirectivaActualizar []WDirectivaActualizar //Modulo de WNomina en API
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "actualizardirectiva"

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

//Calculos Militar
func (p *Militar) Calculo(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "calculo/" + id["id"]
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

//ConsultarNeto Militar
func (p *Militar) ConsultarNeto(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	Cabecera(w, r)
	var pension sssifanb.Pension
	var cedula = mux.Vars(r)

	j, e := pension.ConsultarNetos(cedula["id"], true, "", "")
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = cedula["id"]
	traza.Documento = "Consultando Militar"
	traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//ConsultarNetoSobreviviente Sobrevientes
func (p *Militar) ConsultarNetoSobreviviente(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	Cabecera(w, r)
	var pension sssifanb.Pension
	var cedula = mux.Vars(r)

	j, e := pension.ConsultarNetos(cedula["id"], false, cedula["fam"], "")
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = cedula["id"]
	traza.Documento = "Consultando Militar"
	traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//ConsultarNetoWEB Militar
func (p *Militar) ConsultarNetoWeb(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	Cabecera(w, r)
	var pension sssifanb.Pension
	var cedula = mux.Vars(r)

	j, e := pension.ConsultarNetos(cedula["id"], true, "", " AND sn.esta = 10 ")
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = cedula["id"]
	traza.Documento = "Consultando Militar"
	traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//ConsultarNetoSobrevivienteWeb Sobrevientes
func (p *Militar) ConsultarNetoSobrevivienteWeb(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	Cabecera(w, r)
	var pension sssifanb.Pension
	var cedula = mux.Vars(r)

	j, e := pension.ConsultarNetos(cedula["id"], false, cedula["fam"], " AND sn.esta = 10 ")
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = cedula["id"]
	traza.Documento = "Consultando Militar"
	traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//AplicarDerechoACrecer Militar
func (p *Militar) AplicarDerechoACrecer(w http.ResponseWriter, r *http.Request) {
	//var traza fanb.Traza
	var M sssifanb.Mensaje
	Cabecera(w, r)
	var wderecho sssifanb.WDerechoACrecer
	var pension sssifanb.Pension

	errx := json.NewDecoder(r.Body).Decode(&wderecho)
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

	j, e := pension.AplicarDerechoACrecer(wderecho, UsuarioConectado.Login)
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//AplicarDerechoACrecerUpdate Militar
func (p *Militar) AplicarDerechoACrecerUpdate(w http.ResponseWriter, r *http.Request) {
	//var traza fanb.Traza
	var M sssifanb.Mensaje
	Cabecera(w, r)
	var wderecho sssifanb.WDerechoACrecer
	var pension sssifanb.Pension

	errx := json.NewDecoder(r.Body).Decode(&wderecho)
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

	j, e := pension.AplicarDerechoACrecer(wderecho, UsuarioConectado.Login)
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//SituacionPago Militar
func (p *Militar) SituacionPago(w http.ResponseWriter, r *http.Request) {
	//var traza fanb.Traza
	var M sssifanb.Mensaje
	Cabecera(w, r)
	var wact sssifanb.WActualizarPension
	var pension sssifanb.Pension

	errx := json.NewDecoder(r.Body).Decode(&wact)
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

	e := pension.ActualizarSituacion(wact)
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	M.Mensaje = "Proceso exitoso"
	M.Tipo = 1
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//WRetroactivo Militar
type WRetroactivo struct {
	FechaIngreso       string `json:"fingreso"`
	FechaUltimoAscenso string `json:"fascenso"`
	FechaRetiro        string `json:"fretiro"`
	Grado              string `json:"grado"`
	CodigoGrado        string `json:"codigo"`
	Componente         string `json:"componente"`
	AntiguedadGrado    string `json:"antiguedad"`
	TiempoServicio     string `json:"tiempo"`
	NumeroHijos        string `json:"hijos"`
	Porcentaje         string `json:"porcentaje"`
	FechaInicio        string `json:"inicio"`
	FechaFin           string `json:"fin"`
	Usuario            string `json:"usuario"`
	Situacion          string `json:"situacion"`
}

type WARC struct {
	Cedula         string `json:"cedula"`
	CedulaFamiliar string `json:"cedulafamiliar"`
	Anio           string `json:"anio"`
}

//CalcularRetroactivo Militar
func (p *Militar) CalcularRetroactivo(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wGrado fanb.Grado
	var wRetroactivo WRetroactivo
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "calcularretroactivo"

	errx := json.NewDecoder(r.Body).Decode(&wRetroactivo)
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

	grad, comp := wGrado.RetornarCodigo(wRetroactivo.Componente, wRetroactivo.CodigoGrado)
	wRetroactivo.Componente = strconv.Itoa(comp)
	wRetroactivo.CodigoGrado = grad
	wRetroactivo.Usuario = UsuarioConectado.Login

	jsonW, ex := json.Marshal(wRetroactivo)
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

//ImprimirARC Militar
func (p *Militar) ImprimirARC(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje

	var wArc WARC
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "imprimirarc"
	errx := json.NewDecoder(r.Body).Decode(&wArc)
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
	jsonW, ex := json.Marshal(wArc)
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

//GitAll Paquete de Pension
func (p *Militar) GitAll(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var data interface{}
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "gitall"
	fmt.Println(url)

	errx := json.NewDecoder(r.Body).Decode(&data)
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

	jsonW, ex := json.Marshal(data)

	util.Error(ex)

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
