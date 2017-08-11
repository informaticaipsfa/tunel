package fanb

import (
	"encoding/json"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"gopkg.in/mgo.v2/bson"
)

type Componente struct {
	Codigo string  `json:"codigo" bson:"codigo"`
	Nombre string  `json:"nombre" bson:"nombre"`
	Siglas string  `json:"siglas" bson:"siglas"`
	Grado  []Grado `json:"Grado" bson:"Grado"`
}

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}

//SalvarMGO Guardar
func (comp *Componente) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(BASEDATOS).C(colecion)
		err = c.Insert(comp)
	} else {
		c := sys.MGOSession.DB(BASEDATOS).C("componente")
		err = c.Insert(comp)
	}

	return
}

//Consultar una persona mediante el metodo de MongoDB
func (comp *Componente) Consultar(componente string) (jSon []byte, err error) {
	var msj Mensaje
	c := sys.MGOSession.DB(BASEDATOS).C(COMPONENTE)
	err = c.Find(bson.M{"codigo": componente}).One(&comp)
	if err != nil {
		msj.Tipo = 0
		msj.Mensaje = err.Error()
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(comp)
	}
	return
}
