package sssifanb

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
	"gopkg.in/mgo.v2/bson"
)

func Sincronizar(militar Militar) {
	s := `SELECT p.nropersona FROM personas p
		JOIN pers_dat_militares m on p.nropersona=m.nropersona
		WHERE p.codnip='` + militar.Persona.DatoBasico.Cedula + `' LIMIT 1`
	sq, err := sys.PostgreSQLSAMAN.Query(s)
	if err != nil {
		return
	}
	existe := 0
	for sq.Next() {
		existe = 1
	}
	if existe == 0 {
		m := InsertarMilitarSAMANSN(militar)
		_, e := sys.PostgreSQLSAMAN.Exec(m)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println("INSERTADO: ", militar.Persona.DatoBasico.Cedula)
	} else {
		p := ActualizarPersona(militar.Persona)
		m := ActualizarMilitar(militar)
		_, e := sys.PostgreSQLSAMAN.Exec(p + m)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println("ACTUALIZADO: ", militar.Persona.DatoBasico.Cedula)
	}

}

func SincronizarTest(militar Militar) {
	s := `SELECT personas.nropersona FROM personas JOIN pers_dat_militares on personas.nropersona=pers_dat_militares.nropersona
		WHERE personas.codnip='` + militar.Persona.DatoBasico.Cedula + `' LIMIT 1`
	sq, err := sys.PsqlWEB.Query(s)
	if err != nil {
		fmt.Println("Err", err.Error())

	} else {
		for sq.Next() {
			var ced string
			sq.Scan(&ced)
			fmt.Println("Encontrado... ", ced)

		}
	}

}

func ActualizarPersona(persona Persona) string {
	fecha := time.Now()
	convertedDateString := fecha.Format("2006-01-02")
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := persona.DatoBasico.FechaNacimiento.Format("2006-01-02")
	fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)

	a, _, _ := persona.DatoBasico.FechaDefuncion.Date()
	fechaSlashDefuncion := ""
	if a > 1000 {
		convertirDef := persona.DatoBasico.FechaDefuncion.Format("2006-01-02")
		fechaSlashDefuncion = strings.Replace(convertirDef, "-", "/", -1)
	}

	return `UPDATE personas SET
		tipnip = '` + persona.DatoBasico.Nacionalidad + `',
		codnip = '` + persona.DatoBasico.Cedula + `',
		nombreprimero = '` + strings.TrimSpace(persona.DatoBasico.NombrePrimero) + `',
		nombresegundo ='` + strings.TrimSpace(persona.DatoBasico.NombreSegundo) + `',
		apellidoprimero ='` + strings.TrimSpace(persona.DatoBasico.ApellidoPrimero) + `',
		apellidosegundo ='` + strings.TrimSpace(persona.DatoBasico.ApellidoSegundo) + `',
		nombrecompletoupp = '` + strings.TrimSpace(persona.DatoBasico.ConcatenarNombre()+persona.DatoBasico.ConcatenarApellido()) + `',
		nombrecompleto = '` + strings.TrimSpace(persona.DatoBasico.ConcatenarNombre()+persona.DatoBasico.ConcatenarApellido()) + `',
		nacionalidadcod = '` + persona.DatoBasico.ConvertirNacionalidad() + `',
		sexocod = '` + persona.DatoBasico.Sexo + `',
		edocivilcod = '` + persona.DatoBasico.EstadoCivil + `',
		fechanacimiento = '` + fechaSlashNacimiento + `',
		fechadefuncion = '` + fechaSlashDefuncion + `',
		email1 = '` + persona.Correo.Principal + `',
		email2 = '` + persona.Correo.Alternativo + `',
		auditfechacambio = '` + fechaSlashActual + `',
		auditcodusuario = 'SSSIFANB'
		WHERE nropersona= (SELECT personas.nropersona FROM personas
			JOIN pers_dat_militares on personas.nropersona=pers_dat_militares.nropersona
			WHERE personas.codnip='` + persona.DatoBasico.Cedula + `' LIMIT 1);`
}

//ActualizarMilitar Militares
func ActualizarMilitar(militar Militar) string {
	fecha := time.Now()
	convertedDateString := fecha.Format("2006-01-02")
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := militar.FechaResuelto.Format("2006-01-02")
	fechaSlashResuelto := strings.Replace(convertir, "-", "/", -1)

	convertirC := militar.FechaIngresoComponente.Format("2006-01-02")
	fechaSlashComponente := strings.Replace(convertirC, "-", "/", -1)

	convertirU := militar.FechaAscenso.Format("2006-01-02")
	fechaSlashUltimoAsc := strings.Replace(convertirU, "-", "/", -1)

	convertirE := militar.FechaRetiro.Format("2006-01-02")
	fechaSlashRetiro := strings.Replace(convertirE, "-", "/", -1)

	return `
		UPDATE pers_dat_militares SET
			componentecod = '` + militar.Componente.Abreviatura + `',
			gradocod = '` + militar.Grado.Abreviatura + `',
			perssituaccod = '` + militar.Situacion + `',
			persclasecod = '` + militar.Clase + `',
			fchingcomponente = '` + fechaSlashComponente + `',
			fchultimoascenso = '` + fechaSlashUltimoAsc + `',
			fchpromocion = '` + fechaSlashComponente + `',
			fchegreso = '` + fechaSlashRetiro + `',
			annoreconocido = '` + strconv.Itoa(militar.AnoReconocido) + `',
			mesreconocido = '` + strconv.Itoa(militar.MesReconocido) + `',
			diareconocido = '` + strconv.Itoa(militar.DiaReconocido) + `',
			resueltoreco = '` + militar.NumeroResuelto + `',
			fchresueltoreco = '` + fechaSlashResuelto + `',
			notaresueltoreco = '` + militar.NumeroResuelto + `',
			auditfechacambio = '` + fechaSlashActual + `',
			audithoracambio = '16:04',
			auditfechacreacion = '` + fechaSlashActual + `',
			audithoracreacion = '16:04',
			auditcodusuario = 'SSSIFANB'
			WHERE nropersona = ( SELECT personas.nropersona FROM personas
				JOIN pers_dat_militares on personas.nropersona=pers_dat_militares.nropersona
				WHERE personas.codnip='` + militar.Persona.DatoBasico.Cedula + `' LIMIT 1);
	`

}

func ActualizarPace(militar Militar) string {
	// fecha_ingreso = ' . $this->fecha_ingreso .  '',
	// f_ult_ascenso = ' . $this->fecha_ultimo_ascenso .  '',
	// 	f_ingreso_sistema = ' . $this->fecha_ingreso .  '',
	// 	f_creacion = ' . $this->fecha_creacion .  ',
	// 	f_ult_modificacion = ' . $this->fecha_ultima_modificacion .  ',
	// n_hijos = ` + strconv.Itoa(militar.NumeroHijos()) + `,
	return `UPDATE beneficiario SET
		grado_id = ` + militar.Grado.Abreviatura + `,
		nombres = '` + militar.Persona.DatoBasico.ConcatenarNombre() + `',
		apellidos = '` + militar.Persona.DatoBasico.ConcatenarApellido() + `',
		tiempo_servicio = '` + militar.TiempoSevicio + `',
	  anio_reconocido = ` + strconv.Itoa(militar.AnoReconocido) + ` ,
	  mes_reconocido = ` + strconv.Itoa(militar.MesReconocido) + `,
	 	dia_reconocido = ` + strconv.Itoa(militar.DiaReconocido) + `,
  	st_no_ascenso = ` + strconv.Itoa(militar.Fideicomiso.EstatusNoAscenso) + `,
	 	st_profesion = ` + strconv.Itoa(militar.Fideicomiso.EstatusProfesion) + `,
	 	sexo = '` + militar.Persona.DatoBasico.Sexo + `',
	 	usr_creacion ='tunel',
	 	usr_modificacion ='tunel',
	 	observ_ult_modificacion='SSSIFANB',
	  WHERE cedula = '` + militar.Persona.DatoBasico.Cedula + `';`

	//echo $sActualizar;

}

//InsertarMilitarSAMAN Control
func InsertarMilitarSAMAN(militar *Militar) string {
	fecha := time.Now()
	convertedDateString := fecha.Format("2006-01-02")
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := militar.Persona.DatoBasico.FechaNacimiento.Format("2006-01-02")
	fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)

	convertirC := militar.FechaIngresoComponente.Format("2006-01-02")
	fechaSlashComponente := strings.Replace(convertirC, "-", "/", -1)

	convertirU := militar.FechaAscenso.Format("2006-01-02")
	fechaSlashUltimoAsc := strings.Replace(convertirU, "-", "/", -1)
	return `
		INSERT INTO personas
		(
			ciaopr,
			nropersona,
			tipnip,
			codnip,
			nombreprimero,
			nombresegundo,
			apellidoprimero,
			apellidosegundo,
			nombrecompleto,
			nombrecorto,
			sexocod,
			edocivilcod,
			idiomanativocod,
			nacionalidadcod,
			fechanacimiento,
			auditcodusuario,
			nombrecompletoupp,
			auditfechacambio,
			audithoracambio,
			auditfechacreacion,
			audithoracreacion
		)
		VALUES (
			'1',
			(SELECT MAX(nropersona)+1 FROM personas),
			'V',
			'` + militar.Persona.DatoBasico.Cedula + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.NombrePrimero) + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.NombreSegundo) + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.ApellidoPrimero) + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.ApellidoSegundo) + `',
			'` + militar.Persona.DatoBasico.ConcatenarNombreApellido() + `',
			'',
			'` + militar.Persona.DatoBasico.Sexo + `',
			'` + militar.Persona.DatoBasico.EstadoCivil + `',
			'ESP',
			'VEN',
			'` + fechaSlashNacimiento + `',
			'SSSIFANB',
			'` + militar.Persona.DatoBasico.ConcatenarApellidoNombre() + `',
			'` + fechaSlashActual + `',
			'16:09',
			'` + fechaSlashActual + `',
			'16:09'
			);

		INSERT INTO pers_dat_afiliac values('1',(SELECT MAX(nropersona) FROM personas WHERE codnip = '` + militar.Persona.DatoBasico.Cedula + `'),1,
				'AT',(select max(nropersona) from personas where ciaopr = '1'),'` + fechaSlashComponente + `','',
				'ACT','INCTI','','','','','` + fechaSlashComponente + `','','','',
				'` + fechaSlashActual + `',
				'16:14',
				'` + fechaSlashActual + `',
				'16:14','SSSIFANB','','','','','','',0.00,0.00,0.00,0.00,0.00,0,'TIT');

		INSERT INTO pers_dat_militares values('1',(SELECT MAX(nropersona) FROM personas WHERE codnip = '` + militar.Persona.DatoBasico.Cedula + `'),
			'` + militar.Componente.Abreviatura + `',
			'` + militar.Grado.Abreviatura + `',
			'` + militar.Categoria + `',
			'` + militar.Situacion + `',
			'` + militar.Clase + `',
			/*fecha ingreso componente:*/'` + fechaSlashComponente + `',
			/*fecha ultimo ascenso:*/'` + fechaSlashUltimoAsc + `',
			/*fecha ultimo ascenso:*/'` + fechaSlashUltimoAsc + `',
			'',0,0,0,
			/*años:*/` + strconv.Itoa(militar.AnoReconocido) + `,
			/*meses:*/` + strconv.Itoa(militar.MesReconocido) + `,
			/*dias:*/` + strconv.Itoa(militar.DiaReconocido) + `,
			/*años:*/` + strconv.Itoa(militar.AnoReconocido) + `,
			/*meses:*/` + strconv.Itoa(militar.MesReconocido) + `,
			/*dias*/` + strconv.Itoa(militar.AnoReconocido) + `,
			'','','','','','',
			'` + fechaSlashActual + `', '16:14',
			'` + fechaSlashActual + `','16:14','SSSIFANB','','','','','` + militar.Componente.Abreviatura + `','',0.00,0.00,0.00,0.00,0.00,1)


`

}

func InsertarMilitarSAMANSN(militar Militar) string {
	fecha := time.Now()
	convertedDateString := fecha.Format("2006-01-02")
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := militar.Persona.DatoBasico.FechaNacimiento.Format("2006-01-02")
	fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)

	convertirC := militar.FechaIngresoComponente.Format("2006-01-02")
	fechaSlashComponente := strings.Replace(convertirC, "-", "/", -1)

	convertirU := militar.FechaAscenso.Format("2006-01-02")
	fechaSlashUltimoAsc := strings.Replace(convertirU, "-", "/", -1)
	return `
		INSERT INTO personas
		(
			ciaopr,
			nropersona,
			tipnip,
			codnip,
			nombreprimero,
			nombresegundo,
			apellidoprimero,
			apellidosegundo,
			nombrecompleto,
			nombrecorto,
			sexocod,
			edocivilcod,
			idiomanativocod,
			nacionalidadcod,
			fechanacimiento,
			auditcodusuario,
			nombrecompletoupp,
			auditfechacambio,
			audithoracambio,
			auditfechacreacion,
			audithoracreacion
		)
		VALUES (
			'1',
			(SELECT MAX(nropersona)+1 FROM personas),
			'V',
			'` + militar.Persona.DatoBasico.Cedula + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.NombrePrimero) + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.NombreSegundo) + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.ApellidoPrimero) + `',
			'` + strings.TrimSpace(militar.Persona.DatoBasico.ApellidoSegundo) + `',
			'` + militar.Persona.DatoBasico.ConcatenarNombreApellido() + `',
			'',
			'` + militar.Persona.DatoBasico.Sexo + `',
			'` + militar.Persona.DatoBasico.EstadoCivil + `',
			'ESP',
			'VEN',
			'` + fechaSlashNacimiento + `',
			'SSSIFANB',
			'` + militar.Persona.DatoBasico.ConcatenarApellidoNombre() + `',
			'` + fechaSlashActual + `',
			'16:09',
			'` + fechaSlashActual + `',
			'16:09'
			);

		INSERT INTO pers_dat_afiliac values('1',(SELECT MAX(nropersona) FROM personas WHERE codnip = '` + militar.Persona.DatoBasico.Cedula + `'),1,
				'AT',(select max(nropersona) from personas where ciaopr = '1'),'` + fechaSlashComponente + `','',
				'ACT','INCTI','','','','','` + fechaSlashComponente + `','','','',
				'` + fechaSlashActual + `',
				'16:14',
				'` + fechaSlashActual + `',
				'16:14','SSSIFANB','','','','','','',0.00,0.00,0.00,0.00,0.00,0,'TIT');

		INSERT INTO pers_dat_militares values('1',(SELECT MAX(nropersona) FROM personas WHERE codnip = '` + militar.Persona.DatoBasico.Cedula + `'),
			'` + militar.Componente.Abreviatura + `',
			'` + militar.Grado.Abreviatura + `',
			'` + militar.Categoria + `',
			'` + militar.Situacion + `',
			'` + militar.Clase + `',
			/*fecha ingreso componente:*/'` + fechaSlashComponente + `',
			/*fecha ultimo ascenso:*/'` + fechaSlashUltimoAsc + `',
			/*fecha ultimo ascenso:*/'` + fechaSlashUltimoAsc + `',
			'',0,0,0,
			/*años:*/` + strconv.Itoa(militar.AnoReconocido) + `,
			/*meses:*/` + strconv.Itoa(militar.MesReconocido) + `,
			/*dias:*/` + strconv.Itoa(militar.DiaReconocido) + `,
			/*años:*/` + strconv.Itoa(militar.AnoReconocido) + `,
			/*meses:*/` + strconv.Itoa(militar.MesReconocido) + `,
			/*dias*/` + strconv.Itoa(militar.AnoReconocido) + `,
			'','','','','','',
			'` + fechaSlashActual + `', '16:14',
			'` + fechaSlashActual + `','16:14','SSSIFANB','','','','','` + militar.Componente.Abreviatura + `','',0.00,0.00,0.00,0.00,0.00,1)


`

}

func ActualizarMysqlFT(mil Militar) string {

	dire := obtenerEstado(mil.Persona.Direccion[0].Estado) + " " + mil.Persona.Direccion[0].Ciudad +
		" " + mil.Persona.Direccion[0].Municipio + " " + mil.Persona.Direccion[0].Parroquia +
		" " + mil.Persona.Direccion[0].CalleAvenida + " " + mil.Persona.Direccion[0].Casa +
		" " + strconv.Itoa(mil.Persona.Direccion[0].Numero)

	grad := mil.Grado.Descripcion
	comp := mil.Componente.Descripcion
	situ := obtenerSitiacion(mil.Situacion)
	clas := obtenerClase(mil.Clase)
	cate := obtenerCategoria(mil.Categoria)
	telf := mil.Persona.Telefono.Domiciliario + " " + mil.Persona.Telefono.Movil
	dire += telf + " " +
		obtenerEstadoCivil(mil.Persona.DatoBasico.EstadoCivil) + " " +
		obtenerSexo(mil.Persona.DatoBasico.Sexo)
	fami := ""
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

	return `UPDATE datos SET
		nombre = '` + mil.Persona.DatoBasico.ConcatenarApellidoNombre() + `',
		descripcion = '` + body + `',
		direccion = '` + dire + `',
		familiares = '` + fami + `'
		WHERE cedula='` + mil.Persona.DatoBasico.Cedula + `';`
}

func InsertMysqlFT(mil *Militar) string {

	dire := obtenerEstado(mil.Persona.Direccion[0].Estado) + " " + mil.Persona.Direccion[0].Ciudad +
		" " + mil.Persona.Direccion[0].Municipio + " " + mil.Persona.Direccion[0].Parroquia +
		" " + mil.Persona.Direccion[0].CalleAvenida + " " + mil.Persona.Direccion[0].Casa +
		" " + strconv.Itoa(mil.Persona.Direccion[0].Numero)

	grad := mil.Grado.Descripcion
	comp := mil.Componente.Descripcion
	situ := obtenerSitiacion(mil.Situacion)
	clas := obtenerClase(mil.Clase)
	cate := obtenerCategoria(mil.Categoria)
	telf := mil.Persona.Telefono.Domiciliario + " " + mil.Persona.Telefono.Movil
	dire += telf + " " +
		obtenerEstadoCivil(mil.Persona.DatoBasico.EstadoCivil) + " " +
		obtenerSexo(mil.Persona.DatoBasico.Sexo)
	fami := ""
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
	return `INSERT INTO datos ( cedula, nombre, descripcion, direccion, familiares )
		VALUES ('` + mil.Persona.DatoBasico.Cedula + `','` + mil.Persona.DatoBasico.ConcatenarApellidoNombre() + `','` + body + `','` + dire + `','` + fami + `')`

}

func obtenerPensionados() string {
	return `
		SELECT per.nropersona, per.tipnip, per.codnip,
			per.nombreprimero, per.nombresegundo, per.apellidoprimero, per.apellidosegundo,
			per.sexocod, per.edocivilcod, per.fechanacimiento,
			pdm.componentecod, pdm.gradocod, pdm.perscategcod, pdm.perssituaccod, pdm.persclasecod,
			pdm.fchingcomponente,pdm.fchultimoascenso,pdm.fchpromocion,pdm.fchegreso,
			pdm.annoreconocido,pdm.mesreconocido,pdm.diareconocido,
			pen.nrohijos,pen.tipcuentacod,pen.instfinancod,pen.nrocuenta, pcal.porcentaje,
			pen.estpencod,pen.razestpencod
		FROM personas per
		JOIN pers_dat_militares pdm ON  per.nropersona=pdm.nropersona
		JOIN pension pen ON pdm.nropersona=pen.nropersona
		JOIN (
						SELECT nropersona, MAX(porcprestmonto) AS porcentaje  FROM
							pension_calc
						GROUP BY nropersona
					) pcal ON pen.nropersona=pcal.nropersona

		-- where
		-- codnip='9150043'
		-- pdm.perssituaccod='ACT'`
}

//MGOActualizarPensionados Actualizando datos principales del militar
func (m *Militar) MGOActualizarPensionados() (err error) {

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerPensionados())
	if err != nil {
		return
	}

	for sq.Next() {
		var militar Militar
		var nropersona, tipnip, codnip, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo sql.NullString
		var sexocod, edocivilcod sql.NullString
		var componentecod, gradocod, perscategcod, perssituaccod, persclasecod sql.NullString
		var fchingcomponente, fchultimoascenso, fchpromocion, fchegreso, fechanacimiento sql.NullString
		var annoreconocido, mesreconocido, diareconocido sql.NullString
		var nrohijos, tipcuentacod, instfinancod, nrocuenta sql.NullString
		var porcentaje sql.NullFloat64
		var estpencod, razestpencod sql.NullString

		sq.Scan(&nropersona, &tipnip, &codnip, &nombreprimero, &nombresegundo,
			&apellidoprimero, &apellidosegundo, &sexocod, &edocivilcod, &fechanacimiento,
			&componentecod, &gradocod, &perscategcod, &perssituaccod, &persclasecod,
			&fchingcomponente, &fchultimoascenso, &fchpromocion, &fchegreso,
			&annoreconocido, &mesreconocido, &diareconocido,
			&nrohijos, &tipcuentacod, &instfinancod, &nrocuenta, &porcentaje,
			&estpencod, &razestpencod)

		militar.Pension.GradoCodigo = util.ValidarNullString(gradocod)
		militar.Pension.ComponenteCodigo = util.ValidarNullString(componentecod)
		militar.Pension.Categoria = util.ValidarNullString(perscategcod)
		militar.Pension.Clase = util.ValidarNullString(persclasecod)
		militar.Pension.Situacion = util.ValidarNullString(perssituaccod)

		// militar.Pension.AnoServicio, _ = strconv.Atoi(util.ValidarNullString(annototservicio))
		// militar.Pension.MesServicio, _ = strconv.Atoi(util.ValidarNullString(mestotservicio))
		// militar.Pension.DiaServicio, _ = strconv.Atoi(util.ValidarNullString(diatotservicio))
		// militar.Pension.PensionAsignada = util.ValidarNullFloat64(pensionasignada)
		militar.Pension.PorcentajePrestaciones = util.ValidarNullFloat64(porcentaje)
		militar.Pension.FechaPromocion = util.ValidarNullString(fchpromocion)
		militar.Pension.FechaUltimoAscenso = util.ValidarNullString(fchultimoascenso)

		militar.Pension.DatoFinanciero.Cuenta = util.ValidarNullString(nrocuenta)
		militar.Pension.DatoFinanciero.Tipo = util.ValidarNullString(tipcuentacod)
		militar.Pension.DatoFinanciero.Institucion = util.ValidarNullString(instfinancod)
		militar.Pension.NumeroHijos, _ = strconv.Atoi(util.ValidarNullString(nrohijos))
		pension := make(map[string]interface{})
		c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
		pension["pension"] = militar.Pension
		err = c.Update(bson.M{"id": util.ValidarNullString(codnip)}, bson.M{"$set": pension})
		if err != nil {
			fmt.Println("Err", err.Error())
			return
		}
	}

	return
}
