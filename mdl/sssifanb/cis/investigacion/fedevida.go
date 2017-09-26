package investigacion

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
)

const (
	RETIRO    int32 = 1
	INVALIDEZ int32 = 2
	GRACIA    int32 = 3
)

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
}

type WFedeVida struct {
	ID          string
	IDF         string
	Direccion   Direccion
	DireccionEx string
	FechaEx     time.Time
	Nombre      string
}

//FeDeVida Control de Fe de Vida
type FeDeVida struct {
	Numero        string       `json:"numero" bson:"numero"`
	FechaCreacion time.Time    `json:"fechacreacion" bson:"fechacreacion"`
	DatoBasico    DatoPersonal `json:"DatoBasico" bson:"datobasico"`
	TipoPension   int32        `json:"tipo" bson:"tipo"` //1 retiro, 2 invalidez, 3 gracia
	Estatus       bool         `json:"estatus" bson:"estatus"`
	IDF           string       `json:"idf" bson:"idf"`
	DireccionEx   string       `json:"direccionex" bson:"direccionex"`
	FechaEx       time.Time    `json:"fechaex" bson:"fechaex"`
	PaisEx        string       `json:"paisex" bson:"paisex"`
}

//Crear Creacion de Cuenta
func (fe *WFedeVida) Crear() (jSon []byte, err error) {
	var fevida FeDeVida
	var M Mensaje

	var semillero fanb.Semillero

	i, _ := semillero.Maximo("semillerocis")
	fevida.Numero = util.CompletarCeros(strconv.Itoa(i), 0, 8)
	fevida.DatoBasico.Direccion = fe.Direccion
	fevida.DatoBasico.Cedula = fe.ID
	fevida.DatoBasico.NombreCompleto = fe.Nombre
	fevida.IDF = fe.IDF
	fevida.DireccionEx = fe.DireccionEx
	fevida.FechaEx = fe.FechaEx
	fevida.FechaCreacion = time.Now()
	fevida.Estatus = true
	fevida.TipoPension = 0

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CFEVIDA)
	err = c.Insert(fe)
	if err != nil {
		fmt.Println("Error creando reembolso det: ")
		// return
	}

	cisfe := make(map[string]interface{})

	cisfe["cis.investigacion.fedevida"] = fevida
	co := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = co.Update(bson.M{"id": fe.ID}, bson.M{"$push": cisfe})
	if err != nil {
		fmt.Println("Cedula: " + fe.ID + " -> " + err.Error())
		// return
	}

	M.Mensaje = fevida.Numero
	M.Tipo = 1

	jSon, err = json.Marshal(M)
	return
}
