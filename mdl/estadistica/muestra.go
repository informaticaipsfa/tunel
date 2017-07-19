//motor de estadistica
package estadistica

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/fanb"
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
		var grp, cmp, tipc, inst, cuenta, nrohp sql.NullString     //Datos Financieros de Pension
		var fechf, cuentaf, nrohf, anof, mesf, diaf sql.NullString //Datos Financieros de Pension
		var nro int

		err = sq.Scan(&cedulas, &cedulap, &nro, &nac, &nombp, &nombs, &apellp, &apells, &fnac,
			&sexo, &edoc, &cate, &situ, &clas, &fing, &fult, &fegr, &anor, &mesr, &diar,
			&ccod, &cnom, &csig, &gcod, &gid, &gnom, &gdes, &grp, &cmp, &tipc, &inst, &cuenta, &nrohp,
			&anof, &mesf, &diaf, &fechf, &nrohf, &cuentaf)
		if err != nil {
			fmt.Println(err.Error())
		}
		cedulapace = util.ValidarNullString(cedulap)

		militar.ID = cedulas
		militar.TipoDato = 0
		militar.Situacion = situ
		militar.Clase = clas
		militar.Categoria = util.ValidarNullString(cate)

		militar.Persona.DatoBasico.Cedula = cedulas
		militar.Persona.DatoBasico.NumeroPersona = nro
		militar.Persona.DatoBasico.Nacionalidad = nac
		militar.Persona.DatoBasico.NombrePrimero = strings.ToUpper(util.ValidarNullString(nombp))
		militar.Persona.DatoBasico.NombreSegundo = strings.ToUpper(util.ValidarNullString(nombs))
		militar.Persona.DatoBasico.ApellidoPrimero = strings.ToUpper(util.ValidarNullString(apellp))
		militar.Persona.DatoBasico.ApellidoSegundo = strings.ToUpper(util.ValidarNullString(apells))
		militar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		militar.Persona.DatoBasico.EstadoCivil = util.ValidarNullString(edoc)
		if util.ValidarNullString(grp) != "null" {
			militar.Pension.GradoCodigo = util.ValidarNullString(grp)
			militar.Pension.ComponenteCodigo = util.ValidarNullString(cmp)
			militar.Pension.DatoFinanciero.Cuenta = util.ValidarNullString(cuenta)
			militar.Pension.DatoFinanciero.TipoCuenta = util.ValidarNullString(tipc)
			militar.Pension.DatoFinanciero.Institucion = util.ValidarNullString(inst)
			militar.Pension.NumeroHijos, _ = strconv.Atoi(util.ValidarNullString(nrohp))
		}

		layOut := "2006-01-02"
		fechanacimiento := util.ValidarNullString(fnac)
		if fechanacimiento != "null" {
			dateString := strings.Replace(fechanacimiento, "/", "-", -1)
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				militar.Persona.DatoBasico.FechaNacimiento = dateStamp.UTC()
			}

		}

		fechaingreso := util.ValidarNullString(fing)
		if fechaingreso != "null" {
			dateString := strings.Replace(fechaingreso, "/", "-", -1)
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				militar.FechaIngresoComponente = dateStamp.UTC()
			}
		}

		fechaultimo := util.ValidarNullString(fult)
		if fechaultimo != "null" {
			dateString := strings.Replace(fechaultimo, "/", "-", -1)
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				militar.FechaAscenso = dateStamp.UTC()
			}
		}

		militar.AppNomina = false
		militar.AppSaman = true
		militar.AppPace = true
		fecharetiro := util.ValidarNullString(fegr)
		if fechaultimo != "null" {
			dateString := strings.Replace(fecharetiro, "/", "-", -1)
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				militar.AppNomina = true
				militar.FechaRetiro = dateStamp.UTC()
			}
		}

		militar.Grado.Nombre = strings.ToUpper(util.ValidarNullString(gid))
		militar.Grado.Descripcion = strings.ToUpper(gdes)
		militar.Grado.Abreviatura = strings.ToUpper(gcod)
		militar.Componente.Abreviatura = strings.ToUpper(ccod)
		militar.Componente.Nombre = strings.ToUpper(csig)
		militar.Componente.Descripcion = strings.ToUpper(cnom)

		militar.AnoReconocido, _ = strconv.Atoi(util.ValidarNullString(anor))
		militar.MesReconocido, _ = strconv.Atoi(util.ValidarNullString(mesr))
		militar.DiaReconocido, _ = strconv.Atoi(util.ValidarNullString(diar))

		if cedulapace == "null" {
			militar.AppPace = false
			fmt.Println(cedulas, cedulapace, nro, util.ValidarNullString(nombp), situ, fnac, "->", militar.Persona.DatoBasico.FechaNacimiento, fing, gnom)
		} else {
			militar.Fideicomiso.AnoReconocido, _ = strconv.Atoi(util.ValidarNullString(anof))
			militar.Fideicomiso.MesReconocido, _ = strconv.Atoi(util.ValidarNullString(mesf))
			militar.Fideicomiso.DiaReconocido, _ = strconv.Atoi(util.ValidarNullString(diaf))
			militar.Fideicomiso.NumeroHijos, _ = strconv.Atoi(util.ValidarNullString(nrohf))
			militar.Fideicomiso.CuentaBancaria = util.ValidarNullString(cuentaf)
		}

		militar.Anomalia.Hijo = false
		militar.Anomalia.Ano = false
		militar.Anomalia.Mes = false
		militar.Anomalia.Dia = false

		if militar.Pension.NumeroHijos != militar.Fideicomiso.NumeroHijos {
			militar.Anomalia.Hijo = true
		}
		if militar.AnoReconocido != militar.Fideicomiso.AnoReconocido {
			militar.Anomalia.Ano = true
		}
		if militar.MesReconocido != militar.Fideicomiso.MesReconocido {
			militar.Anomalia.Mes = true
		}
		if militar.DiaReconocido != militar.Fideicomiso.DiaReconocido {
			militar.Anomalia.Dia = true
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
	i := 0
	miliares := 1
	var Historial []interface{}
	for sq.Next() {
		var cedulaAux, cmp, grd, cat, cla, sit, fob, gresu, nresu string
		var frec, fcam, hcam, fcre, hcre, razo sql.NullString
		var historialmilitar sssifanb.HistorialMilitar
		err = sq.Scan(&cedulaAux, &cmp, &grd, &cat, &cla, &sit, &fob,
			&gresu, &nresu, &frec, &fcam, &hcam, &fcre, &hcre, &razo)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if i == 0 {
			cedula = cedulaAux
		}
		if cedula != cedulaAux {
			fmt.Println("# " + strconv.Itoa(i) + " ID_: " + cedula)
			var Militar sssifanb.Militar

			fm := make(map[string]interface{})
			fm["historialmilitar"] = Historial
			Militar.ActualizarMGO(cedula, fm)
			miliares++
			cedula = cedulaAux
			Historial = nil
		}
		historialmilitar.Grado = grd
		historialmilitar.Componente = cmp
		historialmilitar.Categoria = cat
		historialmilitar.Clase = cla
		historialmilitar.Situacion = sit
		historialmilitar.GradoResuelto = gresu
		historialmilitar.NumeroResuelto = nresu
		historialmilitar.FechaCambio = util.ValidarNullString(fcam)
		historialmilitar.HoraCambio = util.ValidarNullString(hcam)
		historialmilitar.FechaCreacion = util.ValidarNullString(fcre)
		historialmilitar.HoraCreacion = util.ValidarNullString(hcre)
		historialmilitar.Razon = util.ValidarNullString(razo)

		Historial = append(Historial, historialmilitar)
		i++
	}
	var Militar sssifanb.Militar

	fm := make(map[string]interface{})
	fm["historialmilitar"] = Historial
	Militar.ActualizarMGO(cedula, fm)

	fmt.Println("CANTIDAD: REGISTROS: ", i, " MILITARES: ", miliares)
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
		layOut := "2006-01-02"
		fecha = util.ValidarNullString(fech)
		if fecha != "null" {
			dateString := strings.Replace(fecha, "/", "-", -1)
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				familiar.Persona.DatoBasico.FechaNacimiento = dateStamp
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

func (e *Estructura) CargarCtaBancaria() (jSon []byte, err error) {
	var msj Mensaje
	sq, err := sys.PostgreSQLSAMAN.Query(obtenerCuentaBancaria())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}

	for sq.Next() {
		var cedula string
		var cuenta, tipo, institucion sql.NullString
		sq.Scan(&cedula, &cuenta, &tipo, &institucion)
		var Militar sssifanb.Militar
		var Historial sssifanb.DatoFinanciero

		Historial.Cuenta = util.ValidarNullString(cuenta)
		Historial.Institucion = util.ValidarNullString(institucion)
		Historial.TipoCuenta = util.ValidarNullString(tipo)

		fm := make(map[string]interface{})
		fm["persona.datofinanciero"] = Historial
		Militar.ActualizarMGO(cedula, fm)
		fmt.Println(cedula)
	}
	return
}

//CargarMilitar Historial militar
func (e *Estructura) CargarComponenteGrado() (jSon []byte, err error) {
	var msj Mensaje
	var codigo string

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerComponenteGrado())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	var lstGrado []fanb.Grado
	var componente fanb.Componente
	for sq.Next() {
		//c.componentecod, componentenombre, componentesiglas, gradocod,gradocodrangoid,gradonombrecorto,
		//gradonombrelargo
		var ccod, cnom, csig, gcod, gran, gnom, gdes string
		var grado fanb.Grado
		err = sq.Scan(&ccod, &cnom, &csig, &gcod, &gran, &gnom, &gdes)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if i == 0 {
			codigo = ccod
			componente.Codigo = codigo
			componente.Nombre = cnom
			componente.Siglas = csig
		}
		if codigo != ccod {
			fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
			componente.Grado = lstGrado
			componente.SalvarMGO("componente")
			lstGrado = nil
			codigo = ccod
			componente.Codigo = codigo
			componente.Nombre = cnom
			componente.Siglas = csig
		}

		grado.Codigo = gcod
		grado.Rango = gran
		grado.Nombre = gnom
		grado.Descripcion = gdes
		lstGrado = append(lstGrado, grado)
		i++
	}
	componente.Grado = lstGrado
	componente.SalvarMGO("componente")
	fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
	return
}

//CargarMilitar Historial militar
func (e *Estructura) CargarEstados() (jSon []byte, err error) {
	var msj Mensaje
	var codigo string

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerEstados())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	var lstCiudad []fanb.Ciudad
	var estado fanb.Estado
	for sq.Next() {

		var cod, iso, ciud string
		var capi int
		var ciudad fanb.Ciudad
		err = sq.Scan(&cod, &iso, &ciud, &capi)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if i == 0 {
			codigo = iso
			estado.Codigo = iso
			estado.Nombre = cod
		}
		if codigo != cod {
			fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
			estado.Ciudad = lstCiudad
			estado.SalvarMGO("estado")
			estado.Codigo = iso
			estado.Nombre = codigo
			lstCiudad = nil
			codigo = iso

		}

		ciudad.Capital = capi
		ciudad.Nombre = ciud
		lstCiudad = append(lstCiudad, ciudad)
		i++
	}
	estado.Ciudad = lstCiudad
	estado.SalvarMGO("estado")
	fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
	return
}

//CargarMilitar Historial militar
func (e *Estructura) CargarMunicipio() (jSon []byte, err error) {
	var msj Mensaje
	var codigo string

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerMunicipiosParroquia())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	var Municipio fanb.Municipio
	var lstParroquia []string
	var estado fanb.Estado
	for sq.Next() {
		var cod, mun, parr string

		err = sq.Scan(&cod, &mun, &parr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if i == 0 {
			codigo = mun
			estado.Nombre = cod
		}
		if codigo != mun {
			Municipio.Nombre = codigo
			Municipio.Parroquia = lstParroquia
			fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
			fm := make(map[string]interface{})
			fm["municipio"] = Municipio
			estado.ActualizarMGO(estado.Nombre, fm)
			estado.Nombre = cod
			lstParroquia = nil
			codigo = mun
		}
		lstParroquia = append(lstParroquia, parr)
		i++
	}
	Municipio.Nombre = codigo
	Municipio.Parroquia = lstParroquia
	fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
	fm := make(map[string]interface{})
	fm["municipio"] = Municipio
	estado.ActualizarMGO(codigo, fm)
	return
}

func errorG(sSQL string) {
	_, err := sys.PostgreSQLSAMAN.Exec(sSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
}
