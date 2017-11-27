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

type WTratamiento struct{}

func (WT *WTratamiento) Registrar(w http.ResponseWriter, r *http.Request) {
	var M sssifanb.Mensaje
	var tp gasto.WTratamiento
	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo("semillerocis")

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&tp)

	util.Error(e)
	fmt.Println("Entrando TP")
	tp.TratamientoProlongado.Numero = strconv.Itoa(i)
	// cis.CrearReembolso(wcis.ID, wcis.Reembolso, wcis.Telefono, wcis.Nombre)
	tp.Crear()
	M.Tipo = 0
	M.Mensaje = tp.TratamientoProlongado.Numero
	j, e := json.Marshal(M)

	var traza fanb.TrazaCIS
	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = tp.TratamientoProlongado.Numero
	traza.Documento = ""
	traza.Crear()

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
