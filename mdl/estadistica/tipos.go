package estadistica

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/sys"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TareasPendientes Pendientes
type TareasPendientes struct {
	Codigo      string    `json:"codigo",bson:"codigo"`
	Observacion string    `json:"observacion",bson:"observacion"`
	FechaInicio time.Time `json:"fechainicio",bson:"fechainicio"`
	FechaFin    time.Time `json:"fechafin",bson:"fechafin"`
	Estatus     int       `json:"estatus",bson:"estatus"`
}

//Reduccion de datos de los familiares
type Reduccion struct {
	Cedula          string    `json:"cedula",bson:"cedula"`
	Nombre          string    `json:"nombre",bson:"nombre"`
	Sexo            string    `json:"sexo",bson:"sexo"`
	Tipo            string    `json:"tipo",bson:"tipo"` //T Titular Militar | F Familiar
	EsMilitar       bool      `json:"esmilitar",bson:"esmilitar"`
	FechaNacimiento time.Time `json:"fecha",bson:"fecha"`
	Parentesco      string    `json:"parentesco",bson:"parentesco"`
	Situacion       string    `json:"situacion",bson:"situacion"`
	Grado           string    `json:"grado",bson:"grado"`
	Componente      string    `json:"componente",bson:"componente"`
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
		"situacion":                   true,
		"grado.abreviatura":           true,
		"componente.abreviatura":      true,
		"persona.datobasico":          true,
		"familiar.persona.datobasico": true,
		"familiar.parentesco":         true,
		"familiar.esmilitar":          true,
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
		prs.Nombre = mil.Persona.DatoBasico.ConcatenarNombreApellido()
		prs.Tipo = "T"
		prs.FechaNacimiento = mil.Persona.DatoBasico.FechaNacimiento
		prs.Sexo = mil.Persona.DatoBasico.Sexo
		prs.Situacion = mil.Situacion
		prs.EsMilitar = true
		prs.Parentesco = "T"
		prs.Grado = mil.Grado.Abreviatura
		prs.Componente = mil.Componente.Abreviatura
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

//ExportarCSV Familiares y Titulares Estadisticas
func (r *Reduccion) ExportarCSV(tipo string) {
	var TP TareasPendientes
	nombrefecha := time.Now().String()[:19]
	TP.Codigo = "CSVMIL-" + nombrefecha
	TP.Estatus = 0
	TP.FechaInicio = time.Now()
	buscar := bson.M{"tipo": "T", "situacion": bson.M{"$ne": "FCP"}}
	TP.Observacion = "Creando csv de militares"
	nom := "MIL-"
	if tipo == "F" {
		TP.Observacion = "Creando csv de familiares"
		buscar = bson.M{"tipo": "F"}
		nom = "FAM-"
	}

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
	for _, rd := range reduccion {
		if tipo == "F" {
			a, _, _ := rd.FechaNacimiento.Date()
			au, _, _ := time.Now().Date()
			edad := au - a
			if edad > 15 && edad < 27 {

				i++
				linea := rd.Cedula + ";" + rd.Nombre + ";" + strconv.Itoa(i) + "\n"
				_, e := f.WriteString(linea)
				if e != nil {
					fmt.Println("Error en la linea...")
				}

			} else if rd.Parentesco != "HJ" {

				i++
				linea := rd.Cedula + ";" + rd.Nombre + ";" + strconv.Itoa(i) + "\n"
				_, e := f.WriteString(linea)
				if e != nil {
					fmt.Println("Error en la linea...")
				}

			}

		} else {
			i++
			linea := rd.Cedula + ";" + rd.Nombre + ";" + strconv.Itoa(i) + "\n"
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
