package fanb

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

//Concepto : Control de variables
type Concepto struct {
	Codigo      string    `json:"codigo" bson:"codigo"`
	Descripcion string    `json:"descripcion" bson:"descripcion"`
	Formula     string    `json:"formula" bson:"formula"`
	Tipo        int       `json:"tipo" bson:"tipo"`
	Partida     string    `json:"partida" bson:"partida"`
	Cuenta      string    `json:"cuenta" bson:"cuenta"`
	Estatus     int       `json:"estatus" bson:"estatus"`
	Componente  string    `json:"componente,omitempty" bson:"componente"`
	Grado       string    `json:"grado,omitempty" bson:"grado"`
	Usuario     string    `json:"usuari,omitempty" bson:"usuario"`
	Creado      time.Time `json:"creado,omitempty" bson:"creado"`
}

//Agregar Crear un Concepto
func (Cp *Concepto) Agregar() (jSon []byte, err error) {

	var msj Mensaje
	Cp.Creado = time.Now()
	query := `INSERT INTO space.conceptos (codigo, descripcion, forumula, partida, cuenta, tipo, estatus,
  componente, grado, usuario, creado)	VALUES (
    '` + Cp.Codigo + `','` + Cp.Descripcion + `',
    '` + Cp.Formula + `','` + Cp.Partida + `',
		'` + Cp.Cuenta + `',` + strconv.Itoa(Cp.Tipo) + `,
    '` + strconv.Itoa(Cp.Estatus) + `','` + Cp.Componente + `',
    '` + Cp.Grado + `','` + Cp.Usuario + `', Now())`

	_, err = sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
		Cp.Actualizar()
	} else {
		c := sys.MGOSession.DB(sys.CBASE).C(sys.CCONCEPTO)

		err = c.Insert(Cp)
		msj.Tipo = 0
		if err != nil {
			fmt.Println("Fallo insertar concepto ")
			msj.Tipo = 313
			jSon, err = json.Marshal(msj)
			return
		}
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
	concepto["cuenta"] = Cp.Cuenta
	concepto["partida"] = Cp.Partida
	concepto["estatus"] = Cp.Estatus
	concepto["usuario"] = Cp.Usuario

	err = c.Update(seleccion, bson.M{"$set": concepto})
	if err != nil {
		fmt.Println(err.Error())
	}

	query := `UPDATE space.conceptos SET
		descripcion='` + Cp.Descripcion + `',
		forumula='` + Cp.Formula + `',
		partida ='` + Cp.Partida + `',
		cuenta='` + Cp.Cuenta + `',
		tipo=` + strconv.Itoa(Cp.Tipo) + `,
		estatus='` + strconv.Itoa(Cp.Estatus) + `',
		usuario='` + Cp.Usuario + `',
		creado=Now() WHERE codigo='` + Cp.Codigo + `'`

	_, err = sys.PostgreSQLPENSION.Exec(query)
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
