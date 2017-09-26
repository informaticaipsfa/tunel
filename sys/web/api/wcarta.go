package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

<<<<<<< HEAD
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/tramitacion"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/util"
=======
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/tramitacion"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/util"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
	"github.com/gorilla/mux"
)

type WCisCarta struct {
	ID     string
	Carta  tramitacion.CartaAval
	Nombre string
}

//Consultar Militares
func (wcis *WCisCarta) Registrar(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	// var M sssifanb.Mensaje
=======
	var M sssifanb.Mensaje
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
	var cis cis.CuidadoIntegral
	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo("semillerocis")

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&wcis)
	// fmt.Println(wcis.Nombre, "---")
	wcis.Carta.FechaCreacion = time.Now()
	wcis.Carta.Usuario = UsuarioConectado.Login
	wcis.Carta.Numero = util.CompletarCeros(strconv.Itoa(i), 0, 8)
	util.Error(e)
<<<<<<< HEAD
	j, e := cis.CrearCarta(wcis.ID, wcis.Carta, wcis.Nombre)
	// M.Tipo = 0
	// j, e := json.Marshal(M)
=======
	cis.CrearCarta(wcis.ID, wcis.Carta, wcis.Nombre)
	M.Tipo = 0
	j, e := json.Marshal(M)
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
	w.WriteHeader(http.StatusOK)

	w.Write(j)
}

//Consultar Militares
func (wcis *WCisCarta) Listar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var variable = mux.Vars(r)
	estatus, _ := strconv.Atoi(variable["id"])
	Cabecera(w, r)
	// fmt.Println("Hola Mundo")
	jSon, _ := cis.ListarCarta(estatus)
	M.Tipo = 0

	w.WriteHeader(http.StatusOK)
	w.Write(jSon)
}

//Consultar Militares
func (wcis *WCisCarta) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
}
