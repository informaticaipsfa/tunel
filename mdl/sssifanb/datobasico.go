package sssifanb

import (
	"strings"
	"time"
)

//DatoBasico Datos importantes de la persona
type DatoBasico struct {
	Cedula          string    `json:"cedula" bson:"cedula"`
	NroPersona      int       `json:"nropersona" bson:"nropersona"`
	Nacionalidad    string    `json:"nacionalidad" bson:"nacionalidad"`
	NombrePrimero   string    `json:"nombreprimero" bson:"nombreprimero"`
	NombreSegundo   string    `json:"nombresegundo" bson:"nombresegundo"`
	ApellidoPrimero string    `json:"apellidoprimero" bson:"apellidoprimero"`
	ApellidoSegundo string    `json:"apellidosegundo" bson:"apellidosegundo"`
	FechaNacimiento time.Time `json:"fechanacimiento" bson:"fechanacimiento"` //POR DEFINIR TIPO DE CAMPO
	Sexo            string    `json:"sexo" bson:"sexo"`
	EstadoCivil     string    `json:"estadocivil,omitempty" bson:"estadocivil"`
	FechaDefuncion  time.Time `json:"fechadefuncion,omitempty" bson:"fechadefuncion"`
}

//AplicarReglas Politicas
func (d *DatoBasico) AplicarReglas() {

}

//ConcatenarNombre Unir nombres
func (d *DatoBasico) ConcatenarNombre() string {
	return strings.TrimSpace(d.NombrePrimero) + " " + strings.TrimSpace(d.NombreSegundo)
}

//ConcatenarApellido Nombre y Apellidos
func (d *DatoBasico) ConcatenarApellido() string {
	return strings.TrimSpace(d.ApellidoPrimero) + " " + strings.TrimSpace(d.ApellidoSegundo)
}

//ConcatenarNombreApellido Nombre y apellido
func (d *DatoBasico) ConcatenarNombreApellido() string {
	return d.ConcatenarNombre() + " " + d.ConcatenarApellido()
}

//ConcatenarApellidoNombre Apellido y Nombre
func (d *DatoBasico) ConcatenarApellidoNombre() string {
	return d.ConcatenarApellido() + " " + d.ConcatenarNombre()
}

//ConvertirNacionalidad Devolver nacionalidad
func (d *DatoBasico) ConvertirNacionalidad() string {
	nacionalidad := ""
	switch d.Nacionalidad {
	case "V":
		nacionalidad = `VEN`
	case "M":
		nacionalidad = `MEN`
	case "E":
		nacionalidad = `EXT`

	}
	return nacionalidad

}
