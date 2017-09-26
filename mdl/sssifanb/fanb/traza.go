package fanb

import (
	"time"

	"gopkg.in/mgo.v2/bson"

<<<<<<< HEAD
	"github.com/informaticaipsfa/tunel/sys"
=======
	"github.com/gesaodin/tunel-ipsfa/sys"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
)

//Traza Historico del Usuario
type Traza struct {
	Usuario   string    `json:"usuario" bson:"usuario"`
	Time      time.Time `json:"tiempo" bson:"tiempo"`
	Log       string    `json:"log" bson:"log"`
	Documento string    `json:"documento" bson:"documento"`
	IP        string    `json:"ip" bson:"ip"`
}

<<<<<<< HEAD
//Traza Historico del Usuario
type TrazaCIS struct {
	Usuario   string      `json:"usuario" bson:"usuario"`
	Time      time.Time   `json:"tiempo" bson:"tiempo"`
	Log       string      `json:"log" bson:"log"`
	Documento interface{} `json:"documento" bson:"documento"`
	IP        string      `json:"ip" bson:"ip"`
}

=======
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
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
<<<<<<< HEAD

//Crear Trazabilidad
func (t *TrazaCIS) Crear() (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTRAZACIS)
	err = c.Insert(t)
	return
}
=======
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
