package credito

type Credito struct {
	Prestamo Prestamo `json:"Prestamo,omitempty" bson:"prestamo"`
}

//Solicitud Solicitar Prestamo o credito
type Solicitud struct {
	Monto          float64 `json:"monto,omitempty" bson:"monto"`
	Cantidad       int64   `json:"cantidad,omitempty" bson:"cantidad"`
	Cuotas         int64   `json:"cuotas,omitempty" bson:"cuotas"`
	PorcentajeTasa float64 `json:"porcentajetasa,omitempty" bson:"porcentajetasa"`
	Concepto       string  `json:"concepto,omitempty" bson:"concepto"`
	Tipo           string  `json:"tipo,omitempty" bson:"tipo"` //Aguinaldo, Vacaciones, Especial
	Total          float64 `json:"total,omitempty" bson:"total"`
}

//Prestamo Prestamoes
type Prestamo struct {
	Vacacional         []Vacacional         `json:"Vacacional,omitempty" bson:"vacacional"`
	Educativo          []Educativo          `json:"Educativo,omitempty" bson:"educativo"`
	Parcelas           []Parcelas           `json:"Parcelas,omitempty" bson:"parcelas"`
	Personal           []Personal           `json:"Personal,omitempty" bson:"personal"`
	Articulos          []Articulos          `json:"Articulos,omitempty" bson:"articulos"`
	MiCasaBienEquipada []MiCasaBienEquipada `json:"MiCasaBienEquipada,omitempty" bson:"micasabienequipada"`
}

//Hipotecario viviendas
type Hipotecario struct {
}

//Vehiculo viviendas
type Vehiculo struct {
}

//Nuevo creacion de nuevo prestamo
func (PP *Prestamo) Nuevo() {

}

//Nuevo Credito Para vivienda
func (CH *Hipotecario) Nuevo() {

}
