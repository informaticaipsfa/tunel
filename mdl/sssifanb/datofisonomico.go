package sssifanb

type DatoFisionomico struct {
	GrupoSanguineo string  `json:"gruposanguineo" bson:"gruposanguineo"`
	ColorPiel      string  `json:"colorpiel" bson:"colorpiel"`
	ColorOjos      string  `json:"colorojos" bson:"colorojos"`
	ColorCabello   string  `json:"colorcabello" bson:"colorcabello"`
	Estatura       float32 `json:"estatura" bson:"estatura"`
	SenaParticular string  `json:"senaparticular" bson:"senaparticular"`
}
