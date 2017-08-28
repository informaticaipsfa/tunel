package tramitacion

import (
	"time"
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
	ID            string
	Reembolso     Reembolso
	Numero        string
	Posicion      int
	Observaciones []string
}

type EstatusReembolso struct {
	ID      string
	Numero  string
	Estatus int
}
