package tramitacion

import "time"

type CartaAval struct {
	Numero            string              `json:"numero" bson:"numero"`
	Estatus           int                 `json:"estatus" bson:"estatus"`
	FechaCreacion     time.Time           `json:"fechacreacion" bson:"fechacreacion"`
	MontoSolicitado   float64             `json:"montosolicitado" bson:"montosolicitado"`
	CuentaBancaria    DatoFinanciero      `json:"CuentaBancaria" bson:"CuentaBancaria"`
	Responsable       string              `json:"responsable" bson:"responsable"`
	Concepto          []ConceptoCartaAval `json:"Concepto" bson:"concepto"`
	TipoCaso          int                 `json:"tipo" bson:"tipo"`
	FondoContingencia string              `json:"fondo" bson:"fondo"`
	FechaAprobado     time.Time           `json:"fechaaprobado" bson:"fechaaprobado"`
	MontoAprobado     float64             `json:"montoaprobado" bson:"montoaprobado"`
	Requisitos        []int               `json:"requisitos" bson:"requisitos"`
	Componente        string              `json:"componente" bson:"componente"`
	Grado             string              `json:"grado" bson:"grado"`
	Clase             string              `json:"clase" bson:"clase"`
	Situacion         string              `json:"situacion" bson:"situacion"`
	Direccion         Direccion           `json:"Direccion" bson:"direccion"`
	Telefono          Telefono            `json:"Telefono" bson:"telefono"`
	Correo            Correo              `json:"Correo" bson:"correo"`
	Seguimiento       Seguimiento         `json:"Seguimiento" bson:"seguimiento"`
	Usuario           string              `json:"usuario" bson:"usuario"`
}

type ConceptoCartaAval struct {
<<<<<<< HEAD
	Motivo             string    `json:"motivo" bson:"motivo"`
	Diagnostico        string    `json:"diagnostico" bson:"diagnostico"`
	Descripcion        string    `json:"descripcion" bson:"descripcion"`
	DatoFactura        Factura   `json:"DatoFactura" bson:"datofactura"`
	Afiliado           string    `json:"afiliado" bson:"afiliado"` //Cedula, Guión (-), Nombre
	Requisitos         []int     `json:"requisitos" bson:"requisitos"`
	MontoPresupuesto   float64   `json:"montopresupuesto" bson:"montopresupuesto"`
	FechaPresupuesto   time.Time `json:"fechapresupuesto" bson:"fechapresupuesto"`
	NumeroPresupuesto  string    `json:"numeropresupuesto" bson:"numeropresupuesto"`
	MontoSeguro        float64   `json:"montoseguro" bson:"montoseguro"`
	FechaSeguro        time.Time `json:"fechaseguro" bson:"fechaseguro"`
	MontoAfiliado      float64   `json:"montoafiliado" bson:"montoafiliado"`
	PorcentajeAfiliado float64   `json:"porcentajeafi" bson:"porcentajeafi"`
=======
	Motivo      string  `json:"motivo" bson:"motivo"`
	Diagnostico string  `json:"diagnostico" bson:"diagnostico"`
	Descripcion string  `json:"descripcion" bson:"descripcion"`
	DatoFactura Factura `json:"DatoFactura" bson:"datofactura"`
	Afiliado    string  `json:"afiliado" bson:"afiliado"` //Cedula, Guión (-), Nombre
	Requisitos  []int   `json:"requisitos" bson:"requisitos"`
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
}

type ColeccionCartaAval struct {
	ID                 string    `json:"id" bson:"id"`
	Nombre             string    `json:"nombre" bson:"nombre"`
	Numero             string    `json:"numero" bson:"numero"`
	FechaCreacion      time.Time `json:"fechacreacion" bson:"fechacreacion"`
	MontoSolicitado    float64   `json:"montosolicitado" bson:"montosolicitado"`
	FechaAprobado      time.Time `json:"fechaaprobado" bson:"fechaaprobado"`
	MontoAprobado      float64   `json:"montoaprobado" bson:"montoaprobado"`
	Estatus            int       `json:"estatus" bson:"estatus"`
	EstatusSeguimiento int       `json:"estatusseguimiento" bson:"estatusseguimiento"`
	Carta              CartaAval `json:"Carta,omitempty" bson:"carta"`
	Usuario            string    `json:"usuario" bson:"usuario"`
}

type ActualizarCartaAval struct {
	ID            string
	Carta         CartaAval
	Numero        string
	Posicion      int
	Observaciones []string
}

type EstatusCartaAval struct {
	ID      string
	Numero  string
	Estatus int
}
