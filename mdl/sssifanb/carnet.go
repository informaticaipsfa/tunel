package sssifanb

import (
	"strconv"
	"time"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/util"
)

//Carnet Tarjeta de Identificacion Militar
type Carnet struct {
	ID                      int        `json:"id,omitempty" bson:"id"`
	Tipo                    int        `json:"tipo,omitempty" bson:"tipo"` // 0: Militar 1: Empleado 2: Familiares
	Condicion               int        `json:"condicion,omitempty" bson:"condicion"`
	Serial                  string     `json:"serial,omitempty" bson:"serial"`
	CodigoComponente        string     `json:"codigocomponente,omitempty" bson:"codigocomponente"`
	FechaCreacion           time.Time  `json:"fechacreacion,omitempty" bson:"fechacreacion"`
	FechaVencimiento        time.Time  `json:"fechavencimiento,omitempty" bson:"fechavencimiento"`
	Responsable             string     `json:"responsable,omitempty" bson:"responsable"`
	Componente              Componente `json:"Componente,omitempty" bson:"componente"`
	Grado                   Grado      `json:"Grado,omitempty" bson:"grado"`
	URLSimbolo              string     `json:"simbolo,omitempty" bson:"simbolo"`
	URLFirmaMinistro        string     `json:"fministro,omitempty" bson:"fministro"`
	URLFirmaPresidenteIPSFA string     `json:"fpresidente,omitempty" bson:"fpresidente"`
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
