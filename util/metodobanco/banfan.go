package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

//Banfanb Control de Banco
type Banfanb struct {
	Firma         string
	Cantidad      int
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
	Tabla         string
}

//CabeceraSQL Creando consulta para archivos
func (b *Banfanb) CabeceraSQL(bancos string) string {
	return `
	  SELECT
			pg.cedu, regexp_replace(pg.nomb, '[^a-zA-Y0-9 ]', '', 'g') AS nomb, pg.nume, pg.tipo, pg.banc, pg.neto,
			pg.cfam, pg.caut, regexp_replace(pg.naut, '[^a-zA-Y0-9 ]', '', 'g') AS autor
	  FROM
	    space.nomina nom
	  JOIN space.` + b.Tabla + ` pg ON nom.oid=pg.nomi
  WHERE banc ` + bancos + ` AND llav='` + b.Firma + `'
	ORDER BY banc, pg.cedu ;`

}

//Generar Archivo
func (b *Banfanb) Generar(PostgreSQLPENSIONSIGESP *sql.DB) bool {

	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL("='0177'"))
	util.Error(err)

	i := 0
	valor := ""
	if b.Tabla == "rechazos" {
		valor = "-XR"
	}
	directorio := URLBanco + b.Firma + valor
	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)
	b.Cantidad = 30000
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

		monto := util.ValidarNullFloat64(neto)
		//montocondecimale := strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64)
		montos := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))

		montos = util.CompletarCeros(montos, 0, 12)
		numerocuenta := util.ValidarNullString(numero)

		bancos := util.CompletarCeros(util.EliminarUnderScore(numerocuenta), 0, 20)

		cedu := ""
		if util.ValidarNullString(ceddante) != "" && util.ValidarNullString(ndante) != "" {
			cedu = util.CompletarCeros(util.ValidarNullString(ceddante), 0, 10)[:10]
		} else {
			cedu = util.CompletarCeros(util.ValidarNullString(cedula), 0, 10)[:10]
			if util.ValidarNullString(familia) != "" {
				cedu = util.CompletarCeros(util.ValidarNullString(familia), 0, 10)[:10]
			}
		}

		cerocinco := "00000"
		tipos := "0" // 0: ABONO 1: DEBITO
		filler := "00"
		sumatotal += monto
		sumaparcial += monto
		linea += b.CodigoEmpresa + montos + bancos + cedu + cerocinco + tipos + filler + "\r\n"
		if i == b.Cantidad {
			arch++
			banf, e := os.Create(directorio + "/banfan " + strconv.Itoa(arch) + ".txt")
			util.Error(e)
			sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
			sumas = util.CompletarCeros(sumas, 0, 17)
			fecha := time.Now()
			fechas := util.EliminarGuionesFecha((fecha.String()[0:10]))
			registros := util.CompletarCeros(strconv.Itoa(i), 0, 4)
			cabecera := b.NumeroEmpresa + fechas + sumas + registros + "\r\n"
			fmt.Fprintf(banf, cabecera)
			fmt.Fprintf(banf, linea)
			banf.Close()
			sumaparcial = 0
			linea = ""
			i = 0
		}

	}
	if i > 0 {
		arch++
		banf, e := os.Create(directorio + "/banfan " + strconv.Itoa(arch) + ".txt")
		util.Error(e)
		sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
		sumas = util.CompletarCeros(sumas, 0, 17)
		fecha := time.Now()
		fechas := util.EliminarGuionesFecha((fecha.String()[0:10]))
		registros := util.CompletarCeros(strconv.Itoa(i), 0, 4)
		cabecera := b.NumeroEmpresa + fechas + sumas + registros + "\r\n"
		fmt.Fprintf(banf, cabecera)
		fmt.Fprintf(banf, linea)
		banf.Close()
	}

	return true
}

//Tercero Generando pago
func (b *Banfanb) Tercero(PostgreSQLPENSIONSIGESP *sql.DB, cuenta string) bool {
	fecha := time.Now()
	b.Cantidad = 40000
	dd := fecha.String()[8:10]
	mm := fecha.String()[5:7]
	aa := fecha.String()[2:4]
	fechas := dd + mm + aa
	valor := ""
	if b.Tabla == "rechazos" {
		valor = "-XR"
	}
	directorio := URLBanco + b.Firma + valor
	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)
	//fmt.Println(b.CabeceraSQL("='" + cuenta + "'"))
	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL("='" + cuenta + "'"))
	util.Error(err)

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
		ordenante := util.CompletarEspacios("02G200036923", 1, 18)
		nombreordenante := util.CompletarEspacios("IPSFA", 1, 30)[:30]
		monto := util.ValidarNullFloat64(neto)
		montos := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))

		montos = util.CompletarCeros(montos, 0, 15)
		numerocuenta := util.ValidarNullString(numero)

		bancos := util.CompletarCeros(util.EliminarUnderScore(numerocuenta), 0, 20)[:20]
		nombrecompleto := ""
		cfamilia := util.ValidarNullString(familia)
		cedu := ""
		if util.ValidarNullString(ceddante) != "" && util.ValidarNullString(ndante) != "" {
			cedu = util.CompletarEspacios(util.ValidarNullString(ceddante), 1, 10)[:10]
			nombrecompleto = util.CompletarEspacios(util.ValidarNullString(ndante), 1, 30)[:30]
		} else if cfamilia != "" {
			cedu = util.CompletarEspacios(cfamilia, 1, 10)[:10]
			nombrecompleto = util.CompletarEspacios(util.ValidarNullString(nombre), 1, 30)[:30]
		} else {
			cedu = util.CompletarEspacios(util.ValidarNullString(cedula), 1, 10)[:10]
			nombrecompleto = util.CompletarEspacios(util.ValidarNullString(nombre), 1, 30)[:30]
		}

		cuatro := "    "
		cinco := "     "
		sesentayocho := util.CompletarEspacios(" ", 1, 68)

		linea += ordenante + nombreordenante + cuatro + montos + b.NumeroEmpresa + "V" + cedu + cuatro
		linea += nombrecompleto + cinco + bancos + "transferenci" + sesentayocho + "0212905452300000000000mediosdepago"
		linea += sesentayocho + "2\r\n"

		sumatotal += monto
		sumaparcial += monto
		if i == b.Cantidad {
			arch++
			banftercero, e := os.Create(directorio + "/banfan terceros ( " + cuenta + " ) " + strconv.Itoa(arch) + ".txt")
			util.Error(e)
			cabecera := "01FINANZAS CARACAS DE FECHA" + fechas + "\r\n"
			fmt.Fprintf(banftercero, cabecera)
			fmt.Fprintf(banftercero, linea)
			banftercero.Close()
			sumaparcial = 0
			linea = ""
			i = 0
		}

	}
	if i > 0 {
		arch++
		banftercero, e := os.Create(directorio + "/banfan terceros ( " + cuenta + " ) " + strconv.Itoa(arch) + ".txt")
		util.Error(e)
		cabecera := "01FINANZAS CARACAS DE FECHA" + fechas + "\r\n"
		fmt.Fprintf(banftercero, cabecera)
		fmt.Fprintf(banftercero, linea)
		banftercero.Close()
	}

	return true
}
