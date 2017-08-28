package sssifanb

import (
	"fmt"
	"strconv"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"gopkg.in/mgo.v2/bson"
)

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
var lstMilitares []Militar
var lstComponente []fanb.Componente

func (P *Pension) Exportar() {
	fmt.Println("Cargando Componente")
	consultarComponentes()
	fmt.Println("Cargando Militares")
	consultarPensionados()
	//
	for _, v := range lstMilitares {
		grado, componente := obtenerGrado(v.Pension.ComponenteCodigo, v.Pension.GradoCodigo)
		ins := `INSERT INTO beneficiario (cedula,nombres,apellidos, grado_id, componente_id)
			VALUES (
				'` + v.Persona.DatoBasico.Cedula + `',
				'` + v.Persona.DatoBasico.NombrePrimero + `',
				'` + v.Persona.DatoBasico.ApellidoPrimero + `',
				` + grado + `,` + strconv.Itoa(componente) + `)`

		_, err := sys.PostgreSQLPENSION.Exec(ins)
		if err != nil {
			fmt.Println("Error en el query: ", ins, " ", err.Error())
		}

		// fmt.Println(" Situacion: ", v.Situacion, " Componente: ", v.Pension.ComponenteCodigo, " Grado Codigo: ", v.Pension.GradoCodigo)
	}

}

func consultarPensionados() {
	//Listado de Militares Pensionados
	// var lst []Militar{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccion := bson.M{
		"persona.datobasico":      true,
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
		componente = c
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
