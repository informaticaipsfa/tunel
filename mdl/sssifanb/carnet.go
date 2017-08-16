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
	ID                      string     `json:"id" bson:"id"`
	IDF                     string     `json:"idf" bson:"idf"`
	Tipo                    int        `json:"tipo,omitempty" bson:"tipo"` // 0: Militar 1: Empleado 2: Familiares
	Nombre                  string     `json:"nombre,omitempty" bson:"nombre"`
	Apellido                string     `json:"apellido,omitempty" bson:"apellido"`
	Condicion               bool       `json:"condicion,omitempty" bson:"condicion"`
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
	Motivo                  string     `json:"motivo" bson:"motivo"`
	Usuario                 string     `json:"usuario" bson:"usuario"`
}

//AplicarReglas Basicas
func (c *Carnet) AplicarReglas() {
	//Generar serial
	//Generar CodigoComponente

}

//GenerarSerial Seriales de Carnet
func (c *Carnet) GenerarSerial() string {
	var Semillero fanb.Semillero
	i, _ := Semillero.Maximo("semillero")
	return util.CompletarCeros(strconv.Itoa(i), 0, 8)
}

//Salvar Guardar
func (tim *Carnet) Salvar() (err error) {
	var militar Militar
	fmt.Println("Salvar " + tim.Usuario)
	militar.ConsultarMGO(tim.ID)
	militar.TIM.ID = tim.ID
	militar.TIM.IDF = tim.IDF
	militar.TIM.IP = tim.IP
	militar.TIM.Motivo = tim.Motivo
	militar.TIM.Usuario = tim.Usuario

	if tim.ID == tim.IDF { // Carnet Titulares
		militar.TIM, _ = militar.GenerarCarnet()
		c := sys.MGOSession.DB(sys.CBASE).C(sys.CTIM)
		err = c.Insert(militar.TIM)
	} else { //Carnet de Familiares
		var TIMS Carnet
		var Parenstesco string
		for _, v := range militar.Familiar {
			if v.Persona.DatoBasico.Cedula == tim.IDF {
				Parenstesco = v.Parentesco

				switch v.Parentesco {
				case "PD":
					TIMS = v.AplicarReglasCarnetPadres()
					fmt.Println("Entrando, Padre...")
				case "HJ":
					TIMS = v.AplicarReglasCarnetHijos()
					fmt.Println("Entrando, Hijos...")
				case "EA":
					TIMS = v.AplicarReglasCarnetEsposa()
					fmt.Println("Entrando, Esposa...")
				case "VI":
					TIMS = v.AplicarReglasCarnetEsposa()
					fmt.Println("Entrando, Esposa...")
				case "HO":
					TIMS = v.AplicarReglasCarnetHermanos()
					fmt.Println("Entrando, Hermano...")
				}
			}
		}

		TIMS.Motivo = tim.Motivo
		TIMS.IP = tim.IP
		TIMS.ID = tim.ID
		TIMS.IDF = tim.IDF
		TIMS.Usuario = tim.Usuario
		TIMS.Componente.Abreviatura = militar.Componente.Abreviatura
		TIMS.Componente.Descripcion = militar.Componente.Descripcion
		TIMS.Grado.Abreviatura = militar.Grado.Abreviatura
		TIMS.Grado.Descripcion = militar.Grado.Descripcion
		TIMS.Grado.Nombre = Parenstesco
		TIMS.Serial = tim.Usuario + TIMS.Serial
		c := sys.MGOSession.DB(sys.CBASE).C(sys.CTIM)
		err = c.Insert(TIMS)
	}

	return
}

// CambiarEstado Seleccionar estados
func (tim *Carnet) CambiarEstado(serial string, estatus int) (err error) {
	carnet := make(map[string]interface{})
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTIM)

	carnet["estatus"] = estatus
	fmt.Println(serial, " ", estatus)
	err = c.Update(bson.M{"serial": serial}, bson.M{"$set": carnet})
	if estatus == 3 {
		err = tim.CambiarEstadoMilitar(serial)
	}
	return
}

//Consultar Carnets
func (tim *Carnet) CambiarEstadoMilitar(serial string) (err error) {
	var TIM Carnet
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Find(bson.M{"serial": serial}).One(&TIM)
	if err != nil {
		return
	}
	if TIM.ID != "" && TIM.IDF == "" {
		carnet := make(map[string]interface{})
		carnet["estatuscarnet"] = 0
		err = c.Update(bson.M{"id": TIM.ID}, bson.M{"$set": carnet})
	}
	return
}

//Listar Carnet Propios
func (tim *Carnet) Listar(estatus int, usuario string) (jSon []byte, err error) {
	var lst []Carnet
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTIM)
	consulta := usuario
	err = c.Find(bson.M{"estatus": estatus, "usuario": bson.M{"$regex": consulta}}).All(&lst)

	if err != nil {
		fmt.Println("No se encontraron registros")
		return
	}
	jSon, err = json.Marshal(lst)
	return
}
