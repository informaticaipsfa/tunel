package sys

import "gopkg.in/mgo.v2/bson"

type Mongo struct{}

//Salvar documentos
func (m *Mongo) Salvar(mgo interface{}, coleccion string) (err error) {
	if coleccion != "" {
		c := MGOSession.DB("ipsfa_test").C(coleccion)
		err = c.Insert(mgo)
	} else {
		c := MGOSession.DB("ipsfa_test").C("persona")
		err = c.Insert(mgo)
	}
	return
}

//ConsultarID Permite realizar consultas puntuales
func (m *Mongo) ConsultarID(id string) (obj interface{}, err error) {
	c := MGOSession.DB("ipsfa_test").C("persona")
	err = c.Find(bson.M{"id": id}).One(&m)
	return
}

//ActualizarID Actualizar DatoBasico
func (m *Mongo) ActualizarID(id string, obj map[string]interface{}) bool {

	return true
}
