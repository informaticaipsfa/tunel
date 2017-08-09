package fanb

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gesaodin/tunel-ipsfa/sys"
)

//Semillero Estructura
type Semillero struct {
	Codigo        int       `json:"codigo" bson:"codigo"`
	Serial        string    `json:"serial" bson:"serial"`
	Autor         string    `json:"autor" bson:"autor"`
	ResponsableID string    `json:"responsableid" bson:"responsableid"`
	Vencimiento   time.Time `json:"vencimiento" bson:"vencimiento"`
	Estatus       int       `json:"estatus" bson:"estatus"`
}

//SalvarMGO Guardar
func (s *Semillero) SalvarMGO() (err error) {

	c := sys.MGOSession.DB(BASEDATOS).C(SEMILLERO)
	err = c.Insert(s)

	return
}

//Consultar una persona mediante el metodo de MongoDB
func (s *Semillero) Consultar(estado string) (jSon []byte, err error) {
	var msj Mensaje
	c := sys.MGOSession.DB(BASEDATOS).C(SEMILLERO)
	err = c.Find(bson.M{"codigo": estado}).One(&s)
	if err != nil {
		msj.Tipo = 0
		msj.Mensaje = err.Error()
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(s)
	}
	return
}

//Maximo una persona mediante el metodo de MongoDB
func (s *Semillero) Maximo() (maximo int, err error) {
	c := sys.MGOSession.DB(BASEDATOS).C(SEMILLERO)
	orden := bson.M{"$sort": bson.M{"codigo": -1}}
	limite := bson.M{"$limit": 1}
	err = c.Pipe([]bson.M{orden, limite}).One(&s)
	if err == nil {
		maximo = s.Codigo + 1
		s.Codigo = maximo
	}
	s.SalvarMGO()
	return
}
