package sssifanb

type Persona struct {
	DatoBasico      DatoBasico       `json:"DatoBasico,omitempty" bson:"datobasico"`
	Correo          Correo           `json:"Correo,omitempty" bson:"correo"`
	DatoFisico      DatoFisico       `json:"DatoFisico,omitempty" bson:"datofisico"`
	DatoFisionomico DatoFisionomico  `json:"DatoFisionomico,omitempty" bson:"datofisionomico"`
	RedSocial       RedSocial        `json:"RedSocial,omitempty" bson:"redsocial"`
	Telefono        []Telefono       `json:"Telefono,omitempty" bson:"telefono"`
	Direccion       []Direccion      `json:"Direccion,omitempty" bson:"direccion"`
	HistoriaMedica  []HistoriaMedica `json:"HistoriaMedica,omitempty" bson:"historiamedica"`
	DatoFinanciero  DatoFinanciero   `json:"DatoFinanciero,omitempty" bson:"datofinanciero"`
	URLFoto         string           `json:"foto,omitempty" bson:"foto"`
	URLHuella       string           `json:"huella,omitempty" bson:"huella"`
	URLFirma        string           `json:"firma,omitempty" bson:"firma"`
	URLCedula       string           `json:"urlcedula,omitempty" bson:"urlcedula"`

	// Militar         Militar          `json:"Militar,omitempty" bson:"militar"`
}
