package sssifanb

import "time"

//DatoBasico Datos importantes de la persona
type DatoBasico struct {
	Cedula          string    `json:"cedula" bson:"cedula"`
	NumeroPersona   int       `json:"nropersona" bson:"nropersona"`
	Nacionalidad    string    `json:"nacionalidad" bson:"nacionalidad"`
	NombrePrimero   string    `json:"nombreprimero" bson:"nombreprimero"`
	NombreSegundo   string    `json:"nombresegundo" bson:"nombresegundo"`
	ApellidoPrimero string    `json:"apellidoprimero" bson:"apellidoprimero"`
	ApellidoSegundo string    `json:"apellidosegundo" bson:"apellidosegundo"`
	FechaNacimiento time.Time `json:"fechanacimiento" bson:"fechanacimiento"` //POR DEFINIR TIPO DE CAMPO
	Sexo            string    `json:"sexo" bson:"sexo"`
	EstadoCivil     string    `json:"estadocivil,omitempty" bson:"estadocivil"`
}

func (d *DatoBasico) AplicarReglas() {

}
