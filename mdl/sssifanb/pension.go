package sssifanb

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"gopkg.in/mgo.v2/bson"
)

//Militar militares
type PensionMilitar struct {
	ID                     string     `json:"id,omitempty" bson:"id"`
	TipoDato               int        `json:"tipodato,omitempty" bson:"tipodato"`
	Persona                Persona    `json:"Persona,omitempty" bson:"persona"`
	Categoria              string     `json:"categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion              string     `json:"situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                  string     `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente time.Time  `json:"fingreso,omitempty" bson:"fingreso"`
	FechaAscenso           time.Time  `json:"fascenso,omitempty" bson:"fascenso"`
	FechaRetiro            time.Time  `json:"fretiro,omitempty" bson:"fretiro"`
	AnoReconocido          int        `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido          int        `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido          int        `json:"dreconocido,omitempty" bson:"dreconocido"`
	NumeroResuelto         string     `json:"nresuelto,omitempty" bson:"nresuelto"`
	FechaResuelto          string     `json:"fresuelto,omitempty" bson:"fresuelto"`
	Posicion               int        `json:"posicion,omitempty" bson:"posicion"`
	DescripcionHistorica   string     `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Componente             Componente `json:"Componente,omitempty" bson:"componente"`
	Grado                  Grado      `json:"Grado,omitempty" bson:"grado"` //grado
	Familiar               []Familiar `json:"Familiar" bson:"familiar"`
	Pension                Pension    `json:"Pension,omitempty" bson:"pension"`
}

type Pension struct {
	GradoCodigo            string                   `json:"grado" bson:"grado"`
	ComponenteCodigo       string                   `json:"componente" bson:"componente"`
	Clase                  string                   `json:"clase" bson:"clase"`
	Categoria              string                   `json:"categoria" bson:"categoria"`
	Situacion              string                   `json:"situacion" bson:"situacion"`
	FechaPromocion         string                   `json:"fpromocion" bson:"fpromocion"`
	FechaUltimoAscenso     string                   `json:"fultimoascenso" bson:"fultimoascenso"`
	AnoServicio            int                      `json:"aservicio" bson:"aservicio"`
	MesServicio            int                      `json:"mservicio" bson:"mservicio"`
	DiaServicio            int                      `json:"dservicio" bson:"dservicio"`
	NumeroHijos            int                      `json:"numerohijos" bson:"numerohijos"`
	DatoFinanciero         DatoFinanciero           `json:"DatoFinanciero" bson:"datofinanciero"`
	PensionAsignada        float64                  `json:"pensionasignada" bson:"pensionasignada"`
	HistorialSueldo        []HistorialPensionSueldo `json:"HistorialSueldo" bson:"historialsueldo"`
	PorcentajePrestaciones float64                  `json:"pprestaciones" bson:"pprestaciones"`
}

type HistorialPensionSueldo struct {
	Directiva       string  `json:"directiva" bson:"directiva"`
	Sueldo          float64 `json:"sueldo" bson:"sueldo"`
	Prima           Prima   `json:"Prima" bson:"prima"`
	PensionAsignada float64 `json:"pensionasignada" bson:"pensionasignada"`
	BonoVacacional  float64 `json:"bonovacacional" bson:"bonovacacional"`
	BonoAguinaldo   float64 `json:"bonoaguinaldo" bson:"bonoaguinaldo"`
}

type Prima struct {
	Transporte          float64 `json:"transporte" bson:"transporte"`
	Descendencia        float64 `json:"descendencia" bson:"descendencia"`
	NoAscenso           float64 `json:"noascenso" bson:"noascenso"`
	PorcentajeNoAscenso float64 `json:"pnoascenso" bson:"pnoascenso"`
	Especial            float64 `json:"especial" bson:"especial"`
	SubTotal            float64 `json:"subtotal" bson:"subtotal"`
}

type Beneficiario struct {
	Persona Persona
	Pension Pension

	// Cedula             string
	// Nombre             string
	// Apellido           string
	// Componente         int
	// Grado              int
	// FechaIngreso       string
	// FechaUltimoAscenso string
	// FechaRetiro        string
	// Estatus            int
	// AnoRecoconocido    int
	// MesReconocido      int
	// DiaReconocido      int
	// NumeroHijos        int
	// Sueldo             float64
}

//Listado de Componentes Por Grados
var lstMilitares []PensionMilitar
var lstComponente []fanb.Componente

func (P *Pension) Exportar() {
	fmt.Println("Cargando Componente")
	consultarComponentes()
	fmt.Println("Cargando Militares")
	consultarPensionados()
	//
	i := 0
	coma := ""
	cuerpo := ""
	insert := `INSERT INTO beneficiario (cedula,nombres,apellidos, grado_id, componente_id, fecha_ingreso, f_ult_ascenso, f_retiro,
		f_retiro_efectiva, porcentaje)	VALUES `
	fmt.Println("Creando lote...")
	for _, v := range lstMilitares {
		if i > 0 {
			coma = ","
		}

		grado, componente := obtenerGrado(v.Pension.ComponenteCodigo, v.Pension.GradoCodigo)
		np := v.Persona.DatoBasico.NombrePrimero
		ap := v.Persona.DatoBasico.ApellidoPrimero
		porcentaje := strconv.FormatFloat(v.Pension.PorcentajePrestaciones, 'f', 2, 64)
		cuerpo += coma + `(
				'` + v.Persona.DatoBasico.Cedula + `',
				'` + strings.Replace(np, "'", " ", -1) + `',
				'` + strings.Replace(ap, "'", " ", -1) + `',
				` + grado + `,` + strconv.Itoa(componente) + `,
				'` + v.FechaIngresoComponente.String()[0:10] + `',
				'` + v.FechaAscenso.String()[0:10] + `',
				'` + v.FechaRetiro.String()[0:10] + `','` + v.FechaRetiro.String()[0:10] + `',` + porcentaje + `)`
		i++

		// fmt.Println(" Situacion: ", v.Situacion, " Componente: ", v.Pension.ComponenteCodigo, " Grado Codigo: ", v.Pension.GradoCodigo)
	}
	fmt.Println("Preparando para insertar: ", i)
	query := insert + cuerpo
	// fmt.Println("Consultar ", query)
	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}

}

func consultarPensionados() {
	//Listado de Militares Pensionados
	// var lst []Militar{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccion := bson.M{
		"persona.datobasico":      true,
		"fascenso":                true,
		"fingreso":                true,
		"fretiro":                 true,
		"situacion":               true,
		"pension.pensionasignada": true,
		"pension.grado":           true,
		"pension.componente":      true,
		"pension.numerohijo":      true,
		"pension.datofinanciero":  true,
		"pension.areconocido":     true,
		"pension.pprestaciones":   true,
	}
	err := c.Find(bson.M{"pension.pensionasignada": bson.M{"$gt": 0}}).Select(seleccion).All(&lstMilitares)
	if err != nil {
		fmt.Println("Error en la consulta de Pensionados Militares")
		//return
	}

}

func InsertPension() {

}
func consultarComponentes() {
	comp := sys.MGOSession.DB(sys.CBASE).C(sys.CCOMPONENTE)
	err := comp.Find(bson.M{}).All(&lstComponente)
	if err != nil {
		fmt.Println("Err. Cargando Componentes")
		//
	}
}

func obtenerGrado(codigo string, gradocodigo string) (grado string, componente int) {

	for c, v := range lstComponente {
		componente = c + 1
		if v.Codigo == codigo {
			for _, g := range v.Grado {
				if g.Codigo == gradocodigo {
					grado = g.Cpace
					return
				}
			}
		}
	}
	return "0", 0
}
