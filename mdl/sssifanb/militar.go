package sssifanb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
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
	ID                     string             `json:"id,omitempty" bson:"id"`
	TipoDato               int                `json:"tipodato,omitempty" bson:"tipodato"`
	Persona                Persona            `json:"Persona,omitempty" bson:"persona"`
	Categoria              string             `json:"categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion              string             `json:"situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                  string             `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente time.Time          `json:"fingreso,omitempty" bson:"fingreso"`
	FechaAscenso           time.Time          `json:"fascenso,omitempty" bson:"fascenso"`
	FechaRetiro            time.Time          `json:"fretiro,omitempty" bson:"fretiro"`
	AnoReconocido          int                `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido          int                `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido          int                `json:"dreconocido,omitempty" bson:"dreconocido"`
	NumeroResuelto         string             `json:"nresuelto,omitempty" bson:"nresuelto"`
	Posicion               int                `json:"posicion,omitempty" bson:"posicion"`
	DescripcionHistorica   string             `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Componente             Componente         `json:"Componente,omitempty" bson:"componente"`
	Grado                  Grado              `json:"Grado,omitempty" bson:"grado"` //grado
	TIM                    Carnet             `json:"Tim,omitempty" bson:"tim"`     //Tarjeta de Identificacion Militar
	Familiar               []Familiar         `json:"Familiar" bson:"familiar"`
	HistorialMilitar       []HistorialMilitar `json:"HistorialMilitar" bson:"historialmilitar"`
	AppSaman               bool               `json:"appsaman" bson:"appsaman"`
	AppPace                bool               `json:"apppace" bson:"apppace"`
	AppNomina              bool               `json:"appnomina" bson:"appnomina"`
	TiempoSevicio          string             `json:"tiemposervicio,omitempty" bson:"tiemposervicio,omitempty"`
	Pension                Pension            `json:"Pension,omitempty" bson:"pension"`
	Fideicomiso            Fideicomiso        `json:"Fideicomiso,omitempty" bson:"fideicomiso"`
	Anomalia               Anomalia           `json:"Anomalia,omitempty" bson:"anomalia"`
}

type Anomalia struct {
	Hijo bool `json:"hijo,omitempty" bson:"hijo"`
	Ano  bool `json:"ano,omitempty" bson:"ano"`
	Mes  bool `json:"mes,omitempty" bson:"mes"`
	Dia  bool `json:"dia,omitempty" bson:"dia"`
}

type HistorialMilitar struct {
	Componente     string    `json:"componente,omitempty" bson:"componente"`
	Grado          string    `json:"grado,omitempty" bson:"grado"`         //grado
	Categoria      string    `json:"categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion      string    `json:"situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase          string    `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaResuelto  time.Time `json:"fresuelto,omitempty" bson:"fresuelto"`
	GradoResuelto  string    `json:"gradoresuelto,omitempty" bson:"gradoresuelto"`
	NumeroResuelto string    `json:"numeroresuelto,omitempty" bson:"numeroresuelto"`
	FechaCambio    string    `json:"dreconocido,omitempty" bson:"dreconocido"`
	HoraCambio     string    `json:"nresuelto,omitempty" bson:"nresuelto"`
	FechaCreacion  string    `json:"posicion,omitempty" bson:"posicion"`
	HoraCreacion   string    `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Razon          string    `json:"razon,omitempty" bson:"razon"`           //codigo
	//TIM                    Carnet    `json:"Tim,omitempty" bson:"tim"`               //Tarjeta de Identificacion Militar
}

type Componente struct {
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
}

type Grado struct {
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
}

type Pension struct {
	GradoCodigo      string         `json:"grado" bson:"grado"`
	ComponenteCodigo string         `json:"Componente" bson:"componente"`
	NumeroHijos      int            `json:"numerohijos" bson:"numerohijos"`
	DatoFinanciero   DatoFinanciero `json:"DatoFinanciero" bson:"datofinanciero"`
}

type Fideicomiso struct {
	GradoCodigo      string `json:"grado" bson:"grado"`
	ComponenteCodigo string `json:"Componente" bson:"componente"`
	NumeroHijos      int    `json:"numerohijos" bson:"numerohijos"`
	AnoReconocido    int    `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido    int    `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido    int    `json:"dreconocido,omitempty" bson:"dreconocido"`
	CuentaBancaria   string `json:"cuentabancaria" bson:"cuentabancaria"`
}

//Listar sistemas
func (m *Militar) Listar() {
	//gesaodin@gmail.com
}

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}

//AplicarReglas Reglas Generales
func (m *Militar) AplicarReglas() {
	m.Conversion()
	m.ConversionGrado()
	a, mes, d := util.CalcularTiempo(m.FechaIngresoComponente)

	m.TiempoSevicio = strconv.Itoa(a) + " A, " + mes.String() + " M, " + strconv.Itoa(d) + " D"
}

//Conversion de Grados
func (m *Militar) Conversion() {

}

//ConversionGrado Grados
func (m *Militar) ConversionGrado() {
	if m.Situacion == "RCP" {

	}
}

//Consultar una persona mediante el metodo de MongoDB
func (m *Militar) Consultar() (jSon []byte, err error) {
	var militar Militar
	var msj Mensaje
	c := sys.MGOSession.DB("ipsfa_test").C("militar")
	err = c.Find(bson.M{"id": m.Persona.DatoBasico.Cedula}).One(&militar)
	if militar.Persona.DatoBasico.Cedula == "" {
		msj.Tipo = 0
		jSon, err = json.Marshal(msj)
	} else {
		militar.Persona.DatoBasico.FechaNacimiento = militar.Persona.DatoBasico.FechaNacimiento.UTC()
		militar.FechaIngresoComponente = militar.FechaIngresoComponente.UTC()
		militar.FechaAscenso = militar.FechaAscenso.UTC()
		militar.AplicarReglas()
		jSon, err = json.Marshal(militar)
	}
	return
}

//ConsultarSAMAN Militar
func (m *Militar) ConsultarSAMAN() (jSon []byte, err error) {
	var msj Mensaje
	var lst []Militar
	var estatus bool
	s := `SELECT codnip,tipnip, nropersona,nombreprimero, nombresegundo,apellidoprimero,apellidosegundo,sexocod
	FROM personas
	WHERE codnip='` + m.Persona.DatoBasico.Cedula + `' AND tipnip != 'P'`
	sq, err := sys.PostgreSQLSAMAN.Query(s)
	if err != nil {
		msj.Mensaje = "Error: Consulta ya existe."
		msj.Tipo = 2
		msj.Pgsql = err.Error()
		jSon, err = json.Marshal(msj)
		fmt.Println(err.Error())
		return
	}
	estatus = true
	for sq.Next() {
		var m Militar
		var cedula, tipnip string
		var nombp, nombs, apellp, apells, sexo sql.NullString
		var numero int

		sq.Scan(&cedula, &tipnip, &numero, &nombp, &nombs, &apellp, &apells, &sexo)
		m.Persona.DatoBasico.Cedula = cedula
		m.Persona.DatoBasico.NumeroPersona = numero
		m.Persona.DatoBasico.NombrePrimero = util.ValidarNullString(nombp)
		m.Persona.DatoBasico.NombreSegundo = util.ValidarNullString(nombs)
		m.Persona.DatoBasico.ApellidoPrimero = util.ValidarNullString(apellp)
		m.Persona.DatoBasico.ApellidoSegundo = util.ValidarNullString(apells)
		m.Persona.DatoBasico.Nacionalidad = tipnip
		m.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		if m.Persona.DatoBasico.NombrePrimero != "null" {
			estatus = false
		} else {
			estatus = true
		}

		lst = append(lst, m)

	}
	if estatus == true {
		msj.Mensaje = "Afiliado no existe."
		msj.Tipo = 0
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(lst)
	}

	return

}

//Actualizar Vida Militar
func (m *Militar) Actualizar() (jSon []byte, err error) {
	var msj Mensaje
	m.TipoDato = 0

	s := `UPDATE personas SET nombreprimero='` +
		m.Persona.DatoBasico.NombrePrimero +
		`', nombresegundo='` +
		m.Persona.DatoBasico.NombreSegundo +
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
func (m *Militar) ActualizarMGO(oid string, familiar map[string]interface{}) (err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("militar")
	err = c.Update(bson.M{"id": oid}, bson.M{"$set": familiar})
	if err != nil {
		fmt.Println("Cedula: " + oid + " -> " + err.Error())
		return
	}
	return
}

//ActualizarMGO Actualizar
func (m *Militar) ActualizarMGOObjeto(oid string, Obj interface{}) (err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("militar")
	err = c.Update(bson.M{"id": oid}, bson.M{"$set": Obj})
	if err != nil {
		fmt.Println("Cedula: " + oid + " -> " + err.Error())
		return
	}
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

	//fmt.Println(err)

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
