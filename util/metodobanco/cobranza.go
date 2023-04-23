package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

// Cobranza Control de Banco
type Cobranza struct {
	Firma         string
	Cantidad      int
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
	Componente    string
}

// CobranzaDetalle Detalles de contcuotrol
type CobranzaDetalle struct {
	Cantidad   int     `json:"cantidad"`
	Monto      float64 `json:"monto"`
	Componente string  `json:"componente"`
}

// CabeceraSQL Creando consulta para archivos
func (CB *Cobranza) CabeceraSQL(desde string, hasta string, componente string) string {
	CB.Componente = componente
	return `  SELECT crd.cedula, cot.cuot, nume, creid, cot.fech, cot.tipo, crd.esta, crd.cant
  FROM space.credito crd
  JOIN space.cuota cot on crd.oid=cot.creid
  WHERE
     -- crd.esta = 3 AND cot.esta = 0 AND tipo = 0
     -- comp = 'GN' AND cot.fech BETWEEN '2020-09-01' AND '2020-12-30'
	 crd.esta = 3 AND comp = '` + componente + `' AND cot.fech BETWEEN '` + desde + `' AND '` + hasta + `'`
}

// GenerarCobranza Creando consulta para archivos
func (CB *Cobranza) GenerarCobranza(PostgreSQLPENSION *sql.DB, desde string, hasta string, componente string) (wcob CobranzaDetalle) {
	CB.Componente = componente
	//fmt.Println(CB.CabeceraSQL(desde, hasta, componente))
	sq, err := PostgreSQLPENSION.Query(CB.CabeceraSQL(desde, hasta, componente))
	util.Error(err)
	i := 0
	suma := 0.0
	directorio := URLCobranza + "cobranza/"
	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)
	linea := ""
	for sq.Next() {
		i++
		var cedula, numero, credito, fecha, tipo, esta, plaz sql.NullString
		var couta sql.NullFloat64
		e := sq.Scan(&cedula, &couta, &numero, &credito, &fecha, &tipo, &esta, &plaz)
		util.Error(e)
		ced := util.CompletarCeros(util.ValidarNullString(cedula), 0, 8)[:8]
		cod := "00557"
		mon := util.ValidarNullFloat64(couta)
		montos := strconv.FormatFloat(mon, 'f', 2, 64)
		montos = util.CompletarCeros(montos, 0, 10)
		num := util.CompletarCeros(util.ValidarNullString(numero), 0, 3)
		plazo := util.ValidarNullString(plaz)
		oid := credito
		cred := "CR" + util.CompletarCeros(util.ValidarNullString(oid), 0, 10)
		condicion := "01"
		estatus := "00"
		linea += ced + "|" + cod + "|" + montos + "|" + num + "|" + plazo + "|" + cred + "|" + condicion + "|" + estatus + "\r\n"
		suma += mon
	}

	cobr, e := os.Create(directorio + "/" + CB.Componente + ".txt")
	util.Error(e)
	fmt.Fprintf(cobr, linea+"")

	cobr.Close()
	wcob.Cantidad = i
	wcob.Monto = suma
	wcob.Componente = CB.Componente

	return
}

// CabeceraSQL Creando consulta para archivos
func (CB *Cobranza) CabeceraSQLBanco(desde string, hasta string) string {
	return `  SELECT crd.cedula, cot.cuot, crd.ncue, creid, cot.fech, cot.tipo, crd.esta, crd.cant
	FROM space.credito crd
  JOIN space.cuota cot on crd.oid=cot.creid
  WHERE
		crd.fini > '2023-04-01' AND crd.inst='0177'  
	 	AND cot.fech BETWEEN '` + desde + `' AND '` + hasta + `'`
}

// GenerarCobranza Creando consulta para archivos
func (CB *Cobranza) GenerarCobranzaBanfanb(psql *sql.DB, desde string, hasta string) (wcob CobranzaDetalle) {
	//fmt.Println(CB.CabeceraSQL(desde, hasta, componente))
	sq, err := psql.Query(CB.CabeceraSQLBanco(desde, hasta))
	util.Error(err)
	i := 0
	suma := 0.0
	directorio := URLCobranza + "cobranza/"
	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)
	linea := ""
	for sq.Next() {
		i++
		var cedula, numero, credito, fecha, tipo, esta, plaz sql.NullString
		var couta sql.NullFloat64
		e := sq.Scan(&cedula, &couta, &numero, &credito, &fecha, &tipo, &esta, &plaz)
		util.Error(e)
		cod := "0131" //Codigo de la empresa
		mon := util.ValidarNullFloat64(couta)
		montos := util.EliminarPuntoDecimal(strconv.FormatFloat(mon, 'f', 2, 64))
		montos = util.CompletarCeros(montos, 0, 12)                            //Monto cuota
		num := util.ValidarNullString(numero)                                  //Numero de Cuenta
		ced := util.CompletarCeros(util.ValidarNullString(cedula), 0, 10)[:10] //cedula
		cinco := "00000"
		cred := "1" // debito
		estatus := "00"
		linea += cod + montos + num + ced + cinco + cred + estatus + "\r\n"
		suma += mon
	}
	cuenta := "01770001481100024031" //20digist
	fecha := time.Now()
	fechas := util.EliminarGuionesFecha((fecha.String()[0:10]))
	//fecha := "20230329"
	cantidad := util.EliminarPuntoDecimal(strconv.FormatFloat(suma, 'f', 2, 64)) //8digit
	monto := util.CompletarCeros(cantidad, 0, 17)                                //17digist
	registros := util.CompletarCeros(strconv.Itoa(i), 0, 4)                      //4digist
	cabecera := cuenta + fechas + monto + registros                              //' 000000000143236230792'
	cobr, e := os.Create(directorio + "/Banfanb.txt")
	util.Error(e)
	fmt.Fprintf(cobr, cabecera+"\n")
	fmt.Fprintf(cobr, linea+"")

	cobr.Close()
	wcob.Cantidad = i
	wcob.Monto = suma
	wcob.Componente = "Banfanb"

	return
}
