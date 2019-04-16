package estadistica

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/sys"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TareasPendientes Pendientes
type TareasPendientes struct {
	Codigo      string    `json:"codigo" bson:"codigo"`
	Observacion string    `json:"observacion" bson:"observacion"`
	FechaInicio time.Time `json:"fechainicio" bson:"fechainicio"`
	FechaFin    time.Time `json:"fechafin" bson:"fechafin"`
	Estatus     int       `json:"estatus" bson:"estatus"`
	Tipo        string    `json:"tipo" bson:"tipo"`
}

//Reduccion de datos Generales
type Reduccion struct {
	Cedula                 string           `json:"cedula,omitempty" bson:"cedula"`
	Persona                sssifanb.Persona `json:"Persona,omitempty" bson:"persona"`
	IDT                    string           `json:"idt,omitempty" bson:"idt"`
	Nombre                 string           `json:"nombre,omitempty" bson:"nombre"`
	Sexo                   string           `json:"sexo,omitempty" bson:"sexo"`
	Tipo                   string           `json:"tipo,omitempty" bson:"tipo"`           //T Titular Militar | F Familiar
	EsMilitar              bool             `json:"esmilitar,omitempty" bson:"esmilitar"` //
	FechaNacimiento        time.Time        `json:"fecha,omitempty" bson:"fecha"`
	Parentesco             string           `json:"parentesco,omitempty" bson:"parentesco"`
	Categoria              string           `json:"categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion              string           `json:"situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                  string           `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente time.Time        `json:"fingreso,omitempty" bson:"fingreso"`
	FechaAscenso           time.Time        `json:"fascenso,omitempty" bson:"fascenso"`
	FechaRetiro            time.Time        `json:"fretiro,omitempty" bson:"fretiro"`
	AnoReconocido          int              `json:"areconocido" bson:"areconocido"`
	MesReconocido          int              `json:"mreconocido" bson:"mreconocido"`
	DiaReconocido          int              `json:"dreconocido" bson:"dreconocido"`
	NumeroResuelto         string           `json:"nresuelto,omitempty" bson:"nresuelto"`
	FechaResuelto          time.Time        `json:"fresuelto,omitempty" bson:"fresuelto"`
	Grado                  string           `json:"grado" bson:"grado"`
	Componente             string           `json:"componente" bson:"componente"`
	FechaCreacion          time.Time        `json:"fcreacion,omitempty" bson:"fcreacion"`
	FechaVencimiento       time.Time        `json:"fvencimiento,omitempty" bson:"fvencimiento"`
}

func Inferencia() {

}

func Descriptiva() {

}

//ListarColecciones Listado
func (r *Reduccion) ListarColecciones() (jSon []byte, err error) {

	db := sys.MGOSession.DB(sys.CBASE)
	nombres, err := db.CollectionNames()
	if err != nil {
		log.Printf("Fallo la conexión para las coleccion: %v", err)
	}
	jSon, err = json.Marshal(nombres)
	return
}

//ListarPendientes Pendientes
func (r *Reduccion) ListarPendientes() (jSon []byte, err error) {
	var tp []TareasPendientes
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CTAREASPENDIENDTE)
	seleccion := bson.M{"estatus": bson.M{"$ne": 2}}
	err = c.Find(seleccion).All(&tp)
	if err != nil {
		fmt.Println(err.Error())
	}
	jSon, err = json.Marshal(tp)
	return
}

//ValidarColeccion Validaciones
func (r *Reduccion) ValidarColeccion(coleccion string) (valor bool) {
	valor = false
	db := sys.MGOSession.DB(sys.CBASE)
	nombres, err := db.CollectionNames()
	if err != nil {
		// Handle error
		log.Printf("Fallo la conexión para las coleccion: %v", err)

	}

	// Simply search in the names slice, e.g.
	for _, nombre := range nombres {
		// fmt.Println(nombre)
		if nombre == coleccion {
			log.Printf("La colección existe!")
			valor = true
			break
		}
	}

	return valor
}

//CrearColeccion Crear Coleccion de Mongo para la Reduccion
func (r *Reduccion) CrearColeccion(coleccion string) {

	var TP TareasPendientes
	var prs Reduccion
	prs.Cedula = "0"
	prs.Nombre = "X"
	prs.Tipo = "X"

	TP.Codigo = "XC-" + time.Now().String()[:19]
	TP.Estatus = 0
	TP.FechaInicio = time.Now()
	TP.Observacion = "Creando colección de cédula"
	cpendiente := sys.MGOSession.DB(sys.CBASE).C("tareaspendientes")
	cpendiente.Insert(TP)

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	err := c.Insert(prs)
	if err != nil {
		panic(err)
	}

	index := mgo.Index{
		Key:        []string{"cedula"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		fmt.Println("No se logró crear el indice de la cedula")
	}
	r.MilitarTitular()
	tarea := make(map[string]interface{})
	tarea["estatus"] = 1
	tarea["fechafin"] = time.Now()
	err = cpendiente.Update(bson.M{"codigo": TP.Codigo}, bson.M{"$set": tarea})
	if err != nil {
		fmt.Println("Error al finalizar la tarea pendiente")
	}
	fmt.Println("Proceso finalizado.")
}

//MilitarTitular Familiares y Titulares Estadisticas
func (r *Reduccion) MilitarTitular() (valor bool) {
	fmt.Println("Inciando Creación...")
	var militar []sssifanb.Militar
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccion := bson.M{
		"categoria":                   true,
		"situacion":                   true,
		"clase":                       true,
		"fingreso":                    true,
		"fascenso":                    true,
		"fretiro":                     true,
		"areconocido":                 true,
		"mreconocido":                 true,
		"dreconocido":                 true,
		"nresuelto":                   true,
		"fresuelto":                   true,
		"grado.abreviatura":           true,
		"componente.abreviatura":      true,
		"persona.datobasico":          true,
		"persona.datofisico":          true,
		"persona.datofisionomico":     true,
		"persona.direccion":           true,
		"persona.telefono":            true,
		"persona.correo":              true,
		"familiar.persona.datobasico": true,
		"familiar.parentesco":         true,
		"familiar.esmilitar":          true,
		"tim.fechacreacion":           true,
		"tim.fechavencimiento":        true,
	}
	// err := c.Find(bson.M{}).Select(seleccion).Limit(4).All(&militar)
	fmt.Println("Preparando los datos...")
	err := c.Find(bson.M{}).Select(seleccion).All(&militar)
	if err != nil {
		fmt.Println(err.Error())
	}
	repetidos := 0
	fmt.Println("Lista la carga...")
	creduccion := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	for _, mil := range militar { //Introducir Militares
		var prs Reduccion
		prs.Cedula = mil.Persona.DatoBasico.Cedula
		prs.Persona.DatoBasico = mil.Persona.DatoBasico
		prs.Persona.DatoFisico = mil.Persona.DatoFisico
		prs.Persona.DatoFisionomico = mil.Persona.DatoFisionomico
		prs.Persona.Direccion = mil.Persona.Direccion
		prs.Persona.Correo = mil.Persona.Correo
		prs.Persona.Telefono = mil.Persona.Telefono
		prs.IDT = mil.Persona.DatoBasico.Cedula
		prs.Nombre = mil.Persona.DatoBasico.ConcatenarNombreApellido()
		prs.Tipo = "T"
		prs.FechaNacimiento = mil.Persona.DatoBasico.FechaNacimiento
		prs.Sexo = mil.Persona.DatoBasico.Sexo
		prs.Categoria = mil.Categoria
		prs.Situacion = mil.Situacion
		prs.Clase = mil.Clase
		prs.FechaIngresoComponente = mil.FechaIngresoComponente
		prs.FechaAscenso = mil.FechaAscenso
		prs.FechaRetiro = mil.FechaRetiro
		prs.AnoReconocido = mil.AnoReconocido
		prs.MesReconocido = mil.MesReconocido
		prs.DiaReconocido = mil.DiaReconocido
		prs.NumeroResuelto = mil.NumeroResuelto
		prs.FechaResuelto = mil.FechaResuelto
		prs.EsMilitar = true
		prs.Parentesco = "T"
		prs.Grado = mil.Grado.Abreviatura
		prs.Componente = mil.Componente.Abreviatura
		prs.FechaCreacion = mil.TIM.FechaCreacion
		prs.FechaVencimiento = mil.TIM.FechaVencimiento
		err := creduccion.Insert(prs)
		if err != nil {
			fmt.Println(err.Error())
			repetidos++
		}
	}
	fmt.Println("Procesando datos militares. Por favor espere.")
	time.Sleep(time.Minute * 1)
	fmt.Println("Preparando datos familiares.")
	for _, mili := range militar {
		for _, Familia := range mili.Familiar {
			var prsf Reduccion
			prsf.Cedula = Familia.Persona.DatoBasico.Cedula
			prsf.IDT = mili.Persona.DatoBasico.Cedula
			prsf.Nombre = Familia.Persona.DatoBasico.ConcatenarNombreApellido()
			prsf.Tipo = "F"
			prsf.FechaNacimiento = Familia.Persona.DatoBasico.FechaNacimiento
			prsf.Sexo = Familia.Persona.DatoBasico.Sexo
			prsf.EsMilitar = Familia.EsMilitar
			prsf.Parentesco = Familia.Parentesco
			prsf.Situacion = mili.Situacion
			prsf.Grado = mili.Grado.Abreviatura
			prsf.Componente = mili.Componente.Abreviatura
			ad, _, _ := Familia.Persona.DatoBasico.FechaDefuncion.Date()
			if ad < 1900 {
				err := creduccion.Insert(prsf)
				if err != nil {
					repetidos++
				}
			}
		}
	}

	fmt.Println("Existen ( ", repetidos, " ) repetidos.")
	time.Sleep(time.Minute * 1)
	fmt.Println("Procesando datos familiares. Por favor espere.")
	return true
}

var ListadoEstados []fanb.Estado

func ConsultarEstado(cod string) string {
	var est string
	for _, estado := range ListadoEstados {
		if cod == estado.Codigo {
			est = estado.Nombre
		}
	}
	return est
}

//ExportarCSV Familiares y Titulares Estadisticas
func (r *Reduccion) ExportarCSV(tipo string) {
	var TP TareasPendientes
	var Estados fanb.Estado

	ListadoEstados = Estados.ConsultarTodo() //Cargando todos los estados

	nombrefecha := time.Now().String()[:19]
	TP.Estatus = 0
	TP.FechaInicio = time.Now()
	// buscar := bson.M{"tipo": "T", "situacion": bson.M{"$ne": "FCP"}}
	buscar := bson.M{"tipo": "T"}
	TP.Observacion = "Creando csv de militares"
	TP.Tipo = "CSV"
	nom := "MIL-"
	if tipo == "F" {
		TP.Observacion = "Creando csv de familiares"
		buscar = bson.M{"tipo": "F"}
		nom = "FAM-"
	}
	TP.Codigo = nom + nombrefecha

	fmt.Println("Inciando Creación...")
	var reduccion []Reduccion
	f, err := os.Create("public_web/SSSIFANB/panel/tmp/" + nom + nombrefecha + ".csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	fmt.Println("Preparando los datos...")
	cpendiente := sys.MGOSession.DB(sys.CBASE).C("tareaspendientes")
	cpendiente.Insert(TP)
	err = c.Find(buscar).All(&reduccion)
	if err != nil {
		fmt.Println(err.Error())
	}
	i := 0

	if tipo == "T" {
		cabecera := "#;cedula;nacionalidad;apellido;nombre;" +
			"sexo;fecha nacimiento;fecha defuncion;peso;talla;cabello;ojos;piel;" +
			//"estatura;" +
			"grupo sanguineo;sena particular;correo;telefono;ciudad;estado;municipio;direccion;" +
			"categoria;situacion;clase;fecha ingreso;fecha ascenso;fecha resuelto;numero resuelto;" +
			"fecha retiro;grado;componente;fecha creacion;fecha vencimiento\n"
		_, e := f.WriteString(cabecera)
		if e != nil {
			fmt.Println("Error en la linea...")
		}
	} else {
		cabecera := "#;cedula;nombre;parentesco;sexo;fecha nacimiento;titular\n"
		_, e := f.WriteString(cabecera)
		if e != nil {
			fmt.Println("Error en la linea...")
		}
	}
	for _, rd := range reduccion {
		if tipo == "F" {
			// a, _, _ := rd.FechaNacimiento.Date()
			// au, _, _ := time.Now().Date()
			// edad := au - a
			// if edad > 15 && edad < 27 {

			// 	i++
			// 	linea := rd.Cedula + ";" + rd.Nombre + ";" + strconv.Itoa(i) + "\n"
			// 	_, e := f.WriteString(linea)
			// 	if e != nil {
			// 		fmt.Println("Error en la linea...")
			// 	}

			// } else if rd.Parentesco != "HJ" {
			convertir := rd.FechaNacimiento.Format("2006-01-02")
			fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)
			i++
			linea := strconv.Itoa(i) + ";" + rd.Cedula + ";" +
				rd.Nombre + ";" + rd.Parentesco + ";" + rd.Sexo + ";" + fechaSlashNacimiento + ";" + rd.IDT +
				"\n"
			_, e := f.WriteString(linea)
			if e != nil {
				fmt.Println("Error en la linea...")
			}

			// }

		} else {
			ciudad := ""
			estado := ""
			municipio := ""
			direccion := ""
			if len(rd.Persona.Direccion) > 0 {
				ciudad = rd.Persona.Direccion[0].Ciudad
				estado = rd.Persona.Direccion[0].Estado
				municipio = rd.Persona.Direccion[0].Municipio
				direccion = rd.Persona.Direccion[0].CalleAvenida + " Casa " + rd.Persona.Direccion[0].Casa
			}

			convertir := rd.Persona.DatoBasico.FechaNacimiento.Format("2006-01-02")
			fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)

			a, _, _ := rd.Persona.DatoBasico.FechaDefuncion.Date() // Fecha de Defunción en caso de poseerla
			fechaSlashDefuncion := ""
			if a > 1000 {
				convertirDef := rd.Persona.DatoBasico.FechaDefuncion.Format("2006-01-02")
				fechaSlashDefuncion = strings.Replace(convertirDef, "-", "/", -1)
			}
			fechai := rd.FechaIngresoComponente.Format("2006-01-02") // Fecha de Ingreso al componente
			fechaIngreso := strings.Replace(fechai, "-", "/", -1)
			fechaa := rd.FechaAscenso.Format("2006-01-02") // Fecha de Ascenso al componente
			fechaAscenso := strings.Replace(fechaa, "-", "/", -1)

			aa, _, _ := rd.FechaResuelto.Date() // Fecha de Defunción en caso de poseerla
			fechaSlashResuelto := ""
			if aa > 1000 {
				convertirRes := rd.FechaResuelto.Format("2006-01-02")
				fechaSlashResuelto = strings.Replace(convertirRes, "-", "/", -1)
			}

			aaa, _, _ := rd.FechaRetiro.Date() // Fecha de Defunción en caso de poseerla
			fechaSlashRetiro := ""
			if aaa > 1000 {
				convertirRet := rd.FechaRetiro.Format("2006-01-02")
				fechaSlashRetiro = strings.Replace(convertirRet, "-", "/", -1)
			}
			fechaCreacion := ""
			if aaa > 1000 {
				convertirRet := rd.FechaCreacion.Format("2006-01-02")
				fechaCreacion = strings.Replace(convertirRet, "-", "/", -1)
			}
			fechaVencimiento := ""
			if aaa > 1000 {
				convertirRet := rd.FechaVencimiento.Format("2006-01-02")
				fechaVencimiento = strings.Replace(convertirRet, "-", "/", -1)
			}

			i++
			linea := strconv.Itoa(i) +
				";" + rd.Cedula +
				";" + rd.Persona.DatoBasico.Nacionalidad +
				";" + rd.Persona.DatoBasico.ApellidoPrimero +
				";" + rd.Persona.DatoBasico.NombrePrimero +
				";" + rd.Persona.DatoBasico.Sexo +
				";" + fechaSlashNacimiento +
				";" + fechaSlashDefuncion +
				";" + rd.Persona.DatoFisico.Peso + ";" + rd.Persona.DatoFisico.Talla +
				";" + rd.Persona.DatoFisionomico.ColorCabello + ";" + rd.Persona.DatoFisionomico.ColorOjos +
				";" + rd.Persona.DatoFisionomico.ColorPiel +
				//";" + rd.Persona.DatoFisionomico.Estatura +
				";" + rd.Persona.DatoFisionomico.GrupoSanguineo + ";" + rd.Persona.DatoFisionomico.SenaParticular +
				";" + rd.Persona.Correo.Principal +
				";" + rd.Persona.Telefono.Movil + "|" + rd.Persona.Telefono.Domiciliario +
				";" + ciudad + ";" + ConsultarEstado(estado) + ";" + municipio + ";" + direccion +
				";" + rd.Categoria +
				";" + rd.Situacion +
				";" + rd.Clase +
				";" + fechaIngreso +
				";" + fechaAscenso +
				";" + fechaSlashResuelto +
				";" + rd.NumeroResuelto +
				";" + fechaSlashRetiro +
				";" + rd.Grado +
				";" + rd.Componente +
				";" + fechaCreacion +
				";" + fechaVencimiento +
				"\n"
			_, e := f.WriteString(linea)
			if e != nil {
				fmt.Println("Error en la linea...")
			}
		}

	}
	fmt.Println("Archivo creado con: ", i)
	f.Sync()
	tarea := make(map[string]interface{})
	tarea["estatus"] = 1
	tarea["fechafin"] = time.Now()
	err = cpendiente.Update(bson.M{"codigo": TP.Codigo}, bson.M{"$set": tarea})
	if err != nil {
		panic(err)
	}

}
