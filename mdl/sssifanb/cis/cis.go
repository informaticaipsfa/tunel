package cis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/gasto"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/investigacion"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/tramitacion"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
	"gopkg.in/mgo.v2/bson"
)

type CuidadoIntegral struct {
	Investigacion  investigacion.Investigacion `json:"Investigacion" bson:"investigacion"`
	ServicioMedico tramitacion.ServicioMedico  `json:"ServicioMedico" bson:"serviciomedico"`
	Gasto          gasto.GastoFarmaceutico     `json:"Gasto" bson:"gasto"`
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
	util.Error(err)

	// **** Actualizando direccion del militar ****
	direccion := reembolso.Direccion
	dir := make(map[string]interface{})
	dir["persona.direccion.0"] = direccion
	err = c.Update(bson.M{"id": id}, bson.M{"$set": dir})
	util.Error(err)

	teldomi := make(map[string]interface{})
	teldomi["persona.telefono.domiciliario"] = reembolso.Telefono.Domiciliario
	err = c.Update(bson.M{"id": id}, bson.M{"$set": teldomi})
	util.Error(err)

	telmovil := make(map[string]interface{})
	telmovil["persona.telefono.movil"] = reembolso.Telefono.Movil
	err = c.Update(bson.M{"id": id}, bson.M{"$set": telmovil})
	util.Error(err)

	tel := make(map[string]interface{})
	tel["persona.telefono.emergencia"] = reembolso.Telefono.Emergencia
	err = c.Update(bson.M{"id": id}, bson.M{"$set": tel})
	util.Error(err)

	corr := make(map[string]interface{})
	corr["persona.correo"] = reembolso.Correo
	err = c.Update(bson.M{"id": id}, bson.M{"$set": corr})
	util.Error(err)

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
	util.Error(err)

	coleccionfactura := sys.MGOSession.DB(sys.CBASE).C(sys.CFACTURA)
	for _, v := range reembolso.Concepto {
		var factura tramitacion.Factura
		factura = v.DatoFactura
		err = coleccionfactura.Insert(factura)
		util.Error(err)
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
	M.Mensaje = "Actualizando Reembolso"
	M.Tipo = 1
	seguir := make(map[string]interface{})
	valor := "cis.serviciomedico.programa.reembolso." + strconv.Itoa(AReembolso.Posicion)
	fmt.Println(valor)
	seguir[valor] = AReembolso.Reembolso
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": AReembolso.ID}, bson.M{"$set": seguir})
	if err != nil {
		// return
	}
	var rmb tramitacion.ColeccionReembolso

	co := sys.MGOSession.DB(sys.CBASE).C(sys.CREEMBOLSO)
	err = co.Find(bson.M{"id": AReembolso.ID, "numero": AReembolso.Numero}).One(&rmb)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Reembolso", AReembolso.Numero, "Estatus: ", rmb.Estatus)
	rmb.Reembolso.Responsable = AReembolso.Reembolso.Responsable
	rmb.EstatusSeguimiento = AReembolso.Reembolso.Seguimiento.Estatus
	rmb.MontoAprobado = AReembolso.Reembolso.MontoAprobado
	rmb.MontoSolicitado = AReembolso.Reembolso.MontoSolicitado
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
	direccion := apoyo.Direccion
	dir := make(map[string]interface{})

	cedulaf := strings.Split(apoyo.Concepto[0].Afiliado, "-")[0]
	if cedulaf == id {
		dir["persona.direccion.0"] = direccion
		err = c.Update(bson.M{"id": id}, bson.M{"$set": dir})
		util.Error(err)
	} else {
		dir["familiar.$.persona.direccion.0"] = direccion
		_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": cedulaf}, bson.M{"$set": dir})
		util.Error(err)
	}

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
	err = co.Find(bson.M{"id": AAPoyo.ID, "numero": AAPoyo.Numero}).One(&rmb)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// rmb.Estatus = AAPoyo.Apoyo.Estatus
	rmb.Apoyo.Responsable = AAPoyo.Apoyo.Responsable
	rmb.EstatusSeguimiento = AAPoyo.Apoyo.Seguimiento.Estatus
	rmb.MontoAprobado = AAPoyo.Apoyo.MontoAprobado
	rmb.MontoSolicitado = AAPoyo.Apoyo.MontoSolicitado
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

// CrearCarta Actualizando
func (cuidado *CuidadoIntegral) CrearCarta(id string, carta tramitacion.CartaAval, nombre string) (jSon []byte, err error) {
	var M Mensaje

	apo := make(map[string]interface{})

	apo["cis.serviciomedico.programa.cartaaval"] = carta
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": id}, bson.M{"$push": apo})
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	var cartaaval tramitacion.ColeccionCartaAval
	cartaaval.ID = id
	cartaaval.Numero = carta.Numero
	cartaaval.Nombre = nombre
	cartaaval.Usuario = carta.Usuario
	cartaaval.Estatus = 0
	cartaaval.Carta = carta
	cartaaval.FechaCreacion = carta.FechaCreacion
	cartaaval.MontoSolicitado = carta.MontoSolicitado
	cartaaval.FechaAprobado = carta.FechaAprobado
	cartaaval.MontoAprobado = carta.MontoAprobado

	coleccion := sys.MGOSession.DB(sys.CBASE).C(sys.CCARTAAVAL)
	err = coleccion.Insert(cartaaval)
	if err != nil {
		fmt.Println("Error creando reembolso det: ", id)
		// return
	}
	M.Mensaje = carta.Numero
	M.Tipo = 1
	jSon, err = json.Marshal(M)
	return
}

// ListarCarta Actualizando
func (cuidado *CuidadoIntegral) ListarCarta(estatus int) (jSon []byte, err error) {
	// var result []tramitacion.ColeccionReembolso
	var result []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCARTAAVAL)
	err = c.Find(bson.M{"estatus": estatus}).Select(bson.M{"carta": false, "_id": false}).All(&result)
	if err != nil {
		fmt.Println("Err. Carta")
		//return
	}
	// var Lista interface{}
	// Lista = interface{
	// 	Reembolso: result,
	// }

	jSon, err = json.Marshal(result)
	return
}
