package gasto

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

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

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
}

type WTratamiento struct {
	ID                    string
	TratamientoProlongado TratamientoProlongado
	Nombre                string
	Observaciones         string
}

//Crear Creacion de Cuenta
func (tp *WTratamiento) Crear() (jSon []byte, err error) {
	var M Mensaje
	fmt.Println("Control...")
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTRATAMIENTO)
	err = c.Insert(tp)
	if err != nil {
		fmt.Println("Error creando reembolso det: ")
		// return
	}

	tra := make(map[string]interface{})
	tra["cis.gasto.tratamientoprolongado"] = tp.TratamientoProlongado
	co := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = co.Update(bson.M{"id": tp.ID}, bson.M{"$push": tra})
	if err != nil {
		fmt.Println("Cedula: " + tp.ID + " -> " + err.Error())
		// return
	}

	M.Mensaje = tp.TratamientoProlongado.Numero
	M.Tipo = 1

	jSon, err = json.Marshal(M)
	return
}
