package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

type Venezuela struct {
	Firma         string
	Cantidad      int
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
	Tabla         string
}

//CabeceraSQL Creando consulta para archivos
func (b *Venezuela) CabeceraSQL(bancos string, tipocuenta string) string {
	return `
  SELECT
		pg.cedu, regexp_replace(pg.nomb, '[^a-zA-Y0-9 ]', '', 'g') AS nomb, pg.nume, pg.tipo, pg.banc, pg.neto,
		pg.cfam, pg.caut, regexp_replace(pg.naut, '[^a-zA-Y0-9 ]', '', 'g') AS autor
  FROM
    space.nomina nom
  JOIN space.` + b.Tabla + ` pg ON nom.oid=pg.nomi
  WHERE banc ` + bancos + ` AND llav='` + b.Firma + `' AND pg.tipo='` + tipocuenta + `' ORDER BY banc, pg.cedu;`
}

//Generar Generando pago
func (b *Venezuela) Generar(PostgreSQLPENSIONSIGESP *sql.DB, tipocuenta string) bool {
	fecha := time.Now()
	dd := fecha.String()[8:10]
	mm := fecha.String()[5:7]
	aa := fecha.String()[2:4]
	fechas := dd + "/" + mm + "/" + aa
	//fmt.Println(b.CabeceraSQL("='0102'", tipocuenta))
	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL("='0102'", tipocuenta))
	util.Error(err)

	valor := ""
	if b.Tabla == "rechazos" {
		valor = "-XR"
	}
	directorio := URLBanco + b.Firma + valor

	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)
	i := 0
	var sumatotal float64
	var sumaparcial float64
	arch := 0
	linea := ""
	for sq.Next() {
		i++
		var cedula, nombre, numero, tipo, banco, familia, ceddante, ndante sql.NullString

		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto, &familia, &ceddante, &ndante)
		util.Error(e)
		//1
		stipo := "0"
		codigocuenta := "0770"
		if util.ValidarNullString(tipo) == "CA" {
			stipo = "1"
			codigocuenta = "1770"
		}
		//20
		numerocuenta := util.CompletarCeros(util.ValidarNullString(numero), 0, 20)
		monto := util.ValidarNullFloat64(neto)
		pagar := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))
		montoapagar := util.CompletarCeros(pagar, 0, 11)
		nombrecompleto := ""
		cedu := ""
		evaluar := util.ValidarNullString(ceddante)
		nombeval := util.ValidarNullString(nombre)
		if evaluar != "" && evaluar != "0" && nombeval != "" {
			nombrecompleto = util.CompletarEspacios(util.ValidarNullString(ndante), 1, 40)[:40]
			cedu = util.CompletarCeros(util.ValidarNullString(ceddante), 0, 10)[:10]
		} else {
			nombrecompleto = util.CompletarEspacios(util.ValidarNullString(nombre), 1, 40)[:40]
			cedu = util.CompletarCeros(util.ValidarNullString(cedula), 0, 10)[:10]
			if util.ValidarNullString(familia) != "" {
				cedu = util.CompletarCeros(util.ValidarNullString(familia), 0, 10)[:10]
			}

		}

		sumatotal += monto
		sumaparcial += monto
		linea += stipo + numerocuenta + montoapagar + codigocuenta + nombrecompleto + cedu + "003291  \r\n"
		if i == b.Cantidad {
			arch++
			venz, e := os.Create(directorio + "/venezuela " + tipocuenta + " " + strconv.Itoa(arch) + ".txt")
			util.Error(e)
			sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
			sumas = util.CompletarCeros(sumas, 0, 13)

			registros := "03291" //util.CompletarCeros(strconv.Itoa(i), 0, 4)
			cabecera := "HINSTITUTO DE PREVISION SOCIAL DE LA FUER" + b.NumeroEmpresa + "01" + fechas + sumas + registros + "  \r\n"
			fmt.Fprintf(venz, cabecera)
			fmt.Fprintf(venz, linea)
			venz.Close()
			sumaparcial = 0
			linea = ""
			i = 0
		}

	}

	if i > 0 {
		arch++
		venz, e := os.Create(directorio + "/venezuela " + tipocuenta + " " + strconv.Itoa(arch) + ".txt")
		util.Error(e)
		sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
		sumas = util.CompletarCeros(sumas, 0, 13)
		fechas := dd + "/" + mm + "/" + aa
		registros := "03291"
		cabecera := "HINSTITUTO DE PREVISION SOCIAL DE LA FUER" + b.NumeroEmpresa + "01" + fechas + sumas + registros + "  \r\n"
		fmt.Fprintf(venz, cabecera)
		fmt.Fprintf(venz, linea)
		venz.Close()
	}

	return true
}
