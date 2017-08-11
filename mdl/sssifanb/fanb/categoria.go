package fanb

import (
	"encoding/json"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"gopkg.in/mgo.v2/bson"
)

type Categoria struct {
	Codigo string `json:"codigo" bson:"codigo"`
	Nombre string `json:"nombre" bson:"nombre"`
	Siglas string `json:"siglas" bson:"siglas"`
}

//Mensaje del sistema
/*type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}*/

//SalvarMGO Guardar
func (cate *Categoria) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(BASEDATOS).C(colecion)
		err = c.Insert(cate)
	} else {
		c := sys.MGOSession.DB(BASEDATOS).C("categoria")
		err = c.Insert(cate)
	}

	return
}

//Consultar una persona mediante el metodo de MongoDB
func (cate *Categoria) Consultar(categoria string) (jSon []byte, err error) {
	var msj Mensaje
	c := sys.MGOSession.DB(BASEDATOS).C(CATEGORIA)
	err = c.Find(bson.M{"codigo": categoria}).One(&cate)
	if err != nil {
		msj.Tipo = 0
		msj.Mensaje = err.Error()
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(cate)
	}
	return
}
