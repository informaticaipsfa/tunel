//motor de estadistica
package estadistica

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

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
		var situ, clas string
		var ccod, cnom, csig, gcod, gnom, gdes string
		var fnac, fing, fult, fegr, cate sql.NullString
		var nro int

		err = sq.Scan(&cedulas, &cedulap, &nro, &nac, &nombp, &nombs, &apellp, &apells,
			&fnac,
			&sexo, &edoc, &cate, &situ, &clas, &fing, &fult, &fegr, &anor, &mesr, &diar,
			&ccod, &cnom, &csig, &gcod, &gid, &gnom, &gdes)
		if err != nil {
			fmt.Println(err.Error())
		}
		cedulapace = util.ValidarNullString(cedulap)
		militar.TipoDato = 0
		militar.Situacion = situ
		militar.Categoria = util.ValidarNullString(cate)
		militar.Persona.DatoBasico.Cedula = cedulas
		militar.Persona.DatoBasico.NumeroPersona = nro
		militar.Persona.DatoBasico.Nacionalidad = nac
		militar.Persona.DatoBasico.NombrePrimero = util.ValidarNullString(nombp)
		militar.Persona.DatoBasico.NombreSegundo = util.ValidarNullString(nombs)
		militar.Persona.DatoBasico.ApellidoPrimero = util.ValidarNullString(apellp)
		militar.Persona.DatoBasico.ApellidoSegundo = util.ValidarNullString(apells)
		militar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		militar.Persona.DatoBasico.EstadoCivil = util.ValidarNullString(edoc)
		//militar.Persona.DatoBasico.FechaNacimiento = strings.Replace(util.ValidarNullString(fnac), "/", "-", -1)

		fechanacimiento := util.ValidarNullString(fnac)
		if fechanacimiento != "null" {
			dateString := strings.Replace(fechanacimiento, "/", "-", -1)
			layOut := "2006-01-02"
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				militar.Persona.DatoBasico.FechaNacimiento = dateStamp.Local()
			}
		}

		fechaingreso := util.ValidarNullString(fing)
		if fechaingreso != "null" {
			dateString := strings.Replace(fechaingreso, "/", "-", -1)
			layOut := "2006-01-02"
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				militar.FechaIngresoComponente = dateStamp.Local()
			}
		}

		fechaultimo := util.ValidarNullString(fult)
		if fechaultimo != "null" {
			dateString := strings.Replace(fechaultimo, "/", "-", -1)
			layOut := "2006-01-02"
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				militar.FechaAscenso = dateStamp.Local()
			}
		}

		//militar.FechaIngresoComponente =
		militar.Grado.Nombre = util.ValidarNullString(gid)
		militar.Grado.Descripcion = gdes
		militar.Componente.Abreviatura = ccod
		militar.Componente.Nombre = cnom
		militar.AppNomina = false
		militar.AppSaman = true
		militar.AppPace = true
		if cedulapace == "null" {
			militar.AppPace = false
			fmt.Println(cedulas, cedulapace, nro, util.ValidarNullString(nombp), situ, fing, gnom)
		}

		militar.SalvarMGO("militar")
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

//CargarMilitar Historial militar
func (e *Estructura) CargarMilitar() (jSon []byte, err error) {
	var msj Mensaje
	var cedula string
	sq, err := sys.PostgreSQLSAMAN.Query(obtenerHistorialMilitar())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	// i := 0
	// miliares := 1
	// var Familiares []interface{}
	for sq.Next() {
		sq.Scan(&cedula)
	}

	return
}

//CargarFamiliar Actualizar a los familiares
func (e *Estructura) CargarFamiliar() (jSon []byte, err error) {
	var msj Mensaje
	var cedula string
	sq, err := sys.PostgreSQLSAMAN.Query(obtenerHistorialFamiliares())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	miliares := 1
	var Familiares []interface{}

	for sq.Next() {
		var cedulaAux, codnip, fecha string
		var familiar sssifanb.Familiar
		var nro int
		var paren, nac, nombp, nombs, apelp, apels, sexo, edoc, fech, nmil sql.NullString
		err = sq.Scan(&cedulaAux, &codnip, &nro, &paren, &nac, &nombp, &nombs, &apelp, &apels, &fech, &sexo, &edoc, &nmil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if i == 0 {
			cedula = cedulaAux
		}
		if cedula != cedulaAux {
			fmt.Println("OTRO --- " + cedula)
			var Militar sssifanb.Militar

			fm := make(map[string]interface{})
			fm["familiar"] = Familiares
			Militar.ActualizarMGO(cedula, fm)
			miliares++
			cedula = cedulaAux
			Familiares = nil
		}
		familiar.Parentesco = util.ValidarNullString(paren)
		familiar.Persona.DatoBasico.Cedula = codnip
		familiar.Persona.DatoBasico.NumeroPersona = nro
		familiar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		familiar.Persona.DatoBasico.Nacionalidad = util.ValidarNullString(nac)

		familiar.Persona.DatoBasico.NombrePrimero = strings.ToUpper(util.ValidarNullString(nombp))
		familiar.Persona.DatoBasico.NombreSegundo = strings.ToUpper(util.ValidarNullString(nombs))
		familiar.Persona.DatoBasico.ApellidoPrimero = strings.ToUpper(util.ValidarNullString(apelp))
		familiar.Persona.DatoBasico.ApellidoSegundo = strings.ToUpper(util.ValidarNullString(apels))
		if util.ValidarNullString(nmil) != "null" {
			familiar.EsMilitar = true
		}
		fecha = util.ValidarNullString(fech)
		if fecha != "null" {
			dateString := strings.Replace(fecha, "/", "-", -1)
			layOut := "2006-01-02"
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				familiar.Persona.DatoBasico.FechaNacimiento = dateStamp.Local()
			}
		}

		familiar.AplicarReglasBeneficio()
		//familiar.DocumentoPadre = cedula
		//fmt.Println(familiar)
		// var Militar sssifanb.Militar
		//var Fam map[string]interface{}
		// fm := make(map[string]interface{})
		// fm["familiar"] = familiar
		// Militar.ActualizarMGO(cedula, fm)

		Familiares = append(Familiares, familiar)
		i++
	}
	var Militar sssifanb.Militar

	fm := make(map[string]interface{})
	fm["familiar"] = Familiares
	Militar.ActualizarMGO(cedula, fm)

	fmt.Println("CANTIDAD: REGISTROS: ", i, " MILITARES: ", miliares)
	return
}

func errorG(sSQL string) {
	_, err := sys.PostgreSQLSAMAN.Exec(sSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
}
