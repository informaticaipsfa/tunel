package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/gasto"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/util"
)

type WMedicina struct{}

//Registrar Militares
func (wcis *WMedicina) Registrar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis gasto.WAltoCosto

	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo("semillerocis")

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&cis)
	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println("Control...")
	fmt.Println(cis.Telefono.Domiciliario)

	cis.Numero = util.CompletarCeros(strconv.Itoa(i), 0, 8)
	cis.Usuario = UsuarioConectado.Login
	cis.Fecha = time.Now()
	cis.Crear()
	var traza fanb.TrazaCIS
	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = "Medicina|" + cis.Numero
	traza.Documento = cis
	traza.Crear()
	M.Tipo = 0
	M.Mensaje = cis.Numero
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
