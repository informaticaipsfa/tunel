package api

import (
	"encoding/json"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/tramitacion"
	"github.com/informaticaipsfa/tunel/util"
)

//WFactura Familiares
type WFactura struct {
	Rif    string
	Numero string
}

//Consultar Militares
func (wfactura *WFactura) Consultar(w http.ResponseWriter, r *http.Request) {

	var Factura tramitacion.Factura
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&wfactura)
	util.Error(e)
	j, _ := Factura.Consultar(wfactura.Rif, wfactura.Numero)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
