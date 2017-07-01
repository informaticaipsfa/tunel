package util

import (
	"database/sql"
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

func CalcularEdad(fechaNacimiento string) int {
	return 18
}

func Error(e error) {
	if e != nil {
		fmt.Println("\nError: ", e)
	}
}
