package credito

//Personal Prestamo:
type Personal struct {
	Solicitud Solicitud `json:"Solicitud,omitempty" bson:"solicitud"`
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
