package sssifanb

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
)

//Carnet Tarjeta de Identificacion Militar
type Carnet struct {
	ID                      string     `json:"id,omitempty" bson:"id"`
	Tipo                    int        `json:"tipo,omitempty" bson:"tipo"` // 0: Militar 1: Empleado 2: Familiares
	Nombre                  string     `json:"nombre,omitempty" bson:"nombre"`
	Apellido                string     `json:"apellido,omitempty" bson:"apellido"`
	Condicion               int        `json:"condicion,omitempty" bson:"condicion"`
	Serial                  string     `json:"serial,omitempty" bson:"serial"`
	FechaCreacion           time.Time  `json:"fechacreacion,omitempty" bson:"fechacreacion"`
	FechaVencimiento        time.Time  `json:"fechavencimiento,omitempty" bson:"fechavencimiento"`
	Responsable             string     `json:"responsable,omitempty" bson:"responsable"`
	Componente              Componente `json:"Componente,omitempty" bson:"componente"`
	Grado                   Grado      `json:"Grado,omitempty" bson:"grado"`
	URLSimbolo              string     `json:"simbolo,omitempty" bson:"simbolo"`
	URLFirmaMinistro        string     `json:"fministro,omitempty" bson:"fministro"`
	URLFirmaPresidenteIPSFA string     `json:"fpresidente,omitempty" bson:"fpresidente"`
	Estatus                 int        `json:"estatus,omitempty" bson:"estatus"`
	IP                      string     `json:"ip" bson:"ip"`
}

//AplicarReglas Basicas
func (c *Carnet) AplicarReglas() {
	//Generar serial
	//Generar CodigoComponente

}

//GenerarSerial Seriales de Carnet
func (c *Carnet) GenerarSerial() string {
	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo()
	return util.CompletarCeros(strconv.Itoa(i), 0, 8)
}

//Salvar Guardar
func (tim *Carnet) Salvar() (err error) {
	var militar Militar
	militar.ConsultarMGO(tim.ID)
	militar.TIM, _ = militar.GenerarCarnet()
	militar.TIM.IP = tim.IP
	c := sys.MGOSession.DB(CBASE).C(CTIM)
	err = c.Insert(militar.TIM)
	return
}

//Consultar Carnets
func (tim *Carnet) Consultar(id string) (err error) {

	return
}

//Listar Carnet Propios
func (tim *Carnet) Listar(estatus int) (jSon []byte, err error) {
	var lst []Carnet
	c := sys.MGOSession.DB(CBASE).C("tim")
	err = c.Find(bson.M{"estatus": estatus}).All(&lst)

	if err != nil {
		fmt.Println("No se encontraron registros")
		return
	}
	jSon, err = json.Marshal(lst)
	return
}
