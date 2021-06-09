//package metodometodo  Generador de codigo para archivos del banco venezuela partiiendo del
//origen de los codigos bancarios Venezuela codigo 0102
//estos archivos generan codigo para realizar test unit mediante go test
package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/informaticaipsfa/tunel/util"
)

type Venezuela struct {
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
func (b *Venezuela) Generar(psqlPension *sql.DB, tipocuenta string) error {

	sq, err := psqlPension.Query(b.CabeceraSQL("='0102'", tipocuenta))
	util.Error(err)

	b.crearDirectorio()

	b.Registro = 0
	b.Contenido = ""

	fechas := generarFecha()

	for sq.Next() {
		b.Registro++
		var cedu, cedula, nombre, numero, tipo, banco, familia, ceddante, ndante sql.NullString
		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto, &familia, &ceddante, &ndante)
		util.Error(e)
		codigocuenta, stipo := evaluarTipoCuenta(util.ValidarNullString(tipo)) //1 Codigo Empresa
		numerocuenta := generarCuentaBancaria(numero)                          //Cuenta Bancaria 20 Digitos
		monto, montoapagar := generarMonto(neto, 0, 11)

		cedulaAutorizado := util.ValidarNullString(ceddante) //Evaluamos cedula del titular de la cuenta
		nombeval := util.ValidarNullString(nombre)

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

		b.Total += monto
		b.SumaParcial += monto
		b.Contenido += stipo + numerocuenta + montoapagar + codigocuenta +
			generarNombre(nombrecomp, 1, 40) + generarCedula(cedu, 0, 10) + "003291  \r\n"

		if b.Registro == b.Cantidad { //Pendiente si existen mas personas por escribir en el archivo
			b.generarArchivo(fechas)
		}
	}

	if b.Registro > 0 { //Pendiente si existen mas personas por escribir en el archivo
		b.generarArchivo(fechas)
	}

	return err
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
//func titularCuenta() string {
//	return ""
//}

//generarArchivo Permite generar archivo del sistema para los bancos
func (b *Venezuela) generarArchivo(fecha string) {
	if !b.DesactivarArchivo { //Desactiva la generacion de los archivos
		b.Archivo++
		venz, e := os.Create(b.Directorio + "/venezuela " + b.TipoCuenta + " " + strconv.Itoa(b.Archivo) + ".txt")
		util.Error(e)
		sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(b.SumaParcial, 'f', 2, 64))
		sumas = util.CompletarCeros(sumas, 0, 13)
		codigo := "03291" //codigo banco final de la linea
		cabecera := "HINSTITUTO DE PREVISION SOCIAL DE LA FUER" + b.NumeroEmpresa + "01" + fecha + sumas + codigo + "  \r\n"
		fmt.Fprintf(venz, cabecera)
		fmt.Fprintf(venz, b.Contenido) //insertar contenido del archivo
		venz.Close()
		b.Contenido = ""
		b.Registro = 0
		b.SumaParcial = 0
	}
}

//crearDirectorio permite iniciar la carpeta donde se crearan los documentos
func (b *Venezuela) crearDirectorio() {
	if !b.DesactivarArchivo {
		if b.Directorio == "" {
			b.Directorio = URLBanco + b.Firma + definirArchivo(b.Tabla)
		}

		err := os.Mkdir(b.Directorio, 0777)
		util.Error(err)
	}
}

//definirArchivo para su asignacion y creacion en los documentos
func definirArchivo(tabla string) (valor string) {

	valor = ""
	if tabla == "rechazos" {
		valor = "-XR"
	}
	return
}
