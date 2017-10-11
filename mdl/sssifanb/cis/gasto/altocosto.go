package gasto

import (
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
	"gopkg.in/mgo.v2/bson"
)

type AltoCosto struct {
	NombreComercial  string    `json:"nombrecomercial" bson:"nombrecomercial"`
	Presentacion     string    `json:"presentacion" bson:"presentacion"`
	Dosis            string    `json:"dosis" bson:"dosis"`
	Cantidad         string    `json:"cantidad" bson:"cantidad"`
	FechaInicio      time.Time `json:"fechainicio" bson:"fechainicio"`
	FechaVencimiento time.Time `json:"fechavencimiento" bson:"fechavencimiento"`
}

type WAltoCosto struct {
	Numero    string      `json:"numero" bson:"numero"`
	ID        string      `json:"id" bson:"id"`
	IDF       string      `json:"idf" bson:"idf"`
	Medicina  []AltoCosto `json:"Medicina" bson:"medicina"`
	Afiliado  string      `json:"afiliado" bson:"afiliado"`
	Usuario   string      `json:"usuario" bson:"usuario"`
	Fecha     time.Time   `json:"fecha" bson:"fecha"`
	Direccion Direccion   `json:"Direccion" bson:"direccion"`
	Telefono  Telefono    `json:"Telefono" bson:"telefono"`
	Correo    Correo      `json:"Correo" bson:"correo"`
}

//Direccion ruta y secciones
type Direccion struct {
	Tipo         int    `json:"tipo,omitempty" bson:"tipo"` //domiciliaria, trabajo, emergencia
	Ciudad       string `json:"ciudad,omitempty" bson:"ciudad"`
	Estado       string `json:"estado,omitempty" bson:"estado"`
	Municipio    string `json:"municipio,omitempty" bson:"municipio"`
	Parroquia    string `json:"parroquia,omitempty" bson:"parroquia"`
	CalleAvenida string `json:"calleavenida" bson:"calleavenida"`
	Casa         string `json:"casa" bson:"casa"`
	Apartamento  string `json:"apartamento" bson:"apartamento"`
	Numero       int    `json:"numero,omitempty" bson:"numero"`
}

type Telefono struct {
	Movil        string `json:"movil,omitempty" bson:"movil"`
	Domiciliario string `json:"domiciliario,omitempty" bson:"domiciliario"`
	Emergencia   string `json:"emergencia,omitempty" bson:"emergencia"`
}

//Correo Direcciones electronicas
type Correo struct {
	Principal     string `json:"principal,omitempty" bson:"principal"`
	Alternativo   string `json:"alternativo,omitempty" bson:"alternativo"`
	Institucional string `json:"institucional,omitempty" bson:"institucional"`
}

//Crear Registrando
func (ac *WAltoCosto) Crear() (jSon []byte, err error) {
	var M fanb.Mensaje
	M.Mensaje = "Creando Medicina Alto Costo"
	M.Tipo = 1
	altoc := make(map[string]interface{})
	altoc["cis.gasto.medicinaaltocosto"] = ac
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)

	// fmt.Println(altocosto.Medicina.Afiliado)
	err = c.Update(bson.M{"id": ac.ID}, bson.M{"$push": altoc})
	util.Error(err)
	return
}
