package sssifanb

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
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

//Reduccion de datos de los familiares
type Reduccion struct {
	Cedula          string    `json:"cedula" bson:"cedula"`
	IDT             string    `json:"idt" bson:"idt"`
	Nombre          string    `json:"nombre" bson:"nombre"`
	Sexo            string    `json:"sexo" bson:"sexo"`
	Tipo            string    `json:"tipo" bson:"tipo"` //T Titular Militar | F Familiar
	EsMilitar       bool      `json:"esmilitar" bson:"esmilitar"`
	FechaNacimiento time.Time `json:"fecha" bson:"fecha"`
	Parentesco      string    `json:"parentesco" bson:"parentesco"`
	Situacion       string    `json:"situacion" bson:"situacion"`
	Grado           string    `json:"grado" bson:"grado"`
	Componente      string    `json:"componente" bson:"componente"`
}

//GArchivo Generar CSV
type GArchivo struct {
	Codigo     string `json:"codigo" bson:"codigo"`
	Tipo       string `json:"tipo" bson:"tipo"`
	Situacion  string `json:"situacion" bson:"situacion"`
	Componente string `json:"componente" bson:"componente"`
}

//ExportarCSV Familiares y Titulares Estadisticas
func (r *GArchivo) ExportarCSV() (jSon []byte, err error) {
	var reduccion []Reduccion
	var M Mensaje
	buscar := bson.M{"tipo": "F", "situacion": r.Situacion, "componente": r.Componente}
	nom := "FAM-" + r.Componente + r.Situacion + "-" + time.Now().String()[:19] + ".csv"

	f, err := os.Create("public_web/SSSIFANB/afiliacion/tmp/" + nom)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	err = c.Find(buscar).All(&reduccion)
	if err != nil {
		fmt.Println(err.Error())
	}
	i := 0
	for _, rd := range reduccion {

		i++
		linea := rd.Cedula + ";" + rd.Nombre + ";" + strconv.Itoa(i) + "\n"
		_, e := f.WriteString(linea)
		if e != nil {
			fmt.Println("Error en la linea...")
		}
	}
	fmt.Println("Archivo creado con: ", i)
	f.Sync()
	M.Mensaje = nom + ";" + strconv.Itoa(i)
	M.Tipo = 1
	jSon, err = json.Marshal(M)
	return
}
