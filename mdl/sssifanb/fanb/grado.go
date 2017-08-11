package fanb
package sssifanb

import (
	"encoding/json"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"gopkg.in/mgo.v2/bson"
)

type Grado struct {
	Codigo      string `json:"codigo" bson:"codigo"`
	Rango       string `json:"rango" bson:"rango"`
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
}

//Mensaje del sistema
/*type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}*/

//SalvarMGO Guardar
func (grad *Grado) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(BASEDATOS).C(colecion)
		err = c.Insert(grad)
	} else {
		c := sys.MGOSession.DB(BASEDATOS).C("grado")
		err = c.Insert(grad)
	}

	return
}

//Consultar una persona mediante el metodo de MongoDB
func (grad *Grado) Consultar(grado string) (jSon []byte, err error) {
	var msj Mensaje
	c := sys.MGOSession.DB(BASEDATOS).C(GRADO)
	err = c.Find(bson.M{"codigo": grado}).One(&grad)
	if err != nil {
		msj.Tipo = 0
		msj.Mensaje = err.Error()
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(grad)
	}
	return
}

//ConversionGrado Grados
func (m *Militar) ConversionGrado() {
	if m.Situacion == "RCP" {

	}
}
