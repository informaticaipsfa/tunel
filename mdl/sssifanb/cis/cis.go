package cis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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

// CrearReembolso Actualizando
func (cuidado *CuidadoIntegral) CrearReembolso(id string, reembolso tramitacion.Reembolso, telefono tramitacion.Telefono, nombre string) (jSon []byte, err error) {
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

	var creembolso tramitacion.ColeccionReembolso
	creembolso.ID = id
	creembolso.Numero = reembolso.Numero
	creembolso.Nombre = nombre
	creembolso.Usuario = reembolso.Usuario
	creembolso.Estatus = 0
	creembolso.Reembolso = reembolso
	creembolso.FechaCreacion = reembolso.FechaCreacion
	creembolso.MontoSolicitado = reembolso.MontoSolicitado
	creembolso.FechaAprobado = reembolso.FechaAprobado
	creembolso.MontoAprobado = reembolso.MontoAprobado

	coleccion := sys.MGOSession.DB(sys.CBASE).C(sys.CREEMBOLSO)
	err = coleccion.Insert(creembolso)
	if err != nil {
		fmt.Println("Error creando reembolso det: ", id)
		// return
	}

	jSon, err = json.Marshal(M)
	return
}

// ListarReembolso Actualizando
func (cuidado *CuidadoIntegral) ListarReembolso(estatus int) (jSon []byte, err error) {
	// var result []tramitacion.ColeccionReembolso
	var result []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CREEMBOLSO)
	err = c.Find(bson.M{"estatus": estatus}).Select(bson.M{"reembolso": false, "_id": false}).All(&result)
	if err != nil {
		fmt.Println("Err")
		//return
	}
	// var Lista interface{}
	// Lista = interface{
	// 	Reembolso: result,
	// }

	jSon, err = json.Marshal(result)
	return
}

// ActualizarReembolso Actualizando
func (cuidado *CuidadoIntegral) ActualizarReembolso(AReembolso tramitacion.ActualizarReembolso) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Creando Reembolso"
	M.Tipo = 1
	seguir := make(map[string]interface{})
	valor := "cis.serviciomedico.programa.reembolso." + strconv.Itoa(AReembolso.Posicion)
	seguir[valor] = AReembolso.Reembolso
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": AReembolso.ID}, bson.M{"$set": seguir})
	if err != nil {
		// return
	}
	var rmb tramitacion.ColeccionReembolso

	co := sys.MGOSession.DB(sys.CBASE).C(sys.CREEMBOLSO)
	err = co.Find(bson.M{"id": AReembolso.ID}).One(&rmb)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rmb.Estatus = AReembolso.Reembolso.Estatus
	rmb.EstatusSeguimiento = AReembolso.Reembolso.Seguimiento.Estatus
	rmb.MontoAprobado = AReembolso.Reembolso.MontoAprobado
	rmb.FechaAprobado = time.Now()

	err = co.Update(bson.M{"id": AReembolso.ID, "numero": AReembolso.Numero}, bson.M{"$set": rmb})
	if err != nil {
		// return
	}
	fmt.Println("Actualizando")
	return
}

// EstatusReembolso Cambiar de Estado
func (cuidado *CuidadoIntegral) EstatusReembolso(E tramitacion.EstatusReembolso) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Estatus del Reembolso"
	M.Tipo = 1

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	estat := make(map[string]interface{})
	estat["cis.serviciomedico.programa.reembolso.$.estatus"] = E.Estatus
	err = c.Update(bson.M{"id": E.ID, "cis.serviciomedico.programa.reembolso.numero": E.Numero}, bson.M{"$set": estat})
	if err != nil {
		fmt.Println(err.Error())
		// return
	}

	co := sys.MGOSession.DB(sys.CBASE).C(sys.CREEMBOLSO)
	esta := make(map[string]interface{})
	esta["estatus"] = E.Estatus
	err = co.Update(bson.M{"id": E.ID, "numero": E.Numero}, bson.M{"$set": esta})
	if err != nil {
		fmt.Println(err.Error())
		// return
	}

	return
}

// CrearApoyo Actualizando
func (cuidado *CuidadoIntegral) CrearApoyo(id string, apoyo tramitacion.Apoyo, nombre string) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Creando Apoyo"
	M.Tipo = 1
	apo := make(map[string]interface{})

	apo["cis.serviciomedico.programa.apoyo"] = apoyo
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": id}, bson.M{"$push": apo})
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	// **** Actualizando direccion del militar ****
	//
	// direccion := reembolso.Direccion
	// dir := make(map[string]interface{})
	// dir["persona.direccion.0"] = direccion
	//
	// fmt.Println("Direccion", direccion)
	// err = c.Update(bson.M{"id": id}, bson.M{"$set": dir})
	// if err != nil {
	// 	fmt.Println("Cedula: " + id + " -> " + err.Error())
	// 	return
	// }

	var capoyo tramitacion.ColeccionApoyo
	capoyo.ID = id
	capoyo.Numero = apoyo.Numero
	capoyo.Nombre = nombre
	capoyo.Usuario = apoyo.Usuario
	capoyo.Estatus = 0
	capoyo.Apoyo = apoyo
	capoyo.FechaCreacion = apoyo.FechaCreacion
	capoyo.MontoSolicitado = apoyo.MontoSolicitado
	capoyo.FechaAprobado = apoyo.FechaAprobado
	capoyo.MontoAprobado = apoyo.MontoAprobado

	coleccion := sys.MGOSession.DB(sys.CBASE).C(sys.CAPOYO)
	err = coleccion.Insert(capoyo)
	if err != nil {
		fmt.Println("Error creando reembolso det: ", id)
		// return
	}

	jSon, err = json.Marshal(M)
	return
}

// ListarApoyo Actualizando
func (cuidado *CuidadoIntegral) ListarApoyo(estatus int) (jSon []byte, err error) {
	// var result []tramitacion.ColeccionReembolso
	var result []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CAPOYO)
	err = c.Find(bson.M{"estatus": estatus}).Select(bson.M{"apoyo": false, "_id": false}).All(&result)
	if err != nil {
		fmt.Println("Err. Apoyo")
		//return
	}
	// var Lista interface{}
	// Lista = interface{
	// 	Reembolso: result,
	// }

	jSon, err = json.Marshal(result)
	return
}

// ActualizarApoyo Actualizando
func (cuidado *CuidadoIntegral) ActualizarApoyo(AAPoyo tramitacion.ActualizarApoyo) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Actualizando Apoyo"
	M.Tipo = 1
	seguir := make(map[string]interface{})
	valor := "cis.serviciomedico.programa.apoyo." + strconv.Itoa(AAPoyo.Posicion)
	seguir[valor] = AAPoyo.Apoyo
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": AAPoyo.ID}, bson.M{"$set": seguir})
	if err != nil {
		// return
	}
	var rmb tramitacion.ColeccionApoyo

	co := sys.MGOSession.DB(sys.CBASE).C(sys.CAPOYO)
	err = co.Find(bson.M{"id": AAPoyo.ID}).One(&rmb)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rmb.Estatus = AAPoyo.Apoyo.Estatus
	rmb.EstatusSeguimiento = AAPoyo.Apoyo.Seguimiento.Estatus
	rmb.MontoAprobado = AAPoyo.Apoyo.MontoAprobado
	rmb.FechaAprobado = time.Now()

	err = co.Update(bson.M{"id": AAPoyo.ID, "numero": AAPoyo.Numero}, bson.M{"$set": rmb})
	if err != nil {
		// return
	}
	fmt.Println("Actualizando")
	return
}

// EstatusApoyo Cambiar de Estado
func (cuidado *CuidadoIntegral) EstatusApoyo(E tramitacion.EstatusApoyo) (jSon []byte, err error) {
	var M Mensaje
	M.Mensaje = "Estatus de los Apoyos"
	M.Tipo = 1

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	estat := make(map[string]interface{})
	estat["cis.serviciomedico.programa.apoyo.$.estatus"] = E.Estatus
	err = c.Update(bson.M{"id": E.ID, "cis.serviciomedico.programa.apoyo.numero": E.Numero}, bson.M{"$set": estat})
	if err != nil {
		fmt.Println(err.Error())
		// return
	}

	co := sys.MGOSession.DB(sys.CBASE).C(sys.CAPOYO)
	esta := make(map[string]interface{})
	esta["estatus"] = E.Estatus
	err = co.Update(bson.M{"id": E.ID, "numero": E.Numero}, bson.M{"$set": esta})
	if err != nil {
		fmt.Println(err.Error())
		// return
	}

	return
}
