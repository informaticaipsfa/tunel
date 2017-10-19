//motor de estadistica
package estadistica

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/sys/seguridad"
	"github.com/informaticaipsfa/tunel/util"
)

//VNULL Validar
const VNULL string = "null"

// Estructura Generacion de Grupos de Datos
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
		var pensioncategoria, pensionsituacion, pensionclase sql.NullString
		var annototservicio, mestotservicio, diatotservicio sql.NullString
		var pensionasignada, pensionporc sql.NullFloat64
		var pensionfpromo, pensionfultimo, compid, grdid sql.NullString
		err = sq.Scan(&cedulas, &cedulap, &nro, &nac, &nombp, &nombs, &apellp, &apells, &fnac,
			&sexo, &edoc, &cate, &situ, &clas, &fing, &fult, &fegr, &anor, &mesr, &diar,
			&ccod, &cnom, &csig, &gcod, &gid, &gnom, &gdes, &grp, &cmp, &tipc, &inst, &cuenta, &nrohp,
			&anof, &mesf, &diaf, &fechf, &nrohf, &cuentaf, &pensioncategoria, &pensionsituacion, &pensionclase,
			&annototservicio, &mestotservicio, &diatotservicio, &pensionasignada, &pensionporc, &pensionfpromo,
			&pensionfultimo, &compid, &grdid)
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
		militar.Persona.DatoBasico.NroPersona = nro
		militar.Persona.DatoBasico.Nacionalidad = nac
		militar.Persona.DatoBasico.NombrePrimero = strings.ToUpper(util.ValidarNullString(nombp)) + " " + strings.ToUpper(util.ValidarNullString(nombs))
		militar.Persona.DatoBasico.ApellidoPrimero = strings.ToUpper(util.ValidarNullString(apellp)) + " " + strings.ToUpper(util.ValidarNullString(apells))
		militar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		militar.Persona.DatoBasico.EstadoCivil = util.ValidarNullString(edoc)
		if util.ValidarNullString(grp) != VNULL {
			militar.Pension.GradoCodigo = util.ValidarNullString(grp)
			militar.Pension.ComponenteCodigo = util.ValidarNullString(cmp)
			militar.Pension.Categoria = util.ValidarNullString(pensioncategoria)
			militar.Pension.Clase = util.ValidarNullString(pensionclase)
			militar.Pension.Situacion = util.ValidarNullString(pensionsituacion)

			militar.Pension.AnoServicio, _ = strconv.Atoi(util.ValidarNullString(annototservicio))
			militar.Pension.MesServicio, _ = strconv.Atoi(util.ValidarNullString(mestotservicio))
			militar.Pension.DiaServicio, _ = strconv.Atoi(util.ValidarNullString(diatotservicio))
			militar.Pension.PensionAsignada = util.ValidarNullFloat64(pensionasignada)
			militar.Pension.PorcentajePrestaciones = util.ValidarNullFloat64(pensionporc)
			militar.Pension.FechaPromocion = util.ValidarNullString(pensionfpromo)
			militar.Pension.FechaUltimoAscenso = util.ValidarNullString(pensionfultimo)

			militar.Pension.DatoFinanciero.Cuenta = util.ValidarNullString(cuenta)
			militar.Pension.DatoFinanciero.Tipo = util.ValidarNullString(tipc)
			militar.Pension.DatoFinanciero.Institucion = util.ValidarNullString(inst)
			militar.Pension.NumeroHijos, _ = strconv.Atoi(util.ValidarNullString(nrohp))
		}

		layOut := "2006-01-02"
		fechanacimiento := util.ValidarNullString(fnac)
		if fechanacimiento != VNULL {
			dateString := strings.Replace(fechanacimiento, "/", "-", -1)
			dateStamp, er := time.Parse(layOut, dateString)
			if er == nil {
				militar.Persona.DatoBasico.FechaNacimiento = dateStamp
			}

		}

		fechaingreso := util.ValidarNullString(fing)
		if fechaingreso != VNULL {
			dateString := strings.Replace(fechaingreso, "/", "-", -1)
			dateStamp, er := time.Parse(layOut, dateString)
			if er == nil {
				militar.FechaIngresoComponente = dateStamp
			}
		}

		fechaultimo := util.ValidarNullString(fult)
		if fechaultimo != VNULL {
			dateString := strings.Replace(fechaultimo, "/", "-", -1)
			dateStamp, er := time.Parse(layOut, dateString)
			if er == nil {
				militar.FechaAscenso = dateStamp
			}
		}

		militar.AppNomina = false
		militar.AppSaman = true
		militar.AppPace = true
		fecharetiro := util.ValidarNullString(fegr)
		if fechaultimo != VNULL {
			dateString := strings.Replace(fecharetiro, "/", "-", -1)
			dateStamp, er := time.Parse(layOut, dateString)
			if er == nil {
				militar.AppNomina = true
				militar.FechaRetiro = dateStamp
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

		if cedulapace == VNULL {
			militar.AppPace = false
			fmt.Println(cedulas, cedulapace, nro, util.ValidarNullString(nombp), situ, fnac, "->", militar.Persona.DatoBasico.FechaNacimiento, fing, gnom)
		} else {
			militar.Fideicomiso.GradoCodigo, _ = strconv.Atoi(util.ValidarNullString(grdid))
			militar.Fideicomiso.ComponenteCodigo, _ = strconv.Atoi(util.ValidarNullString(compid))
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

		militar.SalvarMGO()
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
		familiar.Persona.DatoBasico.NroPersona = nro
		familiar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		familiar.Persona.DatoBasico.Nacionalidad = util.ValidarNullString(nac)

		familiar.Persona.DatoBasico.NombrePrimero = strings.ToUpper(util.ValidarNullString(nombp)) + " " + strings.ToUpper(util.ValidarNullString(nombs))
		familiar.Persona.DatoBasico.ApellidoPrimero = strings.ToUpper(util.ValidarNullString(apelp)) + " " + strings.ToUpper(util.ValidarNullString(apels))
		if util.ValidarNullString(nmil) != VNULL {
			familiar.EsMilitar = true
		}
		layOut := "2006-01-02"
		fecha = util.ValidarNullString(fech)
		if fecha != VNULL {
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

//CargarCtaBancaria Cuenta
func (e *Estructura) CargarCtaBancaria() (jSon []byte, err error) {
	var msj Mensaje
	var cedula string
	var Militar sssifanb.Militar
	sq, err := sys.PostgreSQLSAMAN.Query(obtenerCuentaBancaria())
	fmt.Println(obtenerCuentaBancaria())
	if err != nil {
		fmt.Println(err.Error())
		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	miliares := 1
	var Historial []interface{}
	for sq.Next() {
		var cedulaAux string
		var cuenta, tipo, institucion, prioridad sql.NullString
		var DatoFinanciero sssifanb.DatoFinanciero

		sq.Scan(&cedulaAux, &cuenta, &tipo, &institucion, &prioridad)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if i == 0 {
			cedula = cedulaAux
		}
		if cedula != cedulaAux {
			fmt.Println("# " + strconv.Itoa(i) + " ID_: " + cedula)
			fm := make(map[string]interface{})
			fm["persona.datofinanciero"] = Historial
			Militar.ActualizarMGO(cedula, fm)

			miliares++
			cedula = cedulaAux
			Historial = nil
		}

		DatoFinanciero.Cuenta = util.ValidarNullString(cuenta)
		DatoFinanciero.Institucion = util.ValidarNullString(institucion)
		DatoFinanciero.Tipo = util.ValidarNullString(tipo)
		DatoFinanciero.Prioridad = util.ValidarNullString(prioridad)
		Historial = append(Historial, DatoFinanciero)
		i++
	}
	fmt.Println("# " + strconv.Itoa(i) + " ID_: " + cedula)
	fm := make(map[string]interface{})
	fm["persona.datofinanciero"] = Historial
	Militar.ActualizarMGO(cedula, fm)
	return
}

//CargarComponenteGrado Historial militar
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

//CargarEstados Historial militar
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
			estado.Nombre = strings.ToUpper(cod)
		}
		if codigo != iso {
			fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
			estado.Ciudad = lstCiudad
			estado.Codigo = codigo

			estado.SalvarMGO("estado")
			lstCiudad = nil
			estado.Nombre = strings.ToUpper(cod)
			codigo = iso

		}

		ciudad.Capital = capi
		ciudad.Nombre = strings.ToUpper(ciud)
		lstCiudad = append(lstCiudad, ciudad)
		i++
	}
	estado.Ciudad = lstCiudad
	estado.SalvarMGO("estado")
	fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
	return
}

//CargarMunicipio Historial militar
func (e *Estructura) CargarMunicipio() (jSon []byte, err error) {
	var msj Mensaje
	var codigo string

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerMunicipios())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	var lstMunicipio []fanb.Municipio
	var estado fanb.Estado
	for sq.Next() {
		var cod, mun string
		var Mun fanb.Municipio
		err = sq.Scan(&cod, &mun)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		cod = strings.ToUpper(cod)
		mun = strings.ToUpper(mun)
		if i == 0 {
			codigo = cod
		}
		if codigo != cod {
			fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
			fm := make(map[string]interface{})
			fm["municipio"] = lstMunicipio
			donde := bson.M{"nombre": codigo}
			estado.ActualizarMGO(donde, fm)
			lstMunicipio = nil
			codigo = cod
		}
		Mun.Nombre = mun
		lstMunicipio = append(lstMunicipio, Mun)
		i++
	}
	fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
	fm := make(map[string]interface{})
	fm["municipio"] = lstMunicipio
	donde := bson.M{"nombre": codigo}
	estado.ActualizarMGO(donde, fm)
	return
}

//CargarParroquia Historial militar
func (e *Estructura) CargarParroquia() (jSon []byte, err error) {
	var msj Mensaje
	var codigo string

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerMunicipiosParroquia())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	var lstParroquia []string
	var estado fanb.Estado
	for sq.Next() {
		var cod, mun, parr string
		err = sq.Scan(&cod, &mun, &parr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		parr = strings.ToUpper(parr)
		cod = strings.ToUpper(cod)
		mun = strings.ToUpper(mun)
		if i == 0 {
			codigo = mun
			estado.Nombre = cod
		}
		if codigo != mun {
			fmt.Println("{\"nombre\":\"" + estado.Nombre + "\", \"municipio.nombre\":\"" + codigo + "\"}")
			fm := make(map[string]interface{})
			fm["municipio.$.parroquia"] = lstParroquia
			donde := bson.M{"nombre": estado.Nombre, "municipio.nombre": codigo}
			estado.ActualizarMGO(donde, fm)
			lstParroquia = nil
			codigo = mun
			estado.Nombre = cod
		}
		lstParroquia = append(lstParroquia, parr)
		i++
	}
	fmt.Println("# " + strconv.Itoa(i) + " ID_: " + codigo)
	fm := make(map[string]interface{})
	fm["municipio.$.parroquia"] = lstParroquia
	donde := bson.M{"nombre": estado.Nombre, "municipio.nombre": codigo}
	estado.ActualizarMGO(donde, fm)
	return
}

func errorG(sSQL string) {
	_, err := sys.PostgreSQLSAMAN.Exec(sSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//CargarMilitar Historial militar
func (e *Estructura) CargarPensiones() (jSon []byte, err error) {
	var msj Mensaje
	var cedula string

	sq, err := sys.PostgreSQLSAMAN.Query(HistoriaPension())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	miliares := 1
	var Historial []interface{}
	//var Pension sssifanb.Pension
	for sq.Next() {
		var cedulaAux string
		var vigente, direc, finsc sql.NullString
		var sueldob, ptransporte, pdescenc, pannoserv sql.NullFloat64
		var pnoascenso, ppnoascenso, pespecial, pprofesional, ppprof, subtotal, pprestacion, pasignada sql.NullFloat64
		var bonovac, bonovacaguinaldo sql.NullFloat64

		var historialpension sssifanb.HistorialPensionSueldo
		err = sq.Scan(&cedulaAux, &vigente, &direc, &finsc, &sueldob, &ptransporte, &pdescenc, &pannoserv,
			&pnoascenso, &ppnoascenso, &pespecial, &pprofesional, &ppprof, &subtotal, &pprestacion, &pasignada,
			&bonovac, &bonovacaguinaldo)
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
			fm["pension.historialsueldo"] = Historial
			Militar.ActualizarMGO(cedula, fm)
			miliares++
			cedula = cedulaAux
			Historial = nil
		}
		var prima sssifanb.Prima
		historialpension.Directiva = util.ValidarNullString(direc)
		historialpension.Sueldo = util.ValidarNullFloat64(sueldob)
		historialpension.PensionAsignada = util.ValidarNullFloat64(pasignada)
		prima.Descendencia = util.ValidarNullFloat64(pdescenc)
		prima.Especial = util.ValidarNullFloat64(pespecial)
		prima.NoAscenso = util.ValidarNullFloat64(pnoascenso)
		prima.PorcentajeNoAscenso = util.ValidarNullFloat64(ppnoascenso)
		prima.SubTotal = util.ValidarNullFloat64(subtotal)
		prima.Transporte = util.ValidarNullFloat64(ptransporte)
		historialpension.Prima = prima
		historialpension.BonoVacacional = util.ValidarNullFloat64(bonovac)
		historialpension.BonoAguinaldo = util.ValidarNullFloat64(bonovacaguinaldo)

		Historial = append(Historial, historialpension)
		i++

	}
	var Militar sssifanb.Militar

	fm := make(map[string]interface{})
	fm["pension.historialsueldo"] = Historial
	Militar.ActualizarMGO(cedula, fm)

	fmt.Println("CANTIDAD: REGISTROS: ", i, " MILITARES: ", miliares)
	return
}

func (e *Estructura) CrearUsuarios() (jSon []byte, err error) {
	var msj Mensaje

	sq, err := sys.PostgreSQLSAMAN.Query(HistorialUsuario())
	if err != nil {

		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 0
		jSon, err = json.Marshal(msj)
	}
	i := 0
	for sq.Next() {
		var usua, cedula string
		var nombp, nombs, apellp, apellis sql.NullString
		var Usuario seguridad.Usuario
		var privilegio seguridad.Privilegio
		var lst []seguridad.Privilegio

		sq.Scan(&usua, &cedula, &nombp, &nombs, &apellp, &apellis)
		//

		Usuario.ID = bson.NewObjectId()
		Usuario.Cedula = cedula
		nombre := util.ValidarNullString(nombp) + " " + util.ValidarNullString(nombs) +
			" " + util.ValidarNullString(apellp) + " " + util.ValidarNullString(apellis)
		Usuario.Nombre = strings.ToUpper(nombre)

		Usuario.Login = usua
		Usuario.Sucursal = "Principal"
		Usuario.Clave = util.GenerarHash256([]byte(cedula))

		Usuario.Rol.Descripcion = seguridad.GAFILIACION
		Usuario.Perfil.Descripcion = seguridad.ANALISTA

		privilegio.Metodo = "afiliacion.salvar"
		privilegio.Descripcion = "Crear Usuario"
		privilegio.Accion = "Insert()" // ES6 Metodos
		lst = append(lst, privilegio)

		privilegio.Metodo = "afiliacion.modificar"
		privilegio.Descripcion = "Modificar Usuario"
		privilegio.Accion = "Update()"
		lst = append(lst, privilegio)
		Usuario.Perfil.Privilegios = lst
		Usuario.Salvar()
		i++
		fmt.Println("Insertando", i)
	}
	msj.Tipo = 1
	jSon, err = json.Marshal(msj)
	return
}

func (e *Estructura) ActualizarFechaCarnet() (jSon []byte, err error) {
	var msj Mensaje

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerFechaVencimiento())
	if err != nil {
		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0

	for sq.Next() {

		var ced string
		var fech sql.NullString
		var fechavence time.Time
		err = sq.Scan(&ced, &fech)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		layOut := "2006-01-02"
		fecha := util.ValidarNullString(fech)
		if fecha != VNULL {
			dateString := strings.Replace(fecha, "/", "-", -1)
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				fechavence = dateStamp
			}
		}
		militar := make(map[string]interface{})
		militar["tim.fechavencimiento"] = fechavence
		c := sys.MGOSession.DB("sssifanb").C("militar")
		err = c.Update(bson.M{"id": ced}, bson.M{"$set": militar})

		i++
	}
	return
}

func (e *Estructura) ActualizarFechaDefuncion() (jSon []byte, err error) {
	var msj Mensaje

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerFechaDefuncion())
	if err != nil {
		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0

	for sq.Next() {

		var ced string
		var fech sql.NullString
		var fechavence time.Time
		err = sq.Scan(&ced, &fech)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		layOut := "2006-01-02"
		fecha := util.ValidarNullString(fech)
		if fecha != VNULL {
			dateString := strings.Replace(fecha, "/", "-", -1)
			dateStamp, err := time.Parse(layOut, dateString)
			if err == nil {
				fechavence = dateStamp
			}
		}
		militar := make(map[string]interface{})
		militar["persona.datobasico.fechadefuncion"] = fechavence
		c := sys.MGOSession.DB("sssifanb").C("militar")
		err = c.Update(bson.M{"id": ced}, bson.M{"$set": militar})

		// familiar := make(map[string]interface{})
		// familiar["familiar.$.persona.datobasico.fechadefuncion"] = fechavence
		// c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": ced}, bson.M{"$set": militar})

		i++
		fmt.Println("Actualizando. ", i, " ID_ ", ced)
	}
	return
}

func (e *Estructura) ConvertirGradoGN() (jSon []byte, err error) {
	var msj Mensaje

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerPensionadosAntes2008GN())
	if err != nil {
		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	for sq.Next() {
		var ced string
		militar := make(map[string]interface{})
		sq.Scan(&ced)
		militar["pension.grado"] = "SAY"
		err = c.Update(bson.M{"id": ced}, bson.M{"$set": militar})
		fmt.Println("Cedula", ced)
	}
	return
}

func (e *Estructura) ActualizarPrimaProfesional() {
	sq, err := sys.PostgreSQLSAMAN.Query(obtenerPrimaProfesional())
	util.Error(err)
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	for sq.Next() {
		pprofesional := make(map[string]interface{})
		var ced string
		var pprof sql.NullFloat64
		sq.Scan(&ced, &pprof)
		pprofesional["pension.pprofesional"] = util.ValidarNullFloat64(pprof)
		fmt.Println("Cedula", ced)
		err = c.Update(bson.M{"id": ced}, bson.M{"$set": pprofesional})
		util.Error(err)
	}
}
