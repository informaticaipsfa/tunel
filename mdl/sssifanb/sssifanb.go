package sssifanb

import (
	"strconv"
	"strings"
	"time"
)

func ActualizarPersona(persona Persona) string {
	fecha := time.Now()
	convertedDateString := fecha.Format("2006-01-02")
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := persona.DatoBasico.FechaNacimiento.Format("2006-01-02")
	fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)
	return `UPDATE personas SET
		tipnip = '` + persona.DatoBasico.Nacionalidad + `',
		codnip = '` + persona.DatoBasico.Cedula + `',
		nombreprimero = '` + strings.TrimSpace(persona.DatoBasico.NombrePrimero) + `',
		nombresegundo ='` + strings.TrimSpace(persona.DatoBasico.NombreSegundo) + `',
		apellidoprimero ='` + strings.TrimSpace(persona.DatoBasico.ApellidoPrimero) + `',
		apellidosegundo ='` + strings.TrimSpace(persona.DatoBasico.ApellidoSegundo) + `',
		nombrecompletoupp = '` + strings.TrimSpace(persona.DatoBasico.ConcatenarNombre()+persona.DatoBasico.ConcatenarApellido()) + `',
		nombrecompleto = '` + strings.TrimSpace(persona.DatoBasico.ConcatenarNombre()+persona.DatoBasico.ConcatenarApellido()) + `',
		nacionalidadcod = '` + persona.DatoBasico.ConvertirNacionalidad() + `',
		sexocod = '` + persona.DatoBasico.Sexo + `',
		edocivilcod = '` + persona.DatoBasico.EstadoCivil + `',
		fechanacimiento = '` + fechaSlashNacimiento + `',
		fechadefuncion = '` + persona.DatoBasico.FechaDefuncion + `',
		email1 = '` + persona.Correo.Principal + `',
		email2 = '` + persona.Correo.Alternativo + `',
		auditfechacambio = '` + fechaSlashActual + `',
		auditcodusuario = 'SSSIFANB'
		WHERE nropersona=` + strconv.Itoa(persona.DatoBasico.NroPersona)

}
