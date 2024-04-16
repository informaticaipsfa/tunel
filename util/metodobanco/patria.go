package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

type Patria struct {
	DesactivarArchivo bool
	Firma             string
	Cantidad          int
	CodigoEmpresa     string
	NumeroEmpresa     string
	Fecha             string
	Tabla             string
	Archivo           int
	TipoCuenta        string
	Directorio        string
	Registros         int
	Nombre            string
	Registro          int
	Contenido         string //representa la linea
	SumaParcial       float64
	Total             float64
}

// CabeceraSQL Creando consulta para archivos
func (P *Patria) CabeceraSQL() string {
	P.NumeroEmpresa = "G200036923"
	return SQL_QUERY_PATRIA(P.Firma)
}

// Generar Archivos de Patria
func (P *Patria) Generar(psqlPension *sql.DB) {

	//strConsulta := P.CabeceraSQL()
	//fmt.Println(strConsulta)
	sq, err := psqlPension.Query(P.CabeceraSQL())
	util.Error(err)

	//b.Directorio = crearDirectorio(b.Directorio, b.DesactivarArchivo, b.Firma, b.Tabla)

	//b.Cantidad = 500000
	for sq.Next() {

		var cedula, numero, nombre, tipo sql.NullString
		var neto sql.NullFloat64

		e := sq.Scan(&cedula, &numero, &neto, &nombre, &tipo)
		util.Error(e)

		monto, montos := generarMonto(neto, 0, 11)
		banc := generarCuentaBancaria(numero)
		nomb := util.CompletarEspacios(util.ValidarNullString(nombre), 1, 40)
		cedu := generarCedula(cedula, 0, 8)
		tipo_cedula := util.ValidarNullString(tipo)

		P.Total += monto
		P.SumaParcial += monto
		P.Contenido += tipo_cedula + cedu + banc + montos + nomb + "\r\n"
		P.Registro++
		//fmt.Println(cedu, banc, montos, nombre)
	}

	fecha := time.Now()
	fechas := util.EliminarGuionesFecha((fecha.String()[0:10]))
	P.generarArchivo(fechas)

}

// generarArchivo Permite generar archivo del sistema para los bancos
func (P *Patria) generarArchivo(fecha string) {

	P.Directorio = crearDirectorio(P.Directorio, P.DesactivarArchivo, P.Firma, P.Tabla)

	fpatria, e := os.Create(P.Directorio + "/patria.txt")
	util.Error(e)
	sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(P.SumaParcial, 'f', 2, 64))
	sumas = util.CompletarCeros(sumas, 0, 15)
	registros := util.CompletarCeros(strconv.Itoa(P.Registro), 0, 7)
	cabecera := "ONTNOM" + P.NumeroEmpresa + registros + sumas + "VES" + fecha + "\r\n"
	fmt.Fprintf(fpatria, cabecera+"")
	fmt.Fprintf(fpatria, P.Contenido+"")
	fpatria.Close()
	P.Contenido = ""
	P.Registro = 0
	P.SumaParcial = 0
}
