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

//Validar los campos nulos de la base de datos y retornar su valor original
func ValidarNullString(b sql.NullString) (s string) {
	if b.Valid {
		s = b.String
	} else {
		s = "null"
	}
	return
}

//Convertir de (YYYY-MM-DD) a (DD/MM/YYYY) Humano
func ConvertirFechaSlash(fecha string) string {
	return "23/07/2016"
}

//Calcular los dias de un mes
func DiasDelMes(fecha time.Time) int {
	return 0
}

//Permite llenar con ceros antes y despues de una cadena
func CompletarCeros(cadena string, orientacion int, cantidad int) string {
	return "000"
}

func Fatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//CalcularEdad Calculos
func CalcularEdad(fechaNacimiento time.Time) (Edad int) {
	fechaActual := time.Now()
	AnnoA, MesA, DiaA := fechaActual.Date()

	AnoN, MesN, DiaM := fechaNacimiento.Date()

	Edad = AnnoA - AnoN
	Mes := MesA - MesN
	Dia := DiaA - DiaM
	if Dia < 0 {
		Mes++
	}
	if Mes < 0 {
		Edad++
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
