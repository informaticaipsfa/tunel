package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

type Banfan struct {
	Firma         string
	Cantidad      int
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
}

//CabeceraSQL Creando consulta para archivos
func (b *Banfan) CabeceraSQL(bancos string) string {
	return `
  SELECT
    pg.cedu, pg.nomb, pg.nume, pg.tipo, pg.banc, pg.neto
  FROM
    space.nomina nom
  JOIN space.pagos pg ON nom.oid=pg.nomi
  WHERE banc ` + bancos + ` AND llav='` + b.Firma + `' ORDER BY banc, pg.cedu;`
}

//Generar Archivo
func (b *Banfan) Generar(PostgreSQLPENSIONSIGESP *sql.DB) bool {

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

//Tercero Generando pago a terceros
func (b *Banfan) Tercero(PostgreSQLPENSIONSIGESP *sql.DB) bool {
	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL(" IN ('0134', '0175', '0105', '0108')"))
	util.Error(err)

	i := 0
	// var sumatotal float64
	// var sumaparcial float64
	// arch := 0
	// linea := ""
	for sq.Next() {
		i++
		var cedula, nombre, numero, tipo, banco sql.NullString

		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto)
		util.Error(e)

	}

	return true
}
