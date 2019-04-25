package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

//Banfan Control de Banco
type Banfanb struct {
	Firma         string
	Cantidad      int
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
}

//CabeceraSQL Creando consulta para archivos
func (b *Banfanb) CabeceraSQL(bancos string) string {
	return `
  SELECT
    pg.cedu, pg.nomb, pg.nume, pg.tipo, pg.banc, pg.neto
  FROM
    space.nomina nom
  JOIN space.pagos pg ON nom.oid=pg.nomi
  WHERE banc ` + bancos + ` AND llav='` + b.Firma + `' ORDER BY banc, pg.cedu;`
}

//Generar Archivo
func (b *Banfanb) Generar(PostgreSQLPENSIONSIGESP *sql.DB) bool {

	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL("='0177'"))
	util.Error(err)

	i := 0
	directorio := "./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma
	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)

	var sumatotal float64
	var sumaparcial float64
	arch := 0
	linea := ""
	for sq.Next() {
		i++
		var cedula, nombre, numero, tipo, banco sql.NullString

		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto)
		util.Error(e)

		monto := util.ValidarNullFloat64(neto)
		montos := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))

		montos = util.CompletarCeros(montos, 0, 12)
		bancos := util.CompletarCeros(util.ValidarNullString(numero), 0, 20)
		cedu := util.CompletarCeros(util.ValidarNullString(cedula), 0, 10)
		cerocinco := "00000"
		tipos := "0" // 0: ABONO 1: DEBITO
		filler := "00"
		sumatotal += monto
		sumaparcial += monto
		linea += b.CodigoEmpresa + montos + bancos + cedu + cerocinco + tipos + filler + "\n"
		if i == b.Cantidad {
			arch++
			banf, e := os.Create("./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma + "/banfan " + strconv.Itoa(arch) + ".txt")
			util.Error(e)
			sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
			sumas = util.CompletarCeros(sumas, 0, 17)
			fecha := time.Now()
			fechas := util.EliminarGuionesFecha((fecha.String()[0:10]))
			registros := util.CompletarCeros(strconv.Itoa(i), 0, 4)
			cabecera := b.NumeroEmpresa + fechas + sumas + registros + "\n"
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
		banf, e := os.Create("./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma + "/banfan " + strconv.Itoa(arch) + ".txt")
		util.Error(e)
		sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
		fecha := ""
		registros := util.CompletarCeros(strconv.Itoa(i), 0, 4)
		cabecera := b.NumeroEmpresa + b.NumeroEmpresa + fecha + sumas + registros + "\n"
		fmt.Fprintf(banf, cabecera)
		fmt.Fprintf(banf, linea)
		banf.Close()

	}

	return true
}

//Tercero Generando pago
func (b *Banfanb) Tercero(PostgreSQLPENSIONSIGESP *sql.DB) bool {
	fecha := time.Now()
	dd := fecha.String()[8:10]
	mm := fecha.String()[5:7]
	aa := fecha.String()[2:4]
	fechas := dd + mm + aa
	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL(" IN ('0134', '0175', '0105', '0108')"))
	util.Error(err)

	i := 0
	var sumatotal float64
	var sumaparcial float64
	arch := 0
	linea := ""
	for sq.Next() {
		i++
		var cedula, nombre, numero, tipo, banco sql.NullString
		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto)
		util.Error(e)
		ordenante := util.CompletarEspacios("02G200036923", 1, 16)
		nombreordenante := util.CompletarEspacios("IPSFA", 1, 30)
		monto := util.ValidarNullFloat64(neto)
		montos := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))

		montos = util.CompletarCeros(montos, 0, 12)
		bancos := util.CompletarCeros(util.ValidarNullString(numero), 0, 20)
		cedu := util.CompletarEspacios(util.ValidarNullString(cedula), 1, 10)
		nombrecompleto := util.CompletarEspacios(util.ValidarNullString(nombre), 1, 35)
		cuatro := "    "
		cinco := "     "
		sesentayocho := util.CompletarEspacios(" ", 1, 68)

		linea += ordenante + nombreordenante + cuatro + montos + b.NumeroEmpresa + "V" + cedu + cinco
		linea += nombrecompleto + bancos + "transferenci" + sesentayocho + "0212905452300000000000mediosdepago"
		linea += sesentayocho + "2\n"

		sumatotal += monto
		sumaparcial += monto
		if i == b.Cantidad {
			arch++
			banftercero, e := os.Create("./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma + "/banfan terceros " + strconv.Itoa(arch) + ".txt")
			util.Error(e)
			cabecera := "01FINANZAS CARACAS DE FECHA" + fechas + "\n"
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
		banftercero, e := os.Create("./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma + "/banfan terceros " + strconv.Itoa(arch) + ".txt")
		util.Error(e)
		cabecera := "01FINANZAS CARACAS DE FECHA" + fechas + "\n"
		fmt.Fprintf(banftercero, cabecera)
		fmt.Fprintf(banftercero, linea)
		banftercero.Close()
	}

	return true
}
