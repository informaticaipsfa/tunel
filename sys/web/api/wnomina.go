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
	ID          string      `json:"id" bson:"id"`
	Nombre      string      `json:"nombre" bson:"nombre"`
	Tipo        string      `json:"tipo" bson:"tipo"`
	Directiva   string      `json:"directiva" bson:"directiva"`
	FechaInicio string      `json:"fechainicio" bson:"fechainicio"`
	FechaFin    string      `json:"fechafin" bson:"fechafin"`
	Codigo      string      `json:"codigo" bson:"codigo"`
	Mes         string      `json:"mes" bson:"mes"`
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
	//fmt.Println(UsuarioConectado.Login)
	concepto.Usuario = UsuarioConectado.Login
	j, _ := concepto.Agregar()
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

//ListarPHP Militar
func (N *WNomina) ListarPHP(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "listarconceptos/"
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

//ListarContable Militar
func (N *WNomina) ListarContable(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var codigo = mux.Vars(r)
	oid := codigo["id"]

	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "listarnominadt/" + oid
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

//Gestionar un concepto nuevo
func (N *WNomina) Gestionar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var codigo = mux.Vars(r)
	oid := codigo["id"]
	estatus := codigo["estatus"]
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "nominacerrar/" + oid + "/" + estatus

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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "nominaprocesar"
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

type WLstPagoDetalle struct {
	Llave  string `json:"llave"`
	Codigo string `json:"codigo"`
	Tipo   string `json:"tipo"` //Cuenta de ahorro (CA) y Corriente (CC) o Cheque ()
}

//ListarPagosDetalle un concepto nuevo
func (N *WNomina) ListarPagosDetalle(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var wProcesar WLstPagoDetalle
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "listarpagosdetalles"
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

//VerPartidas Militar
func (N *WNomina) VerPartidas(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "verpartida/" + id["id"]
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
	var banfanb metodobanco.Banfanb
	var bicentenario metodobanco.Bicentenario
	var venzuela metodobanco.Venezuela

	var archivos metodobanco.Archivos

	var id = mux.Vars(r)
	llave := id["id"]
	banfanb.Tabla = "pagos"
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
	bicentenario.Tabla = "pagos"
	bicentenario.CodigoEmpresa = "0651"
	bicentenario.NumeroEmpresa = "01750484310076626369"
	bicentenario.Firma = llave
	bicentenario.Cantidad, _ = strconv.Atoi(id["cant"])
	bicentenario.Generar(sys.PostgreSQLPENSION)

	//
	venzuela.Tabla = "pagos"
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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "listartpendientes/" + id["mes"] + "/" + id["id"]
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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "listarpagos/" + id["id"]
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
func (N *WNomina) VerPagosIndividual(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var id = mux.Vars(r)
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "verpagosindividual/" + id["llav"] + "/" + id["cedu"]
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
	url := "http://" + sys.HostIPPension + sys.HostUrlPension + "cuadrebanco/" + id["id"] + "/" + id["tbl"]
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
