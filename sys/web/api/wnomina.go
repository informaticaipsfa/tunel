package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
)

type WConcepto struct {
	Codigo  string `json:"codigo,omitempty" bson:"codigo"`
	Nombre  string `json:"nombre,omitempty" bson:"nombre"`
	Partida string `json:"partida,omitempty" bson:"partida"`
	Formula string `json:"formula,omitempty" bson:"formula"`
}

type WNomina struct {
	ID        string      `json:"id,omitempty" bson:"id"`
	Directiva string      `json:"directiva,omitempty" bson:"directiva"`
	Fecha     string      `json:"fecha,omitempty" bson:"fecha"`
	Concepto  []WConcepto `json:"Concepto,omitempty" bson:"Concepto"`
}

//Agregar un concepto nuevo
func (N *WNomina) Agregar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	var concepto fanb.Concepto
	err := json.NewDecoder(r.Body).Decode(&concepto)
	M.Tipo = 1
	if err != nil {
		fmt.Println("Estoy en un error al insertar", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	fmt.Println(UsuarioConectado.Login)

	j, _ := concepto.Agregar(UsuarioConectado.Login)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//Consultar un concepto nuevo
func (N *WNomina) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M sssifanb.Mensaje
	M.Tipo = 1
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Listar Ver
func (N *WNomina) Listar(w http.ResponseWriter, r *http.Request) {
	var concepto fanb.Concepto
	Cabecera(w, r)
	j, _ := concepto.Listar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Opciones Militar
func (N *WNomina) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")
	//fmt.Println(w, "Saludos")

}
