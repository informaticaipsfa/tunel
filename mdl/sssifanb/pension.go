package sssifanb

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
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

//Pension Codición del militar
type Pension struct {
	GradoCodigo            string                   `json:"grado,omitempty" bson:"grado"`
	ComponenteCodigo       string                   `json:"componente,omitempty" bson:"componente"`
	Clase                  string                   `json:"clase,omitempty" bson:"clase"`
	Categoria              string                   `json:"categoria,omitempty" bson:"categoria"`
	Situacion              string                   `json:"situacion,omitempty" bson:"situacion"`
	Tipo                   string                   `json:"tipo,omitempty" bson:"tipo"`
	Estatus                string                   `json:"estatus,omitempty" bson:"estatus"`
	Razon                  string                   `json:"razon,omitempty" bson:"razon"`
	FechaPromocion         string                   `json:"fpromocion,omitempty" bson:"fpromocion"`
	FechaUltimoAscenso     string                   `json:"fultimoascenso,omitempty" bson:"fultimoascenso"`
	AnoServicio            int                      `json:"aservicio,omitempty" bson:"aservicio"`
	MesServicio            int                      `json:"mservicio,omitempty" bson:"mservicio"`
	DiaServicio            int                      `json:"dservicio,omitempty" bson:"dservicio"`
	NumeroHijos            int                      `json:"numerohijos,omitempty" bson:"numerohijos"`
	DatoFinanciero         DatoFinanciero           `json:"DatoFinanciero,omitempty" bson:"datofinanciero"`
	PensionAsignada        float64                  `json:"pensionasignada,omitempty" bson:"pensionasignada"`
	HistorialSueldo        []HistorialPensionSueldo `json:"HistorialSueldo,omitempty" bson:"historialsueldo"`
	PorcentajePrestaciones float64                  `json:"pprestaciones,omitempty" bson:"pprestaciones"`
	PrimaProfesional       float64                  `json:"pprofesional,omitempty" bson:"pprofesional"`
	PrimaNoAscenso         float64                  `json:"pnoascenso,omitempty" bson:"pnoascenso"`
	PrimaEspecial          float64                  `json:"pespecial,omitempty" bson:"pespecial"`
	MedidaJudicial         []MedidaJudicial         `json:"MedidaJudicial,omitempty" bson:"medidajudicial"`
	Descuentos             []Descuentos             `json:"Descuentos,omitempty" bson:"descuentos"`
}

//HistorialPensionSueldo Historico
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
}

//Listado de Componentes Por Grados
var lstMilitares []PensionMilitar
var lstComponente []fanb.Componente

//NumeroHijos Contando numero de hijos
func (m *PensionMilitar) NumeroHijos() int {
	cantidad := 0
	for _, v := range m.Familiar {
		if v.Parentesco == "HJ" && v.Benficio == true {
			cantidad++
		}
	}
	return cantidad
}

//ActulizarPensionadosID Por Cedula
func ActulizarPensionadosID(id string) {
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccion := bson.M{
		"persona.datobasico":      true,
		"fascenso":                true,
		"fingreso":                true,
		"fretiro":                 true,
		"situacion":               true,
		"grado.abreviatura":       true,
		"componente.abreviatura":  true,
		"pension.pensionasignada": true,
		"pension.grado":           true,
		"pension.componente":      true,
		"pension.numerohijo":      true,
		"pension.datofinanciero":  true,
		"pension.areconocido":     true,
		"pension.pprestaciones":   true,
		"pension.pprofesional":    true,
		"pension.pnoascenso":      true,
		"pension.pespecial":       true,
		"familiar":                true,
	}
	buscar := bson.M{"id": id}
	err := c.Find(buscar).Select(seleccion).All(&lstMilitares)
	if err != nil {
		fmt.Println("Error en la consulta de Pensionados Militares")
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
		"grado.abreviatura":       true,
		"componente.abreviatura":  true,
		"situacion":               true,
		"pension.pensionasignada": true,
		"pension.grado":           true,
		"pension.componente":      true,
		"pension.numerohijo":      true,
		"pension.datofinanciero":  true,
		"pension.areconocido":     true,
		"pension.pprestaciones":   true,
		"pension.pprofesional":    true,
		"pension.pnoascenso":      true,
		"pension.pespecial":       true,
		"familiar":                true,
	}

	buscar := bson.M{"situacion": "RCP"}
	// buscar := bson.M{"id": "26419599"}
	err := c.Find(buscar).Select(seleccion).All(&lstMilitares)
	if err != nil {
		fmt.Println("Error en la consulta de Pensionados Militares")
		//return
	}

}

//Exportar Controlando Datos
func (P *Pension) Exportar(cedula string, tipo int32) {
	fmt.Println("Cargando Componente")
	consultarComponentes()

	fmt.Println("Cargando Militares")
	if tipo == 0 {
		consultarPensionados()
	} else {
		ActulizarPensionadosID(cedula)
	}
	//
	i := 0
	coma := ""
	cuerpo := ""
	linea := ""
	insert := `INSERT INTO beneficiario (cedula,nombres,apellidos, grado_id, componente_id, fecha_ingreso, f_ult_ascenso, f_retiro,
		f_retiro_efectiva, st_no_ascenso, st_profesion, monto_especial, status_id, n_hijos, porcentaje, numero_cuenta, tipo, banco, situacion)	VALUES `
	fmt.Println("Creando lote...")
	j := 0
	k := 0
	l := 0
	for _, v := range lstMilitares {
		if i > 0 {
			coma = ","
		}

		grado, componente := obtenerGrado(v.Componente.Abreviatura, v.Grado.Abreviatura)
		if grado == "0" {
			l++
			j++
			linea += strconv.Itoa(j) + " : " + v.Persona.DatoBasico.Cedula
			grado, componente = obtenerGrado(v.Pension.ComponenteCodigo, v.Pension.GradoCodigo)
			if grado == "0" {
				k++
				linea += "  --->  " + strconv.Itoa(k) + ": " + grado + " :: " + strconv.Itoa(componente)
			} else {
				linea += "  --->  OK => Resuelto... " + v.Pension.ComponenteCodigo + "|" + v.Pension.GradoCodigo + " == " + grado + " :: " + strconv.Itoa(componente)
			}
			linea += " -- \n"
		}
		np := v.Persona.DatoBasico.NombrePrimero
		ap := v.Persona.DatoBasico.ApellidoPrimero
		porcentaje := strconv.FormatFloat(v.Pension.PorcentajePrestaciones, 'f', 2, 64)
		pprofesional := strconv.FormatFloat(v.Pension.PrimaProfesional, 'f', 2, 64)
		pnoascenso := strconv.FormatFloat(v.Pension.PrimaNoAscenso, 'f', 2, 64)
		pespecial := strconv.FormatFloat(v.Pension.PrimaEspecial, 'f', 2, 64)
		numero := v.Pension.DatoFinanciero.Cuenta
		cuenta := v.Pension.DatoFinanciero.Institucion
		tipo := v.Pension.DatoFinanciero.Tipo
		fRetiro := v.FechaRetiro.String()[0:10]
		fAscenso := v.FechaAscenso.String()[0:10]
		if len(fRetiro) < 10 {
			fRetiro = fAscenso
		}
		if len(fAscenso) < 10 {
			fAscenso = fRetiro
		}

		cuerpo += coma + `(
				'` + v.Persona.DatoBasico.Cedula + `',
				'` + strings.Replace(np, "'", " ", -1) + `',
				'` + strings.Replace(ap, "'", " ", -1) + `',
				` + grado + `,` + strconv.Itoa(componente) + `,
				'` + v.FechaIngresoComponente.String()[0:10] + `',
				'` + fAscenso + `',
				'` + fRetiro + `',
				'` + v.Persona.DatoBasico.FechaDefuncion.String()[0:10] + `',
				` + pnoascenso + `,
				` + pprofesional + `,
				` + pespecial + `,
				201,
				` + strconv.Itoa(v.NumeroHijos()) + `,
				` + porcentaje + `,
				'` + numero + `',
				'` + cuenta + `',
				'` + tipo + `',
				'` + v.Situacion + `')`
		i++

		//fmt.Println("#", strconv.Itoa(i), " Cedula: ", v.Persona.DatoBasico.Cedula, " Componente: ", v.Pension.ComponenteCodigo, " Grado Codigo: ", v.Pension.GradoCodigo)
	}

	fmt.Println("Preparando para insertar: ", i)
	query := insert + cuerpo
	// fmt.Println("Consultar ", query)
	fmt.Println(linea)
	fmt.Println("Cantidad de errores .-> ", l)
	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}

}

// consultarPensionadosFamiliares Familiares
func consultarPensionadosFamiliares() {
	//Listado de Militares Pensionados
	// var lst []Militar{}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccion := bson.M{
		"persona.datobasico":      true,
		"fascenso":                true,
		"fingreso":                true,
		"fretiro":                 true,
		"grado.abreviatura":       true,
		"componente.abreviatura":  true,
		"situacion":               true,
		"pension.pensionasignada": true,
		"pension.grado":           true,
		"pension.componente":      true,
		"pension.numerohijo":      true,
		"pension.datofinanciero":  true,
		"pension.areconocido":     true,
		"pension.pprestaciones":   true,
		"pension.pprofesional":    true,
		"pension.pnoascenso":      true,
		"pension.pespecial":       true,
		"familiar":                true,
	}

	buscar := bson.M{"situacion": "FCP"}
	// buscar := bson.M{"id": "26419599"}
	err := c.Find(buscar).Select(seleccion).All(&lstMilitares)
	if err != nil {
		fmt.Println("Error en la consulta de Pensionados Militares")
		//return
	}

}

//ExportarFamiliares Controlando Datos
func (P *Pension) ExportarFamiliares() {
	fmt.Println("Cargando Componente")
	consultarComponentes()

	fmt.Println("Cargando Militares")
	consultarPensionadosFamiliares()

	i := 0
	coma := ""
	cuerpo := ""
	linea := ""
	insert := `INSERT INTO beneficiario (cedula,nombres,apellidos, grado_id, componente_id, fecha_ingreso, f_ult_ascenso, f_retiro,
		f_retiro_efectiva, st_no_ascenso, st_profesion, monto_especial, status_id, n_hijos, porcentaje, numero_cuenta, tipo, banco, situacion)	VALUES `

	familiar := `INSERT INTO familiar (titular,cedula, nombres, apellidos,sexo,fecha_nacimiento,edo_civil,f_defuncion,
		autorizado,tipo,banco,numero,situacion,estatus,motivo,f_reincorporacion, porcentaje)	VALUES `
	fmt.Println("Creando lote...")
	j := 0
	k := 0
	l := 0
	for _, v := range lstMilitares {
		if i > 0 {
			coma = ","
		}

		grado, componente := obtenerGrado(v.Pension.ComponenteCodigo, v.Pension.GradoCodigo)
		if grado == "0" {
			l++
			j++
			linea += strconv.Itoa(j) + " : " + v.Persona.DatoBasico.Cedula
			grado, componente = obtenerGrado(v.Componente.Abreviatura, v.Grado.Abreviatura)
			if grado == "0" {
				k++
				linea += "  --->  " + strconv.Itoa(k) + ": " + grado + " :: " + strconv.Itoa(componente)
			} else {
				linea += "  --->  OK => Resuelto... " + v.Pension.ComponenteCodigo + "|" + v.Pension.GradoCodigo + " == " + grado + " :: " + strconv.Itoa(componente)
			}
			linea += " -- \n"
		}
		np := v.Persona.DatoBasico.NombrePrimero
		ap := v.Persona.DatoBasico.ApellidoPrimero
		porcentaje := strconv.FormatFloat(v.Pension.PorcentajePrestaciones, 'f', 2, 64)
		pprofesional := strconv.FormatFloat(v.Pension.PrimaProfesional, 'f', 2, 64)
		pnoascenso := strconv.FormatFloat(v.Pension.PrimaNoAscenso, 'f', 2, 64)
		pespecial := strconv.FormatFloat(v.Pension.PrimaEspecial, 'f', 2, 64)
		numero := v.Pension.DatoFinanciero.Cuenta
		cuenta := v.Pension.DatoFinanciero.Institucion
		tipo := v.Pension.DatoFinanciero.Tipo
		fRetiro := v.FechaRetiro.String()[0:10]
		fAscenso := v.FechaAscenso.String()[0:10]
		if len(fRetiro) < 10 {
			fRetiro = fAscenso
		}
		if len(fAscenso) < 10 {
			fAscenso = fRetiro
		}
		estatus := "201" //ACTIVO PARA GENERAR LA PENSION A SUS SOBREVIVIENTES
		if len(v.Familiar) == 0 {
			estatus = "202" //DESACTIVANDO PENSION Y CALCULOS DE LA PENSION SIN FAMILIARES
		}
		cuerpo += coma + `('` + v.Persona.DatoBasico.Cedula + `',
				'` + strings.Replace(np, "'", " ", -1) + `','` + strings.Replace(ap, "'", " ", -1) + `',
				` + grado + `,` + strconv.Itoa(componente) + `,'` + v.FechaIngresoComponente.String()[0:10] + `',
				'` + fAscenso + `','` + fRetiro + `','` + v.Persona.DatoBasico.FechaDefuncion.String()[0:10] + `',
				` + pnoascenso + `,` + pprofesional + `,` + pespecial + `,` + estatus + `,` + strconv.Itoa(v.NumeroHijos()) + `,
				` + porcentaje + `,'` + numero + `','` + cuenta + `','` + tipo + `','` + v.Situacion + `')`
		i++
		familiar += coma + insertFamiliares(v.Familiar)
		//fmt.Println("#", strconv.Itoa(i), " Cedula: ", v.Persona.DatoBasico.Cedula, " Componente: ", v.Pension.ComponenteCodigo, " Grado Codigo: ", v.Pension.GradoCodigo)
	}

	fmt.Println("Preparando para insertar: ", i)
	query := insert + cuerpo
	// fmt.Println("Consultar ", query)
	fmt.Println(linea)
	fmt.Println("Cantidad de errores .-> ", l)
	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("Error en el query: ", err.Error())
	}

}

func insertFamiliares(f []Familiar) (linea string) {
	for _, v := range f {
		linea = v.Persona.DatoBasico.Cedula
	}
	return
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

					grado = strconv.Itoa(g.Cpace)
					return
				}
			}
		}
	}
	return "0", 0
}

//
// func (P *Pension) CargarCodigoPace() {
// 	var componenteCodigo fanb.Componente
// 	sq, err := sys.PostgreSQLPENSION.Query(`
// 		select c.descripcion, g.nombre, g.id, g.codigo from componente c
// 		JOIN grado g ON c.id=g.componente_id
// 		ORDER BY c.id
// 	`)
//
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	for sq.Next() {
// 		var des, nom string
// 		var id, cod int
// 		sq.Scan(&des, &nom, &id, &cod)
// 		c := sys.MGOSession.DB(sys.CBASE).C(sys.CCOMPONENTE)
// 		fmt.Println(componenteCodigo.ComponenteCodigo(des), nom)
// 		comp := make(map[string]interface{})
// 		comp["Grado.$.cpace"] = cod
// 		buscar := bson.M{"codigo": componenteCodigo.ComponenteCodigo(des), "Grado.codigo": nom}
// 		err = c.Update(buscar, bson.M{"$set": comp})
// 		if err != nil {
// 			fmt.Println("Err", err.Error())
// 			//return
// 		}
// 	}
// }
