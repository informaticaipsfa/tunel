package sssifanb

import (
	"fmt"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

var Estados []fanb.Estado

//ExportarMysql Exportar Full Text
func ExportarMysql() {
	var TP TareasPendientes

	var militar []Militar
	var Estado fanb.Estado

	TP.Codigo = "MYSQL-" + time.Now().String()[:19]
	TP.Estatus = 0
	TP.FechaInicio = time.Now()
	TP.Observacion = "Extrayendo datos MYSQL - FULLTEXT"
	cpendiente := sys.MGOSession.DB(sys.CBASE).C("tareaspendientes")
	cpendiente.Insert(TP)

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	seleccion := bson.M{
		"categoria":                   true,
		"situacion":                   true,
		"clase":                       true,
		"grado":                       true,
		"componente":                  true,
		"persona.telefono":            true,
		"persona.datobasico":          true,
		"persona.direccion":           true,
		"familiar.persona.datobasico": true,
		"familiar.parentesco":         true,
		"familiar.esmilitar":          true,
	}

	Estados = Estado.ConsultarTodo()
	// err := c.Find(bson.M{"id": "10107698"}).Select(seleccion).All(&militar)
	err := c.Find(bson.M{}).Select(seleccion).All(&militar)
	if err != nil {
		fmt.Println(err.Error())
	}
	//repetidos := 0
	//creduccion := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	for _, mil := range militar {
		var cedu, nomb, dire, grad, comp, situ, clas, cate, telf, fami string
		cedu = mil.Persona.DatoBasico.Cedula
		nomb = mil.Persona.DatoBasico.ConcatenarApellidoNombre()

		if len(mil.Persona.Direccion) > 0 {
			dire = obtenerEstado(mil.Persona.Direccion[0].Estado) + " " + mil.Persona.Direccion[0].Ciudad +
				" " + mil.Persona.Direccion[0].Municipio + " " + mil.Persona.Direccion[0].Parroquia +
				" " + mil.Persona.Direccion[0].CalleAvenida + " " + mil.Persona.Direccion[0].Casa +
				" " + strconv.Itoa(mil.Persona.Direccion[0].Numero)
		} else {
			dire = ""
		}

		grad = mil.Grado.Descripcion
		comp = mil.Componente.Descripcion
		situ = obtenerSitiacion(mil.Situacion)
		clas = obtenerClase(mil.Clase)
		cate = obtenerCategoria(mil.Categoria)
		telf = mil.Persona.Telefono.Domiciliario + " " + mil.Persona.Telefono.Movil
		dire += telf + " " +
			obtenerEstadoCivil(mil.Persona.DatoBasico.EstadoCivil) + " " +
			obtenerSexo(mil.Persona.DatoBasico.Sexo)
		for _, familiar := range mil.Familiar {
			var direr string
			if len(familiar.Persona.Direccion) > 0 {
				direr = obtenerEstado(familiar.Persona.Direccion[0].Estado) + " " + familiar.Persona.Direccion[0].Ciudad +
					" " + familiar.Persona.Direccion[0].Municipio + " " + familiar.Persona.Direccion[0].Parroquia +
					" " + familiar.Persona.Direccion[0].CalleAvenida + " " + familiar.Persona.Direccion[0].Casa +
					" " + strconv.Itoa(familiar.Persona.Direccion[0].Numero)
			}
			fami += " | " + obtenerParentesco(familiar.Parentesco, familiar.Persona.DatoBasico.Sexo) + " " +
				familiar.Persona.DatoBasico.Cedula + " " + familiar.Persona.DatoBasico.ConcatenarApellidoNombre() + " " +
				obtenerEstadoCivil(familiar.Persona.DatoBasico.EstadoCivil) + " " +
				obtenerSexo(familiar.Persona.DatoBasico.Sexo) + " " + direr
		}

		body := grad + " " + comp + " " + situ + " " + clas + " " + cate
		insert := `INSERT INTO datos ( cedula, nombre, descripcion, direccion, familiares ) 
		VALUES ('` + cedu + `','` + nomb + `','` + body + `','` + dire + `','` + fami + `')`
		// fmt.Println(insert)
		sys.MysqlFullText.Exec(insert)
	}
	tarea := make(map[string]interface{})
	tarea["estatus"] = 1
	tarea["fechafin"] = time.Now()
	err = cpendiente.Update(bson.M{"codigo": TP.Codigo}, bson.M{"$set": tarea})
	if err != nil {
		fmt.Println("Error al finalizar la tarea pendiente mysql")
	}
	fmt.Println("Proceso MYSQL-FULLTEXT finalizado.")
}

func obtenerSitiacion(situacion string) (situ string) {
	switch situacion {
	case "ACT":
		situ = "ACTIVO"
		break
	case "RCP":
		situ = "RETIRADO CON PENSION"
		break
	case "RSP":
		situ = "RETIRADO SIN PENSION"
		break
	case "FCP":
		situ = "FALLECIDO CON PENSION"
		break
	default:
		break
	}
	return
}

func obtenerClase(clase string) (clas string) {
	switch clase {
	case "TPROF":
		clas = "TROPA PROFESIONAL"
		break
	case "OFI":
		clas = "OFICIAL"
		break
	case "OFIT":
		clas = "OFICIAL TECNICO"
		break
	case "SUBOFI":
		clas = "SUBOFICIAL"
		break
	case "OFITR":
		clas = "OFICIAL TROPA"
		break
	default:
		break
	}
	return
}

func obtenerCategoria(categoria string) (cate string) {
	switch categoria {
	case "EFE":
		cate = "EFECTIVO"
		break
	case "TRP":
		cate = "TROPA"
		break
	case "RES":
		cate = "RESERVA"
		break
	case "ASI":
		cate = "ASIMILADO"
		break
	default:
		break
	}
	return
}

func obtenerParentesco(parentesco string, sexo string) (pare string) {
	switch parentesco {
	case "HJ":
		pare = "HIJO"
		if sexo == "F" {
			pare = "HIJA"
		}
		break
	case "PD":
		pare = "PADRE"
		if sexo == "F" {
			pare = "MADRE"
		}
		break
	case "EA":
		pare = "ESPOS0"
		if sexo == "F" {
			pare = "ESPOSA"
		}
		break
	case "VI":
		pare = "VIUDO"
		if sexo == "F" {
			pare = "VIUDA"
		}
		break
	default:
		break
	}
	return
}

func obtenerEstado(codigo string) string {
	var nombre string
	for _, v := range Estados {
		if v.Codigo == codigo {
			nombre = v.Nombre
		}
	}
	return nombre
}

func obtenerSexo(sexo string) (sex string) {
	switch sexo {
	case "F":
		sex = "FEMENINO"
		break
	case "M":
		sex = "MASCULINO"
		break
	default:
		break
	}
	return
}

func obtenerEstadoCivil(estado string) (civil string) {
	switch estado {
	case "S":
		civil = "SOLTERO"
		break
	case "C":
		civil = "CASADO"
		break
	case "D":
		civil = "DIVORSIADO"
		break
	case "V":
		civil = "VIUDO"
		break
	default:
		break
	}
	return
}
