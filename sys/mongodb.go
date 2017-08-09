package sys

type Mongo struct {
}

//Salvar documentos
func (m *Mongo) Salvar(mgo interface{}, coleccion string) (err error) {

	c := MGOSession.DB(CBASE).C(coleccion)
	err = c.Insert(mgo)

	return
}
