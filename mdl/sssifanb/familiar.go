package sssifanb

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
)

//Familiar Busquedas
type Familiar struct {
	ID              int       `json:"id" bson:"id"`
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

func (f *Familiar) AplicarReglasCarnetHijos() (fechaActual time.Time, fechaVencimientoCarnet time.Time) {
	fechaActual = time.Now()
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
	fechaVencimientoCarnet, _ = time.Parse(layOut, fecha)
	return
}

//IncluirFamiliar Agregar
func (f *Familiar) IncluirFamiliar(cedmilitar string) (err error) {
	c := sys.MGOSession.DB(BASEDEDATOS).C(COLECCION)
	err = c.Update(bson.M{"id": cedmilitar}, bson.M{"$push": f})

	if err != nil {
		fmt.Println(" " + err.Error())
		return
	}
	return
}

//ContarFamiliar Contando Familiares
func (f *Familiar) ContarFamiliar() {

}
