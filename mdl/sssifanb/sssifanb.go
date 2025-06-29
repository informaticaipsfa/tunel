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

const layout string = "2006-01-02"

//
// func Sincronizar(militar Militar) {
// 	s := `SELECT p.nropersona FROM personas p
// 		JOIN pers_dat_militares m on p.nropersona=m.nropersona
// 		WHERE p.codnip='` + militar.Persona.DatoBasico.Cedula + `' LIMIT 1`
// 	sq, err := sys.PostgreSQLSAMAN.Query(s)
// 	if err != nil {
// 		return
// 	}
// 	existe := 0
// 	for sq.Next() {
// 		existe = 1
// 	}
// 	if existe == 0 {
// 		m := InsertarMilitarSAMANSN(militar)
// 		_, e := sys.PostgreSQLSAMAN.Exec(m)
// 		if e != nil {
// 			fmt.Println(e.Error())
// 			return
// 		}
// 		fmt.Println("INSERTADO: ", militar.Persona.DatoBasico.Cedula)
// 	} else {
// 		p := ActualizarPersona(militar.Persona)
// 		m := ActualizarMilitar(militar)
// 		_, e := sys.PostgreSQLSAMAN.Exec(p + m)
// 		if e != nil {
// 			fmt.Println(e.Error())
// 			return
// 		}
// 		fmt.Println("ACTUALIZADO: ", militar.Persona.DatoBasico.Cedula)
// 	}
//
// }

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
	convertedDateString := fecha.Format(layout)
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := persona.DatoBasico.FechaNacimiento.Format(layout)
	fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)

	a, _, _ := persona.DatoBasico.FechaDefuncion.Date()
	fechaSlashDefuncion := ""
	if a > 1000 {
		convertirDef := persona.DatoBasico.FechaDefuncion.Format(layout)
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

// ActualizarMilitar Militares
func ActualizarMilitar(militar Militar) string {
	fecha := time.Now()
	convertedDateString := fecha.Format(layout)
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := militar.FechaResuelto.Format(layout)
	fechaSlashResuelto := strings.Replace(convertir, "-", "/", -1)

	convertirC := militar.FechaIngresoComponente.Format(layout)
	fechaSlashComponente := strings.Replace(convertirC, "-", "/", -1)

	convertirU := militar.FechaAscenso.Format(layout)
	fechaSlashUltimoAsc := strings.Replace(convertirU, "-", "/", -1)

	convertirE := militar.FechaRetiro.Format(layout)
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

// InsertarMilitarSAMAN Control
func InsertarMilitarSAMAN(militar *Militar) string {
	fecha := time.Now()
	convertedDateString := fecha.Format(layout)
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := militar.Persona.DatoBasico.FechaNacimiento.Format(layout)
	fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)

	convertirC := militar.FechaIngresoComponente.Format(layout)
	fechaSlashComponente := strings.Replace(convertirC, "-", "/", -1)

	convertirU := militar.FechaAscenso.Format(layout)
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
	convertedDateString := fecha.Format(layout)
	fechaSlashActual := strings.Replace(convertedDateString, "-", "/", -1)

	convertir := militar.Persona.DatoBasico.FechaNacimiento.Format(layout)
	fechaSlashNacimiento := strings.Replace(convertir, "-", "/", -1)

	convertirC := militar.FechaIngresoComponente.Format(layout)
	fechaSlashComponente := strings.Replace(convertirC, "-", "/", -1)

	convertirU := militar.FechaAscenso.Format(layout)
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

//func ActualizarMysqlFT(mil Militar, familiar Familiar) string {

func ActualizarMysqlFT(mil Militar, familiar Familiar) []string {
	var queries []string
	escape := func(s string) string {
		return strings.ReplaceAll(s, "'", "''")
	}
	dire := obtenerEstado(mil.Persona.Direccion[0].Estado) + " " + mil.Persona.Direccion[0].Ciudad +
		" " + mil.Persona.Direccion[0].Municipio + " " + mil.Persona.Direccion[0].Parroquia +
		" " + mil.Persona.Direccion[0].CalleAvenida + " " + mil.Persona.Direccion[0].Casa +
		" " + strconv.Itoa(mil.Persona.Direccion[0].Numero)

	convertir := mil.TIM.FechaVencimiento.Format(layout)
	fechaSlashVencimiento := strings.Replace(convertir, "/", "-", -1)

	convertirf := familiar.Persona.DatoBasico.FechaNacimiento.Format(layout)
	fechaSlashNacimiento := strings.Replace(convertirf, "-", "/", -1)

	convertirfv := familiar.TIF.FechaVencimiento.Format(layout)
	fechaSlashVencimientof := strings.Replace(convertirfv, "-", "/", -1)

	serial_carnetf := familiar.TIF.Serial
	donante := familiar.Donante
	grupo_sanguineoF := familiar.GrupoSanguineo
	historial_clinicoF := familiar.HistoriaMedica
	nombreM1 := mil.Persona.DatoBasico.NombrePrimero
	apellidoM1 := mil.Persona.DatoBasico.ApellidoPrimero
	parentesco := obtenerParentesco(familiar.Parentesco, familiar.Persona.DatoBasico.Sexo)
	abreviatura_comp := mil.Componente.Abreviatura
	color_piel := mil.Persona.DatoFisionomico.ColorPiel
	ojos := mil.Persona.DatoFisionomico.ColorOjos
	estaturaStr := fmt.Sprintf("%.2f", mil.Persona.DatoFisionomico.Estatura)
	grupo_sanguineo := mil.Persona.DatoFisionomico.GrupoSanguineo
	historial_clinico := mil.NumeroHistoria
	cabello := mil.Persona.DatoFisionomico.ColorCabello
	codigo_comp := mil.CodigoComponente
	serial := mil.TIM.Serial

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
	afiliado := grad + " " + apellidoM1 + " " + nombreM1 + "" + mil.Persona.DatoBasico.Cedula

	queries = append(queries, /*`UPDATE datos SET
		nombre = '`+mil.Persona.DatoBasico.ConcatenarApellidoNombre()+`',
		descripcion = '`+body+`',
		direccion = '`+dire+`',
		familiares = '`+fami+`'
		WHERE cedula='`+mil.Persona.DatoBasico.Cedula+`');*/

		`UPDATE  sssifanb.carp_militar SET 
					nombreprimero = '`+escape(strings.TrimSpace(mil.Persona.DatoBasico.NombrePrimero))+`',
					nombresegundo = '`+escape(strings.TrimSpace(mil.Persona.DatoBasico.NombreSegundo))+`',
					apellidoprimero = '`+escape(strings.TrimSpace(mil.Persona.DatoBasico.ApellidoPrimero))+`',
					apellidosegundo = '`+escape(strings.TrimSpace(mil.Persona.DatoBasico.ApellidoSegundo))+`',
					grado = '`+escape(grad)+`',
					fecha_vencimiento = '`+fechaSlashVencimiento+`',
					serial_carnet = '`+escape(serial)+`',
					codigo_comp = '`+escape(codigo_comp)+`',
					cabello='`+cabello+`',
					historial_clinico = '`+escape(historial_clinico)+`',
					grupo_sanguineo = '`+escape(grupo_sanguineo)+`',
					estatura = '`+estaturaStr+`',
					ojos = '`+ojos+`',
					color_piel ='`+color_piel+`',
					componente = '`+abreviatura_comp+`',
					categoria ='`+escape(cate)+`',
					huella = '`+escape(abreviatura_comp)+`',
					foto = '`+escape(grupo_sanguineo)+`'
					WHERE cedula =' `+escape(mil.Persona.DatoBasico.Cedula)+`';

						
					UPDATE sssifanb.carp_familiar SET
    nombreprimerof = `+familiar.Persona.DatoBasico.NombrePrimero+`,
    apellidoprimerof = `+familiar.Persona.DatoBasico.ApellidoPrimero+`,
    fecha_nacimiento = '`+escape(fechaSlashNacimiento)+`',
    parentesco = '`+escape(parentesco)+`',
    afiliado = '`+escape(afiliado)+`',
    f_vencimiento = '`+escape(fechaSlashVencimientof)+`',
    historial_clinicof = '`+escape(historial_clinicoF)+`',
    grupo_sanguineof = '`+escape(grupo_sanguineoF)+`',
    donante = `+donante+`,
    serial_carnet = '`+escape(serial_carnetf)+`',
    huella = '`+escape(serial_carnetf)+`',
    foto = '`+escape(body)+`'
WHERE cedula_militar = '`+escape(mil.Persona.DatoBasico.Cedula)+`' 
AND cedula_familiar = '`+escape(familiar.Persona.DatoBasico.Cedula)+`'
					
			);`)

	return queries

}

func InsertMysqlFT(mil *Militar, familiar *Familiar) string {

	dire := obtenerEstado(mil.Persona.Direccion[0].Estado) + " " + mil.Persona.Direccion[0].Ciudad +
		" " + mil.Persona.Direccion[0].Municipio + " " + mil.Persona.Direccion[0].Parroquia +
		" " + mil.Persona.Direccion[0].CalleAvenida + " " + mil.Persona.Direccion[0].Casa +
		" " + strconv.Itoa(mil.Persona.Direccion[0].Numero)

	convertir := mil.TIM.FechaVencimiento.Format(layout)
	fechaSlashVencimiento := strings.Replace(convertir, "/", "-", -1)

	convertirf := familiar.Persona.DatoBasico.FechaNacimiento.Format(layout)
	fechaSlashNacimiento := strings.Replace(convertirf, "-", "/", -1)

	convertirfv := familiar.TIF.FechaVencimiento.Format(layout)
	fechaSlashVencimientof := strings.Replace(convertirfv, "-", "/", -1)

	serial_carnetf := familiar.TIF.Serial
	donante := familiar.Donante
	grupo_sanguineoF := familiar.GrupoSanguineo
	historial_clinicoF := familiar.HistoriaMedica
	nombreM1 := mil.Persona.DatoBasico.NombrePrimero
	apellidoM1 := mil.Persona.DatoBasico.ApellidoPrimero
	parentesco := obtenerParentesco(familiar.Parentesco, familiar.Persona.DatoBasico.Sexo)
	abreviatura_comp := mil.Componente.Abreviatura
	ojos := mil.Persona.DatoFisionomico.ColorOjos
	estaturaStr := fmt.Sprintf("%.2f", mil.Persona.DatoFisionomico.Estatura)
	grupo_sanguineo := mil.Persona.DatoFisionomico.GrupoSanguineo
	historial_clinico := mil.NumeroHistoria
	codigo_comp := mil.CodigoComponente
	serial := mil.TIM.Serial
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
	afiliado := grad + " " + apellidoM1 + " " + nombreM1 + "" + mil.Persona.DatoBasico.Cedula

	/*return /*`INSERT INTO datos ( cedula, nombre, descripcion, direccion, familiares )
	VALUES ('` + mil.Persona.DatoBasico.Cedula + `','` + mil.Persona.DatoBasico.ConcatenarApellidoNombre() + `','` + body + `','` + dire + `','` + fami + `');
	*/
	return ` INSERT INTO sssifanb.carp_militar
( cedula, 
 nombreprimero,
  nombresegundo, 
  apellidoprimero,
   apellidosegundo,
    grado, 
	fecha_vencimiento,
	 serial_carnet,
	  codigo_comp, 
	  cabello, 
	  historial_clinico,
	   grupo_sanguineo,
	    estatura, 
		ojos,
		 color_piel, 
		 componente, 
		 categoria,
		  huella,
		   foto)
VALUES('` + mil.Persona.DatoBasico.Cedula + `',
 '` + strings.TrimSpace(mil.Persona.DatoBasico.NombrePrimero) + `', 
 '` + strings.TrimSpace(mil.Persona.DatoBasico.NombreSegundo) + `',
  '` + strings.TrimSpace(mil.Persona.DatoBasico.ApellidoPrimero) + `', 
  '` + strings.TrimSpace(mil.Persona.DatoBasico.ApellidoSegundo) + `',
   '` + grad + `', 
   '` + fechaSlashVencimiento + `',
    '` + serial + `',
	 '` + serial + `',
	  '` + codigo_comp + `',
	   ' ` + historial_clinico + `',
	    '` + grupo_sanguineo + `',
		 '` + grupo_sanguineo + `',
		  '` + estaturaStr + `',
		  '` + ojos + `',
		   '` + abreviatura_comp + `',
		    '` + cate + `', 
			'` + body + `',
			 '` + abreviatura_comp + `');	

			INSERT INTO sssifanb.carp_familiar (
					cedula_militar,
					cedula_familiar,
					nombreprimero,
					nombresegundo,
					apellidoprimero,
					apellidosegundo,
					fecha_nacimiento,
					parentesco,
					afiliado,
					f_vencimiento,
					historial_clinicof,
					grupo_sanguineof,
					donante,
					serial_carnet,
					huella,
					foto
					) VALUES ('` + mil.Persona.DatoBasico.Cedula + `',
			       '` + familiar.Persona.DatoBasico.Cedula + `',
				    '` + familiar.Persona.DatoBasico.NombrePrimero + `',
					 '` + familiar.Persona.DatoBasico.NombreSegundo + `',
					  '` + familiar.Persona.DatoBasico.ApellidoPrimero + `',
					   '` + familiar.Persona.DatoBasico.ApellidoSegundo + `',
					    '` + fechaSlashNacimiento + `,
						 '` + parentesco + `', 
						 '` + afiliado + `',
						  '` + fechaSlashVencimientof + `',
						  '` + historial_clinicoF + `',
						  '` + grupo_sanguineoF + `',
						  '` + donante + `',
						  '` + serial_carnetf + `',
						  '` + serial_carnetf + `',
						  '` + donante + `');`
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
		pen.estpencod,pen.razestpencod, per.fechadefuncion
	FROM personas per
	JOIN pers_dat_militares pdm ON  per.nropersona=pdm.nropersona
	JOIN pension pen ON pdm.nropersona=pen.nropersona
	JOIN (
		SELECT nropersona, MAX(porcprestmonto) AS porcentaje  FROM
			pension_calc
		GROUP BY nropersona
	) pcal ON pen.nropersona=pcal.nropersona
	-- LIMIT 1
	WHERE
	-- pdm.fchultimoascenso = '//' '2759874' "15236250"
	codnip='18554859'
  AND
	pen.estpencod = 'ACT'
	AND
	pen.razestpencod ='INI'
	AND
	pdm.perssituaccod IN ('FCP')`
}

// MGOActualizarPensionados Actualizando datos principales del militar desde SAMAN
func (m *Militar) MGOActualizarPensionados() (err error) {

	sq, err := sys.PostgreSQLSAMAN.Query(obtenerPensionados())
	if err != nil {
		return
	}
	i := 0
	log := ""
	for sq.Next() {
		var militar Militar
		var nropersona, tipnip, codnip, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo sql.NullString
		var sexocod, edocivilcod sql.NullString
		var componentecod, gradocod, perscategcod, perssituaccod, persclasecod sql.NullString
		var fchingcomponente, fchultimoascenso, fchpromocion, fchegreso, fechanacimiento, fechadefuncion sql.NullString
		var annoreconocido, mesreconocido, diareconocido sql.NullString
		var nrohijos, tipcuentacod, instfinancod, nrocuenta sql.NullString
		var porcentaje sql.NullFloat64
		var estpencod, razestpencod sql.NullString
		i++
		er := sq.Scan(&nropersona, &tipnip, &codnip, &nombreprimero, &nombresegundo,
			&apellidoprimero, &apellidosegundo, &sexocod, &edocivilcod, &fechanacimiento,
			&componentecod, &gradocod, &perscategcod, &perssituaccod, &persclasecod,
			&fchingcomponente, &fchultimoascenso, &fchpromocion, &fchegreso,
			&annoreconocido, &mesreconocido, &diareconocido,
			&nrohijos, &tipcuentacod, &instfinancod, &nrocuenta, &porcentaje,
			&estpencod, &razestpencod, &fechadefuncion)
		if er != nil {
			fmt.Println("Scan ", i, er.Error())
			return
		}
		militar.Persona.DatoBasico.Cedula = util.ValidarNullString(codnip)
		militar.Persona.DatoBasico.Nacionalidad = "V"
		militar.Persona.DatoBasico.NombrePrimero = strings.ToUpper(util.ValidarNullString(nombreprimero)) + " " + strings.ToUpper(util.ValidarNullString(nombresegundo))
		militar.Persona.DatoBasico.ApellidoPrimero = strings.ToUpper(util.ValidarNullString(apellidoprimero)) + " " + strings.ToUpper(util.ValidarNullString(apellidosegundo))
		militar.Persona.DatoBasico.Sexo = util.ValidarNullString(sexocod)
		militar.Persona.DatoBasico.EstadoCivil = util.ValidarNullString(edocivilcod)
		militar.Persona.DatoBasico.FechaDefuncion = util.GetFechaConvert(fechadefuncion)
		//
		militar.Grado.Abreviatura = util.ValidarNullString(gradocod)
		militar.Componente.Abreviatura = util.ValidarNullString(componentecod)
		militar.Categoria = util.ValidarNullString(perscategcod)
		militar.Clase = util.ValidarNullString(persclasecod)
		militar.Situacion = util.ValidarNullString(perssituaccod)

		militar.AnoReconocido, _ = strconv.Atoi(util.ValidarNullString(annoreconocido))
		militar.MesReconocido, _ = strconv.Atoi(util.ValidarNullString(mesreconocido))
		militar.DiaReconocido, _ = strconv.Atoi(util.ValidarNullString(diareconocido))

		militar.Pension.GradoCodigo = util.ValidarNullString(gradocod)
		militar.Pension.ComponenteCodigo = util.ValidarNullString(componentecod)
		militar.Pension.Categoria = util.ValidarNullString(perscategcod)
		militar.Pension.Clase = util.ValidarNullString(persclasecod)
		militar.Pension.Situacion = util.ValidarNullString(perssituaccod)
		militar.Pension.Estatus = util.ValidarNullString(estpencod)
		militar.Pension.Razon = util.ValidarNullString(razestpencod)

		switch militar.Pension.Estatus {
		case "ACT":
			militar.SituacionPago = "201"
			break
		case "INACT":
			if militar.Pension.Razon == "VIDA" {
				militar.SituacionPago = "207"
			} else if militar.Pension.Razon == "FA" {
				militar.SituacionPago = "209"
			} else {
				militar.SituacionPago = "210"
			}
		}

		militar.Pension.AnoServicio, _ = strconv.Atoi(util.ValidarNullString(annoreconocido))
		militar.Pension.MesServicio, _ = strconv.Atoi(util.ValidarNullString(mesreconocido))
		militar.Pension.DiaServicio, _ = strconv.Atoi(util.ValidarNullString(diareconocido))

		militar.Pension.PensionAsignada = 0
		militar.Pension.PorcentajePrestaciones = util.ValidarNullFloat64(porcentaje)
		militar.Pension.FechaPromocion = util.ValidarNullString(fchpromocion)
		militar.Pension.FechaUltimoAscenso = util.ValidarNullString(fchultimoascenso)

		militar.Pension.DatoFinanciero.Cuenta = util.ValidarNullString(nrocuenta)
		militar.Pension.DatoFinanciero.Tipo = util.ValidarNullString(tipcuentacod)
		militar.Pension.DatoFinanciero.Institucion = util.ValidarNullString(instfinancod)
		militar.Pension.NumeroHijos, _ = strconv.Atoi(util.ValidarNullString(nrohijos))

		reduc := make(map[string]interface{})
		cred := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)

		reduc["persona.datobasico.fechadefuncion"] = militar.Persona.DatoBasico.FechaDefuncion
		// reduc["fingreso"] = getFechaConvert(fchingcomponente)
		// reduc["fretiro"] = getFechaConvert(fchegreso)
		// reduc["fascenso"] = getFechaConvert(fchultimoascenso)

		reduc["areconocido"] = militar.Pension.AnoServicio
		reduc["mreconocido"] = militar.Pension.MesServicio
		reduc["dreconocido"] = militar.Pension.DiaServicio
		reduc["categoria"] = militar.Pension.Categoria
		reduc["situacion"] = militar.Pension.Situacion
		reduc["situacionpago"] = militar.SituacionPago

		reduc["clase"] = militar.Pension.Clase
		reduc["pension"] = militar.Pension

		err = cred.Update(bson.M{"id": util.ValidarNullString(codnip)}, bson.M{"$set": reduc})
		if err != nil {
			fmt.Println("# ", i, codnip)
			c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
			err = c.Insert(militar)
			if err != nil {
				//fmt.Println("Err: Insertando cedula ", militar.Persona.DatoBasico.Cedula, " Descripción: ", err.Error())
				log += militar.Persona.DatoBasico.Cedula + " Descripción: " + err.Error()
			}
			// else {
			// 	fmt.Println("Insertando cedula ", militar.Persona.DatoBasico.Cedula)
			// }
		}
		fmt.Println(i, " : ", militar.Persona.DatoBasico.Cedula)
	}
	fmt.Println("Cantidad ", i, "  LOG \n", log)
	return
}

func sqlSobrevivientes() string {
	return `

		SELECT codnip, tipcuentacod, prin.nrocuenta, instfinancod, benefporcentaje, canalliquidcod,
				tipnip, familia, nombreprimero, nombresegundo,apellidoprimero,apellidosegundo, sexocod,
				edocivilcod, fechanacimiento, cedula_autoriz FROM (
		SELECT b.codnip, pq.tipcuentacod, pq.nrocuenta, pq.instfinancod, pq.benefporcentaje, pq.canalliquidcod,
				pq.tipnip, pq.familia, pq.nombreprimero, pq.nombresegundo, pq.apellidoprimero, pq.apellidosegundo, pq.sexocod,
				pq.edocivilcod, pq.fechanacimiento
		FROM personas b JOIN (
			SELECT tb.nropersonatitular, tb.tipcuentacod, tb.nrocuenta, tb.instfinancod, tb.benefporcentaje, tb.canalliquidcod,
				tipnip, codnip as familia, nombreprimero,nombresegundo, apellidoprimero,apellidosegundo,sexocod,edocivilcod,fechanacimiento
			FROM (
				SELECT nropersona,nropersonatitular,
					tipcuentacod,nrocuenta,instfinancod, benefporcentaje, canalliquidcod
				FROM benef_montos -- limit 1
				WHERE

					estbenefmoncod is null
					AND
					benefconcepcod = 'PS'
					OR
					estbenefmoncod IN ('ACT', 'INIAC')
				) tb
		JOIN personas p ON tb.nropersona=p.nropersona ) AS pq ON pq.nropersonatitular=b.nropersona ) prin
		LEFT JOIN cuentas_bancarias cb on prin.nrocuenta=cb.nrocuenta AND prin.familia=cb.cedula_familiar
		ORDER BY codnip
		--WHERE codnip='15236250'
	`
}

// MGOActualizarSobrevivientes Actualizando datos principales del militar
func (m *Militar) MGOActualizarSobrevivientes() (err error) {

	sq, err := sys.PostgreSQLSAMAN.Query(sqlSobrevivientes())
	if err != nil {
		return
	}
	// i := 0
	// log := ""
	for sq.Next() {
		var f Familiar
		var codnip, tipcuentacod, nrocuenta, instfinancod, canalliquidcod sql.NullString
		var tipnip, familia, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo, sexocod sql.NullString
		var edocivilcod, fechanacimiento sql.NullString
		var porcentaje sql.NullFloat64
		var cautoriz sql.NullString
		var direc DatoFinanciero
		err = sq.Scan(&codnip, &tipcuentacod, &nrocuenta, &instfinancod,
			&porcentaje, &canalliquidcod, &tipnip, &familia, &nombreprimero, &nombresegundo,
			&apellidoprimero, &apellidosegundo, &sexocod, &edocivilcod, &fechanacimiento, &cautoriz,
		)
		obtenerErr(err, "")
		f.DocumentoPadre = util.ValidarNullString(codnip)
		f.Persona.DatoBasico.Cedula = util.ValidarNullString(familia)
		f.PorcentajePrestaciones = util.ValidarNullFloat64(porcentaje)
		f.Persona.DatoBasico.FechaNacimiento = util.GetFechaConvert(fechanacimiento)
		f.CondicionPago = util.ValidarNullString(canalliquidcod)

		c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
		familiar := make(map[string]interface{})
		if f.CondicionPago == "null" {
			f.CondicionPago = "CHQ"
		} else if f.CondicionPago == "BANCO" {

			direc.Cuenta = util.EliminarGuionesFecha(util.ValidarNullString(nrocuenta))
			direc.Institucion = util.ValidarNullString(instfinancod)
			direc.Tipo = util.ValidarNullString(tipcuentacod)
		}

		familiar["familiar.$.pprestaciones"] = f.PorcentajePrestaciones
		familiar["familiar.$.condicionpago"] = f.CondicionPago
		if direc.Cuenta != "" {
			familiar["familiar.$.persona.datofinanciero.0"] = direc
		}
		err = c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": familiar})
		obtenerErr(err, " id: "+f.Persona.DatoBasico.Cedula+" PADRE: "+f.DocumentoPadre)

		mil := make(map[string]interface{})
		mil["pprestaciones"] = f.PorcentajePrestaciones
		err = c.Update(bson.M{"id": f.Persona.DatoBasico.Cedula}, bson.M{"$set": mil})
		fmt.Println("Cedula ", f.Persona.DatoBasico.Cedula, " PADRE ", f.DocumentoPadre, " PORC: ", f.PorcentajePrestaciones, f.CondicionPago)
	}
	return
}

func sqlFeDeVida() string {
	return `SELECT codnip, A.familiar, razestbenefcod FROM (
		select p.codnip AS familiar, pf.nropersonatitular, pf.nropersonaautor, pf.razestbenefcod from pers_dat_benef pf
		JOIN personas p ON p.nropersona=pf.nropersona
		 WHERE pf.estbenefcod ='INACT' --AND pf.razestbenefcod IN ('VIDA', 'MAY26', 'FA', 'CAR')
		) AS A
		JOIN personas ON personas.nropersona=A.nropersonatitular
		-- where codnip='5216292'
		ORDER BY codnip`
}

// MGOActualizarFEVIDA Control de la fe de vida
func (m *Militar) MGOActualizarFEVIDA() (err error) {
	sq, err := sys.PostgreSQLSAMAN.Query(sqlFeDeVida())
	util.Error(err)
	i := 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	for sq.Next() {
		var codnip, familiar, razon sql.NullString
		i++
		err = sq.Scan(&codnip, &familiar, &razon)
		obtenerErr(err, "")
		ced := util.ValidarNullString(codnip)
		fam := util.ValidarNullString(familiar)
		raz := util.ValidarNullString(razon)
		situacionpago := make(map[string]interface{})
		situa := "299"
		if raz == "CAR" {
			situa = "206"
		} else if raz == "VIDA" {
			situa = "207"
		}
		situacionpago["familiar.$.situacionpago"] = situa
		situacionpago["familiar.$.razonpago"] = raz

		err = c.Update(bson.M{"familiar.persona.datobasico.cedula": fam, "id": ced}, bson.M{"$set": situacionpago})
		util.Error(err)
		fmt.Println("#", i, "(Cedula): "+ced)

	}
	return
}

func sqlFideicomitentesSssifanb() string {
	return `
		SELECT
			cedula, fecha_ingreso, f_ult_ascenso, f_retiro, anio_reconocido,mes_reconocido, dia_reconocido
  	FROM beneficiario
		-- WHERE cedula='12517973'
	`
}

// MGOActualizarSobrevivientesFideicomiso Actualizar segun Fideicomiso
func (m *Militar) MGOActualizarSobrevivientesFideicomiso() (err error) {

	sq, err := sys.PostgreSQLPACE.Query(sqlFideicomitentesSssifanb())
	if err != nil {
		return
	}
	i := 0
	j := 0
	for sq.Next() {
		var cedula, ingreso, ascenso, retiro, anio, mes, dia sql.NullString
		i++
		err = sq.Scan(&cedula, &ingreso, &ascenso, &retiro, &anio, &mes, &dia)
		obtenerErr(err, "")
		reduc := make(map[string]interface{})
		cred := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)

		reduc["fingreso"] = getFechaConvertGuiones(ingreso)
		reduc["fascenso"] = getFechaConvertGuiones(ascenso)

		f_retiro := util.ValidarNullString(retiro)
		if f_retiro != "" {
			reduc["fretiro"] = getFechaConvertGuiones(retiro)
		}

		reduc["areconocido"], _ = strconv.Atoi(util.ValidarNullString(anio))
		reduc["mreconocido"], _ = strconv.Atoi(util.ValidarNullString(mes))
		reduc["dreconocido"], _ = strconv.Atoi(util.ValidarNullString(dia))
		err = cred.Update(bson.M{"id": util.ValidarNullString(cedula)}, bson.M{"$set": reduc})
		if err != nil {
			j++
			fmt.Println("Update ALL ", err.Error(), i, " UFF -> ", util.ValidarNullString(cedula))
		} else {
			fmt.Println(i, " -> ", util.ValidarNullString(cedula), "Insertado")
		}

	}
	fmt.Println("Insertados: ", i, " Errados: ", j)
	return
}

// getFechaConvertGuiones Con guioness
func getFechaConvertGuiones(f sql.NullString) (dateStamp time.Time) {
	fecha := util.ValidarNullString(f)
	if fecha != "null" {
		dateString := fecha[0:10]
		dateStamp, _ = time.Parse(layout, dateString)
	}
	return
}

func obtenerErr(r error, msj string) bool {
	if r != nil {
		fmt.Println(msj, " : ", r.Error())
	}
	return true
}
