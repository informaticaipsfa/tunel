package util

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
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

//ConvertirFechaSlash de (YYYY-MM-DD) a (DD/MM/YYYY) Humano
func ConvertirFechaSlash(fecha string) string {
	return "23/07/2016"
}

//DiasDelMes los dias de un mes
func DiasDelMes(fecha time.Time) int {
	return 0
}

//CompletarCeros llenar con ceros antes y despues de una cadena
func CompletarCeros(cadena string, orientacion int, cantidad int) string {
	var result string
	for i := 0; i < cantidad; i++ {
		result += "0"
	}
	if orientacion == 0 {
		result += cadena
	} else {
		result = cadena + result
	}
	return result
}

func Fatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//CalcularTiempo Calculos
func CalcularTiempo(fechaNacimiento time.Time) (Ano int, Mes time.Month, Dia int) {
	fechaActual := time.Now()
	AnnoA, MesA, DiaA := fechaActual.Date()
	AnoN, MesN, DiaM := fechaNacimiento.Date()

	Ano = AnnoA - AnoN
	Mes = MesA - MesN
	Dia = DiaA - DiaM

	if Dia < 0 {
		Dia = 0
		Mes++
	}
	if Mes <= 0 {
		Ano--
	}

	return
}

func GenerarHash256(password []byte) (encry string) {
	h := sha256.New()
	h.Write(password)
	encry = hex.EncodeToString(h.Sum(nil))
	return

}

func Error(e error) {
	if e != nil {
		fmt.Println("\nError: ", e)
	}
}
