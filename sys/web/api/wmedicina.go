package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/gasto"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
)

type WMedicina struct{}

//Registrar Militares
func (wcis *WMedicina) Registrar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis gasto.WAltoCosto

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&cis)
	if e != nil {
		fmt.Println(e.Error())
	}
	cis.Usuario = UsuarioConectado.Login
	cis.Fecha = time.Now()
	cis.Crear()
	var traza fanb.TrazaCIS
	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = "Medicina"
	traza.Documento = cis
	traza.Crear()
	M.Tipo = 0
	M.Mensaje = "Creado..."
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
