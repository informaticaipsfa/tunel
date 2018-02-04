package estadistica

import (
	"fmt"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/sys"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Reduccion de datos de los familiares
type Reduccion struct {
	Cedula          string    `json:"cedula", bson:"cedula"`
	Nombre          string    `json:"nombre", bson:"nombre"`
	Sexo            string    `json:"sexo", bson:"sexo"`
	Tipo            string    `json:"tipo",  bson:"tipo"` //T Titular Militar | F Familiar
	EsMilitar       bool      `json:"esmilitar",  bson:"esmilitar"`
	FechaNacimiento time.Time `json:"fecha",  bson:"fecha"`
}

func Inferencia() {

}

func Descriptiva() {

}

//CrearColeccion Crear Coleccion de Mongo para la Reduccion
func CrearColeccion() {
	var prs Reduccion
	prs.Cedula = "0"
	prs.Nombre = "X"
	prs.Tipo = "X"

	c := sys.MGOSession.DB(sys.CBASE).C("reduccion")
	err := c.Insert(prs)
	if err != nil {
		panic(err)
	}

	index := mgo.Index{
		Key:        []string{"cedula"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

//MilitarTitular Familiares y Titulares Estadisticas
func (r *Reduccion) MilitarTitular() (jSon []byte, err error) {
	var militar []sssifanb.Militar
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Find(bson.M{}).All(&militar)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, mil := range militar {
		fmt.Println(mil.Persona.DatoBasico.Cedula)
		for _, Familia := range mil.Familiar {
			fmt.Println(Familia.Persona.DatoBasico.Cedula)
		}
	}
	return
}
