package gasto

import "time"

type TratamientoProlongado struct {
	Numero            string        `json:"numero" bson:"numero"`
	FechaInforme      time.Time     `json:"fechainforme" bson:"fechainforme"`
	TipoCentro        string        `json:"tipocentro" bson:"tipocentro"`
	NombreCentro      string        `json:"nombrecentro" bson:"nombrecentro"`
	MedicoContratante Medico        `json:"Medico" bson:"medico"`
	MedicoAvala       MedicoAvala   `json:"MedicoAvala" bson:"medicoavala"`
	PrestadorServicio string        `json:"prestadorservicio" bson:"prestadorservicio"`
	Zona              string        `json:"zona" bson:"zona"`
	Patologia         []Patologia   `json:"Patologia" bson:"patologia"`
	Tratamiento       []Tratamiento `json:"Tratamiento" bson:"tratamiento"`
	Grado             string        `json:"grado" bson:"grado"`
	Componente        string        `json:"componente" bson:"componente"`
	Direccion         Direccion     `json:"Direccion" bson:"direccion"`
	Telefono          Telefono      `json:"Telefono" bson:"telefono"`
	Correo            Correo        `json:"Correo" bson:"correo"`
	FechaCreacion     time.Time     `json:"fechacreado" bson:"fechacreado"`
	Estatus           int           `json:"estatus" bson:"estatus"`
}

//Medico: datos personales del medico
type Medico struct {
	Nombre       string `json:"nombre" bson:"nombre"`
	Cedula       string `json:"cedula" bson:"cedula"`
	Especialidad string `json:"especialidad" bson:"especialidad"`
	Codigo       string `json:"codigo" bson:"codigo"`
	CodigoMPPS   string `json:"codigompps" bson:"codigompps"`
}

//Medico: datos personales del medico
type MedicoAvala struct {
	NombreCentro string `json:"nombrecentro" bson:"nombrecentro"`
	Nombre       string `json:"nombre" bson:"nombre"`
	Cedula       string `json:"cedula" bson:"cedula"`
	Especialidad string `json:"especialidad" bson:"especialidad"`
	Codigo       string `json:"codigo" bson:"codigo"`
	CodigoMPPS   string `json:"codigompps" bson:"codigompps"`
}

//Patologia Control Medico
type Patologia struct {
	Nombre string `json:"nombre" bson:"nombre"`
}

//Tratamiento Arreglo de Medicamentos
type Tratamiento struct {
	Nombre           string    `json:"nombre" bson:"nombre"`
	Principio        string    `json:"principio" bson:"principio"`
	Presentacion     string    `json:"presentacion" bson:"presentacion"`
	Dosis            string    `json:"dosis" bson:"dosis"`
	Cantidad         string    `json:"cantidad" bson:"cantidad"`
	FechaVencimiento time.Time `json:"fechavencimiento" bson:"fechavencimiento"`
}
