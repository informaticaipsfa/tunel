package sssifanb

//DatoBasico Datos importantes de la persona
type DatoBasico struct {
	Cedula          string `json:"cedula,omitempty" bson:"cedula"`
	NumeroPersona   int    `json:"nropersona,omitempty" bson:"nropersona"`
	Nacionalidad    string `json:"nacionalidad,omitempty" bson:"nacionalidad"`
	NombrePrimero   string `json:"nombreprimero,omitempty" bson:"nombreprimero"`
	NombreSegundo   string `json:"nombresegundo,omitempty" bson:"nombresegundo"`
	ApellidoPrimero string `json:"apellidoprimero,omitempty" bson:"apellidoprimero"`
	ApellidoSegundo string `json:"apellidosegundo,omitempty" bson:"apellidosegundo"`
	FechaNacimiento string `json:"fechanacimiento,omitempty" bson:"fechanacimiento"` //POR DEFINIR TIPO DE CAMPO
	Sexo            string `json:"sexo,omitempty" bson:"sexo"`
	EstadoCivil     string `json:"estadocivil,omitempty" bson:"estadocivil"`
}
