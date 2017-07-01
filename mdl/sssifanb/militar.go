package sssifanb

import (
	"encoding/json"
	"time"

	"github.com/gesaodin/bdse/sys"
	"gopkg.in/mgo.v2/bson"
)

const (
	MILITAR int = 0
)

type Militar struct {
	//Persona                DatoBasico  `json:"Persona,omitempty" bson:"persona"`
	// Direccion              []Direccion `json:"Direccion,omitempty" bson:"direccion"`
	// Telefono               []Telefono  `json:"Telefono,omitempty" bson:"telefono"`a
	// Correo                 Correo      `json:"Correo,omitempty" bson:"correo"`
	ID                     int        `json:"id,omitempty" bson:"id"`
	TipoDato               int        `json:"tipodato,omitempty" bson:"tipodato"`
	Persona                Persona    `json:"Persona,omitempty" bson:"persona"`
	Categoria              int        `json:"Categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion              string     `json:"Situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                  int        `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente time.Time  `json:"fingreso,omitempty" bson:"fing"`
	FechaAscenso           time.Time  `json:"fascenso,omitempty" bson:"fascenso"`
	AnoReconocido          string     `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido          string     `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido          string     `json:"dreconocido,omitempty" bson:"dreconocido"`
	NumeroResuelto         string     `json:"nresuelto,omitempty" bson:"nresuelto"`
	Posicion               int        `json:"posicion,omitempty" bson:"posicion"`
	DescripcionHistorica   string     `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Componente             Componente `json:"Componente,omitempty" bson:"componente"`
	Grado                  Grado      `json:"Grado,omitempty" bson:"grado"` //grado
	TIM                    Carnet     `json:"Tim,omitempty" bson:"tim"`     //Tarjeta de Identificacion Militar
	Familiar               []Familiar `json:"Familiar,omitempty" bson:"familiar"`
}

type Componente struct {
	ID          int
	Nombre      string
	Descripcion string
	Abreviatura string
}

type Grado struct {
	ID          int
	Nombre      string
	Descripcion string
	Abreviatura string
}

//
func (m *Militar) Listar() {
	//gesaodin@gmail.com
}

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj,omitempty"`
	Tipo    int    `json:"tipo,omitempty"`
	Pgsql   string `json:"pgsql,omitempty"`
}

//Consultar Militar
func (m *Militar) Consultar() (jSon []byte, err error) {
	var msj Mensaje
	var lst []Militar
	s := `SELECT codnip,nropersona,nombreprimero FROM personas WHERE codnip='` + m.Persona.DatoBasico.Cedula + `'`
	sq, err := sys.PostgreSQLSAMAN.Query(s)
	if err != nil {
		msj.Mensaje = "Error: Consulta ya existe."
		msj.Tipo = 2
		msj.Pgsql = err.Error()
		jSon, err = json.Marshal(m)
		//fmt.Println(err.Error())
		return
	}
	for sq.Next() {
		var m Militar
		var cedula, nombre string
		var numero int
		sq.Scan(&cedula, &numero, &nombre)
		m.Persona.DatoBasico.Cedula = cedula
		m.Persona.DatoBasico.NumeroPersona = numero
		m.Persona.DatoBasico.NombrePrimero = nombre
		lst = append(lst, m)
	}
	jSon, err = json.Marshal(lst)
	return

}

//Actualizar Vida Militar
func (m *Militar) Actualizar() (jSon []byte, err error) {
	var msj Mensaje
	m.TipoDato = 0

	s := `UPDATE personas SET nombreprimero='` +
		m.Persona.DatoBasico.NombrePrimero +
		`' WHERE codnip='` + m.Persona.DatoBasico.Cedula + `'`
	_, err = sys.PostgreSQLSAMAN.Exec(s)
	if err != nil {
		msj.Mensaje = "Error: Consulta ya existe."
		msj.Tipo = 2
		msj.Pgsql = err.Error()
		jSon, err = json.Marshal(msj)
		return
	}
	msj.Mensaje = "Su data ha sido actualizada."
	msj.Tipo = 2
	jSon, err = json.Marshal(msj)
	m.SalvarMGO("")
	return
}

//ActualizarMGO Actualizar
func (m *Militar) ActualizarMGO(persona map[string]interface{}) (err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("persona")
	err = c.Update(bson.M{"cedula": persona["cedula"]}, bson.M{"$set": persona})

	return
}

//SalvarMGO Guardar
func (m *Militar) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB("ipsfa_test").C(colecion)
		err = c.Insert(m)
	} else {
		c := sys.MGOSession.DB("ipsfa_test").C("persona")
		err = c.Insert(m)
	}

	//fmt.Println(p)

	return
}

//ConsultarMGO una persona mediante el metodo de MongoDB
func (m *Militar) ConsultarMGO(cedula string) (err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("persona")
	err = c.Find(bson.M{"cedula": cedula}).One(&m)
	return
}

//ListarMGO Listado General
func (m *Militar) ListarMGO(cedula string) (lst []Militar, err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("persona")
	err = c.Find(bson.M{}).All(&lst)
	return
}
