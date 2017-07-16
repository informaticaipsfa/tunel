package sssifanb

type DatoFinanciero struct {
	TipoCuenta  string `json:"tipocuenta" bson:"tipocuenta"`
	Institucion string `json:"institucion" bson:"institucion"`
	Cuenta      string `json:"cuenta" bson:"cuenta"`
}
