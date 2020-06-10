package credito

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
	"gopkg.in/mgo.v2/bson"
)

//DatoFinanciero Establecer un modulo de datos bancarios
type DatoFinanciero struct {
	Tipo        string `json:"tipo" bson:"tipo"`
	Institucion string `json:"institucion" bson:"institucion"`
	Cuenta      string `json:"cuenta" bson:"cuenta"`
	Prioridad   string `json:"prioridad" bson:"prioridad"`
	Autorizado  string `json:"autorizado" bson:"autorizado"`
	Titular     string `json:"titular" bson:"titular"`
}

//Credito formales
type Credito struct {
	Prestamo Prestamo `json:"Prestamo,omitempty" bson:"prestamo"`
}

//Solicitud Solicitar Prestamo o credito
type Solicitud struct {
	Oid              string         `json:"oid,omitempty" bson:"oid"`
	Cedula           string         `json:"cedula,omitempty" bson:"cedula"`                 //Monto total del credito solicitado
	Nombre           string         `json:"nombre,omitempty" bson:"nombre"`                 //Nombre total del credito solicitado
	Capital          float64        `json:"capital,omitempty" bson:"capital"`               //Monto total del credito solicitado
	MontoAprobado    float64        `json:"montoaprobado,omitempty" bson:"montoaprobado"`   //Monto Aprobado
	Cantidad         int            `json:"cantidad,omitempty" bson:"cantidad"`             //Cantidad por cuota
	PorcentajeTasa   float64        `json:"porcentajetasa,omitempty" bson:"porcentajetasa"` //Porcentaje de la Tasa
	Concepto         string         `json:"concepto,omitempty" bson:"concepto"`             //Detalle del prestamo
	Periodo          string         `json:"periodo,omitempty" bson:"periodo"`               //Aguinaldo, Vacaciones, Especial
	Estatus          int            `json:"total,omitempty" bson:"total"`
	Banco            DatoFinanciero `json:"Banco,omitempty" bson:"banco"`
	Cuota            float64        `json:"cuota,omitempty" bson:"cuota"`
	Cuotas           []Cuota        `json:"cuotas,omitempty" bson:"cuotas"`
	TotalInteres     float64        `json:"totalinteres,omitempty" bson:"totalinteres"`         //Monto Aprobado
	Intereses        float64        `json:"intereses,omitempty" bson:"intereses"`               //Monto Aprobado
	PorcentajeSeguro float64        `json:"porcentajeseguro,omitempty" bson:"porcentajeseguro"` //Monto Aprobado
	TotalDepositar   float64        `json:"totaldepositar,omitempty" bson:"totaldepositar"`     //Monto Aprobado
	FechaAprobado    time.Time      `json:"fechaaprobado,omitempty" bson:"fechaaprobado"`
	FechaCreacion    time.Time      `json:"fechacreacion,omitempty" bson:"fechacreacion"`
}

//Prestamo Prestamoes
type Prestamo struct {
	Vacacional         []Vacacional         `json:"Vacacional,omitempty" bson:"vacacional"`
	Educativo          []Educativo          `json:"Educativo,omitempty" bson:"educativo"`
	Hipotecario        []Hipotecario        `json:"Hipotecario,omitempty" bson:"hipotecario"`
	Parcelas           []Parcelas           `json:"Parcelas,omitempty" bson:"parcelas"`
	Personal           []Personal           `json:"Personal,omitempty" bson:"personal"`
	Articulos          []Articulos          `json:"Articulos,omitempty" bson:"articulos"`
	MiCasaBienEquipada []MiCasaBienEquipada `json:"MiCasaBienEquipada,omitempty" bson:"micasabienequipada"`
}

//Cuota Prestamos
type Cuota struct {
	ID      string  `json:"id,omitempty" bson:"id"`
	Balance float64 `json:"balance,omitempty" bson:"balance"`
	Cuota   float64 `json:"cuota,omitempty" bson:"cuota"`
	Interes float64 `json:"interes,omitempty" bson:"interes"`
	Capital float64 `json:"capital,omitempty" bson:"capital"`
	Saldo   float64 `json:"saldo,omitempty" bson:"saldo"`
	Fecha   string  `json:"fecha,omitempty" bson:"fecha"`
	Estatus int     `json:"estatus,omitempty" bson:"estatus"`
}

//Hipotecario viviendas
type Hipotecario struct {
	Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
}

//Vehiculo viviendas
type Vehiculo struct {
}

//NuevoPrestamo creacion de nuevo prestamo
func (PP *Solicitud) NuevoPrestamo(usuario string) string {
	var query string
	var coma string
	var oid string

	query = `INSERT INTO space.credito(
            cedula, nomb, capi, monta, cant, cuot, porc, conc, peri, esta, inst,
            tcue, ncue, fini, toti, inte, pors, totd, crea, usua)
    VALUES (
			'` + PP.Cedula + `',
			'` + PP.Nombre + `',
			` + strconv.FormatFloat(PP.Capital, 'f', 2, 64) + `,` + strconv.FormatFloat(PP.MontoAprobado, 'f', 2, 64) + `,
			` + strconv.Itoa(PP.Cantidad) + `,` + strconv.FormatFloat(PP.Cuota, 'f', 2, 64) + `,` + strconv.FormatFloat(PP.PorcentajeTasa, 'f', 2, 64) + `,
			'` + PP.Concepto + `','` + PP.Periodo + `', ` + strconv.Itoa(PP.Estatus) + `,
			'` + PP.Banco.Institucion + `','` + PP.Banco.Tipo + `',
			'` + strings.Replace(strings.Trim(PP.Banco.Cuenta, " "), "-", "", -1) + `',
			'` + PP.FechaAprobado.String()[0:10] + `',
			` + strconv.FormatFloat(PP.TotalInteres, 'f', 2, 64) + `,
			` + strconv.FormatFloat(PP.Intereses, 'f', 2, 64) + `,
			` + strconv.FormatFloat(PP.PorcentajeSeguro, 'f', 2, 64) + `,
			` + strconv.FormatFloat(PP.TotalDepositar, 'f', 2, 64) + `,
			 Now(), '` + usuario + `')  RETURNING oid;`
	//fmt.Println(query)
	sq, err := sys.PostgreSQLPENSION.Query(query)
	if err != nil {
		fmt.Println("Error en el query Credito: ", err.Error())
		return "0"
	}
	for sq.Next() {
		sq.Scan(&oid)
	}

	query = `INSERT INTO space.cuota ( creid, cedula, bala, cuot, inte, capi, sald, fech, esta) VALUES `
	i := 0
	for _, lst := range PP.Cuotas {

		if i > 0 {
			coma = `,`
		}
		query += coma + ` ( ` + oid + `,
			'` + PP.Cedula + `',
			` + strconv.FormatFloat(lst.Balance, 'f', 2, 64) + `,
			` + strconv.FormatFloat(lst.Cuota, 'f', 2, 64) + `,
			` + strconv.FormatFloat(lst.Interes, 'f', 2, 64) + `,
			` + strconv.FormatFloat(lst.Capital, 'f', 2, 64) + `,
			` + strconv.FormatFloat(lst.Saldo, 'f', 2, 64) + `,
			'` + lst.Fecha + `',` + strconv.Itoa(lst.Estatus) + `) `

		i++
	}

	//fmt.Println(query)
	_, err = sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query "+oid+" Credito-Cuota: ", err.Error())
		return "0"
	}

	PP.Oid = oid
	creditopersonal := make(map[string]interface{})
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	creditopersonal["credito.prestamo.personal"] = PP
	err = c.Update(bson.M{"id": PP.Cedula}, bson.M{"$push": creditopersonal})
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	return util.CompletarCeros(oid, 0, 8)
}

//Nuevo Credito Para vivienda
func (CH *Hipotecario) Nuevo() {

}

//wCredito Control
type WCredito struct {
	Cedula    string  `json:"cedula"`
	Nombre    string  `json:"nombre"`
	Concepto  string  `json:"concepto"`
	Instituto string  `json:"instituto"`
	Tipo      string  `json:"tipo"`
	Cuenta    string  `json:"cuenta"`
	Fecha     string  `json:"fecha"`
	Monto     float64 `json:"monto"`
}

//Listar consultando
func (CR *Credito) Listar(fecha string) (jSon []byte, err error) {
	var lst []WCredito
	s := `SELECT cedula, nomb, conc, inst, tcue, ncue, totd, fini FROM space.credito crd`
	sq, err := sys.PostgreSQLPENSION.Query(s)
	util.Error(err)

	for sq.Next() {
		var ced, nomb, conc, inst, tcue, ncue, fini sql.NullString
		var totd sql.NullFloat64
		var credito WCredito
		err = sq.Scan(&ced, &nomb, &conc, &inst, &tcue, &ncue, &totd, &fini)
		credito.Cedula = util.ValidarNullString(ced)
		credito.Nombre = util.ValidarNullString(nomb)
		credito.Concepto = util.ValidarNullString(conc)
		credito.Instituto = util.ValidarNullString(inst)
		credito.Tipo = util.ValidarNullString(tcue)
		credito.Cuenta = util.ValidarNullString(ncue)
		credito.Monto = util.ValidarNullFloat64(totd)
		credito.Fecha = util.ValidarNullString(fini)
		lst = append(lst, credito)
	}

	jSon, err = json.Marshal(lst)
	return
}
