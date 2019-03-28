package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
)

//WDescuentos Medidas
type WDescuentos struct {
}

//Agregar un concepto nuevo
func (WM *WDescuentos) Agregar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var descuentos sssifanb.Descuentos
	err := json.NewDecoder(r.Body).Decode(&descuentos)
	M.Tipo = 1
	if err != nil {
		fmt.Println("Estoy en un error al insertar", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	descuentos.Usuario = UsuarioConectado.Login
	j, _ := descuentos.Agregar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Consultar Medida Judica
func (WM *WDescuentos) Consultar(w http.ResponseWriter, r *http.Request) {
	var M fanb.Mensaje
	var descuentos sssifanb.Descuentos
	Cabecera(w, r)

	descuentos.Agregar()
	w.WriteHeader(http.StatusOK)
	M.Tipo = 1
	M.Mensaje = "Impresion"
	j, _ := json.Marshal(M)
	w.Write(j)
}
