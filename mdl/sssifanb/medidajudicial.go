package sssifanb

/*
DROP TABLE space.medidajudicial;

CREATE TABLE space.medidajudicial
(
  oid serial NOT NULL,
  nume character varying(128) NOT NULL,
  expe character varying(256) NOT NULL,
  tipo integer,
  obse character varying(256) NOT NULL,
  tpag integer,
  fnxm character varying(512) NOT NULL,
  fpag character varying(32) NOT NULL,
  inst character varying(256) NOT NULL,
  tcue character varying(2) NOT NULL,
  ncue character varying(20) NOT NULL,
  autoridad character varying(256) NOT NULL,
  esta character varying(256) NOT NULL,
  ciud character varying(256) NOT NULL,
  muni character varying(256) NOT NULL,
  dins character varying(256) NOT NULL,
  cben character varying(256) NOT NULL,
  bene character varying(256) NOT NULL,
  pare character varying(256) NOT NULL,
  caut character varying(256) NOT NULL,
  auto character varying(128) NOT NULL,
  creado timestamp without time zone,
  ffin timestamp without time zone,
  usua character varying(128) NOT NULL,
  estatus integer,
  cedula character varying(32) NOT NULL,
  CONSTRAINT medidajudicial_pkey PRIMARY KEY (oid)
)
*/
import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

//MedidaJudicial militares
type MedidaJudicial struct {
	ID                     string    `json:"id,omitempty" bson:"id"`
	Numero                 string    `json:"numero,omitempty" bson:"numero"`
	Expediente             string    `json:"expediente,omitempty" bson:"expediente"`
	Tipo                   int       `json:"tipo,omitempty" bson:"tipo"`               // efectivo,asimilado,invalidez, reserva activa, tropa
	Observacion            string    `json:"observacion,omitempty" bson:"observacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	FormaPago              string    `json:"formapago,omitempty" bson:"formapago"`
	TipoPago               string    `json:"tipopago,omitempty" bson:"tipopago"`
	Formula                string    `json:"formula,omitempty" bson:"formula"`
	Institucion            string    `json:"institucion,omitempty" bson:"institucion"`
	TipoCuenta             string    `json:"tipocuenta,omitempty" bson:"tipocuenta"`
	NumeroCuenta           string    `json:"numerocuenta,omitempty" bson:"numerocuenta"`
	Autoridad              string    `json:"autoridad,omitempty" bson:"autoridad"`
	Cargo                  string    `json:"cargo,omitempty" bson:"cargo"`
	Estado                 string    `json:"estado,omitempty" bson:"estado"` //codigo
	Ciudad                 string    `json:"ciudad,omitempty" bson:"ciudad"`
	Municipio              string    `json:"municipio,omitempty" bson:"municipio"` //grado
	DescripcionInstitucion string    `json:"descripcion" bson:"descripcion"`
	CedulaBeneficiario     string    `json:"cedbeneficiario,omitempty" bson:"cedbeneficiario"`
	Beneficiario           string    `json:"beneficiario,omitempty" bson:"beneficiario"`
	Parentesco             string    `json:"parentesco,omitempty" bson:"parentesco"`
	CedulaAutorizado       string    `json:"cedautorizado,omitempty" bson:"cedautorizado"`
	Autorizado             string    `json:"autorizado,omitempty" bson:"autorizado"`
	Fecha                  time.Time `json:"fecha,omitempty" bson:"fecha"`
	FechaFin               time.Time `json:"fechafin,omitempty" bson:"fechafin"`
	Usuario                string    `json:"usuario,omitempty" bson:"usuario"`
}

//Agregar Sistema
func (MJ *MedidaJudicial) Agregar() (jSon []byte, err error) {
	var msj Mensaje

	medida := make(map[string]interface{})

	medida["pension.medidajudicial"] = MJ
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": MJ.ID}, bson.M{"$push": medida})
	msj.Tipo = 0
	if err != nil {
		fmt.Println("Fallo insertar Medida Judicial")
		msj.Tipo = 313
		jSon, err = json.Marshal(msj)
		return
	}
	InsertarPension(MJ)
	msj.Mensaje = "Proceso exitoso"
	msj.Tipo = 1
	jSon, err = json.Marshal(msj)

	return
}

//InsertarPension Cargar medidas
func InsertarPension(CMJ *MedidaJudicial) {
	query := `
	INSERT INTO space.medidajudicial (
		nume, expe, tipo, obse,
		tpag, fnxm,
		fpag, inst, tcue, ncue, autoridad, esta,
		ciud, muni, dins, cben,	bene, pare,
		caut, auto, creado, ffin, usua, estatus, cedula
	) VALUES `
	query += `('` + CMJ.Numero + `','` + CMJ.Expediente + `',` + strconv.Itoa(CMJ.Tipo) + `,'` + CMJ.Observacion + `',
						'` + CMJ.TipoPago + `','` + CMJ.Formula + `','` + CMJ.FormaPago + `','` + CMJ.Institucion + `',
						'` + CMJ.TipoCuenta + `','` + CMJ.NumeroCuenta + `',
						'` + CMJ.Autoridad + `','` + CMJ.Estado + `','` + CMJ.Ciudad + `','` + CMJ.Municipio + `',
						'` + CMJ.DescripcionInstitucion + `','` + CMJ.CedulaBeneficiario + `','` + CMJ.Beneficiario + `','` + CMJ.Parentesco + `',
						'` + CMJ.CedulaAutorizado + `','` + CMJ.Autorizado + `','` + CMJ.Fecha.String()[:10] + `','` + CMJ.FechaFin.String()[:10] + `','` + CMJ.Usuario + `',1,
						'` + CMJ.ID + `')`

	//fmt.Println(query)
	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}

	//jSon, err = json.Marshal(msj)
}

//Actualizar Nomina
func (MJ *MedidaJudicial) Actualizar() {
	fmt.Println("09876")

}
