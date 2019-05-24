package util

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"
)

type Mensajes struct {
	Msj         string    `json:"msj,onmitempty"`
	Tipo        int       `json:"tipo,onmitempty"`
	Fecha       time.Time `json:"fecha,onmitempty"`
	Responsable int       `json:"responsable,onmitempty"`
}

//NullTime Tiempo nulo
type NullTime struct {
	Time  time.Time
	Valid bool
}

const layout string = "2006-01-02"

//ValidarNullFloat64 los campos nulos de la base de datos y retornar su valor original
func ValidarNullFloat64(b sql.NullFloat64) (f float64) {
	if b.Valid {
		f = b.Float64
	} else {
		f = 0
	}
	return
}

//ValidarNullTime los campos nulos de la base de datos y retorna fecha
func ValidarNullTime(b interface{}) (t time.Time) {
	t, e := b.(time.Time)
	if !e {
		return time.Now()
	}
	return
}

//ValidarNullString Validar los campos nulos de la base de datos y retornar su valor original
func ValidarNullString(b sql.NullString) (s string) {
	if b.Valid {
		s = b.String
	} else {
		s = "null"
	}
	return
}

func GetFechaConvert(f sql.NullString) (dateStamp time.Time) {
	fecha := ValidarNullString(f)
	if fecha != "null" {
		dateString := strings.Replace(fecha, "/", "-", -1)
		dateStamp, _ = time.Parse(layout, dateString)
	}
	return
}

//DiasDelMes los dias de un mes
func DiasDelMes(fecha time.Time) int {
	return 0
}

//CompletarCeros llenar con ceros antes y despues de una cadena
func CompletarCeros(cadena string, orientacion int, cantidad int) string {
	var result string
	cant := len(cadena)
	total := cantidad - cant
	for i := 0; i < total; i++ {
		result += "0"
	}
	if orientacion == 0 {
		result += cadena
	} else {
		result = cadena + result
	}
	return result
}

//CompletarEspacios llenar con ceros antes y despues de una cadena
func CompletarEspacios(cadena string, orientacion int, cantidad int) string {
	var result string
	cant := len(cadena)
	total := cantidad - cant
	for i := 0; i < total; i++ {
		result += " "
	}
	if orientacion == 0 {
		result += cadena
	} else {
		result = cadena + result
	}
	return result
}

//Fatal Error
func Fatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//CalcularTiempo Calculos
func CalcularTiempo(fecha time.Time) (Ano int, Mes time.Month, Dia int) {
	fechaActual := time.Now()

	AnnoA, MesA, DiaA := fechaActual.Date() //
	AnoN, MesN, DiaN := fecha.Date()        //Fecha de ingreso

	Ano = AnnoA - AnoN
	Mes = MesA - MesN
	Dia = DiaA - DiaN

	if Dia < 0 {
		Dia = 30
		Mes--
	}
	if Mes < 0 {
		Mes = 11
		Ano--
	}

	return
}

//CalcularTiempoServicio Calculos
func CalcularTiempoServicio(fechaActual time.Time, fecha time.Time) (Ano int, Mes time.Month, Dia int) {

	AnnoA, MesA, DiaA := fechaActual.Date() //
	AnoN, MesN, DiaN := fecha.Date()        //Fecha de ingreso

	Ano = AnnoA - AnoN
	Mes = MesA - MesN
	Dia = DiaA - DiaN

	if Dia < 0 {
		Dia = (30 + DiaA) - DiaN
		Mes--
	}
	if Mes < 0 {
		Mes = (12 + MesA) - MesN
		Ano--
	}
	return
}

//GenerarHash256 Generar Claves 256 para usuarios
func GenerarHash256(password []byte) (encry string) {
	h := sha256.New()
	h.Write(password)
	encry = hex.EncodeToString(h.Sum(nil))
	return

}

//EliminarPuntoDecimal Reemplazando coma por puntos
func EliminarPuntoDecimal(cadena string) string {
	return strings.Replace(strings.Trim(cadena, " "), ".", "", -1)
}

//EliminarEspacioBlanco Reemplazando coma por puntos
func EliminarEspacioBlanco(cadena string) string {
	return strings.Replace(strings.Trim(cadena, " "), " ", "", -1)
}

//EliminarUnderScore Reemplazando UnderScore por 0
func EliminarUnderScore(cadena string) string {
	return strings.Replace(strings.Trim(cadena, " "), "_", "0", -1)
}

//EliminarGuionesFecha Reemplazando coma por puntos
func EliminarGuionesFecha(cadena string) string {
	return strings.Replace(strings.Trim(cadena, " "), "-", "", -1)
}

//ReemplazarGuionesPorSlah Reemplazando coma por puntos
func ReemplazarGuionesPorSlah(cadena string) string {
	return strings.Replace(strings.Trim(cadena, " "), "-", "/", -1)
}

//Error Procesa errores del sistema
func Error(e error) {
	if e != nil {
		fmt.Println("\n Utilidad Error: ", e.Error())
	}
}
