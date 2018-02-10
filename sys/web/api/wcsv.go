package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
)

type CSV struct {
	Codigo string
}

//GCSVSC Generar CSV por Situación y Componente
func (c *CSV) GCSVSC(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var archivo sssifanb.GArchivo
	err := json.NewDecoder(r.Body).Decode(&archivo)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error ", err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}
	j, _ := archivo.ExportarCSV()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
