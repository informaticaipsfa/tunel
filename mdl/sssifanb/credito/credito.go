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
	Cedula           string         `json:"cedula,omitempty" bson:"cedula"` //Monto total del credito solicitado
	Nombre           string         `json:"nombre,omitempty" bson:"nombre"` //Nombre total del credito solicitado
	Grado            string         `json:"grado,omitempty" bson:"grado"`
	Componente       string         `json:"componente,omitempty" bson:"componente"`
	Situacion        string         `json:"situacion,omitempty" bson:"situacion"`
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
	Tipo    int     `json:"tipo,omitempty" bson:"tipo"`
	Dias    int     `json:"dias,omitempty" bson:"dias"`
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
            tcue, ncue, fini, toti, inte, pors, totd, crea, usua, grad, comp, situa)
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
			 Now(), '` + usuario + `', '` + PP.Grado + `', '` + PP.Componente + `', '` + PP.Situacion + `')  RETURNING oid;`
	//fmt.Println(query)
	sq, err := sys.PostgreSQLPENSION.Query(query)
	if err != nil {
		fmt.Println("Error en el query Credito: ", err.Error())
		return "0"
	}
	for sq.Next() {
		sq.Scan(&oid)
	}

	query = `INSERT INTO space.cuota ( creid, cedula, bala, cuot, inte, capi, sald, fech, esta, tipo, dias) VALUES `
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
			'` + lst.Fecha + `',` + strconv.Itoa(lst.Estatus) + `,` + strconv.Itoa(lst.Tipo) + `,` + strconv.Itoa(lst.Dias) + `) `

		i++
	}

	//fmt.Println(query)
	_, err = sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query "+oid+" Credito-Cuota: ", err.Error())
		return "0"
	}

	PP.Oid = "CR" + util.CompletarCeros(oid, 0, 10)
	creditopersonal := make(map[string]interface{})
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	creditopersonal["credito.prestamo.personal"] = PP
	err = c.Update(bson.M{"id": PP.Cedula}, bson.M{"$push": creditopersonal})
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	return "CR" + util.CompletarCeros(oid, 0, 10)
}

//Nuevo Credito Para vivienda
func (CH *Hipotecario) Nuevo() {

}

//wCredito Control
type WCredito struct {
	OID        string  `json:"oid"`
	Cedula     string  `json:"cedula"`
	Componente string  `json:"componente"`
	Grado      string  `json:"grado"`
	Situacion  string  `json:"situacion"`
	Codigo     string  `json:"codigo"`
	Nombre     string  `json:"nombre"`
	Concepto   string  `json:"concepto"`
	Instituto  string  `json:"instituto"`
	Tipo       string  `json:"tipo"`
	Cuenta     string  `json:"cuenta"`
	Fecha      string  `json:"fecha"`
	Monto      float64 `json:"monto"`
}

//Listar consultando
func (CR *Credito) Listar(fecha string, desde string, hasta string, estatus int) (jSon []byte, err error) {
	var lst []WCredito
	s := `SELECT oid, cedula, nomb, conc, inst, tcue, ncue, totd, fini, comp, grad, situa
	FROM space.credito crd WHERE esta = ` + strconv.Itoa(estatus) + ` AND crea BETWEEN '` + desde + `' AND '` + hasta + `'`

	sq, err := sys.PostgreSQLPENSION.Query(s)

	util.Error(err)

	for sq.Next() {
		var oid, ced, nomb, conc, inst, tcue, ncue, fini, comp, grad, situa sql.NullString
		var totd sql.NullFloat64

		var credito WCredito
		err = sq.Scan(&oid, &ced, &nomb, &conc, &inst, &tcue, &ncue, &totd, &fini, &comp, &grad, &situa)
		credito.OID = util.ValidarNullString(oid)
		credito.Codigo = "CR" + util.CompletarCeros(util.ValidarNullString(oid), 0, 10)
		credito.Cedula = util.ValidarNullString(ced)
		credito.Componente = util.ValidarNullString(comp)
		credito.Grado = util.ValidarNullString(grad)
		credito.Situacion = util.ValidarNullString(situa)
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

//WCreditoActualizar Creditos
type WCreditoActualizar struct {
	Estatus     int      `json:"estatus"`
	Serie       []string `json:"serie"`
	Cantidad    int      `json:"cantidad"`
	Total       float64  `json:"total"`
	Llave       string   `json:"llave"`
	Observacion string   `json:"Observacion"`
}

//ActualizarLote credito lotes
func (CR *Credito) ActualizarLote(wca WCreditoActualizar, usuario string) (jSon []byte, err error) {

	query := `INSERT INTO space.credito_detalle(
            llav, obse, fech, esta, cant, totd, crea, usua)
    VALUES ('` + wca.Llave + `', 'CRED', Now(), 1,  '` + strconv.Itoa(wca.Cantidad) + `',
		'` + strconv.FormatFloat(wca.Total, 'f', 2, 64) + `', Now(), '` + usuario + `');`

	_, err = sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query crédito ", err.Error())

	}

	for i := 0; i < len(wca.Serie); i++ {
		fmt.Println("Actualizando credito: ", wca.Serie[i])
		s := `UPDATE space.credito SET esta=1, llav='` + wca.Llave + `' WHERE oid=` + wca.Serie[i]
		_, err = sys.PostgreSQLPENSION.Exec(s)
		if err != nil {
			fmt.Println("Error en el query crédito ", err.Error())
		}
	}

	jSon, err = json.Marshal(wca)
	return
}
