package sssifanb

import "github.com/gesaodin/bdse/util"

type Familiar struct {
	ID             int
	Persona        Persona
	Parentesco     string //0:Mama, 1:papa, 2: Esposa  3: hijo
	EsMilitar      bool
	Condicion      int //Sano o Condicion especial
	Estudia        int
	Benficio       bool //
	Documento      int
	DocumentoPadre string
}

//AplicarReglas OJO SEGUROS HORIZONTES
func (f *Familiar) AplicarReglasBeneficio() {
	f.Benficio = false
	if f.Condicion == 1 {
		f.Benficio = true
	} else {
		if util.CalcularEdad("2001-01-01") < 18 {
			f.Benficio = true
		} else if f.Estudia == 1 && util.CalcularEdad("2001-01-01") < 27 {
			f.Benficio = true
		}
	}

}

func (f *Familiar) AplicarReglasParentesco() {

}
