package sssifanb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

// UpsertMysqlFT inserta o actualiza un registro de Militar en la base de datos MySQL
/*func UpsertMysqlFT(mil Militar, fam Familiar) []string {

	// Mapa de códigos de color de cabello a descripciones
	colorCabelloMap := map[string]string{

		"BA": "BLANCO",
		"MA": "MARRON",
		"CA": "CASTAÑO",
		"AM": "AMARILLO",
		"AZ": "AZUL",
		"VI": "VIOLETA",
		"CV": "CALVO",
		"GR": "GRIS",
		"NE": "NEGRO",
	}

	// Mapa de códigos de color de ojos a descripciones (igual estructura que cabello)
	colorOjosMap := map[string]string{
		"CA": "CASTAÑO",
		"PA": "PARDO",
		"AM": "ÁMBAR",
		"AV": "AVELLANA",
		"VE": "VERDE",
		"AZ": "AZUL",
		"GR": "GRIS",
		"NE": "NEGRO",
		"MA": "MARRÓN",
	}

	// Mapa de códigos de color de piel a descripciones
	colorPielMap := map[string]string{
		"NE": "NEGRA",
		"TR": "TRIGUEÑA",
		"BL": "BLANCA",
		"CA": "CANELA",
		"MO": "MORENA",
		"RO": "ROSADA",
	}

	componenteMap := map[string]string{
		"EJ": "EJB",
		"AR": "ARB",
		"AV": "AMB",
		"GN": "GNB",
		"MI": "MIL",
	}

	// Obtener la descripción del color de cabello
	var descCabello string
	if mil.Persona.DatoFisionomico.ColorCabello != "" {
		if desc, ok := colorCabelloMap[mil.Persona.DatoFisionomico.ColorCabello]; ok {
			descCabello = desc
		} else {
			// Si el código no está en el mapa, usar el valor original
			descCabello = mil.Persona.DatoFisionomico.ColorCabello
		}
	}

	// Obtener la descripción del color de ojos
	var descOjos string
	if mil.Persona.DatoFisionomico.ColorOjos != "" {
		if desc, ok := colorOjosMap[mil.Persona.DatoFisionomico.ColorOjos]; ok {
			descOjos = desc
		} else {
			// Si el código no está en el mapa, usar el valor original
			descOjos = mil.Persona.DatoFisionomico.ColorOjos
		}
	}

	// Obtener la descripción del color de piel
	var descPiel string
	if mil.Persona.DatoFisionomico.ColorPiel != "" {
		if desc, ok := colorPielMap[mil.Persona.DatoFisionomico.ColorPiel]; ok {
			descPiel = desc
		} else {
			descPiel = mil.Persona.DatoFisionomico.ColorPiel
		}
	}

	var abrevComponente string
	if mil.Componente.Abreviatura != "" {
		if abrev, ok := componenteMap[mil.Componente.Abreviatura]; ok {
			abrevComponente = abrev
		} else {
			abrevComponente = mil.Componente.Abreviatura
		}
	}

	// Ejemplo de impresión de los structs mil y familiar como JSON

	bmil, _ := json.MarshalIndent(mil, "", "  ")
	bfam, _ := json.MarshalIndent(fam, "", "  ")
	fmt.Println("Militar JSON:", string(bmil))
	fmt.Println("Familiar JSON:", string(bfam))
	var queries []string
	escape := func(s string) string {
		if s == "" {
			return ""
		}
		return strings.ReplaceAll(s, "'", "''")
	}

	// Validación básica de datos requeridos
	if mil.Persona.DatoBasico.Cedula == "" {
		return queries
	}

	// Formateo de fechas seguro
	formatDate := func(t time.Time) string {
		if t.IsZero() {
			return "NULL"
		}
		return "'" + t.Format("02-01-2006") + "'"
	}

	// Preparación de datos esenciales
	datos := struct {
		cedula, nombre1, nombre2, apellido1, apellido2  string
		grado, serial, codigoComp, cabello, histClinico string
		grupoSanguineo, estatura, ojos, colorPiel       string
		abrevComp, categoria                            string
		fechaVencimiento                                string
	}{
		cedula:    escape(mil.Persona.DatoBasico.Cedula),
		nombre1:   escape(strings.TrimSpace(mil.Persona.DatoBasico.NombrePrimero)),
		nombre2:   escape(strings.TrimSpace(mil.Persona.DatoBasico.NombreSegundo)),
		apellido1: escape(strings.TrimSpace(mil.Persona.DatoBasico.ApellidoPrimero)),
		apellido2: escape(strings.TrimSpace(mil.Persona.DatoBasico.ApellidoSegundo)),
		grado:     escape(mil.Grado.Descripcion),
		serial: func() string {
			s := strings.TrimSpace(mil.TIM.Serial)
			if s == "" {
				return "N/A"
			}
			return escape(s)
		}(),
		codigoComp:       escape(mil.CodigoComponente),
		cabello:          escape(descCabello),
		histClinico:      escape(mil.NumeroHistoria),
		grupoSanguineo:   escape(mil.Persona.DatoFisionomico.GrupoSanguineo),
		estatura:         fmt.Sprintf("%.2f", mil.Persona.DatoFisionomico.Estatura),
		ojos:             escape(descOjos),
		colorPiel:        escape(descPiel),
		abrevComp:        escape(abrevComponente),
		categoria:        escape(obtenerCategoria(mil.Categoria)),
		fechaVencimiento: formatDate(mil.TIM.FechaVencimiento),
	}

	// Construcción de la query con validación de valores NULL
	queryMilitar := fmt.Sprintf(`
    INSERT INTO sssifanb.carp_militar (
        cedula, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo,
        grado, fecha_vencimiento, serial_carnet, codigo_comp, cabello,
        historial_clinico, grupo_sanguineo, estatura, ojos, color_piel,
        componente, categoria, huella, foto, QR, firma
    ) VALUES (
        '%s', %s, %s, %s, %s,
        %s, %s, %s, %s, %s,
        %s, %s, %s, %s, %s,
        %s, %s, %s, %s, %s, %s
    ) ON DUPLICATE KEY UPDATE
        nombreprimero = COALESCE(VALUES(nombreprimero), nombreprimero),
        nombresegundo = COALESCE(VALUES(nombresegundo), nombresegundo),
        apellidoprimero = COALESCE(VALUES(apellidoprimero), apellidoprimero),
        apellidosegundo = COALESCE(VALUES(apellidosegundo), apellidosegundo),
        grado = COALESCE(VALUES(grado), grado),
        fecha_vencimiento = COALESCE(VALUES(fecha_vencimiento), fecha_vencimiento),
		serial_carnet = VALUES(serial_carnet),
        codigo_comp = COALESCE(VALUES(codigo_comp), codigo_comp),
        cabello = COALESCE(VALUES(cabello), cabello),
        historial_clinico = COALESCE(VALUES(historial_clinico), historial_clinico),
        grupo_sanguineo = COALESCE(VALUES(grupo_sanguineo), grupo_sanguineo),
        estatura = COALESCE(VALUES(estatura), estatura),
        ojos = COALESCE(VALUES(ojos), ojos),
        color_piel = COALESCE(VALUES(color_piel), color_piel),
        componente = COALESCE(VALUES(componente), componente),
        categoria = COALESCE(VALUES(categoria), categoria),
        huella = COALESCE(VALUES(huella), huella),
        foto = COALESCE(VALUES(foto), foto),
		QR = COALESCE(VALUES (QR), QR),
		firma = COALESCE(VALUES (firma), firma)`,
		datos.cedula,
		wrapValue(datos.nombre1),
		wrapValue(datos.nombre2),
		wrapValue(datos.apellido1),
		wrapValue(datos.apellido2),
		wrapValue(datos.grado),
		datos.fechaVencimiento,
		wrapValue(datos.serial),
		wrapValue(datos.codigoComp),
		wrapValue(datos.cabello),
		wrapValue(datos.histClinico),
		wrapValue(datos.grupoSanguineo),
		datos.estatura,
		wrapValue(datos.ojos),
		wrapValue(datos.colorPiel),
		wrapValue(datos.abrevComp),
		wrapValue(datos.categoria),
		wrapValue(""),
		wrapValue(""),
		wrapValue(""),
		wrapValue(""),
	)

	// Limpieza de espacios y nueva línea
	queryMilitar = strings.NewReplacer(
		"\n", " ",
		"\t", "",
		"  ", " ",
	).Replace(queryMilitar)
	queryMilitar = strings.TrimSpace(queryMilitar)

	queries = append(queries, queryMilitar)
	return queries
}*/
func UpsertMysqlFT(mil Militar, fam Familiar) []string {
	// Mapa de códigos de color de cabello a descripciones
	colorCabelloMap := map[string]string{
		"BA": "BLANCO",
		"MA": "MARRON",
		"CA": "CASTAÑO",
		"AM": "AMARILLO",
		"AZ": "AZUL",
		"VI": "VIOLETA",
		"CV": "CALVO",
		"GR": "GRIS",
		"NE": "NEGRO",
	}

	// Mapa de códigos de color de ojos a descripciones
	colorOjosMap := map[string]string{
		"CA": "CASTAÑO",
		"PA": "PARDO",
		"AM": "ÁMBAR",
		"AV": "AVELLANA",
		"VE": "VERDE",
		"AZ": "AZUL",
		"GR": "GRIS",
		"NE": "NEGRO",
		"MA": "MARRÓN",
	}

	// Mapa de códigos de color de piel a descripciones
	colorPielMap := map[string]string{
		"NE": "NEGRA",
		"TR": "TRIGUEÑA",
		"BL": "BLANCA",
		"CA": "CANELA",
		"MO": "MORENA",
		"RO": "ROSADA",
	}

	componenteMap := map[string]string{
		"EJ": "EJB",
		"AR": "ARB",
		"AV": "AMB",
		"GN": "GNB",
		"MI": "MIL",
	}

	// Obtener descripciones de los mapas
	descCabello := getMappedValue(mil.Persona.DatoFisionomico.ColorCabello, colorCabelloMap)
	descOjos := getMappedValue(mil.Persona.DatoFisionomico.ColorOjos, colorOjosMap)
	descPiel := getMappedValue(mil.Persona.DatoFisionomico.ColorPiel, colorPielMap)
	abrevComponente := getMappedValue(mil.Componente.Abreviatura, componenteMap)

	// Log de datos de entrada
	bmil, _ := json.MarshalIndent(mil, "", "  ")
	bfam, _ := json.MarshalIndent(fam, "", "  ")
	fmt.Println("Militar JSON:", string(bmil))
	fmt.Println("Familiar JSON:", string(bfam))

	var queries []string

	// Validación básica de datos requeridos
	if mil.Persona.DatoBasico.Cedula == "" {
		log.Println("Error: Cédula vacía - no se puede realizar la operación")
		return queries
	}

	// Preparación de datos
	datos := prepareData(mil, descCabello, descOjos, descPiel, abrevComponente)

	// Construcción de la query con la expresión QR
	queryMilitar := buildUpsertQuery(datos, mil.Persona.DatoBasico.Cedula)
	queries = append(queries, queryMilitar)

	return queries
}

// Función auxiliar para obtener valores mapeados
func getMappedValue(code string, mapping map[string]string) string {
	if code == "" {
		return ""
	}
	if desc, ok := mapping[code]; ok {
		return desc
	}
	return code
}

// Función auxiliar para preparar los datos
func prepareData(mil Militar, descCabello, descOjos, descPiel, abrevComponente string) struct {
	cedula, nombre1, nombre2, apellido1, apellido2  string
	grado, serial, codigoComp, cabello, histClinico string
	grupoSanguineo, estatura, ojos, colorPiel       string
	abrevComp, categoria                            string
	fechaVencimiento                                string
} {
	return struct {
		cedula, nombre1, nombre2, apellido1, apellido2  string
		grado, serial, codigoComp, cabello, histClinico string
		grupoSanguineo, estatura, ojos, colorPiel       string
		abrevComp, categoria                            string
		fechaVencimiento                                string
	}{
		cedula:           escape(mil.Persona.DatoBasico.Cedula),
		nombre1:          escape(strings.TrimSpace(mil.Persona.DatoBasico.NombrePrimero)),
		nombre2:          escape(strings.TrimSpace(mil.Persona.DatoBasico.NombreSegundo)),
		apellido1:        escape(strings.TrimSpace(mil.Persona.DatoBasico.ApellidoPrimero)),
		apellido2:        escape(strings.TrimSpace(mil.Persona.DatoBasico.ApellidoSegundo)),
		grado:            escape(mil.Grado.Descripcion),
		serial:           getSerial(mil.TIM.Serial),
		codigoComp:       escape(mil.CodigoComponente),
		cabello:          escape(descCabello),
		histClinico:      escape(mil.NumeroHistoria),
		grupoSanguineo:   escape(mil.Persona.DatoFisionomico.GrupoSanguineo),
		estatura:         fmt.Sprintf("%.2f", mil.Persona.DatoFisionomico.Estatura),
		ojos:             escape(descOjos),
		colorPiel:        escape(descPiel),
		abrevComp:        escape(abrevComponente),
		categoria:        escape(obtenerCategoria(mil.Categoria)),
		fechaVencimiento: formatDateForMySQL(mil.TIM.FechaVencimiento),
	}
}

// Función auxiliar para formatear fecha
func formatDateForMySQL(t time.Time) string {
	if t.IsZero() {
		return "NULL"
	}
	return "'" + t.Format("2006-01-02") + "'"
}

// Función auxiliar para manejar el serial
func getSerial(serial string) string {
	s := strings.TrimSpace(serial)
	if s == "" {
		return "NULL"
	}
	return escape(s)
}

// Función auxiliar para escapar strings
func escape(s string) string {
	if s == "" {
		return ""
	}
	return strings.ReplaceAll(s, "'", "''")
}

// Función auxiliar para envolver valores

// Función auxiliar para construir la query de upsert
func buildUpsertQuery(datos struct {
	cedula, nombre1, nombre2, apellido1, apellido2  string
	grado, serial, codigoComp, cabello, histClinico string
	grupoSanguineo, estatura, ojos, colorPiel       string
	abrevComp, categoria                            string
	fechaVencimiento                                string
}, cedula string) string {
	// Generar la expresión QR
	qrExpression := fmt.Sprintf("CONCAT('https://apps.ipsfa.gob.ve/app/#/certificado/', MD5(CONCAT('CI-','%s')))", cedula)

	/*query := fmt.Sprintf(`
	    INSERT INTO sssifanb.carp_militar (
	        cedula, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo,
	        grado, fecha_vencimiento, serial_carnet, codigo_comp, cabello,
	        historial_clinico, grupo_sanguineo, estatura, ojos, color_piel,
	        componente, categoria, huella, foto, QR, firma
	    ) VALUES (
	        '%s', %s, %s, %s, %s,
	        %s, %s, %s, %s, %s,
	        %s, %s, %s, %s, %s,
	        %s, %s, %s, %s, %s, %s
	    ) ON DUPLICATE KEY UPDATE
	        nombreprimero = VALUES(nombreprimero),
	        nombresegundo = VALUES(nombresegundo),
	        apellidoprimero = VALUES(apellidoprimero),
	        apellidosegundo = VALUES(apellidosegundo),
	        grado = VALUES(grado),
	        fecha_vencimiento = VALUES(fecha_vencimiento),
	        serial_carnet = VALUES(serial_carnet),
	        codigo_comp = VALUES(codigo_comp),
	        cabello = VALUES(cabello),
	        historial_clinico = VALUES(historial_clinico),
	        grupo_sanguineo = VALUES(grupo_sanguineo),
	        estatura = VALUES(estatura),
	        ojos = VALUES(ojos),
	        color_piel = VALUES(color_piel),
	        componente = VALUES(componente),
	        categoria = VALUES(categoria),
	        huella = VALUES(huella),
	        foto = VALUES(foto),
	        QR = VALUES(QR),
	        firma = VALUES(firma)`,
			datos.cedula,
			wrapValue(datos.nombre1),
			wrapValue(datos.nombre2),
			wrapValue(datos.apellido1),
			wrapValue(datos.apellido2),
			wrapValue(datos.grado),
			datos.fechaVencimiento,
			wrapValue(datos.serial),
			wrapValue(datos.codigoComp),
			wrapValue(datos.cabello),
			wrapValue(datos.histClinico),
			wrapValue(datos.grupoSanguineo),
			datos.estatura,
			wrapValue(datos.ojos),
			wrapValue(datos.colorPiel),
			wrapValue(datos.abrevComp),
			wrapValue(datos.categoria),
			wrapValue(""),
			wrapValue(""),
			wrapValue(""),
			wrapValue(""),
		)*/
	// Construcción de la query corregida
	query := fmt.Sprintf(`
    INSERT INTO sssifanb.carp_militar (
        cedula, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo,
        grado, fecha_vencimiento, serial_carnet, codigo_comp, cabello,
        historial_clinico, grupo_sanguineo, estatura, ojos, color_piel,
        componente, categoria, huella, foto, QR, firma
    ) VALUES (
        '%s', %s, %s, %s, %s,
        %s, %s, %s, %s, %s,
        %s, %s, %s, %s, %s,
        %s, %s, %s, %s, %s, %s
    ) ON DUPLICATE KEY UPDATE
        nombreprimero = IF(VALUES(serial_carnet) != serial_carnet, VALUES(nombreprimero), nombreprimero),
        nombresegundo = IF(VALUES(serial_carnet) != serial_carnet, VALUES(nombresegundo), nombresegundo),
        apellidoprimero = IF(VALUES(serial_carnet) != serial_carnet, VALUES(apellidoprimero), apellidoprimero),
        apellidosegundo = IF(VALUES(serial_carnet) != serial_carnet, VALUES(apellidosegundo), apellidosegundo),
        grado = IF(VALUES(serial_carnet) != serial_carnet, VALUES(grado), grado),
        fecha_vencimiento = IF(VALUES(serial_carnet) != serial_carnet, VALUES(fecha_vencimiento), fecha_vencimiento),
        serial_carnet = IF(VALUES(serial_carnet) != serial_carnet, VALUES(serial_carnet), serial_carnet),
        codigo_comp = IF(VALUES(serial_carnet) != serial_carnet, VALUES(codigo_comp), codigo_comp),
        cabello = IF(VALUES(serial_carnet) != serial_carnet, VALUES(cabello), cabello),
        historial_clinico = IF(VALUES(serial_carnet) != serial_carnet, VALUES(historial_clinico), historial_clinico),
        grupo_sanguineo = IF(VALUES(serial_carnet) != serial_carnet, VALUES(grupo_sanguineo), grupo_sanguineo),
        estatura = IF(VALUES(serial_carnet) != serial_carnet, VALUES(estatura), estatura),
        ojos = IF(VALUES(serial_carnet) != serial_carnet, VALUES(ojos), ojos),
        color_piel = IF(VALUES(serial_carnet) != serial_carnet, VALUES(color_piel), color_piel),
        componente = IF(VALUES(serial_carnet) != serial_carnet, VALUES(componente), componente),
        categoria = IF(VALUES(serial_carnet) != serial_carnet, VALUES(categoria), categoria),
        huella = IF(VALUES(serial_carnet) != serial_carnet, VALUES(huella), huella),
        foto = IF(VALUES(serial_carnet) != serial_carnet, VALUES(foto), foto),
        QR = IF(VALUES(serial_carnet) != serial_carnet, VALUES(QR), QR),
        firma = IF(VALUES(serial_carnet) != serial_carnet, VALUES(firma), firma)`,
		datos.cedula,
		wrapValue(datos.nombre1),
		wrapValue(datos.nombre2),
		wrapValue(datos.apellido1),
		wrapValue(datos.apellido2),
		wrapValue(datos.grado),
		datos.fechaVencimiento,
		wrapValue(datos.serial),
		wrapValue(datos.codigoComp),
		wrapValue(datos.cabello),
		wrapValue(datos.histClinico),
		wrapValue(datos.grupoSanguineo),
		datos.estatura,
		wrapValue(datos.ojos),
		wrapValue(datos.colorPiel),
		wrapValue(datos.abrevComp),
		wrapValue(datos.categoria),
		wrapValue(""),
		wrapValue(""),
		qrExpression,
		wrapValue(""),
	)
	// Limpieza de la query
	return strings.NewReplacer("\n", " ", "\t", "", "  ", " ").Replace(strings.TrimSpace(query))
}

/* func UpsertMysqlFTFamiliar(fam Familiar) []string {
	var queries []string
	escape := func(s string) string {
		return strings.ReplaceAll(s, "'", "''")
	}

	// Validación básica de datos requeridos
	if fam.Persona.DatoBasico.Cedula == "" || fam.DocumentoPadre == "" {
		fmt.Println("Error: Cédula del familiar o documento del padre vacío")
		return queries
	}

	// Preparación de datos
	fechaNacimiento := fam.Persona.DatoBasico.FechaNacimiento.Format("2006-01-02")
	parentesco := obtenerParentesco(fam.Parentesco, fam.Persona.DatoBasico.Sexo)

	// Construir campo afiliado (ej: "CAPITAN PEREZ JUAN 12345678")
	afiliado := fam.DocumentoPadre

	// Preparar fecha de vencimiento (usar fecha actual + 1 año si no viene)
	fechaVencimiento := fam.TIF.FechaVencimiento.Format("2006-01-02")
	if fam.TIF.FechaVencimiento.IsZero() {
		fechaVencimiento = time.Now().AddDate(1, 0, 0).Format("2006-01-02")
	}

	// Convertir donante a valor numérico (0 o 1)
	donante := 0
	if strings.ToUpper(fam.Donante) == "S" {
		donante = 1
	}

	// Query UPSERT completa con todos los campos
	query := fmt.Sprintf(`
    INSERT INTO sssifanb.carp_familiar (
        id, cedula_familiar, cedula_militar, nombreprimerof, apellidoprimerof,
        fecha_nacimiento, parentesco, afiliado, f_vencimiento,
        historial_clinicof, grupo_sanguineof, donante, serial_carnet,
        huella, foto
    ) VALUES (
        0, '%s', '%s', '%s', '%s',
        '%s', '%s', '%s', '%s',
        '%s', '%s', %d, '%s',
        '%s', '%s'
    ) ON DUPLICATE KEY UPDATE
        cedula_familiar = VALUES(cedula_familiar),
        cedula_militar = VALUES(cedula_militar),
        nombreprimerof = VALUES(nombreprimerof),
        apellidoprimerof = VALUES(apellidoprimerof),
        fecha_nacimiento = VALUES(fecha_nacimiento),
        parentesco = VALUES(parentesco),
        afiliado = VALUES(afiliado),
        f_vencimiento = VALUES(f_vencimiento),
        historial_clinicof = VALUES(historial_clinicof),
        grupo_sanguineof = VALUES(grupo_sanguineof),
        donante = VALUES(donante),
        serial_carnet = VALUES(serial_carnet),
        huella = VALUES(huella),
        foto = VALUES(foto)`,
		escape(fam.Persona.DatoBasico.Cedula),
		escape(fam.DocumentoPadre),
		escape(strings.TrimSpace(fam.Persona.DatoBasico.NombrePrimero)),
		escape(strings.TrimSpace(fam.Persona.DatoBasico.ApellidoPrimero)),
		fechaNacimiento,
		escape(parentesco),
		escape(afiliado),
		fechaVencimiento,
		escape(fam.HistoriaMedica),
		escape(fam.Persona.DatoFisionomico.GrupoSanguineo),
		donante,
		escape(fam.TIF.ID),
		escape(fam.TIF.ID), // Usar el mismo valor para huella
		escape("foto_"+fam.Persona.DatoBasico.Cedula), // Ejemplo de nombre de foto
	)

	// Limpieza de espacios en blanco
	query = strings.Join(strings.Fields(query), " ")
	queries = append(queries, query)

	return queries
//}*/

func UpsertMysqlFTFamiliar(fam Familiar, mil Militar) []string {
	var queries []string
	escape := func(s string) string {
		// Convertir a mayúsculas ANTES de escapar las comillas
		return strings.ReplaceAll(strings.ToUpper(s), "'", "''")
	}

	// Validación básica de datos requeridos
	if fam.Persona.DatoBasico.Cedula == "" || fam.DocumentoPadre == "" {
		fmt.Println("Error: Cédula del familiar o documento del padre vacío")
		return queries
	}

	// Construir campo afiliado
	afiliado := fmt.Sprintf("%s - %s %s CI:%s",
		strings.ToUpper(mil.Grado.Abreviatura),
		escape(strings.TrimSpace(mil.Persona.DatoBasico.NombrePrimero)),
		escape(strings.TrimSpace(mil.Persona.DatoBasico.ApellidoPrimero)),
		escape(mil.Persona.DatoBasico.Cedula))

	// Preparar fechas
	fechaNacimiento := fam.Persona.DatoBasico.FechaNacimiento.Format("2006-01-02")
	fechaVencimiento := fam.TIF.FechaVencimiento.Format("2006-01-02")
	if fam.TIF.FechaVencimiento.IsZero() {
		fechaVencimiento = time.Now().AddDate(1, 0, 0).Format("2006-01-02")
	}

	// Determinar valor de donante
	donanteValue := "'NO'"
	if strings.ToUpper(fam.Donante) == "S" {
		donanteValue = "'SI'"
	}

	// Query UPSERT corregida
	query := fmt.Sprintf(`
    INSERT INTO sssifanb.carp_familiar (
        id, cedula_familiar, cedula_militar, nombreprimerof, apellidoprimerof,
        fecha_nacimiento, parentesco, afiliado, f_vencimiento,
        historial_clinicof, grupo_sanguineof, donante, serial_carnet,
        huella, foto, QRF
    ) VALUES (
        0, '%s', '%s', '%s', '%s',
        '%s', '%s', '%s', '%s',
        '%s', '%s', %s, '%s',
        '%s', '%s', '%s'
    ) ON DUPLICATE KEY UPDATE
        nombreprimerof = VALUES(nombreprimerof),
        apellidoprimerof = VALUES(apellidoprimerof),
        fecha_nacimiento = VALUES(fecha_nacimiento),
        parentesco = VALUES(parentesco),
        afiliado = VALUES(afiliado),
        f_vencimiento = VALUES(f_vencimiento),
        historial_clinicof = VALUES(historial_clinicof),
        grupo_sanguineof = VALUES(grupo_sanguineof),
        donante = VALUES(donante),
        serial_carnet = VALUES(serial_carnet),
        huella = VALUES(huella),
        foto = VALUES(foto),
        QRF = VALUES(QRF)`,
		escape(fam.Persona.DatoBasico.Cedula),
		escape(fam.DocumentoPadre),
		escape(strings.TrimSpace(fam.Persona.DatoBasico.NombrePrimero)),
		escape(strings.TrimSpace(fam.Persona.DatoBasico.ApellidoPrimero)),
		fechaNacimiento,
		escape(strings.ToUpper(obtenerParentesco(fam.Parentesco, fam.Persona.DatoBasico.Sexo))),
		escape(afiliado),
		fechaVencimiento,
		escape(fam.HistoriaMedica),
		escape(fam.Persona.DatoFisionomico.GrupoSanguineo),
		donanteValue,
		escape(fam.TIF.Serial), // CORRECCIÓN: Usar Serial en lugar de ID
		wrapValue(""),
		wrapValue(""),
		wrapValue(""),
	)

	query = strings.Join(strings.Fields(query), " ")
	queries = append(queries, query)

	return queries
}

// Función auxiliar para manejar valores vacíos
func wrapValue(s string) string {
	if s == "" {
		return "NULL"
	}
	return "'" + s + "'"
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
