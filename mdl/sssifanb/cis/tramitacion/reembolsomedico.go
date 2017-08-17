package tramitacion

import (
	"time"
)

type Reembolso struct {
	Numero          string         `json:"numero" bson:"numero"`
	Estatus         int            `json:"estatus" bson:"estatus"`
	MontoSolicitado float32        `json:"montosolicitado" bson:"montosolicitado"`
	FechaCreacion   time.Time      `json:"fechacreacion" bson:"fechacreacion"`
	CuentaBancaria  DatoFinanciero `json:"CuentaBancaria" bson:"CuentaBancaria"`
	Responsable     string         `json:"responsable" bson:"responsable"`
	Concepto        []Concepto     `json:"Concepto" bson:"concepto"`
	MontoAprobado   float32        `json:"montoaprobado" bson:"montoaprobado"`
	FechaAprobado   time.Time      `json:"fechaaprobado" bson:"fechaaprobado"`
	Requisitos      []int          `json:"requisitos" bson:"requisitos"`
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
}

type Factura struct {
	Numero       string    `json:"numero" bson:"numero"`
	Control      string    `json:"control" bson:"control"`
	Fecha        time.Time `json:"fecha" bson:"fecha"`
	Monto        float32   `json:"monto" bson:"monto"`
	Beneficiario Proveedor `json:"Beneficiario" bson:"beneficiario"`
}

type Proveedor struct {
	Rif         string         `json:"rif" bson:"rif"`
	RazonSocial string         `json:"razonsocial" bson:"razonsocial"`
	Tipo        string         `json:"tipo" bson:"tipo"`
	Direccion   string         `json:"direccion" bson:"direccion"`
	Banco       DatoFinanciero `json:"Banco" bson:"banco"`
	Descripcion string         `json:"descripcion" bson:"descripcion"`
}
