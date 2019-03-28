package sssifanb

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

type Descuentos struct {
	ID          string    `json:"id,omitempty" bson:"id"`
	Tipo        string    `json:"tipo,omitempty" bson:"tipo"`
	Concepto    string    `json:"concepto,omitempty" bson:"concepto"`
	Formula     string    `json:"formula,omitempty" bson:"formula"`
	Observacion string    `json:"observacion,omitempty" bson:"observacion"`
	Estatus     int       `json:"estatus,omitempty" bson:"estatus"`
	FechaInicio time.Time `json:"fechainicio,omitempty" bson:"fechainicio"`
	FechaFin    time.Time `json:"fechafin,omitempty" bson:"fechafin"`
	Usuario     string    `json:"usuario,omitempty" bson:"usuario"`
}

//Agregar Sistema
func (DC *Descuentos) Agregar() (jSon []byte, err error) {
	var msj Mensaje

	descuento := make(map[string]interface{})

	descuento["pension.descuentos"] = DC
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": DC.ID}, bson.M{"$push": descuento})
	msj.Tipo = 0
	if err != nil {
		fmt.Println("Fallo insertar el descuento")
		msj.Tipo = 313
		jSon, err = json.Marshal(msj)
		return
	}
	InsertarPensionDescuento(DC)
	msj.Mensaje = "Proceso exitoso"
	msj.Tipo = 1
	jSon, err = json.Marshal(msj)

	return
}

//InsertarPensionDescuento Cargar medidas
func InsertarPensionDescuento(CD *Descuentos) {
	query := `
	INSERT INTO space.descuentos (
		tipo, conc, fnxc, obse,
		fini, ffin,
		usua, crea, esta, cedula
	) VALUES `
	query += `('` + CD.Tipo + `','` + CD.Concepto + `','` + CD.Formula + `','` + CD.Observacion + `'
						,'` + CD.FechaInicio.String()[:10] + `','` + CD.FechaFin.String()[:10] + `','` + CD.Usuario + `',Now()
						,'` + strconv.Itoa(CD.Estatus) + `','` + CD.ID + `')`

	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}

	//jSon, err = json.Marshal(msj)
}
