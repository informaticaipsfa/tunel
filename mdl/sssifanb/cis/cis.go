package cis

import (
	"encoding/json"
	"fmt"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/gasto"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/tramitacion"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"gopkg.in/mgo.v2/bson"
)

const (
	CCIS  string = "cis"
	CBASE string = "sssifanb"
)

type CuidadoIntegral struct {
	ServicioMedico tramitacion.ServicioMedico `json:"ServicioMedico" bson:"serviciomedico"`
	Gasto          gasto.GastoFarmaceutico    `json:"Gasto" bson:"gasto"`
}

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
}

func (cuidado *CuidadoIntegral) CrearReembolso(id string) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Creando Reembolso"
	M.Tipo = 1
	reembolso := make(map[string]interface{})
	reembolso["cis.serviciomedico.programa.reembolso"] = cuidado.ServicioMedico.Programa.ReembolsoMedico
	c := sys.MGOSession.DB(CBASE).C("militar")
	err = c.Update(bson.M{"id": id}, bson.M{"$push": reembolso})
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}
	jSon, err = json.Marshal(M)
	return
}
