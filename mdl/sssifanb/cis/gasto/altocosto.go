package gasto

import (
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
	"gopkg.in/mgo.v2/bson"
)

type AltoCosto struct {
	NombreComercial  string    `json:"nombrecomercial" bson:"nombrecomercial"`
	Presentacion     string    `json:"presentacion" bson:"presentacion"`
	Dosis            string    `json:"dosis" bson:"dosis"`
	Cantidad         string    `json:"cantidad" bson:"cantidad"`
	FechaInicio      time.Time `json:"fechainicio" bson:"fechainicio"`
	FechaVencimiento time.Time `json:"fechavencimiento" bson:"fechavencimiento"`
}

type WAltoCosto struct {
	ID       string      `json:"id" bson:"id"`
	IDF      string      `json:"idf" bson:"idf"`
	Medicina []AltoCosto `json:"Medicina" bson:"medicina"`
	Afiliado string      `json:"afiliado" bson:"afiliado"`
	Usuario  string      `json:"usuario" bson:"usuario"`
	Fecha    time.Time   `json:"fecha" bson:"fecha"`
}

//Crear Registrando
func (ac *WAltoCosto) Crear() (jSon []byte, err error) {
	var M fanb.Mensaje
	M.Mensaje = "Creando Medicina Alto Costo"
	M.Tipo = 1
	altoc := make(map[string]interface{})
	altoc["cis.gasto.medicinaaltocosto"] = ac
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)

	// fmt.Println(altocosto.Medicina.Afiliado)
	err = c.Update(bson.M{"id": ac.ID}, bson.M{"$push": altoc})
	util.Error(err)
	return
}
