package fanb

import (
	"encoding/json"
	"fmt"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

//Estado Sitio
type Estado struct {
	Nombre    string      `json:"nombre" bson:"nombre"`
	Codigo    string      `json:"codigo" bson:"codigo"`
	Ciudad    []Ciudad    `json:"Ciudad" bson:"ciudad"`
	Municipio []Municipio `json:"Municipio" bson:"municipio"`
}

//Ciudad Sitio
type Ciudad struct {
	Capital int    `json:"codigo" bson:"capital"`
	Nombre  string `json:"nombre" bson:"nombre"`
}

//Municipio Sitio
type Municipio struct {
	Nombre    string   `json:"nombre" bson:"nombre"`
	Parroquia []string `json:"Parroquia" bson:"parroquia"`
}

//SalvarMGO Guardar
func (e *Estado) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(sys.CBASE).C(colecion)
		err = c.Insert(e)
	} else {
		c := sys.MGOSession.DB("ipsfa_test").C("estado")
		err = c.Insert(e)
	}
	return
}

//ConsultarEstado una persona mediante el metodo de MongoDB
func (e *Estado) ConsultarEstado() (jSon []byte, err error) {
	var msj Mensaje
	var lst []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CESTADO)
	err = c.Find(nil).All(&lst)
	if err != nil {
		msj.Tipo = 0
		msj.Mensaje = err.Error()
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(lst)
	}
	return
}

//Consultar una persona mediante el metodo de MongoDB
func (e *Estado) Consultar(estado string) (jSon []byte, err error) {
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CESTADO)
	err = c.Find(bson.M{"codigo": estado}).One(&e)
	if err != nil {
		msj.Tipo = 0
		msj.Mensaje = err.Error()
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(e)
	}
	return
}

//ActualizarMGO Actualizar
func (e *Estado) ActualizarMGO(donde bson.M, estado map[string]interface{}) (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CESTADO)
	err = c.Update(donde, bson.M{"$set": estado})

	if err != nil {
		fmt.Println("Actualizar: -> " + err.Error())
		return
	}
	return
}
