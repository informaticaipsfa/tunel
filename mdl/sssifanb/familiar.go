package sssifanb

import (
	"encoding/json"
	"fmt"
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

	//
	familiar["familiar.$.persona"] = f.Persona
	c := sys.MGOSession.DB(CBASE).C(CMILITAR)
	if f.ID != id {
		fmt.Println("Cambio de Cedula de un familiar: ", id)
		_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": f.ID}, bson.M{"$set": familiar})
	} else {
		_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": familiar})
	}
	if err != nil {
		fmt.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	parentesco := make(map[string]interface{})
	//
	parentesco["familiar.$.parentesco"] = f.Parentesco
	// c = sys.MGOSession.DB(CBASE).C(CMILITAR)
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": parentesco})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
		return
	}

	beneficio := make(map[string]interface{})
	beneficio["familiar.$.beneficio"] = f.Benficio
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": beneficio})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
		return
	}
	historia := make(map[string]interface{})
	historia["familiar.$.historiamedica"] = f.HistoriaMedica
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": historia})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
		return
	}
	donante := make(map[string]interface{})
	donante["familiar.$.donante"] = f.Donante
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": donante})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
		return
	}
	// grupo := make(map[string]interface{})
	// grupo["familiar.$.beneficio"] = f.GrupoSanguineo
	// err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": grupo})
	// if err != nil {
	// 	fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
	// 	return
	// }

	//fmt.Println(canal)
	return
}

//AplicarReglasCarnetHijos Reglas
func (f *Familiar) AplicarReglasCarnetHijos() (TIM Carnet) {
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
	AnioS := strconv.Itoa(Anio)
	MesS := strconv.Itoa(int(Mes))
	DiaS := strconv.Itoa(Dia)

	fecha := AnioS + "-" + MesS + "-" + DiaS
	fechaVencimientoCarnet, _ := time.Parse(layOut, fecha)
	TIM.Serial = TIM.GenerarSerial()
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimientoCarnet
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	// TIM.Componente.Abreviatura = m.Componente.Abreviatura
	// TIM.Componente.Descripcion = m.Componente.Descripcion
	// TIM.Grado.Abreviatura = m.Grado.Abreviatura
	// TIM.Grado.Descripcion = m.Grado.Descripcion
	// TIM.ID = f.Persona.DatoBasico.Cedula
	TIM.Tipo = 0
	TIM.Estatus = 0
	c := sys.MGOSession.DB(CBASE).C(CMILITAR)
	carnet["estatuscarnet"] = 1
	err := c.Update(bson.M{"id": f.Persona.DatoBasico.Cedula}, bson.M{"$set": carnet})
	if err != nil {
		return
	}
	return
}

//IncluirFamiliar Agregar
func (f *Familiar) IncluirFamiliar() (err error) {
	familiar := make(map[string]interface{})
	familiar["familiar"] = f
	c := sys.MGOSession.DB(CBASE).C(CMILITAR)
	err = c.Update(bson.M{"id": f.DocumentoPadre}, bson.M{"$push": familiar})

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
	return

}
