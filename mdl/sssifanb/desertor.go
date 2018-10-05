package sssifanb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

type Desertor struct {
}

func (d *Desertor) LeerDesertores(nombreArchivo string) bool {
	var TP TareasPendientes

	archivo, err := os.Open("./tmp/" + nombreArchivo)
	if err != nil {
		fmt.Println("Error leyendo el archivo")
		return false
	}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	condicion := make(map[string]interface{})

	TP.Codigo = "DTS-" + time.Now().String()[:19]
	TP.Estatus = 0
	TP.FechaInicio = time.Now()
	TP.Observacion = "Carga de desertores"
	cpendiente := sys.MGOSession.DB(sys.CBASE).C("tareaspendientes")
	cpendiente.Insert(TP)
	fmt.Println("Proceso de carga de desertores Inicializado")

	f, err := os.Create("public_web/SSSIFANB/panel/tmp/" + TP.Codigo + ".log")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	i := 0
	scan := bufio.NewScanner(archivo)
	for scan.Scan() {
		i++
		linea := strings.Split(scan.Text(), ";")
		condicion["condicion"] = 1
		err = c.Update(bson.M{"id": linea[0]}, bson.M{"$set": condicion})
		if err != nil {
			log := "Error linea  # " + strconv.Itoa(i) + ": " + scan.Text() + "\n"
			_, e := f.WriteString(log)
			if e != nil {
				fmt.Println("Error en la linea...")
			}
		}
		//fmt.Println(linea[0])
	}

	tarea := make(map[string]interface{})
	tarea["estatus"] = 1
	tarea["tipo"] = "LOG"
	tarea["fechafin"] = time.Now()
	err = cpendiente.Update(bson.M{"codigo": TP.Codigo}, bson.M{"$set": tarea})
	if err != nil {
		fmt.Println("Error al finalizar la tarea pendiente")
	}
	fmt.Println("Proceso de carga de desertores finalizado")
	return true
}
