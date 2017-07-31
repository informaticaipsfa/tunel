package sssifanb

type Telefono struct {
	Movil        string `json:"movil,omitempty" bson:"movil"`
	Domiciliario string `json:"domiciliario,omitempty" bson:"domiciliario"`
	Emergencia   string `json:"emergencia,omitempty" bson:"emergencia"`
}
