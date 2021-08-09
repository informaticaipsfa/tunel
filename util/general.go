package util

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os/exec"
	"strconv"
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
		str := strconv.FormatFloat(b.Float64, 'f', 2, 64)
		f, _ = strconv.ParseFloat(str, 64)
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
	cant := len(EliminarEspacioBlanco(cadena))
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
	cant := len(EliminarEspacioBlanco(cadena))
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

	fmt.Println("Imprimiendo la fecha de ingreso ", fecha)

	fmt.Println(fechaActual)

	AnnoA, MesA, DiaA := fechaActual.Date() //
	fmt.Println(AnnoA, "-", MesA, "-", DiaA)

	AnoN, MesN, DiaN := fecha.Date() //Fecha de ingreso

	Ano = AnnoA - AnoN
	MesX := int(MesA) - int(MesN)
	Dia = DiaA - DiaN
	//fmt.Println("Imprimiendo Mes de la operacion MESES: ", Ano, Dia, DiaA, DiaN)
	if Dia <= 0 {
		Dia = (30 + DiaA) - DiaN
		//MesX--
	}
	fmt.Println("Imprimiendo Mes de la operacion MESES: ", int(MesX), int(MesA), int(MesN))

	Mes = 1

	if int(MesX) < 0 {
		Mes = (12 + MesA) - MesN
		Ano--
	} else if int(MesX) == 0 {
		Mes = 0
	} else {
		Mes = MesA - MesN
	}
	//fmt.Println("Imprimiendo Mes de la operacion MESES: ", int(Mes), int(MesA), int(MesN))
	return
}

//GenerarHash256 Generar Claves 256 para usuarios
func GenerarHash256(password []byte) (encry string) {
	h := sha256.New()
	h.Write(password)
	encry = hex.EncodeToString(h.Sum(nil))
	return

}

//EliminarPuntoDecimal Reemplazando por nada
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

//ReemplazarPuntoPorComa Reemplazando coma por puntos
func ReemplazarPuntoPorComa(cadena string) string {
	return strings.Replace(strings.Trim(cadena, " "), ".", ",", -1)
}

//ReemplazarPuntoyComaPorComa Reemplazando
func ReemplazarPuntoyComaPorComa(cadena string) string {
	return strings.Replace(strings.Trim(cadena, " "), ";", ",", -1)
}

//Error Procesa errores del sistema
func Error(e error) {
	if e != nil {
		fmt.Println("\n Utilidad Error: ", e.Error())
	}
}

//EjecutarScript ejecucion de comandos
func EjecutarScript() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "/bin/sh", "update.sh").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
	}
	return
}

//GitAll Actualizando proyectos
func GitAll(paquete string, cmd string, origen string) (out []byte, err error) {
	if paquete == "bus" {
		origen = "."
	} else {
		origen = "public_web/SSSIFANB/" + paquete + "/"
	}
	argstr := []string{"gitall.sh", origen, cmd}
	out, err = exec.Command("/bin/sh", argstr...).Output()
	Error(err)
	return
}
