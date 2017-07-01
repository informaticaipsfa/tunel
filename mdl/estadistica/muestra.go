//motor de estadistica
package estadistica

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
)

//Estructura Generacion de Grupos de Datos
type Estructura struct {
	Cedula       string  `json:"cedula,omitempty"` //Identificador
	Nropersona   int     `json:"nropersona,omitempty"`
	Prioridad    int     `json:"prioridad,omitempty"` //Inferencia
	Peso         float64 `json:"peso,omitempty"`
	Caracteres   float64 `json:"caracteres,omitempty"`
	Varianza     int     `json:"varianza,omitempty"`
	Familiares   int     `json:"familiares,omitempty"` //dato que genera la varianza
	Militares    int     `json:"militares,omitempty"`
	Historiales  int     `json:"historiales,omitempty"`
	Creditos     int     `json:"creditos,omitempty"`
	Reembolsos   int     `json:"reembolsos,omitempty"`
	Medicamentos int     `json:"medicamentos,omitempty"`
	Tratamientos int     `json:"tratamientos,omitempty"`
}

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj,omitempty"`
	Tipo    int    `json:"tipo,omitempty"`
	Pgsql   string `json:"pgsql,omitempty"`
}

//Reduccion Evaluación y simplificación
func (e *Estructura) Reduccion() (jSon []byte, err error) {
	personaMilitar()
	historialFamiliares()
	historialCreditos()
	historialPension()
	historialMilitares()
	historialReembolso()
	return
}

//Migracion Relaciones
func (e *Estructura) Migracion() (jSon []byte, err error) {
	var msj Mensaje
	sq, err := sys.PostgreSQLSAMAN.Query(reduccion())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	// i := 0
	for sq.Next() {
		var militar sssifanb.Militar
		var cedulap, anor, mesr, diar, gid, edoc, sexo sql.NullString

		var cedulas, cedulapace, nac string
		var nombp, nombs, apellp, apells sql.NullString
		var cate, situ, clas string
		var ccod, cnom, csig, gcod, gnom, gdes string
		var fing, fult, fegr sql.NullString
		var nro int

		err = sq.Scan(&cedulas, &cedulap, &nro, &nac, &nombp, &nombs, &apellp, &apells,
			&sexo, &edoc, &cate, &situ, &clas, &fing, &fult, &fegr, &anor, &mesr, &diar,
			&ccod, &cnom, &csig, &gcod, &gid, &gnom, &gdes)
		if err != nil {
			fmt.Println(err.Error())
		}
		cedulapace = util.ValidarNullString(cedulap)
		militar.TipoDato = 0
		militar.Situacion = situ
		militar.Persona.DatoBasico.Cedula = cedulas
		militar.Persona.DatoBasico.NumeroPersona = nro
		militar.Persona.DatoBasico.Nacionalidad = nac
		militar.Persona.DatoBasico.NombrePrimero = util.ValidarNullString(nombp)
		militar.Persona.DatoBasico.NombreSegundo = util.ValidarNullString(nombs)
		militar.Persona.DatoBasico.ApellidoPrimero = util.ValidarNullString(apellp)
		militar.Persona.DatoBasico.ApellidoSegundo = util.ValidarNullString(apells)
		militar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)

		militar.Grado.Nombre = util.ValidarNullString(gid)
		militar.Grado.Descripcion = gdes
		militar.Componente.Abreviatura = ccod
		militar.Componente.Nombre = cnom
		//militar.SalvarMGO("militar")

		if cedulapace == "null" {
			fmt.Println(cedulas, cedulapace, nro, util.ValidarNullString(nombp), situ, fing, gnom)
		}

		jSon, err = json.Marshal(militar)
	}
	fmt.Println("Finalizo el proceso...")

	return
}

//EstatusMilitar Situacion de un militar
func EstatusMilitar(valor string) (estatus int) {

	switch valor {
	case "RSP":
		estatus = 0
		break
	}
	return
}

//ActualizarFamiliar Actualizar a los familiares
func (e *Estructura) ActualizarFamiliar() (jSon []byte, err error) {
	var msj Mensaje
	var cedula string
	sq, err := sys.PostgreSQLSAMAN.Query(obtenerFamiliares())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	var Familiares []interface{}

	for sq.Next() {
		var cedulaAux, codnip string
		var familiar sssifanb.Familiar
		var nro int
		var paren, nac, nombp, nombs, apelp, apels, sexo, edoc, fech, nmil sql.NullString
		err = sq.Scan(&cedulaAux, &codnip, &nro, &paren, &nac, &nombp, &nombs, &apelp, &apels, &sexo, &edoc, &fech, &nmil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if i == 0 {
			cedula = cedulaAux
		}
		if cedula != cedulaAux {
			//fmt.Println(bson.M{"familiar": Familiares})
			//fmt.Println("OTRO --- ")
			fmt.Println("CEDULA: ", cedula, familiar.Persona.DatoBasico.Cedula, familiar.Persona.DatoBasico.NombrePrimero)
			cedula = cedulaAux
			Familiares = nil
		}
		familiar.Parentesco = util.ValidarNullString(paren)
		familiar.Persona.DatoBasico.Cedula = codnip
		familiar.Persona.DatoBasico.NumeroPersona = nro
		familiar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		familiar.Persona.DatoBasico.Nacionalidad = util.ValidarNullString(nac)
		familiar.Persona.DatoBasico.FechaNacimiento = util.ValidarNullString(fech)

		familiar.Persona.DatoBasico.NombrePrimero = strings.ToUpper(util.ValidarNullString(nombp))
		familiar.Persona.DatoBasico.NombreSegundo = strings.ToUpper(util.ValidarNullString(nombs))
		familiar.Persona.DatoBasico.ApellidoPrimero = strings.ToUpper(util.ValidarNullString(apelp))
		familiar.Persona.DatoBasico.ApellidoSegundo = strings.ToUpper(util.ValidarNullString(apels))
		if util.ValidarNullString(nmil) != "null" {
			familiar.EsMilitar = true
		}
		if familiar.Persona.DatoBasico.Sexo != "M" {
			familiar.Parentesco = "MD"
		}
		familiar.AplicarReglasBeneficio()
		familiar.DocumentoPadre = cedula
		fmt.Println(familiar)
		Familiares = append(Familiares, familiar)
		i++
	}
	//fmt.Println("CEDULA: ", cedula, familiar.Persona.DatoBasico.Cedula, familiar.Persona.DatoBasico.NombrePrimero)
	//fmt.Println(bson.M{"familiar": Familiares})

	return
}

func errorG(sSQL string) {
	_, err := sys.PostgreSQLSAMAN.Exec(sSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
}
