package fanb

import (
	"encoding/json"
	"fmt"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

type Grado struct {
	Codigo      string `json:"codigo" bson:"codigo"`
	Rango       string `json:"rango" bson:"rango"`
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Cpace       string `json:"cpace" bson:"cpace"`
	nompace     string `json:"nompace" bson:"nompace"`
}

//Mensaje del sistema
/*type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}*/

// SalvarMGO Guardar
func (grad *Grado) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(sys.CBASE).C(colecion)
		err = c.Insert(grad)
	} else {
		c := sys.MGOSession.DB(sys.CBASE).C("grado")
		err = c.Insert(grad)
	}
	return
}

// Consultar una persona mediante el metodo de MongoDB
func (grad *Grado) Consultar(grado string) (jSon []byte, err error) {
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C("grado")
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
func (g *Grado) ConversionGrado() {

	fmt.Println(obtenerGradoFideicomiso())

}

func obtenerGradoFideicomiso() string {
	return `
		SELECT c.componentecod, componentenombre, componentesiglas, gradocod,gradocodrangoid,gradonombrecorto,
		gradonombrelargo
		FROM ipsfa_grados AS g JOIN ipsfa_componentes AS c ON g.componentecod=c.componentecod
		ORDER BY c.componentepriorpt,g.gradocodrangoid
	`
}
