package sssifanb

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/credito"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
	"gopkg.in/mgo.v2/bson"
)

const (
	MILITAR      int = 0
	FAMILIAR     int = 1
	PROVEEDOR    int = 2
	BENEFICIARIO int = 3
)

// Militar militares
type Militar struct {
	ID                           string              `json:"id,omitempty" bson:"id"`
	TipoDato                     int                 `json:"tipodato,omitempty" bson:"tipodato"`
	Persona                      Persona             `json:"Persona,omitempty" bson:"persona"`
	Categoria                    string              `json:"categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion                    string              `json:"situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                        string              `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente       time.Time           `json:"fingreso,omitempty" bson:"fingreso"`
	FechaAscenso                 time.Time           `json:"fascenso,omitempty" bson:"fascenso"`
	FechaRetiro                  time.Time           `json:"fretiro,omitempty" bson:"fretiro"`
	AnoReconocido                int                 `json:"areconocido" bson:"areconocido"`
	MesReconocido                int                 `json:"mreconocido" bson:"mreconocido"`
	DiaReconocido                int                 `json:"dreconocido" bson:"dreconocido"`
	PrimaPorNoAscenso            int                 `json:"pxnoascenso" bson:"pxnoascenso"`
	SituacionPago                string              `json:"situacionpago" bson:"situacionpago"` //Este estatus permite habilitar un pago en pensiones 201 | 202 o paralizar
	PrimaProfesionalizacion      int                 `json:"pprof" bson:"pprof"`
	PrimaEspecialAntiguedadGrado int                 `json:"pespecial" bson:"pespecial"`
	NumeroResuelto               string              `json:"nresuelto,omitempty" bson:"nresuelto"`
	FechaResuelto                time.Time           `json:"fresuelto,omitempty" bson:"fresuelto"`
	Posicion                     int                 `json:"posicion,omitempty" bson:"posicion"`
	Condicion                    int                 `json:"condicion,omitempty" bson:"condicion"`
	PorcentajePrestaciones       float64             `json:"pprestaciones,omitempty" bson:"pprestaciones"`
	DescripcionHistorica         string              `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Componente                   Componente          `json:"Componente,omitempty" bson:"componente"`
	Grado                        Grado               `json:"Grado,omitempty" bson:"grado"` //grado
	TIM                          Carnet              `json:"Tim,omitempty" bson:"tim"`     //Tarjeta de Identificacion Militar
	Familiar                     []Familiar          `json:"Familiar" bson:"familiar"`
	HistorialMilitar             []HistorialMilitar  `json:"HistorialMilitar" bson:"historialmilitar"`
	AppSaman                     bool                `json:"appsaman" bson:"appsaman"`
	AppPace                      bool                `json:"apppace" bson:"apppace"`
	AppNomina                    bool                `json:"appnomina" bson:"appnomina"`
	TiempoSevicio                string              `json:"tiemposervicio,omitempty" bson:"tiemposervicio,omitempty"`
	Pension                      Pension             `json:"Pension,omitempty" bson:"pension"`
	Fideicomiso                  Fideicomiso         `json:"Fideicomiso,omitempty" bson:"fideicomiso"`
	Anomalia                     Anomalia            `json:"Anomalia,omitempty" bson:"anomalia"`
	CodigoComponente             string              `json:"codigocomponente,omitempty" bson:"codigocomponente"`
	NumeroHistoria               string              `json:"numerohistoria,omitempty" bson:"numerohistoria"`
	EstatusCarnet                int                 `json:"estatuscarnet" bson:"estatuscarnet"`
	PaseARetiro                  bool                `json:"pasearetiro" bson:"pasearetiro"`
	CIS                          cis.CuidadoIntegral `json:"CIS" bson:"cis"`
	Credito                      credito.Credito     `json:"Credito" bson:"credito"`
}

// Anomalia Irregularidades
type Anomalia struct {
	Hijo bool `json:"hijo,omitempty" bson:"hijo"`
	Ano  bool `json:"ano,omitempty" bson:"ano"`
	Mes  bool `json:"mes,omitempty" bson:"mes"`
	Dia  bool `json:"dia,omitempty" bson:"dia"`
}

// HistorialMilitar Historico
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

// Componente componente
type Componente struct {
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
}

// Grado Rango / Jerarquia
type Grado struct {
	Nombre      string `json:"nombre,omitempty" bson:"nombre"`
	Descripcion string `json:"descripcion,omitempty" bson:"descripcion"`
	Abreviatura string `json:"abreviatura,omitempty" bson:"abreviatura"`
}

// Listar sistemas
func (m *Militar) Listar() {
	//informaticaipsfa@gmail.com
}

// Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}

// Consultar una persona mediante el metodo de MongoDB
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

// AplicarReglas Reglas Generales
func (m *Militar) AplicarReglas() {
	if m.Situacion != "ACT" {
		m.Conversion()
	}
	fecha := m.FechaIngresoComponente.UTC()
	fechaActual := time.Now()
	if m.Situacion != "ACT" {
		if m.FechaRetiro.String() == "" {
			m.FechaRetiro = m.FechaResuelto
		}
		fechaActual = m.FechaRetiro.UTC()
	}
	a, mes, d := util.CalcularTiempoServicio(fechaActual, fecha)
	//fmt.Println("a , m , d", a, mes, d)
	acr := a + m.AnoReconocido
	mcr := int(mes)
	dcr := d
	if m.DiaReconocido > 29 {
		dcr += m.DiaReconocido - 30
		mcr++
	} else {
		dcr += m.DiaReconocido
	}

	if m.MesReconocido > 11 {
		mcr += m.MesReconocido - 12
		acr++
	} else {
		mcr += m.MesReconocido
	}

	if m.AnoReconocido > 0 || m.MesReconocido > 0 || m.DiaReconocido > 0 {
		if dcr > 29 {
			dcr = dcr - 30
			mcr++
		}
		if mcr > 11 {
			mcr = mcr - 12
			acr++
		}
		a = acr

		mes = time.Month(mcr)
		d = dcr
	}

	m.TiempoSevicio = strconv.Itoa(a) + "A " + strconv.Itoa(int(mes)) + "M " + strconv.Itoa(d) + "D"
}

// NumeroHijos Cantidad de hijos con situacion de beneficiario
func (m *Militar) NumeroHijos() int {
	cantidad := 0
	for _, v := range m.Familiar {
		if v.Parentesco == "HJ" && v.Benficio == true {
			cantidad++
		} else if v.Condicion == 1 {
			cantidad++
		}
	}
	return cantidad
}

// Conversion de Grados
func (m *Militar) Conversion() {
	var comp fanb.Componente
	conver := comp.ConsultarGrado(m.Componente.Abreviatura, m.Grado.Abreviatura)
	m.Pension.ComponenteCodigo = conver.Componente
	m.Pension.GradoCodigo = conver.GradoPace
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

// GenerarCarnet Generacion de Carnet
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
	if m.ID == "1088631" || m.ID == "10275164" {
		fvenc = "2024-07-05"
	}
	fechavece, _ := time.Parse("2006-01-02", fvenc)

	serial := m.TIM.Usuario + m.TIM.GenerarSerial()
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

// Actualizar Vida Militar
func (m *Militar) Actualizar(usuario string, ip string) (jSon []byte, err error) {
	var msj Mensaje
	m.TipoDato = 0
	msj.Mensaje = "Su data ha sido actualizada."
	msj.Tipo = 2
	jSon, err = json.Marshal(msj)
	m.MGOActualizar(usuario, ip)
	return
}

// ActualizarMGO Actualizar familiares del militar con mongodb
func (m *Militar) ActualizarMGO(oid string, familiar map[string]interface{}) (err error) {

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Update(bson.M{"id": oid}, bson.M{"$set": familiar})
	fmt.Println("Actualizando familiar: militar.go ( ActualizarMGO )")
	if err != nil {
		fmt.Println("Error: " + oid + " -> " + err.Error())
		return
	}
	return
}

// MGOActualizar Actualizando datos principales del militar
func (m *Militar) MGOActualizar(usuario string, ip string) (err error) {
	var comp fanb.Componente
	var mOriginal Militar
	var mOriginalf Familiar
	var traza fanb.Traza

	mOriginal, _ = consultarMongo(m.ID)
	mOriginal.Persona = m.Persona

	m.Grado.Nombre = comp.ObtenerGradoID(m.Componente.Abreviatura, m.Grado.Abreviatura)

	fmt.Println("Analizando los datos militar.go (MGOActualizar) ")
	modificar := false
	if mOriginal.Grado.Abreviatura != m.Grado.Abreviatura {
		modificar = true
	}

	mOriginal.Grado = m.Grado
	mOriginal.Componente = m.Componente
	mOriginal.Categoria = m.Categoria
	mOriginal.Clase = m.Clase
	mOriginal.Situacion = m.Situacion
	mOriginal.FechaIngresoComponente = m.FechaIngresoComponente
	mOriginal.FechaAscenso = m.FechaAscenso
	mOriginal.FechaResuelto = m.FechaResuelto
	mOriginal.Posicion = m.Posicion
	mOriginal.Condicion = m.Condicion
	mOriginal.SituacionPago = m.SituacionPago
	mOriginal.NumeroResuelto = m.NumeroResuelto
	mOriginal.CodigoComponente = m.CodigoComponente
	mOriginal.NumeroHistoria = m.NumeroHistoria
	mOriginal.PaseARetiro = m.PaseARetiro
	mOriginal.AnoReconocido = m.AnoReconocido
	mOriginal.MesReconocido = m.MesReconocido
	mOriginal.DiaReconocido = m.DiaReconocido

	if mOriginal.Situacion == "RCP" || mOriginal.Situacion == "FCP" || mOriginal.Situacion == "RSP" || mOriginal.Situacion == "I" || mOriginal.Situacion == "PG" {
		mOriginal.FechaRetiro = m.FechaResuelto
		mOriginal.Pension.DatoFinanciero.Tipo = m.Persona.DatoFinanciero[0].Tipo
		mOriginal.Pension.DatoFinanciero.Cuenta = m.Persona.DatoFinanciero[0].Cuenta
		mOriginal.Pension.DatoFinanciero.Institucion = m.Persona.DatoFinanciero[0].Institucion
	}

	mOriginal.PrimaPorNoAscenso = m.PrimaPorNoAscenso
	mOriginal.PrimaEspecialAntiguedadGrado = m.PrimaEspecialAntiguedadGrado
	mOriginal.PrimaProfesionalizacion = m.PrimaProfesionalizacion

	mOriginal.Pension.PorcentajePrestaciones = m.Pension.PorcentajePrestaciones
	mOriginal.Pension.Causal = m.Pension.Causal
	mOriginal.Pension.Situacion = m.Pension.Situacion

	mOriginal.Pension.GradoCodigo = mOriginal.Grado.Abreviatura
	mOriginal.Pension.ComponenteCodigo = mOriginal.Componente.Abreviatura
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

	if mOriginal.Situacion == "RCP" || mOriginal.Situacion == "FCP" || mOriginal.Situacion == "RSP" || mOriginal.Situacion == "I" || mOriginal.Situacion == "PG" {
		var pension Pension
		mOriginal.FechaRetiro = m.FechaResuelto
		pension.InsertarPensionado(mOriginal, usuario, ip)
	}

	go ActualizarMysqlFullText(ActualizarMysqlFT(mOriginal, mOriginalf))
	//go ActualizarMysqlFullText(ActualizarMysqlFT(mOriginal, mOriginalf))

	//Reducción
	reduc := make(map[string]interface{})
	cred := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	reduc["cedula"] = mOriginal.ID
	reduc["fechanacimiento"] = mOriginal.Persona.DatoBasico.FechaNacimiento
	reduc["nombre"] = mOriginal.Persona.DatoBasico.ConcatenarNombreApellido()
	reduc["situacion"] = mOriginal.Situacion
	err = cred.Update(bson.M{"cedula": mOriginal.ID}, bson.M{"$set": reduc})
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if modificar {
		grado := m.Grado.Abreviatura
		componente := m.Componente.Abreviatura
		for _, fam := range mOriginal.Familiar {
			fam.ActualizarPorReduccion(grado, componente)
		}
	}

	return
}

/*
	func ActualizarMysqlFullText(d string) {
		_, err := sys.MysqlFullText.Exec(d)
		if err != nil {
			fmt.Println("MYSQL FULLTEXT: ", err.Error())
			return
		}
	}
*/
func ActualizarMysqlFullText(queries []string) {
	for _, query := range queries {
		_, err := sys.MysqlFullText.Exec(query)
		if err != nil {
			fmt.Println("MYSQL FULLTEXT: ", err.Error())
			// Decide si quieres continuar con las siguientes consultas o retornar
		}
	}
}

func (m *Militar) ActualizarFoto(id string) {
	persona := make(map[string]interface{})
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	persona["persona.foto"] = id
	err := c.Update(bson.M{"id": id}, bson.M{"$set": persona})
	if err != nil {
		fmt.Println("Err", err.Error())
	}
}

// SalvarMGO Guardar
func (m *Militar) SalvarMGO() (err error) {
	var comp fanb.Componente

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	m.Grado.Nombre = comp.ObtenerGradoID(m.Componente.Abreviatura, m.Grado.Abreviatura)
	err = c.Insert(m)
	if err != nil {
		fmt.Println("Err: Insertando cedula ", m.Persona.DatoBasico.Cedula, " Descripción: ", err.Error())
	}

	go InsertarMysqlFullText(InsertMysqlFT(m, nil))

	//Reducción
	reduc := make(map[string]interface{})
	cred := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	reduc["cedula"] = m.ID
	reduc["fechanacimiento"] = m.Persona.DatoBasico.FechaNacimiento
	reduc["nombre"] = m.Persona.DatoBasico.ConcatenarNombreApellido()
	reduc["situacion"] = m.Situacion
	err = cred.Update(bson.M{"cedula": m.ID}, bson.M{"$set": reduc})
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	return
}

func InsertarMysqlFullText(d string) {
	_, err := sys.MysqlFullText.Exec(d)

	if err != nil {
		fmt.Println("INSERT MYSQL FULLTEXT: ", err.Error())
	}
}

// consultarMongo una persona mediante el metodo de MongoDB
func consultarMongo(cedula string) (m Militar, err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Find(bson.M{"id": cedula}).One(&m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// SalvarMGOI Guardar
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

// ConsultarMGO una persona mediante el metodo de MongoDB
func (m *Militar) ConsultarMGO(cedula string) (err error) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Find(bson.M{"id": cedula}).One(&m)
	return
}

// ListarMGO Listado General
func (m *Militar) ListarMGO(cedula string) (lst []Militar, err error) {
	c := sys.MGOSession.DB(sys.CBASE).C("persona")
	err = c.Find(bson.M{}).All(&lst)
	return
}

// EstadisticasPorComponente Estadisticas Por Componente
func (m *Militar) EstadisticasPorComponente() (jSon []byte, err error) {
	// 	db.militar.aggregate([
	//   {$match:
	//       {"componente.abreviatura": "EJ", "situacion":"ACT"}
	//   },
	//   {$group:
	//     {
	//       _id: { "grado":"$grado.abreviatura", "situacion":"$situacion" },
	//       cantidad:{$sum:1}
	//     }
	//   }
	// ])
	var rs []interface{}
	donde := bson.M{"$match": bson.M{}}
	grupo := bson.M{"$group": bson.M{"_id": bson.M{"componente": "$componente.abreviatura", "situacion": "$situacion"}, "cantidad": bson.M{"$sum": 1}}}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Pipe([]bson.M{donde, grupo}).All(&rs)
	if err != nil {
		return
	}
	// fmt.Println(rs)
	jSon, err = json.Marshal(rs)
	return
}

// EstadisticasPorGrado Estadisticas Por Grado
func (m *Militar) EstadisticasPorGrado(codComponente string) (jSon []byte, err error) {
	// 	db.militar.aggregate([
	//   {$match:
	//       {"componente.abreviatura": "EJ"}
	//
	//   },
	//   {$group:
	//     {
	//       _id: { "codigo":"$grado.nombre", "grado":"$grado.abreviatura", "situacion":"$situacion" },
	//       cantidad:{$sum:1}
	//     }
	//   },
	//   {$sort:
	//    { _id:1}
	//   }
	//
	// ])
	donde := bson.M{"$match": bson.M{"componente.abreviatura": codComponente}}

	grupo := bson.M{"$group": bson.M{"_id": bson.M{"codigo": "$grado.nombre", "grado": "$grado.abreviatura", "situacion": "$situacion"}, "cantidad": bson.M{"$sum": 1}}}
	orden := bson.M{"$sort": bson.M{"_id": 1}}
	var rs []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Pipe([]bson.M{donde, grupo, orden}).All(&rs)
	if err != nil {
		fmt.Println("Error Grados ", err.Error())
		return
	}
	jSon, err = json.Marshal(rs)

	return
}

// EstadisticasFamiliar Estadisticas Por Grado
func (m *Militar) EstadisticasFamiliar() (jSon []byte, err error) {
	donde := bson.M{"$match": bson.M{"tipo": "F"}}
	grupo := bson.M{"$group": bson.M{"_id": bson.M{
		"codigo": "$componente", "situacion": "$situacion"}, "cantidad": bson.M{"$sum": 1}}}
	orden := bson.M{"$sort": bson.M{"_id": 1}}
	var rs []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	err = c.Pipe([]bson.M{donde, grupo, orden}).All(&rs)
	if err != nil {
		fmt.Println("Error Familiares ", err.Error())
		return
	}
	jSon, err = json.Marshal(rs)

	return
}

// ActualizarGradoCodigo Actualizando
func (m *Militar) ActualizarGradoCodigo() {
	var militar []Militar
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err := c.Find(bson.M{"grado.nombre": ""}).All(&militar)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}

	for _, v := range militar {
		var comp fanb.Componente

		v.Grado.Nombre = comp.ObtenerGradoID(v.Componente.Abreviatura, v.Grado.Abreviatura)
		grado := make(map[string]interface{})
		c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
		grado["grado.nombre"] = v.Grado.Nombre
		err = c.Update(bson.M{"id": v.ID}, bson.M{"$set": grado})
		if err != nil {
			fmt.Println("Err", err.Error())
			return
		}
	}
}

// InsertarFullText Insertar
func (m *Militar) InsertarFullText() {
	var militar []Militar
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccionar := bson.M{"persona": true, "componente": true, "grado": true}
	err := c.Find(bson.M{}).Select(seleccionar).All(&militar)
	if err != nil {
		fmt.Println(err.Error())
		msj.Tipo = 0
		// jSon, err = json.Marshal(msj)
	}
	for _, v := range militar {
		fmt.Println("Control ", v.Persona.DatoBasico.Cedula)

	}
}

// InsertMySQL MySQL
func InsertMySQL(mil Militar) {
	cedula := mil.Persona.DatoBasico.Cedula
	direccion := mil.Persona.Direccion
	cuerpo := `a` + direccion[0].Ciudad
	cadena := `INSERT INTO militar ('` + cedula + `','','` + cuerpo + `')`
	fmt.Println(cadena)
}

// MYSQLFULL DATA
type MYSQLFULL struct {
	ID          int    `json:"id,omitempty"`
	Cedula      string `json:"cedula,omitempty"`
	Nombre      string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
	Direccion   string `json:"direccion,omitempty"`
	Familiares  string `json:"familiares,omitempty"`
	Puntuacion  string `json:"puntuacion,omitempty"`
}

// BusquedaFullText Busqueda Avanzada
func (m *Militar) BusquedaFullText(contenido string, tipo int) (jSon []byte, err error) {
	rows, err := sys.MysqlFullText.Query(QueryMysqlText(contenido, tipo))
	if err != nil {
		panic(err.Error())
	}
	var lst []interface{}
	for rows.Next() {
		var mysqldata MYSQLFULL
		var id, cedula, nombre, descripcion, direccion, familiares, puntuacion string
		rows.Scan(&id, &cedula, &nombre, &descripcion, &direccion, &familiares, &puntuacion)
		mysqldata.ID, _ = strconv.Atoi(id)
		mysqldata.Cedula = cedula
		mysqldata.Nombre = nombre
		mysqldata.Descripcion = descripcion
		mysqldata.Direccion = direccion
		mysqldata.Familiares = familiares
		mysqldata.Puntuacion = puntuacion
		lst = append(lst, mysqldata)
	}
	jSon, err = json.Marshal(lst)
	return
}

// QueryMysqlText Consultando
func QueryMysqlText(contenido string, tipo int) string {
	parametro := ""
	switch tipo {
	case 1:
		parametro = "nombre"
		break
	case 2:
		parametro = "descripcion"
		break
	case 3:
		parametro = "familiares"
		break
	case 4:
		parametro = "cedula, nombre, descripcion, direccion, familiares"
		break
	}
	return `SELECT id, cedula, nombre, descripcion, direccion, familiares, MATCH (` + parametro + `)
	AGAINST ('` + contenido + `' IN BOOLEAN MODE) AS puntuacion FROM datos
	WHERE MATCH (` + parametro + `) AGAINST ('` + contenido + `'  IN BOOLEAN MODE)
	ORDER BY puntuacion DESC LIMIT 500`
}

// Consultar una persona mediante el metodo de MongoDB
func (m *Militar) ConsultarCedula(id string) (jSon []byte, err error) {
	var militar Militar
	var msj Mensaje
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)

	err = c.Find(bson.M{"id": id}).One(&militar)

	if err != nil {

		seleccion := bson.M{
			"familiar.persona.datobasico.$": true,
		}
		err = c.Find(bson.M{"familiar.persona.datobasico.cedula": id}).Select(seleccion).One(&militar)
		msj.Tipo = 1
		for _, v := range militar.Familiar {
			msj.Mensaje = v.Persona.DatoBasico.ConcatenarApellidoNombre()
		}
		//militar.Familiar[0].Persona.DatoBasico.ConcatenarApellidoNombre()
		jSon, err = json.Marshal(msj)
	} else {
		msj.Tipo = 1
		msj.Mensaje = militar.Persona.DatoBasico.ConcatenarApellidoNombre()
		jSon, err = json.Marshal(msj)
	}

	return
}
