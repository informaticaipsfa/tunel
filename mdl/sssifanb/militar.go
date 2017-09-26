package sssifanb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

<<<<<<< HEAD
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
=======
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
	"gopkg.in/mgo.v2/bson"
)

const (
	MILITAR      int = 0
	FAMILIAR     int = 1
	PROVEEDOR    int = 2
	BENEFICIARIO int = 3
)

//Militar militares
type Militar struct {
	ID                     string              `json:"id,omitempty" bson:"id"`
	TipoDato               int                 `json:"tipodato,omitempty" bson:"tipodato"`
	Persona                Persona             `json:"Persona,omitempty" bson:"persona"`
	Categoria              string              `json:"categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion              string              `json:"situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                  string              `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente time.Time           `json:"fingreso,omitempty" bson:"fingreso"`
	FechaAscenso           time.Time           `json:"fascenso,omitempty" bson:"fascenso"`
	FechaRetiro            time.Time           `json:"fretiro,omitempty" bson:"fretiro"`
	AnoReconocido          int                 `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido          int                 `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido          int                 `json:"dreconocido,omitempty" bson:"dreconocido"`
	NumeroResuelto         string              `json:"nresuelto,omitempty" bson:"nresuelto"`
	FechaResuelto          time.Time           `json:"fresuelto,omitempty" bson:"fresuelto"`
	Posicion               int                 `json:"posicion,omitempty" bson:"posicion"`
	DescripcionHistorica   string              `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Componente             Componente          `json:"Componente,omitempty" bson:"componente"`
	Grado                  Grado               `json:"Grado,omitempty" bson:"grado"` //grado
	TIM                    Carnet              `json:"Tim,omitempty" bson:"tim"`     //Tarjeta de Identificacion Militar
	Familiar               []Familiar          `json:"Familiar" bson:"familiar"`
	HistorialMilitar       []HistorialMilitar  `json:"HistorialMilitar" bson:"historialmilitar"`
	AppSaman               bool                `json:"appsaman" bson:"appsaman"`
	AppPace                bool                `json:"apppace" bson:"apppace"`
	AppNomina              bool                `json:"appnomina" bson:"appnomina"`
	TiempoSevicio          string              `json:"tiemposervicio,omitempty" bson:"tiemposervicio,omitempty"`
	Pension                Pension             `json:"Pension,omitempty" bson:"pension"`
	Fideicomiso            Fideicomiso         `json:"Fideicomiso,omitempty" bson:"fideicomiso"`
	Anomalia               Anomalia            `json:"Anomalia,omitempty" bson:"anomalia"`
	CodigoComponente       string              `json:"codigocomponente,omitempty" bson:"codigocomponente"`
	NumeroHistoria         string              `json:"numerohistoria,omitempty" bson:"numerohistoria"`
	EstatusCarnet          int                 `json:"estatuscarnet" bson:"estatuscarnet"`
	PaseARetiro            bool                `json:"pasearetiro" bson:"pasearetiro"`
	CIS                    cis.CuidadoIntegral `json:"CIS" bson:"cis"`
}

//Anomalia Irregularidades
type Anomalia struct {
	Hijo bool `json:"hijo,omitempty" bson:"hijo"`
	Ano  bool `json:"ano,omitempty" bson:"ano"`
	Mes  bool `json:"mes,omitempty" bson:"mes"`
	Dia  bool `json:"dia,omitempty" bson:"dia"`
}

//HistorialMilitar Historico
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

//Componente componente
type Componente struct {
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
}

//Grado Rango / Jerarquia
type Grado struct {
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
}

//Listar sistemas
func (m *Militar) Listar() {
<<<<<<< HEAD
	//informaticaipsfa@gmail.com
=======
	//gesaodin@gmail.com
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
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
	// m.ConversionGrado()
	a, mes, d := util.CalcularTiempo(m.FechaIngresoComponente)
	m.TiempoSevicio = strconv.Itoa(a) + "A " + strconv.Itoa(int(mes)) + "M " + strconv.Itoa(d) + "D"
}

func (m *Militar) NumeroHijos() int {
	// hm = m.HistorialMilitar

	return 1
}

//Conversion de Grados
func (m *Militar) Conversion() {
	switch m.Categoria {
	case "ASI":
		m.Clase = "ASI"
	}
}

//Consultar una persona mediante el metodo de MongoDB
func (m *Militar) Consultar() (jSon []byte, err error) {
	var militar Militar
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)

	err = c.Find(bson.M{"id": m.Persona.DatoBasico.Cedula}).One(&militar)
	if err != nil {
		fmt.Println(err.Error())
		msj.Tipo = 0
		jSon, err = json.Marshal(msj)
	} else {
		if militar.Persona.DatoBasico.Cedula == "" {
			msj.Tipo = 0
			jSon, err = json.Marshal(msj)
		} else {
			militar.AplicarReglas()
			jSon, err = json.Marshal(militar)
		}
	}
	return
}

// ConsultarCIS una persona mediante el metodo de MongoDB
func (m *Militar) ConsultarCIS() (jSon []byte, err error) {
	var militar Militar
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccionar := bson.M{"persona": true, "familiar": true, "cis": true}
	err = c.Find(bson.M{"id": m.Persona.DatoBasico.Cedula}).Select(seleccionar).One(&militar)
	if err != nil {
		fmt.Println(err.Error())
		msj.Tipo = 0
		jSon, err = json.Marshal(msj)
	} else {
		if militar.Persona.DatoBasico.Cedula == "" {
			msj.Tipo = 0
			jSon, err = json.Marshal(msj)
		} else {
			militar.AplicarReglas()
			jSon, err = json.Marshal(militar)
		}
	}
	return
}

//GenerarCarnet Generacion de Carnet
func (m *Militar) GenerarCarnet() (TIM Carnet, err error) {
	var mes, dia string

	fecha := time.Now()
	a, me, d := fecha.Date()
	a += 7
	mes = strconv.Itoa(int(me))
	if int(me) < 10 {
		mes = "0" + strconv.Itoa(int(me))
	}
	dia = strconv.Itoa(d)
	if d < 10 {
		dia = "0" + strconv.Itoa(d)
	}
	fvenc := strconv.Itoa(a) + "-" + mes + "-" + dia
	fechavece, _ := time.Parse("2006-01-02", fvenc)

	serial := m.TIM.Usuario + m.TIM.GenerarSerial()
	// fmt.Println("Geenerando Carnet Militar: ", serial)
	TIM.Serial = serial
	TIM.FechaCreacion = fecha
	TIM.FechaVencimiento = fechavece
	TIM.Nombre = m.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = m.Persona.DatoBasico.ApellidoPrimero
	TIM.Componente.Abreviatura = m.Componente.Abreviatura
	TIM.Componente.Descripcion = m.Componente.Descripcion
	TIM.Grado.Abreviatura = m.Grado.Abreviatura
	TIM.Grado.Descripcion = m.Grado.Descripcion
	TIM.ID = m.ID
	TIM.Condicion = m.PaseARetiro
	TIM.Tipo = 0
	TIM.Estatus = 0
	TIM.IP = m.TIM.IP
	TIM.Motivo = m.TIM.Motivo
	TIM.Usuario = m.TIM.Usuario
	// TIM.Usuario = m.TIM.Usuario
	carnet := make(map[string]interface{})
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["estatuscarnet"] = 1
	err = c.Update(bson.M{"id": m.ID}, bson.M{"$set": carnet})

	foto := make(map[string]interface{})
	foto["persona.foto"] = "foto.jpg"
	err = c.Update(bson.M{"id": m.ID}, bson.M{"$set": foto})
	//jSon, err = json.Marshal(TIM)
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
		m.Persona.DatoBasico.NroPersona = numero
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
func (m *Militar) Actualizar(usuario string, ip string) (jSon []byte, err error) {
	var msj Mensaje
	m.TipoDato = 0
	msj.Mensaje = "Su data ha sido actualizada."
	msj.Tipo = 2
	jSon, err = json.Marshal(msj)
	m.MGOActualizar(usuario, ip)
	return
}

//ActualizarMGO Actualizar
func (m *Militar) ActualizarMGO(oid string, familiar map[string]interface{}) (err error) {

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": oid}, bson.M{"$set": familiar})

	if err != nil {
		fmt.Println("Error: " + oid + " -> " + err.Error())
		return
	}
	return
}

//MGOActualizar Actualizando en MONGO
func (m *Militar) MGOActualizar(usuario string, ip string) (err error) {
	var mOriginal Militar
	var traza fanb.Traza

	mOriginal, _ = consultarMongo(m.ID)
	mOriginal.Persona = m.Persona
	mOriginal.Grado = m.Grado
	mOriginal.Componente = m.Componente
	mOriginal.Categoria = m.Categoria
	mOriginal.Clase = m.Clase
	mOriginal.Situacion = m.Situacion
	mOriginal.FechaIngresoComponente = m.FechaIngresoComponente
	mOriginal.FechaAscenso = m.FechaAscenso
	mOriginal.FechaResuelto = m.FechaResuelto
	mOriginal.Posicion = m.Posicion
	mOriginal.NumeroResuelto = m.NumeroResuelto
	mOriginal.CodigoComponente = m.CodigoComponente
	mOriginal.NumeroHistoria = m.NumeroHistoria
	mOriginal.PaseARetiro = m.PaseARetiro
	traza.Time = time.Now()
	traza.Usuario = usuario
	traza.IP = ip
	traza.Log = "Actualizacion: " + mOriginal.Grado.Abreviatura + "|" + mOriginal.Situacion +
		"|" + m.FechaIngresoComponente.String() + "|" + m.FechaAscenso.String()
	traza.Documento = m.ID
	traza.CrearHistoricoConsulta("hmilitar")
	//
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": mOriginal.ID}, &mOriginal)
	if err != nil {
		fmt.Println("Cedula: " + m.ID + " -> " + err.Error())
		return
	}
	s := ActualizarPersona(m.Persona)
	//fmt.Println(m.Persona.DatoBasico.NroPersona)
	go sys.PostgreSQLSAMAN.Exec(s)
	return
}

//SalvarMGO Guardar
func (m *Militar) SalvarMGO() (err error) {

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Insert(m)
	if err != nil {
		fmt.Println("Err: Insertando cedula ", m.Persona.DatoBasico.Cedula, " Descripción: ", err.Error())
	}
	return
}

//consultarMongo una persona mediante el metodo de MongoDB
func consultarMongo(cedula string) (m Militar, err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Find(bson.M{"id": cedula}).One(&m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

//SalvarMGOI Guardar
func (m *Militar) SalvarMGOI(colecion string, objeto interface{}) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(sys.CBASE).C(colecion)
		err = c.Insert(objeto)
	} else {
		c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
		err = c.Insert(objeto)
	}

	//fmt.Println(err)

	return
}

//ConsultarMGO una persona mediante el metodo de MongoDB
func (m *Militar) ConsultarMGO(cedula string) (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Find(bson.M{"id": cedula}).One(&m)
	return
}

//ListarMGO Listado General
func (m *Militar) ListarMGO(cedula string) (lst []Militar, err error) {
	c := sys.MGOSession.DB(sys.CBASE).C("persona")
	err = c.Find(bson.M{}).All(&lst)
	return
}
