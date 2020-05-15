package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/credito"
)

//WCredito API de apoyo a credito
type WCredito struct {
}

//Guardar Salvando datos del credito a un militar
func (wc *WCredito) Guardar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var wPrestamo credito.Solicitud
	Cabecera(w, r)

	err := json.NewDecoder(r.Body).Decode(&wPrestamo)
	if err != nil {
		fmt.Println(err.Error())
		M.Mensaje = "Error de Prestamos " + err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	} else {
		fmt.Println(wPrestamo)
		for _, lst := range wPrestamo.Cuotas {
			fmt.Println(lst.Cuota, " ", lst.Fecha, " ", lst.Interes)
		}
		wPrestamo.NuevoPrestamo()
		M.Mensaje = "Proceso exitoso para el prestamo personal " + wPrestamo.Cedula
		M.Tipo = 1
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}

}
