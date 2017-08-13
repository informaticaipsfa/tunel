package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/tramitacion"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/util"
)

//WRecibo Familiares
type WCis struct {
	ID        string
	Reembolso tramitacion.ReembolsoMedico
}

//Consultar Militares
func (wcis *WCis) RegistrarReembolso(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var cis cis.CuidadoIntegral
	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo("semillerocis")

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&wcis)
	wcis.Reembolso.FechaCreacion = time.Now()
	wcis.Reembolso.Numero = util.CompletarCeros(strconv.Itoa(i), 0, 8)
	util.Error(e)
	cis.ServicioMedico.Programa.ReembolsoMedico = append(cis.ServicioMedico.Programa.ReembolsoMedico, wcis.Reembolso)
	cis.CrearReembolso(wcis.ID)
	fmt.Println(wcis.ID)
	fmt.Println("obj: ", cis)
	M.Tipo = 0
	j, e := json.Marshal(M)
	w.WriteHeader(http.StatusOK)

	w.Write(j)
	//cis.CrearReembolso(wcis.ID)

	//fmt.Println(r)
}
