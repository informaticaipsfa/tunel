package credito

//DatoFinanciero Establecer un modulo de datos bancarios
type DatoFinanciero struct {
	Tipo        string `json:"tipo" bson:"tipo"`
	Institucion string `json:"institucion" bson:"institucion"`
	Cuenta      string `json:"cuenta" bson:"cuenta"`
	Prioridad   string `json:"prioridad" bson:"prioridad"`
	Autorizado  string `json:"autorizado" bson:"autorizado"`
	Titular     string `json:"titular" bson:"titular"`
}

type Credito struct {
	Prestamo Prestamo `json:"Prestamo,omitempty" bson:"prestamo"`
}

//Solicitud Solicitar Prestamo o credito
type Solicitud struct {
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
	TotalInterese    float64        `json:"totalintereses,omitempty" bson:"totalintereses"`     //Monto Aprobado
	Intereses        float64        `json:"intereses,omitempty" bson:"intereses"`               //Monto Aprobado
	PorcentajeSeguro float64        `json:"porcentajeseguro,omitempty" bson:"porcentajeseguro"` //Monto Aprobado
	TotalDepositar   float64        `json:"totaldepositar,omitempty" bson:"totaldepositar"`     //Monto Aprobado

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
func (PP *Solicitud) NuevoPrestamo() {

}

//Nuevo Credito Para vivienda
func (CH *Hipotecario) Nuevo() {

}
