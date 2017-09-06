package sssifanb

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
)

//Familiar Busquedas
type Familiar struct {
	ID              string    `json:"id" bson:"id"`
	Persona         Persona   `json:"Persona" bson:"persona"`
	FechaAfiliacion time.Time `json:"fechaafiliacion" bson:"fechaafiliacion"`
	Parentesco      string    `json:"parentesco" bson:"parentesco"` //0:Mama, 1:papa, 2: Esposa  3: hijo
	EsMilitar       bool      `json:"esmilitar" bson:"esmilitar"`
	Condicion       int       `json:"condicion" bson:"condicion"` //Sano o Condicion especial
	Estudia         int       `json:"estudia" bson:"estudia"`
	Benficio        bool      `json:"beneficio" bson:"beneficio"` //
	Documento       int       `json:"documento" bson:"documento"`
	Adoptado        bool      `json:"adoptado" bson:"adoptado"`
	DocumentoPadre  string    `json:"documentopadre" bson:"documentopadre"`
	HistoriaMedica  string    `json:"historiamedica" bson:"historiamedica"`
	Donante         string    `json:"donante" bson:"donante"`
	EstatusCarnet   int       `json:"estatuscarnet" bson:"estatuscarnet"`
	GrupoSanguineo  string    `json:"gruposanguineo" bson:"gruposanguineo"`
	TIF             Carnet    `json:"Tif" bson:"tif"`

	// EstatusAfiliacion string `json:"estatus" bson:"adoptado"`
	// RazonAfiliacion   string `json:"adoptado" bson:"adoptado"`
}

//AplicarReglasBeneficio OJO SEGUROS HORIZONTES
func (f *Familiar) AplicarReglasBeneficio() {
	edad, _, _ := util.CalcularTiempo(f.Persona.DatoBasico.FechaNacimiento)
	if f.Parentesco == "HJ" {
		f.Benficio = false

		if f.Condicion == 1 {
			f.Benficio = true
		} else {
			if edad < 18 {
				f.Benficio = true
			} else if f.Estudia == 1 && edad < 27 {
				f.Benficio = true
			}
		}
	} else { // ESPOSA Y PADRES
		f.Benficio = true
	}

}

//AplicarReglasParentesco Reglas
func (f *Familiar) AplicarReglasParentesco() {

}

//ConvertirFechaHumano Validacion
func (f *Familiar) ConvertirFechaHumano() {

}

//Actualizar Vida Militar
func (f *Familiar) MGOActualizar() (jSon []byte, err error) {
	var msj Mensaje
	//f.TipoDato = 0

	// s := `UPDATE personas SET nombreprimero='` +
	// 	m.Persona.DatoBasico.NombrePrimero +
	// 	`', nombresegundo='` +
	// 	m.Persona.DatoBasico.NombreSegundo +
	// 	`' WHERE codnip='` + m.Persona.DatoBasico.Cedula + `'`
	// _, err = sys.PostgreSQLSAMAN.Exec(s)
	// if err != nil {
	// 	msj.Mensaje = "Error: Consulta ya existe."
	// 	msj.Tipo = 2
	// 	msj.Pgsql = err.Error()
	// 	jSon, err = json.Marshal(msj)
	// 	return
	// }
	msj.Mensaje = "Su data ha sido actualizada."
	msj.Tipo = 2
	jSon, err = json.Marshal(msj)
	f.MGOActualizar()
	return
}

//MGOActualizar Actualizando en MONGO
func (f *Familiar) Actualizar() (jSon []byte, err error) {

	id := f.Persona.DatoBasico.Cedula
	familiar := make(map[string]interface{})
	familiar["familiar.$.persona"] = f.Persona
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	if f.ID != id {
		fmt.Println("Cambio de Cedula de un familiar: ", id)
		_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": f.ID}, bson.M{"$set": familiar})
	} else {
		_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": familiar})
	}
	if err != nil {
		log.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	fechaafiliacion := make(map[string]interface{})
	fechaafiliacion["familiar.$.fechaafiliacion"] = time.Now()
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": fechaafiliacion})
	if err != nil {
		fmt.Println("Incluyendo parentesco eRR Cedula: " + id + " -> " + err.Error())
	}

	//
	parentesco := make(map[string]interface{})
	parentesco["familiar.$.parentesco"] = f.Parentesco
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": parentesco})
	if err != nil {
		fmt.Println("Incluyendo parentesco eRR Cedula: " + id + " -> " + err.Error())
	}

	beneficio := make(map[string]interface{})
	beneficio["familiar.$.beneficio"] = f.Benficio
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": beneficio})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
	}
	historia := make(map[string]interface{})
	historia["familiar.$.historiamedica"] = f.HistoriaMedica
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": historia})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
	}
	donante := make(map[string]interface{})
	donante["familiar.$.donante"] = f.Donante
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": donante})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
	}

	documentopadre := make(map[string]interface{})
	documentopadre["familiar.$.documentopadre"] = f.DocumentoPadre
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": documentopadre})
	if err != nil {
		fmt.Println("Incluyendo parentesco eRR Cedula: " + id + " -> " + err.Error())
	}

	condicion := make(map[string]interface{})
	condicion["familiar.$.condicion"] = f.Condicion
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": condicion})
	if err != nil {
		fmt.Println("Incluyendo parentesco eRR Cedula: " + id + " -> " + err.Error())
	}

	return
}

//AplicarReglasCarnetHijos Reglas
func (f *Familiar) AplicarReglasCarnetHijos() (TIM Carnet) {
	var mes, dia string
	carnet := make(map[string]interface{})
	fechaActual := time.Now()
	Anio, Mes, Dia := fechaActual.Date()
	edad, _, _ := util.CalcularTiempo(f.Persona.DatoBasico.FechaNacimiento)
	layOut := "2006-01-02"

	switch {
	case edad < 15:
		Anio += 5
	case edad >= 15 && edad <= 18:
		Anio += 3
	case edad > 18 && edad <= 27:
		Anio += 2
	case edad > 18 && f.Condicion == 1:
		Anio += 5
	}

	mes = strconv.Itoa(int(Mes))
	if int(Mes) < 10 {
		mes = "0" + strconv.Itoa(int(Mes))
	}
	dia = strconv.Itoa(Dia)
	if Dia < 10 {
		dia = "0" + strconv.Itoa(Dia)
	}
	fecha := strconv.Itoa(Anio) + "-" + mes + "-" + dia

	fechaVencimientoCarnet, _ := time.Parse(layOut, fecha)
	TIM.Serial = TIM.GenerarSerial()
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimientoCarnet
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.Tipo = 0
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hijos")
	}
	return
}

//AplicarReglasCarnetHijos Reglas
func (f *Familiar) AplicarReglasCarnetHermanos() (TIM Carnet) {
	var mes, dia string
	carnet := make(map[string]interface{})
	fechaActual := time.Now()
	Anio, Mes, Dia := fechaActual.Date()
	layOut := "2006-01-02"
	Anio += 1
	mes = strconv.Itoa(int(Mes))
	if int(Mes) < 10 {
		mes = "0" + strconv.Itoa(int(Mes))
	}
	dia = strconv.Itoa(Dia)
	if Dia < 10 {
		dia = "0" + strconv.Itoa(Dia)
	}
	fecha := strconv.Itoa(Anio) + "-" + mes + "-" + dia
	fechaVencimientoCarnet, _ := time.Parse(layOut, fecha)
	TIM.Serial = TIM.GenerarSerial()
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimientoCarnet
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.Tipo = 0
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hermanos")
	}
	return
}

//IncluirFamiliar Agregar
func (f *Familiar) IncluirFamiliar() (err error) {
	familiar := make(map[string]interface{})
	familiar["familiar"] = f
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	// var fam map[string]interface{}
	var parametro string
	//
	// fmt.Println("cedula", f.DocumentoPadre, " familiar ", f.Persona.DatoBasico.Cedula)
	// donde := bson.M{"id": f.DocumentoPadre, "familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula}
	// e := c.Find(donde).Select(bson.M{"familiar.persona.datobasico.cedula": true, "_id": 0}).One(&fam)
	// if e != nil {
	// 	fmt.Println("Insertando...", e.Error())
	// 	// return
	// }

	// m := fam.(map[string]map[string]map[string]map[string]string)
	// fmt.Println(fam)
	// fmt.Println("Actualizando...")
	// fami := fam["familiar"]
	// for fami {
	// 	fmt.Println(c, v)
	//
	// }

	parametro = "$push"
	// if fam.Persona.DatoBasico.Cedula != "" {
	// 	parametro = "$set"
	// }
	err = c.Update(bson.M{"id": f.DocumentoPadre}, bson.M{parametro: familiar})
	if err != nil {
		fmt.Println(" " + err.Error())
		return
	}
	return
}

//ContarFamiliar Contando Familiares
func (f *Familiar) ContarFamiliar() {

}

//AplicarReglasCarnetPadres Reglas
func (f *Familiar) AplicarReglasCarnetPadres() (TIM Carnet) {
	carnet := make(map[string]interface{})
	var mes, dia string
	var fechaVencimiento time.Time
	fechaActual := time.Now()
	AnnoA, MesA, DiaA := fechaActual.Date()
	layout := "2006-01-02"

	if f.Parentesco == "PD" {
		AnnoA += 10
		mes = strconv.Itoa(int(MesA))
		if int(MesA) < 10 {
			mes = "0" + strconv.Itoa(int(MesA))
		}
		dia = strconv.Itoa(DiaA)
		if DiaA < 10 {
			dia = "0" + strconv.Itoa(DiaA)
		}
		fvenc := strconv.Itoa(AnnoA) + "-" + mes + "-" + dia
		fechaVencimiento, _ = time.Parse(layout, fvenc)
	}

	TIM.Serial = TIM.GenerarSerial()
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimiento
	// TIM.CodigoComponente = m.Componente.Abreviatura
	// TIM.Grado.Abreviatura = m.Grado.Abreviatura
	TIM.Responsable = f.Persona.DatoBasico.Cedula
	TIM.Tipo = 1
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hijos")
	}
	return

}

//AplicarReglasCarnetPadres Reglas
func (f *Familiar) AplicarReglasCarnetEsposa() (TIM Carnet) {
	carnet := make(map[string]interface{})
	var mes, dia string
	var fechaVencimiento time.Time
	fechaActual := time.Now()
	AnnoA, MesA, DiaA := fechaActual.Date()
	layout := "2006-01-02"

	if f.Parentesco == "EA" {
		AnnoA += 3
		mes = strconv.Itoa(int(MesA))
		if int(MesA) < 10 {
			mes = "0" + strconv.Itoa(int(MesA))
		}
		dia = strconv.Itoa(DiaA)
		if DiaA < 10 {
			dia = "0" + strconv.Itoa(DiaA)
		}
		fvenc := strconv.Itoa(AnnoA) + "-" + mes + "-" + dia
		fechaVencimiento, _ = time.Parse(layout, fvenc)
	}

	TIM.Serial = TIM.GenerarSerial()
	TIM.Serial = TIM.Usuario + TIM.Serial
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimiento
	// TIM.CodigoComponente = m.Componente.Abreviatura
	// TIM.Grado.Abreviatura = m.Grado.Abreviatura
	TIM.Responsable = f.Persona.DatoBasico.Cedula
	TIM.Tipo = 1
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hijos")
	}
	return

}
