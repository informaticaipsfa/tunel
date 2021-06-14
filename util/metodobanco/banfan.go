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

//CabeceraSQL Creando consulta para archivos
func (b *Banfanb) CabeceraSQL(bancos string) string {
	return `
	  SELECT
			pg.cedu, regexp_replace(pg.nomb, '[^a-zA-Y0-9 ]', '', 'g') AS nomb, pg.nume, pg.tipo, pg.banc, pg.neto,
			pg.cfam, pg.caut, regexp_replace(pg.naut, '[^a-zA-Y0-9 ]', '', 'g') AS autor
	  FROM
	    space.nomina nom
	  JOIN space.` + b.Tabla + ` pg ON nom.oid=pg.nomi
  	WHERE banc ` + bancos + ` AND llav='` + b.Firma + `' ORDER BY banc, pg.cedu ;`

}

//Generar Archivo
func (b *Banfanb) Generar(psqlPension *sql.DB) (err error) {

	sq, err := psqlPension.Query(b.CabeceraSQL("='0177'"))
	util.Error(err)

	b.Directorio = crearDirectorio(b.Directorio, b.DesactivarArchivo, b.Firma, b.Tabla)

	//b.Cantidad = 500000
	for sq.Next() {
		b.Registro++
		b.Registros++

		var cedula, nombre, numero, tipo, banco, familia, ceddante, ndante sql.NullString
		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto, &familia, &ceddante, &ndante)
		util.Error(e)

		monto, montos := generarMonto(neto, 0, 12)
		bancos := generarCuentaBancaria(numero)
		cedulaAutorizado := util.ValidarNullString(ceddante) //Evaluamos cedula del titular de la cuenta

		p := cedulaAutorizado != ""
		q := cedulaAutorizado != "0"

		cedu := cedula //cedula del titular

		if p == q {
			cedu = ceddante
		} else if util.ValidarNullString(familia) != "" {
			cedu = familia
		}

		cerocinco := "00000"
		tipos := "0" // 0: ABONO 1: DEBITO
		filler := "00"
		b.Total += monto
		b.SumaParcial += monto
		b.Contenido += b.CodigoEmpresa + montos + bancos + generarCedula(cedu, 0, 10) + cerocinco + tipos + filler + "\r\n"
		if b.Registro == b.Cantidad { //Pendiente si existen mas personas por escribir en el archivo
			b.generarArchivo()
		}

	}

	if b.Registro > 0 { //Pendiente si existen mas personas por escribir en el archivo
		b.generarArchivo()
	}
	return
}

func (b *Banfanb) generarArchivo() {

	if !b.DesactivarArchivo { //Desactiva la generacion de los archivos
		b.Archivo++

		banf, e := os.Create(b.Directorio + "/banfan " + strconv.Itoa(b.Archivo) + ".txt")
		util.Error(e)
		sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(b.SumaParcial, 'f', 2, 64))
		sumas = util.CompletarCeros(sumas, 0, 17)
		fecha := time.Now()
		fechas := util.EliminarGuionesFecha((fecha.String()[0:10]))
		registros := util.CompletarCeros(strconv.Itoa(b.Registro), 0, 4)
		fmt.Fprintf(banf, b.NumeroEmpresa+fechas+sumas+registros+"\r\n")
		fmt.Fprintf(banf, b.Contenido+"")
		banf.Close()
		b.SumaParcial = 0
		b.Contenido = ""
		b.Registro = 0
	}
}

//Tercero Generando pago
func (b *Banfanb) Tercero(PostgreSQLPENSIONSIGESP *sql.DB, cuenta string) bool {
	fecha := time.Now()
	b.Cantidad = 50000
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
	// fmt.Println(b.CabeceraSQL("='" + cuenta + "'"))
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
			fmt.Fprintf(banftercero, "01FINANZAS CARACAS DE FECHA"+fechas+"\r\n")
			fmt.Fprintf(banftercero, ""+linea)
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
		fmt.Fprintf(banftercero, "01FINANZAS CARACAS DE FECHA"+fechas+"\r\n")
		fmt.Fprintf(banftercero, ""+linea)
		banftercero.Close()
	}

	return true
}
