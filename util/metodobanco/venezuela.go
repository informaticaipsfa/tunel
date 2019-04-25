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
}

//CabeceraSQL Creando consulta para archivos
func (b *Venezuela) CabeceraSQL(bancos string) string {
	return `
  SELECT
    pg.cedu, pg.nomb, pg.nume, pg.tipo, pg.banc, pg.neto
  FROM
    space.nomina nom
  JOIN space.pagos pg ON nom.oid=pg.nomi
  WHERE banc ` + bancos + ` AND llav='` + b.Firma + `' ORDER BY banc, pg.cedu;`
}

//Generar Generando pago
func (b *Venezuela) Generar(PostgreSQLPENSIONSIGESP *sql.DB) bool {
	fecha := time.Now()
	dd := fecha.String()[8:10]
	mm := fecha.String()[5:7]
	aa := fecha.String()[2:4]
	fechas := dd + "/" + mm + "/" + aa
	sq, err := PostgreSQLPENSIONSIGESP.Query(b.CabeceraSQL("='0102'"))
	util.Error(err)

	directorio := "./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma
	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)
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
		//1
		stipo := "0"
		if util.ValidarNullString(tipo) == "CA" {
			stipo = "1"
		}
		//20
		numerocuenta := util.CompletarCeros(util.ValidarNullString(numero), 0, 20)
		monto := util.ValidarNullFloat64(neto)
		pagar := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))
		montoapagar := util.CompletarCeros(pagar, 0, 11)
		nombrecompleto := util.CompletarEspacios(util.ValidarNullString(nombre), 1, 40)
		cedu := util.CompletarCeros(util.ValidarNullString(cedula), 0, 11)
		sumatotal += monto
		sumaparcial += monto
		linea += stipo + numerocuenta + montoapagar + "0770" + nombrecompleto + cedu + "03291\n"
		if i == b.Cantidad {
			arch++
			venz, e := os.Create("./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma + "/venezuela " + strconv.Itoa(arch) + ".txt")
			util.Error(e)
			sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
			sumas = util.CompletarCeros(sumas, 0, 13)

			registros := "03291" //util.CompletarCeros(strconv.Itoa(i), 0, 4)
			cabecera := "HINSTITUTO DE PREVISION SOCIAL DE LA FUER" + b.NumeroEmpresa + "01" + fechas + sumas + registros + "\n"
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
		venz, e := os.Create("./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma + "/venezuela " + strconv.Itoa(arch) + ".txt")
		util.Error(e)
		sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
		sumas = util.CompletarCeros(sumas, 0, 13)
		fechas := dd + "/" + mm + "/" + aa
		registros := "03291" //util.CompletarCeros(strconv.Itoa(i), 0, 4)
		cabecera := "HINSTITUTO DE PREVISION SOCIAL DE LA FUER" + b.NumeroEmpresa + "01" + fechas + sumas + registros + " \n"
		fmt.Fprintf(venz, cabecera)
		fmt.Fprintf(venz, linea)
		venz.Close()
	}

	return true
}
