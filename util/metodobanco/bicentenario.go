package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

type Bicentenario struct {
	Firma         string
	Cantidad      int
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
}

//CabeceraSQL Creando consulta para archivos
func (b *Bicentenario) CabeceraSQL(bancos string) string {
	return `
  SELECT
		pg.cedu, regexp_replace(pg.nomb, '[^a-zA-Y0-9 ]', '', 'g') AS nomb, pg.nume, pg.tipo, pg.banc, pg.neto,
		pg.cfam, pg.caut, regexp_replace(pg.naut, '[^a-zA-Y0-9 ]', '', 'g') AS autor
  FROM
    space.nomina nom
  JOIN space.pagos pg ON nom.oid=pg.nomi
  WHERE banc ` + bancos + ` AND llav='` + b.Firma + `' ORDER BY banc, pg.cedu;`

	//WHERE banc ` + bancos + ` AND llav='` + b.Firma + `' ORDER BY banc, pg.cedu;`
	//AND cfam != '' AND cedu!=cfam AND naut =''
}

//Generar Archivo
func (b *Bicentenario) Generar(PostgreSQLPENSIONSIGESP *sql.DB) bool {
	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL("='0175'"))
	util.Error(err)

	i := 0
	directorio := URLBanco + b.Firma
	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)

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
		montos := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))
		montos = util.CompletarCeros(montos, 0, 12)
		bancos := util.CompletarCeros(util.ValidarNullString(numero), 0, 20)[:20]

		cedu := ""
		if util.ValidarNullString(ceddante) != "" && util.ValidarNullString(ndante) != "" {
			cedu = util.CompletarCeros(util.ValidarNullString(ceddante), 0, 10)
		} else {
			cedu = util.CompletarCeros(util.ValidarNullString(cedula), 0, 10)
			if util.ValidarNullString(familia) != "" {
				cedu = util.CompletarCeros(util.ValidarNullString(familia), 0, 10)
			}
		}

		cerocinco := "00000"
		tipos := "0" // 0: ABONO 1: DEBITO
		filler := "00"
		sumatotal += monto
		sumaparcial += monto
		linea += b.CodigoEmpresa + montos + bancos + cedu + cerocinco + tipos + filler + "\r\n"

	}
	// if i > 0 {
	arch++
	banf, e := os.Create(URLBanco + b.Firma + "/bicentenario " + strconv.Itoa(arch) + ".txt")
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

	// }

	return true
}
