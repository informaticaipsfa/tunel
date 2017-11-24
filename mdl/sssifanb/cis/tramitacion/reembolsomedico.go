package tramitacion

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

const (
	SACTIVO        int = 0
	SPENDIENTE     int = 1
	SRECOMENDACION int = 2
	RRECHAZADO     int = -1
	RINICIAL       int = 0
	RPENDIENTE     int = 1
	RJEFATURA      int = 2
	RGERENCIA      int = 3
	RPRESIDENCIA   int = 4
	RAPROBADO      int = 5
)

type Reembolso struct {
	Numero          string         `json:"numero" bson:"numero"`
	Estatus         int            `json:"estatus" bson:"estatus"`
	FechaCreacion   time.Time      `json:"fechacreacion" bson:"fechacreacion"`
	MontoSolicitado float64        `json:"montosolicitado" bson:"montosolicitado"`
	CuentaBancaria  DatoFinanciero `json:"CuentaBancaria" bson:"CuentaBancaria"`
	Responsable     string         `json:"responsable" bson:"responsable"`
	Concepto        []Concepto     `json:"Concepto" bson:"concepto"`
	FechaAprobado   time.Time      `json:"fechaaprobado" bson:"fechaaprobado"`
	MontoAprobado   float64        `json:"montoaprobado" bson:"montoaprobado"`
	Requisitos      []int          `json:"requisitos" bson:"requisitos"`
	Componente      string         `json:"componente" bson:"componente"`
	Grado           string         `json:"grado" bson:"grado"`
	Clase           string         `json:"clase" bson:"clase"`
	Situacion       string         `json:"situacion" bson:"situacion"`
	Direccion       Direccion      `json:"Direccion" bson:"direccion"`
	Telefono        Telefono       `json:"Telefono" bson:"telefono"`
	Correo          Correo         `json:"Correo" bson:"correo"`
	Seguimiento     Seguimiento    `json:"Seguimiento" bson:"seguimiento"`
	Usuario         string         `json:"usuario" bson:"usuario"`
}

type DatoFinanciero struct {
	Titular     string `json:"titular" bson:"titular"`
	Cedula      string `json:"cedula" bson:"cedula"`
	Tipo        string `json:"tipo" bson:"tipo"`
	Institucion string `json:"institucion" bson:"institucion"`
	Cuenta      string `json:"cuenta" bson:"cuenta"`
	Prioridad   string `json:"prioridad" bson:"prioridad"`
}

type Concepto struct {
	Descripcion string  `json:"descripcion" bson:"descripcion"`
	DatoFactura Factura `json:"DatoFactura" bson:"datofactura"`
	Afiliado    string  `json:"afiliado" bson:"afiliado"` //Cedula, Guión (-), Nombre
	Requisitos  []int   `json:"requisitos" bson:"requisitos"`
}

type Factura struct {
	Numero        string    `json:"numero" bson:"numero"`
	Control       string    `json:"control" bson:"control"`
	Fecha         time.Time `json:"fecha" bson:"fecha"`
	Monto         float64   `json:"monto" bson:"monto"`
	Porcentaje    float64   `json:"porcentaje" bson:"porcentaje"`
	MontoAprobado float64   `json:"montoaprobado" bson:"montoaprobado"`
	Beneficiario  Proveedor `json:"Beneficiario" bson:"beneficiario"`
}

type Proveedor struct {
	Rif         string         `json:"rif" bson:"rif"`
	RazonSocial string         `json:"razonsocial" bson:"razonsocial"`
	Tipo        string         `json:"tipo" bson:"tipo"`
	Direccion   string         `json:"direccion" bson:"direccion"`
	Banco       DatoFinanciero `json:"Banco" bson:"banco"`
	Descripcion string         `json:"descripcion" bson:"descripcion"`
}

//Direccion ruta y secciones
type Direccion struct {
	Tipo         int    `json:"tipo,omitempty" bson:"tipo"` //domiciliaria, trabajo, emergencia
	Ciudad       string `json:"ciudad,omitempty" bson:"ciudad"`
	Estado       string `json:"estado,omitempty" bson:"estado"`
	Municipio    string `json:"municipio,omitempty" bson:"municipio"`
	Parroquia    string `json:"parroquia,omitempty" bson:"parroquia"`
	CalleAvenida string `json:"calleavenida" bson:"calleavenida"`
	Casa         string `json:"casa" bson:"casa"`
	Apartamento  string `json:"apartamento" bson:"apartamento"`
	Numero       int    `json:"numero,omitempty" bson:"numero"`
}

type Telefono struct {
	Movil        string `json:"movil,omitempty" bson:"movil"`
	Domiciliario string `json:"domiciliario,omitempty" bson:"domiciliario"`
	Emergencia   string `json:"emergencia,omitempty" bson:"emergencia"`
}

//Correo Direcciones electronicas
type Correo struct {
	Principal     string `json:"principal,omitempty" bson:"principal"`
	Alternativo   string `json:"alternativo,omitempty" bson:"alternativo"`
	Institucional string `json:"institucional,omitempty" bson:"institucional"`
}

type Seguimiento struct {
	Estatus       int           `json:"estatus,omitempty" bson:"estatus"`
	Observaciones []Observacion `json:"Observaciones,omitempty" bson:"observaciones"`
}

type Observacion struct {
	FechaCreacion time.Time `json:"fechacreacion,omitempty" bson:"fechacreacion"`
	Usuario       string    `json:"usuario,omitempty" bson:"usuario"`
	Contenido     string    `json:"contenido,omitempty" bson:"contenido"`
}
type ColeccionReembolso struct {
	ID                 string    `json:"id" bson:"id"`
	Nombre             string    `json:"nombre" bson:"nombre"`
	Numero             string    `json:"numero" bson:"numero"`
	FechaCreacion      time.Time `json:"fechacreacion" bson:"fechacreacion"`
	MontoSolicitado    float64   `json:"montosolicitado" bson:"montosolicitado"`
	FechaAprobado      time.Time `json:"fechaaprobado" bson:"fechaaprobado"`
	MontoAprobado      float64   `json:"montoaprobado" bson:"montoaprobado"`
	Estatus            int       `json:"estatus" bson:"estatus"`
	EstatusSeguimiento int       `json:"estatusseguimiento" bson:"estatusseguimiento"`
	Reembolso          Reembolso `json:"Reembolso,omitempty" bson:"reembolso"`
	Usuario            string    `json:"usuario" bson:"usuario"`
}

type ActualizarReembolso struct {
	ID            string    `json:"id" bson:"id"`
	Reembolso     Reembolso `json:"Reembolso" bson:"Reembolso"`
	Numero        string    `json:"numero" bson:"numero"`
	Posicion      int       `json:"posicion" bson:"posicion"`
	Observaciones []string  `json:"observaciones" bson:"observaciones"`
}

type EstatusReembolso struct {
	ID       string `json:"id" bson:"id"`
	Numero   string `json:"numero" bson:"numero"`
	Posicion int    `json:"posicion" bson:"posicion"`
	Estatus  int    `json:"estatus" bson:"estatus"`
}

func (fact *Factura) Consultar(rif string, numero string) (jSon []byte, err error) {
	var result Factura
	var M fanb.Mensaje

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CFACTURA)
	err = c.Find(bson.M{"beneficiario.rif": rif, "numero": numero}).One(&result)
	if err != nil {
		fmt.Println("Err. Factura")
		//return
	}
	M.Tipo = 0
	M.Mensaje = "Factura Disponible"
	if result.Numero != "" {
		M.Tipo = 1
		M.Mensaje = "La factura ya se encuentra registrada"
	}
	jSon, err = json.Marshal(M)
	return
}

//GenerarReporte Control de Listados
func (r *Reembolso) GenerarReporte(estatus string) (jSon []byte, err error) {
	var result []Reembolso
	var lista []interface{}

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CREEMBOLSO)
	err = c.Find(bson.M{"estatus": estatus, "fechaaprobado": "numero"}).All(&result)
	if err != nil {
		fmt.Println("Err. Reporte")
		//return
	}
	for _, v := range result {
		lst := make(map[string]interface{})
		lst["cedula"] = v.Numero
		lista = append(lista, lst)
	}
	return
}
