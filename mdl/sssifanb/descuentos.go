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
	Familiar    string    `json:"familiar,omitempty" bson:"familiar"`
	Modifica    string    `json:"modifica,omitempty" bson:"modifica"`
}

//Agregar Sistema
func (DC *Descuentos) Agregar() (jSon []byte, err error) {
	var msj Mensaje
	strupdate := "$push"
	descuento := make(map[string]interface{})
	if DC.Modifica == "X" {
		descuento["pension.descuentos"] = DC
		InsertarPensionDescuento(DC)
	} else {
		descuento["pension.descuentos."+DC.Modifica] = DC
		strupdate = "$set"
	}

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": DC.ID}, bson.M{strupdate: descuento})
	msj.Tipo = 0
	if err != nil {
		fmt.Println("Fallo insertar el descuento")
		msj.Tipo = 313
		jSon, err = json.Marshal(msj)
		return
	}

	msj.Mensaje = "Proceso exitoso"
	msj.Tipo = 1
	jSon, err = json.Marshal(msj)

	return
}

//InsertarPensionDescuento Cargar medidas
func InsertarPensionDescuento(CD *Descuentos) {
	query := `
	INSERT INTO space.asig_deduc (
		tipo, conc, fnxc, obse,
		fini, ffin,
		usua, crea, esta, cedula, familiar
	) VALUES `
	query += `('` + CD.Tipo + `','` + CD.Concepto + `','` + CD.Formula + `','` + CD.Observacion + `'
						,'` + CD.FechaInicio.String()[:10] + `','` + CD.FechaFin.String()[:10] + `','` + CD.Usuario + `',Now()
						,'` + strconv.Itoa(CD.Estatus) + `','` + CD.ID + `','` + CD.Familiar + `')`

	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}

	//jSon, err = json.Marshal(msj)
}

//ActualizarPensionDescuento Cargar medidas
func ActualizarPensionDescuento(CD *Descuentos) {
	query := `
	UPDATE space.asig_deduc SET
		tipo='` + CD.Tipo + `',
		conc='` + CD.Concepto + `',
		fnxc='` + CD.Formula + `',
		obse='` + CD.Observacion + `',
		fini='` + CD.FechaInicio.String()[:10] + `',
		ffin='` + CD.FechaFin.String()[:10] + `',
		usua='` + CD.Usuario + `',
		crea=Now(),
		esta='` + strconv.Itoa(CD.Estatus) + `',
		cedula='` + CD.ID + `',
		familiar='` + CD.Familiar + `'
	WHERE cedula='` + CD.ID + `' AND familiar='` + CD.Familiar + `'`

	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}

	//jSon, err = json.Marshal(msj)
}
