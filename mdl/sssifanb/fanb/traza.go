package fanb

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gesaodin/tunel-ipsfa/sys"
)

//Traza Historico del Usuario
type Traza struct {
	Usuario   string    `json:"usuario" bson:"usuario"`
	Time      time.Time `json:"tiempo" bson:"tiempo"`
	Log       string    `json:"log" bson:"log"`
	Documento string    `json:"documento" bson:"documento"`
}

//Crear Trazabilidad
func (t *Traza) Crear() (err error) {
	c := sys.MGOSession.DB(BASEDATOS).C(TRAZA)
	err = c.Insert(t)
	return
}

//Consultar Trazabilidad
func (t *Traza) Consultar() (lst []Traza, err error) {
	c := sys.MGOSession.DB(BASEDATOS).C(TRAZA)
	err = c.Find(bson.M{}).All(&lst)
	return
}
