package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/tramitacion"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/util"
	"github.com/gorilla/mux"
)

type WCisCarta struct {
	ID     string
	Carta  tramitacion.CartaAval
	Nombre string
}

//Consultar Militares
func (wcis *WCisCarta) Registrar(w http.ResponseWriter, r *http.Request) {
	// var M sssifanb.Mensaje
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
	j, e := cis.CrearCarta(wcis.ID, wcis.Carta, wcis.Nombre)
	// M.Tipo = 0
	// j, e := json.Marshal(M)
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
