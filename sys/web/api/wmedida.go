package api

import (
	"encoding/json"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
)

//WMedidaJudicial Medidas
type WMedidaJudicial struct {
}

//Consultar Medida Judica
func (WM *WMedidaJudicial) Consultar(w http.ResponseWriter, r *http.Request) {
	var M fanb.Mensaje
	var MedicinaJudicial sssifanb.MedidaJudicial
	Cabecera(w, r)

	MedicinaJudicial.Agregar()
	w.WriteHeader(http.StatusOK)
	M.Tipo = 1
	M.Mensaje = "Impresion"
	j, _ := json.Marshal(M)
	w.Write(j)
}
