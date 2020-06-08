package credito

import "time"

//Personal Prestamo:
type Personal struct {
	// Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
	Oid              string         `json:"oid,omitempty" bson:"oid"`
	Cedula           string         `json:"cedula,omitempty" bson:"cedula"`                 //Monto total del credito solicitado
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

//Vacacional Prestamo:
type Vacacional struct {
	Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
}

//Educativo Prestamo:
type Educativo struct {
	Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
}

//Parcelas Prestamo:
type Parcelas struct {
	Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
}

//Articulos Prestamo:
type Articulos struct {
	Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
}

//MiCasaBienEquipada Prestamo:
type MiCasaBienEquipada struct {
	Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
}
