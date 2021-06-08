package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/informaticaipsfa/tunel/util"
)

type Venezuela struct {
	GenerarArchivo bool
	Firma          string
	Cantidad       int
	CodigoEmpresa  string
	NumeroEmpresa  string
	Fecha          string
	Tabla          string
	Archivo        int
	TipoCuenta     string
	Directorio     string
	Registros      int
	Nombre         string
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
	var sumatotal, sumaparcial float64
	linea := ""

	fechas := generarFecha()

	for sq.Next() {
		i++
		var cedu, cedula, nombre, numero, tipo, banco, familia, ceddante, ndante sql.NullString
		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto, &familia, &ceddante, &ndante)
		util.Error(e)
		codigocuenta, stipo := evaluarTipoCuenta(util.ValidarNullString(tipo)) //1 Codigo Empresa
		numerocuenta := generarCuentaBancaria(numero)                          //Cuenta Bancaria 20 Digitos
		monto := util.ValidarNullFloat64(neto)
		pagar := util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))
		montoapagar := util.CompletarCeros(pagar, 0, 11)

		cedulaAutorizado := util.ValidarNullString(ceddante)
		nombeval := util.ValidarNullString(nombre) //Evaluamos cedula del titular de la cuenta

		p := cedulaAutorizado != ""
		q := cedulaAutorizado != "0"

		nombrecomp := nombre //Nombre del titular
		cedu = cedula        //cedula del titular

		if p == q && nombeval != "" {
			nombrecomp = ndante
			cedu = ceddante //cedula del autorizado
		} else if util.ValidarNullString(familia) != "" {
			cedu = familia //cedula del familiar
		}

		sumatotal += monto
		sumaparcial += monto
		linea += stipo + numerocuenta + montoapagar + codigocuenta + generarNombre(nombrecomp, 1, 40) + generarCedula(cedu, 0, 10) + "003291  \r\n"
		if i == b.Cantidad {
			b.Archivo++
			b.generarArchivo(fechas, linea, sumaparcial)
			i = 0
			linea = ""
			sumaparcial = 0
		}
	}

	if i > 0 { //Pendiente si existen mas personas por escribir en el archivo
		b.generarArchivo(fechas, linea, sumaparcial)
	}

	return true
}

//evaluarTipoCuenta Parte del caso de que sea cuenta de Ahorro o Corriente
func evaluarTipoCuenta(sTipo string) (codigo string, tipo string) {
	codigo = "0770"
	tipo = "0"
	if sTipo == "CA" {
		tipo = "1"
		codigo = "1770"
	}
	return
}

//titularCuenta generar la persona responsable con nombre y cedula
func titularCuenta() string {
	return ""
}

//generarArchivo Permite generar archivo del sistema para los bancos
func (b *Venezuela) generarArchivo(fecha string, contenido string, sumaparcial float64) {
	venz, e := os.Create(b.Directorio + "/venezuela " + b.TipoCuenta + " " + strconv.Itoa(b.Archivo) + ".txt")
	util.Error(e)
	sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(sumaparcial, 'f', 2, 64))
	sumas = util.CompletarCeros(sumas, 0, 13)
	codigo := "03291" //codigo banco final de la linea
	cabecera := "HINSTITUTO DE PREVISION SOCIAL DE LA FUER" + b.NumeroEmpresa + "01" + fecha + sumas + codigo + "  \r\n"
	fmt.Fprintf(venz, cabecera)
	fmt.Fprintf(venz, contenido) //insertar contenido del archivo
	venz.Close()
}
