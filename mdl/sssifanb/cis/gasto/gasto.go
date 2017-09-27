package gasto

type GastoFarmaceutico struct {
	MedicinaAltoCosto AltoCosto `json:"MedicinaAltoCosto" bson:"medicinaaltocosto"`
}
