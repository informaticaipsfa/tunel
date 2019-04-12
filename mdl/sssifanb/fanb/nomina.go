package fanb

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

//Concepto : Control de variables
type Concepto struct {
	Codigo      string `json:"codigo,omitempty" bson:"codigo"`
	Descripcion string `json:"descripcion,omitempty" bson:"descripcion"`
	Formula     string `json:"formula,omitempty" bson:"formula"`
	Tipo        int    `json:"tipo,omitempty" bson:"tipo"`
	Partida     string `json:"partida,omitempty" bson:"partida"`
	Estatus     int    `json:"estatus,omitempty" bson:"estatus"`
	Componente  string `json:"componente,omitempty" bson:"componente"`
	Grado       string `json:"grado,omitempty" bson:"grado"`
}

//Agregar Crear un Concepto
func (Cp *Concepto) Agregar(user string) (jSon []byte, err error) {
	fmt.Println(user)
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCONCEPTO)
	err = c.Insert(Cp)
	msj.Tipo = 0
	if err != nil {
		fmt.Println("Fallo insertar concepto ")
		msj.Tipo = 313
		jSon, err = json.Marshal(msj)
		return
	}

	query := `INSERT INTO space.conceptos (codigo, descripcion, forumula, partida, tipo, estatus,
  componente, grado, usuario, creado)	VALUES (
    '` + Cp.Codigo + `',
    '` + Cp.Descripcion + `',
    '` + Cp.Formula + `',
    '` + Cp.Partida + `',
    ` + strconv.Itoa(Cp.Tipo) + `,
    '` + strconv.Itoa(Cp.Estatus) + `',
    '` + Cp.Componente + `',
    '` + Cp.Grado + `',
    '` + user + `', Now())`
	_, err = sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}
	msj.Mensaje = "Proceso exitoso"
	msj.Tipo = 1
	jSon, err = json.Marshal(msj)
	return

}

//Consultar Conceptos desde mongodb
func (Cp *Concepto) Consultar() (jSon []byte, err error) {

	var concepto Concepto

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCONCEPTO)
	seleccion := bson.M{"codigo": concepto.Codigo}
	err = c.Find(seleccion).One(&concepto)
	if err != nil {
		fmt.Println(err.Error())
	}
	jSon, err = json.Marshal(concepto)
	return
}

//Actualizar Conceptos desde mongodb
func (Cp *Concepto) Actualizar() (jSon []byte, err error) {

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCONCEPTO)
	seleccion := bson.M{"codigo": Cp.Codigo}
	concepto := make(map[string]interface{})
	concepto["descripcion"] = Cp.Descripcion
	concepto["formula"] = Cp.Formula
	concepto["tipo"] = Cp.Tipo
	concepto["partido"] = Cp.Partida
	concepto["estatus"] = Cp.Estatus

	err = c.Update(seleccion, bson.M{"$set": concepto})
	if err != nil {
		fmt.Println(err.Error())
	}
	jSon, err = json.Marshal(Cp)
	return
}

//Listar Conceptos desde mongodb
func (Cp *Concepto) Listar() (jSon []byte, err error) {

	var lst []Concepto
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CCONCEPTO)
	seleccion := bson.M{"estatus": 1}
	err = c.Find(seleccion).All(&lst)
	if err != nil {
		fmt.Println(err.Error())
	}
	jSon, err = json.Marshal(lst)
	return
}
