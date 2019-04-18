package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/informaticaipsfa/tunel/util"
)

type Banfan struct {
	Firma         string
	Cantidad      string
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
}

//Generar Archivo
func (b *Banfan) Generar(PostgreSQLPENSIONSIGESP *sql.DB) bool {
	consulta := `
  SELECT
    pg.cedu, pg.nomb, pg.nume, pg.tipo, pg.banc, pg.neto
  FROM
    space.nomina nom
  JOIN space.pagos pg ON nom.oid=pg.nomi
  WHERE banc='0177' AND llav='` + b.Firma + `';`
	sq, err := PostgreSQLPENSIONSIGESP.Query(consulta)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	i := 0
	directorio := "./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma
	errr := os.Mkdir(directorio, 0777)
	if errr != nil {
		fmt.Println(errr.Error())
	}

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

		cedu := util.CompletarCeros(util.ValidarNullString(cedula), 0, 10)
		monto := util.ValidarNullFloat64(neto)
		montos := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))
		bancos := util.CompletarCeros(util.ValidarNullString(numero), 0, 20)
		cerocinco := "00000"
		tipos := "0"
		filler := "00"
		sumatotal += monto
		sumaparcial += monto
		linea += b.CodigoEmpresa + montos + bancos + cedu + cerocinco + tipos + filler + "\n"
		if i == 500 {
			arch++
			banf, e := os.Create("./public_web/SSSIFANB/afiliacion/temp/banco/" + b.Firma + "/banfan " + strconv.Itoa(arch) + ".txt")
			util.Error(e)
			sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
			fecha := ""
			registros := util.CompletarCeros(strconv.Itoa(i), 0, 4)
			cabecera := b.NumeroEmpresa + fecha + sumas + registros + "\n"
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

//Generar Archivo
func (b *Banfan) Tercero(PostgreSQLPENSIONSIGESP *sql.DB, oidnomina string) bool {
	return true
}
