package fanb

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/informaticaipsfa/tunel/sys"
)

//Traza Historico del Usuario
type Traza struct {
	Usuario   string    `json:"usuario" bson:"usuario"`
	Time      time.Time `json:"tiempo" bson:"tiempo"`
	Log       string    `json:"log" bson:"log"`
	Documento string    `json:"documento" bson:"documento"`
	IP        string    `json:"ip" bson:"ip"`
}

//Traza Historico del Usuario
type TrazaCIS struct {
	Usuario   string      `json:"usuario" bson:"usuario"`
	Time      time.Time   `json:"tiempo" bson:"tiempo"`
	Log       string      `json:"log" bson:"log"`
	Documento interface{} `json:"documento" bson:"documento"`
	IP        string      `json:"ip" bson:"ip"`
}

//Crear Trazabilidad
func (t *Traza) Crear() (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTRAZA)
	err = c.Insert(t)
	return
}

//Consultar Trazabilidad
func (t *Traza) Consultar() (lst []Traza, err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTRAZA)
	err = c.Find(bson.M{}).All(&lst)
	return
}

//Crear Trazabilidad
func (t *Traza) CrearHistoricoConsulta(colecion string) (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(colecion)
	err = c.Insert(t)
	return
}

//Crear Trazabilidad
func (t *TrazaCIS) Crear() (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTRAZACIS)
	err = c.Insert(t)
	return
}
