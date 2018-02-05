package estadistica

import (
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
}

func Inferencia() {

}

func Descriptiva() {

}

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

	c := sys.MGOSession.DB(sys.CBASE).C("reduccion")
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
		panic(err)
	}
	r.MilitarTitular()
	tarea := make(map[string]interface{})
	tarea["estatus"] = 1
	tarea["fechafin"] = time.Now()
	err = cpendiente.Update(bson.M{"codigo": TP.Codigo}, bson.M{"$set": tarea})
	if err != nil {
		panic(err)
	}
}

//MilitarTitular Familiares y Titulares Estadisticas
func (r *Reduccion) MilitarTitular() (valor bool) {
	fmt.Println("Inciando Creación...")
	var militar []sssifanb.Militar
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccion := bson.M{
		"situacion":                   true,
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
	fmt.Println("Lista la carga...")
	creduccion := sys.MGOSession.DB(sys.CBASE).C("reduccion")
	for _, mil := range militar {

		var prs Reduccion
		prs.Cedula = mil.Persona.DatoBasico.Cedula
		prs.Nombre = mil.Persona.DatoBasico.NombrePrimero + " " + mil.Persona.DatoBasico.ApellidoPrimero
		prs.Tipo = "T"
		prs.FechaNacimiento = mil.Persona.DatoBasico.FechaNacimiento
		prs.Sexo = mil.Persona.DatoBasico.Sexo
		prs.Situacion = mil.Situacion
		prs.EsMilitar = true
		prs.Parentesco = "T"
		err := creduccion.Insert(prs)
		if err != nil {
			fmt.Println("Cedula repetida...")
		}

		for _, Familia := range mil.Familiar {
			var prsf Reduccion
			a, _, _ := Familia.Persona.DatoBasico.FechaNacimiento.Date()
			au, _, _ := time.Now().Date()
			edad := au - a
			prsf.Cedula = Familia.Persona.DatoBasico.Cedula
			prsf.Nombre = Familia.Persona.DatoBasico.NombrePrimero + " " + Familia.Persona.DatoBasico.ApellidoPrimero
			prsf.Tipo = "F"
			prsf.FechaNacimiento = Familia.Persona.DatoBasico.FechaNacimiento
			prsf.Sexo = Familia.Persona.DatoBasico.Sexo
			prsf.EsMilitar = Familia.EsMilitar
			prsf.Parentesco = Familia.Parentesco
			if edad > 15 && edad < 27 {
				ad, _, _ := Familia.Persona.DatoBasico.FechaDefuncion.Date()
				if ad < 1900 {
					err := creduccion.Insert(prsf)
					if err != nil {
						fmt.Println("Cedula repetida...")
					}
				}
			} else if Familia.Parentesco != "HJ" {
				ad, _, _ := Familia.Persona.DatoBasico.FechaDefuncion.Date()
				if ad < 1900 {
					err := creduccion.Insert(prsf)
					if err != nil {
						fmt.Println("Cedula repetida...")
					}
				}
			}

			//fmt.Println(Familia.Persona.DatoBasico.Cedula)
		}
	}
	fmt.Println("Proceso finalizado...")
	return true
}

//ExportarCSVMilitar Familiares y Titulares Estadisticas
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
	c := sys.MGOSession.DB(sys.CBASE).C("reduccion")
	fmt.Println("Preparando los datos...")
	cpendiente := sys.MGOSession.DB(sys.CBASE).C("tareaspendientes")
	cpendiente.Insert(TP)
	err = c.Find(buscar).All(&reduccion)
	if err != nil {
		fmt.Println(err.Error())
	}
	i := 0
	for _, rd := range reduccion {
		if rd.Situacion == "FCP" {
			fmt.Println("PD.-> ", rd.Situacion, rd.Cedula)
		}
		i++
		linea := rd.Cedula + ";" + rd.Nombre + ";" + strconv.Itoa(i) + "\n"
		_, e := f.WriteString(linea)
		if e != nil {
			fmt.Println("Error en la linea...")
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
