package sssifanb

import "time"

//Reduccion de datos de los familiares
type Reduccion struct {
	Cedula          string    `json:"cedula",bson:"cedula"`
	Nombre          string    `json:"nombre",bson:"nombre"`
	Sexo            string    `json:"sexo",bson:"sexo"`
	Tipo            string    `json:"tipo",bson:"tipo"` //T Titular Militar | F Familiar
	EsMilitar       bool      `json:"esmilitar",bson:"esmilitar"`
	FechaNacimiento time.Time `json:"fecha",bson:"fecha"`
	Parentesco      string    `json:"parentesco",bson:"parentesco"`
	Situacion       string    `json:"situacion",bson:"situacion"`
}
