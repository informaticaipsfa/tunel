package fanb

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

type Componente struct {
	Codigo string `json:"codigo" bson:"codigo"`
	Nombre string `json:"nombre" bson:"nombre"`
	Siglas string `json:"siglas" bson:"siglas"`
	//Cpace  int     `json:"cpace" bson:"cpace"`
	Grado []Grado `json:"Grado" bson:"Grado"`
}

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}

type Conversion struct {
	Componente string `json:"componente"`
	GradoPace  string `json:"grado"`
	GradoSaman string `json:"gradosaman"`
}

//SalvarMGO Guardar
func (comp *Componente) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(sys.CBASE).C(colecion)
		err = c.Insert(comp)
	} else {
		c := sys.MGOSession.DB(sys.CBASE).C("componente")
		err = c.Insert(comp)
	}

	return
}

//Consultar una Componente mediante el metodo de MongoDB
func (comp *Componente) Consultar(componente string) (jSon []byte, err error) {
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCOMPONENTE)
	err = c.Find(bson.M{"codigo": componente}).One(&comp)
	if err != nil {
		msj.Tipo = 0
		msj.Mensaje = err.Error()
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(comp)
	}
	return
}

//Consultar una Componente mediante el metodo de MongoDB
func (comp *Componente) ConsultarGrado(componente string, grado string) (ComponenteConver Conversion) {
	// var ComponenteConver Conversion
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCOMPONENTE)
	//fmt.Println("codigo: ", componente, "Grado.codigo: ", grado)
	err := c.Find(bson.M{"codigo": componente}).One(&comp)
	ComponenteConver.Componente = ComponenteID(comp.Codigo)
	if err != nil {
		fmt.Println("Err: fanb.componente.ConsultarGrado() ")
	} else {
		for _, v := range comp.Grado {
			if v.Codigo == grado {
				//fmt.Println(v.Codigo, "   ", grado, " --> ", v.Cpace)
				ComponenteConver.GradoPace = strconv.Itoa(v.Cpace)
			}
		}
	}
	return
}

func ComponenteID(abreviatura string) (codigo string) {
	switch abreviatura {
	case "EJ":
		codigo = "1"
		break
	case "AR":
		codigo = "2"
		break
	case "AV":
		codigo = "3"
		break
	case "GN":
		codigo = "4"
		break
	}
	return
}

func (P *Componente) ComponenteCodigo(abreviatura string) (codigo string) {
	switch abreviatura {
	case "EJB":
		codigo = "EJ"
		break
	case "ARB":
		codigo = "AR"
		break
	case "AVB":
		codigo = "AV"
		break
	case "GNB":
		codigo = "GN"
		break
	}
	return
}

func (comp *Componente) ObtenerGradoID(codigo string, grado string) string {
	var compo Componente
	buscar := bson.M{"codigo": codigo, "Grado": bson.M{"$elemMatch": bson.M{"codigo": grado}}}
	valor := bson.M{"Grado": bson.M{"$elemMatch": bson.M{"codigo": grado}}}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCOMPONENTE)
	err := c.Find(buscar).Select(valor).One(&compo)
	if err != nil {
		return "0"
	}
	return compo.Grado[0].Rango
}
