package sssifanb

//Correo Direcciones electronicas
type Correo struct {
	Principal     string `json:"principal,omitempty" bson:"principal"`
	Alternativo   string `json:"alternativo,omitempty" bson:"alternativo"`
	Institucional string `json:"institucional,omitempty" bson:"institucional"`
}
