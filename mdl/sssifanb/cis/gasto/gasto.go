package gasto

type GastoFarmaceutico struct {
	MedicinaAltoCosto []WAltoCosto `json:"MedicinaAltoCosto" bson:"medicinaaltocosto"`
}
