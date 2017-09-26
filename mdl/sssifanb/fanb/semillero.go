package fanb

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/informaticaipsfa/tunel/sys"
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
func (s *Semillero) SalvarMGO(coleccion string) (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(coleccion)
	err = c.Insert(s)
	return
}

//Maximo una persona mediante el metodo de MongoDB
func (s *Semillero) Maximo(coleccion string) (maximo int, err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(coleccion)
	orden := bson.M{"$sort": bson.M{"codigo": -1}}
	limite := bson.M{"$limit": 1}
	err = c.Pipe([]bson.M{orden, limite}).One(&s)
	if err == nil {
		maximo = s.Codigo + 1
		s.Codigo = maximo
	}
	s.SalvarMGO(coleccion)
	return
}
