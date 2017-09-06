package tramitacion

import "time"

type Apoyo struct {
	Numero            string          `json:"numero" bson:"numero"`
	Estatus           int             `json:"estatus" bson:"estatus"`
	FechaCreacion     time.Time       `json:"fechacreacion" bson:"fechacreacion"`
	MontoSolicitado   float64         `json:"montosolicitado" bson:"montosolicitado"`
	CuentaBancaria    DatoFinanciero  `json:"CuentaBancaria" bson:"CuentaBancaria"`
	Responsable       string          `json:"responsable" bson:"responsable"`
	Concepto          []ConceptoApoyo `json:"Concepto" bson:"concepto"`
	TipoCaso          int             `json:"tipo" bson:"tipo"`
	FondoContingencia string          `json:"fondo" bson:"fondo"`
	FechaAprobado     time.Time       `json:"fechaaprobado" bson:"fechaaprobado"`
	MontoAprobado     float64         `json:"montoaprobado" bson:"montoaprobado"`
	Requisitos        []int           `json:"requisitos" bson:"requisitos"`
	Componente        string          `json:"componente" bson:"componente"`
	Grado             string          `json:"grado" bson:"grado"`
	Clase             string          `json:"clase" bson:"clase"`
	Situacion         string          `json:"situacion" bson:"situacion"`
	Direccion         Direccion       `json:"Direccion" bson:"direccion"`
	Telefono          Telefono        `json:"Telefono" bson:"telefono"`
	Correo            Correo          `json:"Correo" bson:"correo"`
	Seguimiento       Seguimiento     `json:"Seguimiento" bson:"seguimiento"`
	Usuario           string          `json:"usuario" bson:"usuario"`
}

type ConceptoApoyo struct {
	Descripcion      string  `json:"descripcion" bson:"descripcion"`
	DatoFactura      Factura `json:"DatoFactura" bson:"datofactura"`
	Afiliado         string  `json:"afiliado" bson:"afiliado"` //Cedula, Guión (-), Nombre
	Requisitos       []int   `json:"requisitos" bson:"requisitos"`
	Patologia        string  `json:"patologia" bson:"patologia"`
	MontoAseguradora float64 `json:"montoaseguradora" bson:"montoaseguradora"`
	MontoAportar     float64 `json:"montoaportar" bson:"montoaportar"`
}

type ColeccionApoyo struct {
	ID                 string    `json:"id" bson:"id"`
	Nombre             string    `json:"nombre" bson:"nombre"`
	Numero             string    `json:"numero" bson:"numero"`
	FechaCreacion      time.Time `json:"fechacreacion" bson:"fechacreacion"`
	MontoSolicitado    float64   `json:"montosolicitado" bson:"montosolicitado"`
	FechaAprobado      time.Time `json:"fechaaprobado" bson:"fechaaprobado"`
	MontoAprobado      float64   `json:"montoaprobado" bson:"montoaprobado"`
	Estatus            int       `json:"estatus" bson:"estatus"`
	EstatusSeguimiento int       `json:"estatusseguimiento" bson:"estatusseguimiento"`
	Apoyo              Apoyo     `json:"Apoyo,omitempty" bson:"apoyo"`
	Usuario            string    `json:"usuario" bson:"usuario"`
}

type ActualizarApoyo struct {
	ID            string
	Apoyo         Apoyo
	Numero        string
	Posicion      int
	Observaciones []string
}

type EstatusApoyo struct {
	ID      string
	Numero  string
	Estatus int
}

// func (ae *ApoyoEconomico) Listar() []ApoyoEconomico {
// 	return
// }
