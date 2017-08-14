package tramitacion

type Programa struct {
	// ApoyoEconomico []Apoyo     `json:"Apoyo" bson:"apoyo"`
	Reembolso []Reembolso `json:"Reembolso" bson:"reembolso"`
	// CartaAval      []CartaAval `json:"CartaAval" bson:"cartaaval"`
}
