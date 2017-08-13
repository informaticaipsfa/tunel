package tramitacion

type Programa struct {
	ApoyoEconomico  []ApoyoEconomico  `json:"Apoyo" bson:"apoyo"`
	ReembolsoMedico []ReembolsoMedico `json:"Reembolso" bson:"reembolso"`
	CartaAval       []CartaAval       `json:"CartaAval" bson:"cartaaval"`
}
