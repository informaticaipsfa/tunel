package gasto

import "time"

type TratamientoProlongado struct {
	FechaInforme      time.Time
	TipoCentro        string
	NombreCentro      string
	MedicoContratante string
	CedulaMedico      string
	Especialidad      string
	CodigoMedico      string
	CodigoMPPS        string
	PrestadorServicio string
	Zona              string
	Patologia         []Patologia
	Tratamiento       []Tratamiento
	Grado             string
	Componente        string
}

//Patologia Control Medico
type Patologia struct {
	Nombre string
}

//Tratamiento Arreglo de Medicamentos
type Tratamiento struct {
	Principio        string
	Nombre           string
	Presentacion     string
	Dosis            string
	Cantidad         string
	FechaVencimiento time.Time
}
