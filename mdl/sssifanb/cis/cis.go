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

func (cuidado *CuidadoIntegral) CrearReembolso(id string, reembolso tramitacion.Reembolso, telefono tramitacion.Telefono) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Creando Reembolso"
	M.Tipo = 1
	reemb := make(map[string]interface{})

	reemb["cis.serviciomedico.programa.reembolso"] = reembolso
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": id}, bson.M{"$push": reemb})
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	// **** Actualizando direccion del militar ****

	direccion := reembolso.Direccion
	dir := make(map[string]interface{})
	dir["persona.direccion.0"] = direccion

	fmt.Println("Direccion", direccion)
	err = c.Update(bson.M{"id": id}, bson.M{"$set": dir})
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	tel := make(map[string]interface{})
	tel["persona.telefono"] = telefono
	err = c.Update(bson.M{"id": id}, bson.M{"$set": tel})
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	corr := make(map[string]interface{})
	corr["persona.correo"] = reembolso.Correo
	err = c.Update(bson.M{"id": id}, bson.M{"$set": corr})
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	jSon, err = json.Marshal(M)
	return
}

func (cuidado *CuidadoIntegral) ListarReembolso(estatus int) (jSon []byte, err error) {
	// var result []tramitacion.ColeccionReembolso
	var result []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CREEMBOLSO)
	err = c.Find(bson.M{"estatus": estatus}).Select(bson.M{"reembolso": false, "_id": false}).All(&result)
	if err != nil {
		fmt.Println("Err")
		//return
	}
	jSon, err = json.Marshal(result)
	return
}

func (cuidado *CuidadoIntegral) CrearSeguimientoReembolso(id string, numero string, Seguimiento tramitacion.Seguimiento) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Creando Reembolso"
	M.Tipo = 1
	seguir := make(map[string]interface{})

	seguir["cis.serviciomedico.programa.reembolso.seguimiento"] = Seguimiento
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": id, "numero": numero}, bson.M{"$push": seguir})
	if err != nil {
		// return
	}

	return
}
