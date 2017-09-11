package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/investigacion"
)

type WFedeVida struct {
}

func (WFe *WFedeVida) Registrar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var fe investigacion.WFedeVida
	e := json.NewDecoder(r.Body).Decode(&fe)
	if e != nil {
		// return
	}

	fmt.Println("Direccion")
	fmt.Println(fe.Direccion)
	j, _ := fe.Crear()

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
