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
func (b *Bicentenario) CabeceraSQL(bancos string) string {
	return `
  SELECT
		pg.cedu, regexp_replace(pg.nomb, '[^a-zA-Y0-9 ]', '', 'g') AS nomb, pg.nume, pg.tipo, pg.banc, pg.neto,
		pg.cfam, pg.caut, regexp_replace(pg.naut, '[^a-zA-Y0-9 ]', '', 'g') AS autor
  FROM
    space.nomina nom
  JOIN space.` + b.Tabla + ` pg ON nom.oid=pg.nomi
  WHERE banc ` + bancos + ` AND pg.tipo='CA' AND llav='` + b.Firma + `' ORDER BY banc, pg.cedu;`
}

//Generar Archivo
func (b *Bicentenario) Generar(psqlPension *sql.DB) (err error) {
	//fmt.Println(b.CabeceraSQL("='0175'"))
	sq, err := psqlPension.Query(b.CabeceraSQL("='0175'"))
	util.Error(err)

	b.Directorio = crearDirectorio(b.Directorio, b.DesactivarArchivo, b.Firma, b.Tabla)

	fecha := time.Now()
	fechas := util.EliminarGuionesFecha((fecha.String()[0:10]))
	for sq.Next() {
		b.Registro++
		b.Registros++

		var cedu, cedula, nombre, numero, tipo, banco, familia, ceddante, ndante sql.NullString

		var neto sql.NullFloat64
		e := sq.Scan(&cedula, &nombre, &numero, &tipo, &banco, &neto, &familia, &ceddante, &ndante)
		util.Error(e)

		monto, montos := generarMonto(neto, 0, 12)
		cuenta := generarCuentaBancaria(numero)

		cedulaAutorizado := util.ValidarNullString(ceddante)
		p := cedulaAutorizado != "" //cedula del autorizado
		q := cedulaAutorizado != "0"

		if p == q {
			cedu = ceddante //cedula del autorizado
		} else if util.ValidarNullString(familia) != "" {
			cedu = familia //cedula del familiar
		} else {
			cedu = cedula //cedula del titular
		}

		cerocinco := "00000"
		tipos := "0" // 0: ABONO 1: DEBITO
		filler := "00"
		b.Total += monto
		b.SumaParcial += monto
		b.Contenido += b.CodigoEmpresa + montos + cuenta +
			generarCedula(cedu, 0, 10) + cerocinco + tipos + filler + "\r\n"

		if b.Registro == b.Cantidad { //Pendiente si existen mas personas por escribir en el archivo
			b.generarArchivo(fechas)
		}
	}

	if b.Registro > 0 { //Pendiente si existen mas personas por escribir en el archivo
		b.generarArchivo(fechas)
	}

	return err
}

//generarArchivo Permite generar archivo del sistema para los bancos
func (b *Bicentenario) generarArchivo(fecha string) {
	if !b.DesactivarArchivo { //Desactiva la generacion de los archivos
		b.Archivo++

		banf, e := os.Create(b.Directorio + "/bicentenario " + strconv.Itoa(b.Archivo) + ".txt")
		util.Error(e)
		sumas := util.EliminarPuntoDecimal(strconv.FormatFloat(b.SumaParcial, 'f', 2, 64))
		sumas = util.CompletarCeros(sumas, 0, 17)
		registros := util.CompletarCeros(strconv.Itoa(b.Registro), 0, 4)
		cabecera := b.NumeroEmpresa + fecha + sumas + registros + "\r\n"
		fmt.Fprintf(banf, cabecera+"")
		fmt.Fprintf(banf, b.Contenido+"")
		banf.Close()
		b.Contenido = ""
		b.Registro = 0
		b.SumaParcial = 0
	}
}
