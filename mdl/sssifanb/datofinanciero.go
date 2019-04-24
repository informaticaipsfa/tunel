package sssifanb

//DatoFinanciero Establecer un modulo de datos bancarios
type DatoFinanciero struct {
	Tipo        string `json:"tipo" bson:"tipo"`
	Institucion string `json:"institucion" bson:"institucion"`
	Cuenta      string `json:"cuenta" bson:"cuenta"`
	Prioridad   string `json:"prioridad" bson:"prioridad"`
	Autorizado  string `json:"autorizado" bson:"autorizado"`
	Titular     string `json:"titular" bson:"titular"`
}
