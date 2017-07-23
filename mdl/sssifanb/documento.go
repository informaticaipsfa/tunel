package sssifanb

//PartidaNacimiento Documentos
type PartidaNacimiento struct {
	Registro string `json:"registro" bson:"registro"`
	Ano      string `json:"ano" bson:"ano"`
	Acta     string `json:"acta" bson:"acta"`
	Folio    string `json:"folio" bson:"folio"`
	Libro    string `json:"libro" bson:"libro"`
	Img      string `json:"img" bson:"img"`
}
