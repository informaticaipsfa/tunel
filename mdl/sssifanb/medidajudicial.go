package sssifanb

import (
	// "fmt"
	// "strconv"
	// "strings"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	//"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	// "github.com/informaticaipsfa/tunel/sys"
	// "gopkg.in/mgo.v2/bson"
)

//Militar militares
type MedidaJudicial struct {
	Numero                 string    `json:"numero,omitempty" bson:"numero"`
	Expediente             string    `json:"expediente,omitempty" bson:"expediente"`
	Fecha                  time.Time `json:"fecha,omitempty" bson:"fecha"`
	Tipo                   int       `json:"tipo,omitempty" bson:"tipo"`               // efectivo,asimilado,invalidez, reserva activa, tropa
	Observacion            string    `json:"observacion,omitempty" bson:"observacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Porcentaje             string    `json:"porcentaje,omitempty" bson:"porcentaje"`   //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	SueldoMinimo           float64   `json:"sueldominimo,omitempty" bson:"sueldominimo"`
	MontoFijo              float64   `json:"montofijo,omitempty" bson:"montofijo"`
	UnidadTributaria       float64   `json:"unidadtributaria,omitempty" bson:"unidadtributaria"`
	MontoTotal             float64   `json:"montototal,omitempty" bson:"montototal"`
	FormaPago              string    `json:"formapago,omitempty" bson:"formapago"`
	Institucion            string    `json:"institucion,omitempty" bson:"institucion"`
	TipoCuenta             string    `json:"tipocuenta,omitempty" bson:"tipocuenta"`
	NumeroCuenta           string    `json:"numerocuenta,omitempty" bson:"numerocuenta"`
	Autoridad              int       `json:"posicion,omitempty" bson:"posicion"`
	Estado                 string    `json:"estado,omitempty" bson:"estado"` //codigo
	Ciudad                 string    `json:"ciudad,omitempty" bson:"ciudad"`
	Municipio              string    `json:"municipio,omitempty" bson:"municipio"` //grado
	DescripcionInstitucion string    `json:"descripcion" bson:"descripcion"`
	CedulaBeneficiario     string    `json:"cedbeneficiario,omitempty" bson:"cedbeneficiario"`
	Beneficiario           string    `json:"beneficiario,omitempty" bson:"beneficiario"`
	Parentesco             string    `json:"parentesco,omitempty" bson:"parentesco"`
	CedulaAutorizado       string    `json:"cedautorizado,omitempty" bson:"cedautorizado"`
	Autorizado             string    `json:"autorizado,omitempty" bson:"autorizado"`
}

//Agregar Sistema
func (MJ *MedidaJudicial) Agregar() {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMEDIDAJUDICIAL)
	err := c.Insert(MJ)
	if err != nil {
		fmt.Println(err)
		return
	}
}


                  time.Time `json:"fecha,omitempty" bson:"fecha"`
                   int       `json:"tipo,omitempty" bson:"tipo"`               // efectivo,asimilado,invalidez, reserva activa, tropa
            string    `json:"observacion,omitempty" bson:"observacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
             float64    `json:"porcentaje,omitempty" bson:"porcentaje"`   //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
           float64   `json:"sueldominimo,omitempty" bson:"sueldominimo"`
              float64   `json:"montofijo,omitempty" bson:"montofijo"`
       float64   `json:"unidadtributaria,omitempty" bson:"unidadtributaria"`
             float64   `json:"montototal,omitempty" bson:"montototal"`
              string    `json:"formapago,omitempty" bson:"formapago"`
            string    `json:"institucion,omitempty" bson:"institucion"`
TipoCuenta             string    `json:"tipocuenta,omitempty" bson:"tipocuenta"`
NumeroCuenta           string    `json:"numerocuenta,omitempty" bson:"numerocuenta"`
              int       `json:"posicion,omitempty" bson:"posicion"`
Estado                 string    `json:"estado,omitempty" bson:"estado"` //codigo
Ciudad                 string    `json:"ciudad,omitempty" bson:"ciudad"`
              string    `json:"municipio,omitempty" bson:"municipio"` //grado





func InsertPension(MJ MedidaJudicial) {
	query := `
	INSERT INTO medida_judicial (
				f_documento,
		nro_oficio,
		nro_expediente,
		total_monto,
		porcentaje,
		desc_embargo,
		forma_pago_id,
		municipio_id,
		institucion,
		desc_institucion,
		ci_beneficiario,
		n_beneficiario,
		n_autorizado,
		status_id,
		parentesco_id,
		tipo_medida_id,
		cantidad_salario,
		unidad_tributaria,
		nombre_autoridad,
		cargo_autoridad,
		motivo_id,
		cedula,
		ci_autorizado,
		f_creacion,
		usr_creacion,
		f_ult_modificacion,
		usr_modificacion,
		observ_ult_modificacion,
		mensualidades,
	  numero_cuenta,
	  tipo_cuenta_id
	) VALUES (`

		query +=`
			'\'' . Fecha . '\',
			\'' . Numero . '\',
			\'' . Expediente . '\',
			\'' . MontoFijo . '\',
			\'' . Porcentaje . '\',
			\'' . Observacion . '\',
			\'' . FormaPago . '\',
			\'' . Municipio . '\',
			\'' . Institucion . '\',
			\'' . DescripcionInstitucion . '\',
			\'' . CedulaBeneficiario . '\',
			\'' . Beneficiario . '\',
			\'' . Autorizado . '\',
			\'' . Estatus . '\',
			\'' . Parentesco . '\',
			\'' . Tipo . '\',
			\'' . MontoTotal . '\',
			\'' . UnidadTributaria . '\',
			\'' . Autoridad . '\',
			\'' . $this->cargo . '\',
			\'' . $this->motivo . '\',
			\'' . $this->cedula . '\',
			\'' . CedulaAutorizado . '\',
			\'' . $this->fecha_creacion . '\',
			\'' . $this->usuario_creacion . '\',
			\'' . $this->fecha_modificacion . '\',
			\'' . $this->usuario_modificacion . '\',
			\'' . $this->ultima_observacion . '\',)
	`
	_, err = sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}
	msj.Mensaje = "Proceso exitoso"
	msj.Tipo = 1
	jSon, err = json.Marshal(msj)
}

//Actualizar Nomina
func (MJ *MedidaJudicial) Actualizar() {
	fmt.Println("09876")

}
