package estadistica

func personaMilitar() {
	sSQL := `DROP TABLE IF EXISTS analisis.personas_militares;
	CREATE TABLE analisis.personas_militares AS (
		SELECT cedula,nropersona,fr FROM (
		SELECT *
		FROM (		SELECT codnip,
			count(codnip) AS fr
		 FROM personas GROUP BY codnip
		 ) AS tb --WHERE tb.fr > 1
		) AS pers JOIN (
		SELECT	codnip AS cedula, prs.nropersona
		FROM personas AS prs	JOIN	pers_dat_militares AS pd ON  prs.nropersona=pd.nropersona
		) AS militar ON pers.codnip=militar.cedula );
		CREATE INDEX cedula_idx ON analisis.personas_militares (cedula);
	`
	errorG(sSQL)
}

//historialMilitares
func historialFamiliares() {
	sSQL := `DROP TABLE IF EXISTS analisis.familiares; CREATE TABLE analisis.familiares AS
	(SELECT nropersona, count(nropersona) as familiar
	FROM pers_relaciones GROUP BY nropersona);`
	errorG(sSQL)
}

//historialMilitares
func historialMilitares() {
	sSQL := `DROP TABLE IF EXISTS analisis.militares; CREATE TABLE analisis.militar AS
	(SELECT nropersona, count(nropersona) as militar
	FROM ipsfa_grado_x_pers GROUP BY nropersona);`
	errorG(sSQL)
}

//historialReembolso
func historialPension() {
	sSQL := `DROP TABLE IF EXISTS analisis.pension; CREATE TABLE analisis.pension AS
	(SELECT nropersona, count(nropersona) as pension
	FROM pension GROUP BY nropersona);`
	errorG(sSQL)
}

//historialCreditos
func historialCreditos() {
	sSQL := `DROP TABLE IF EXISTS analisis.creditos; CREATE TABLE analisis.creditos AS
	(SELECT nropersona, count(nropersona) as creditos
	FROM creditos GROUP BY nropersona);`
	errorG(sSQL)
}

//historialReembolso
func historialReembolso() {
	sSQL := `DROP TABLE IF EXISTS analisis.reembolso; CREATE TABLE analisis.reembolso AS
	(SELECT nropersona, count(nropersona) as creditos_fr
	FROM ci_reembolso_solic GROUP BY nropersona);`
	errorG(sSQL)
}

//reduccion
func reduccion() string {
	return `SELECT
			cedula_saman AS cedulas, cedula_pace AS cedulap,np,
			tipnip,nombreprimero,nombresegundo,apellidoprimero,apellidosegundo,fechanacimiento,
			sexocod,edocivilcod,
			perscategcod, perssituaccod,persclasecod,
			fchingcomponente,fchultimoascenso,fchegreso,
			annoreconocido,mesreconocido,diareconocido,
			componentecod,componentenombre,componentesiglas,
			gradocod,gradocodrangoid,gradonombrecorto,gradonombrelargo,
			gradop,componentep,tipcuentacod,instfinancod,nrocuenta,nrohijos,
			anio_reconocido,mes_reconocido,dia_reconocido,f_retiro,n_hijos,numero_cuenta,
			pensioncategoria, pensionsituacion, pensionclase,
			annototservicio,mestotservicio, diatotservicio,
			pensionasignada,porcprestmonto,pensionpromocion,pensionascenso,
			componente_id, grado_id
	 FROM (
						SELECT DISTINCT ON (A.cedula_saman) cedula_saman, cedula_pace, np FROM (
						SELECT * FROM analisis.reducciones WHERE militar > 0 ORDER BY militar DESC ) AS A
					) AS TBL JOIN (
						SELECT p.codnip,p.nropersona,
							p.tipnip,p.nombreprimero,nombresegundo,apellidoprimero,apellidosegundo,fechanacimiento,
							p.sexocod,p.edocivilcod,
							pm.perscategcod, pm.perssituaccod,pm.persclasecod,
							pm.fchingcomponente,pm.fchultimoascenso,pension.fchegreso,
							pm.annoreconocido,pm.mesreconocido,pm.diareconocido,
							icom.componentecod,icom.componentenombre,icom.componentesiglas,
							igra.gradocod,igra.gradocodrangoid,igra.gradonombrecorto,igra.gradonombrelargo,
							pension.gradocod AS gradop, pension.componentecod AS componentep,
							pension.tipcuentacod, pension.instfinancod, pension.nrocuenta, pension.nrohijos,
							bnf.anio_reconocido, bnf.mes_reconocido,bnf.dia_reconocido, bnf.f_retiro, bnf.n_hijos,
							bnf.numero_cuenta,
							pension.perscategcod AS pensioncategoria, pension.perssituaccod AS pensionsituacion,
							pension.persclasecod AS pensionclase,
							pension.annototservicio, pension.mestotservicio, pension.diatotservicio,
							pension.pensionasignada,porcprestmonto,
							pension.fchpromocion AS pensionpromocion, pension.fchultimoascenso AS pensionascenso,
							bnf.componente_id, bnf.grado_id
						FROM pers_dat_militares AS pm
						LEFT JOIN pension ON pension.nropersona=pm.nropersona
						JOIN personas AS p ON pm.nropersona=p.nropersona
						LEFT JOIN beneficiario AS bnf ON p.codnip=bnf.cedula
						JOIN ipsfa_componentes AS icom ON pm.componentecod=icom.componentecod
						JOIN ipsfa_grados AS igra ON pm.gradocod=igra.gradocod AND pm.componentecod=igra.componentecod
						LEFT JOIN (SELECT DISTINCT ON (nropersona) nropersona,porcprestmonto FROM pension_calc) AS porc
							ON porc.nropersona=pension.nropersona
					) AS B ON B.nropersona = TBL.np  --limit 1 -- WHERE cedula_saman='16872776' --  WHERE B.perssituaccod = 'ACT' --`
}

//obtenerHistorialFamiliares
func obtenerHistorialFamiliares() string {
	return `SELECT
						AR.cedula_saman,
						p.codnip,pr.nropersonarel,pr.persrelstipcod,p.tipnip,
						p.nombreprimero,p.nombresegundo,
						p.apellidoprimero,p.apellidosegundo,
						p.fechanacimiento,p.sexocod,p.edocivilcod,
						pm.nropersona AS militar
					FROM
						analisis.reducciones as AR
					JOIN pers_relaciones pr ON AR.np=pr.nropersona
					JOIN personas p ON pr.nropersonarel=p.nropersona
					LEFT JOIN pers_dat_militares AS pm ON pm.nropersona=pr.nropersonarel
					 ORDER BY AR.cedula_saman --LIMIT 100`
	//WHERE pr.nropersona IN (1393199,79227)

}

//obtenerHistorialMilitar
func obtenerHistorialMilitar() string {
	return `
			SELECT
				cedula_saman, ipg.componentecod, ipg.gradocod, ipg.perscategcod,
				ipg.persclasecod, ipg.perssituaccod, ipg.gradofchobten, ipg.gradoresuelto,
				ipg.gradonroenresuelto,ipg.gradofchrecipsfa,
				ipg.auditfechacambio,ipg.audithoracambio,
				ipg.auditfechacreacion,ipg.audithoracreacion,ipg.razonhistcod
				FROM
				analisis.reducciones as AR
				JOIN personas p ON AR.np=p.nropersona
				JOIN ipsfa_grado_x_pers ipg ON ipg.nropersona = AR.np
			ORDER BY AR.cedula_saman,p.nropersona, gradofchrecipsfa
			--LIMIT 10`
	//WHERE p.nropersona IN (1393199,79227)

}

func obtenerPrimaProfesional() string {
	return `SELECT pr.codnip, pen.porcprimaprof
	FROM pension_calc AS pen
	JOIN personas AS pr ON pen.nropersona=pr.nropersona
	WHERE pensionvigente ='S' AND pensionasignada > 0`
}

func obtenerPrimaNoAscenso() string {
	return `SELECT pr.codnip, pen.porcprimanoascenso
	FROM pension_calc AS pen
	JOIN personas AS pr ON pen.nropersona=pr.nropersona
	WHERE pensionvigente ='S' AND pensionasignada > 0`
}

func obtenerPrimaEspecial() string {
	return `SELECT pr.codnip, pen.primaespecial
	FROM pension_calc AS pen
	JOIN personas AS pr ON pen.nropersona=pr.nropersona
	WHERE pensionvigente ='S' AND pensionasignada > 0`
}

func obtenerPensionadosAntes2008GN() string {
	return `SELECT pr.codnip, pen.gradocod FROM pension AS pen
	JOIN personas AS pr ON pen.nropersona=pr.nropersona
	WHERE gradocod IN ('SAY','S1','S2','C1','C2','DTGDO','GN')  AND  componentecod = 'GN'`
}

func obtenerPensionadosAntes2008AV() string {
	return `SELECT pr.codnip, pen.gradocod FROM pension AS pen
	JOIN personas AS pr ON pen.nropersona=pr.nropersona
	WHERE gradocod IN ('ATG','ATA','ATM','AT1','AT2','AT3','AT')  AND  componentecod = 'AV'`
}

func obtenerFechaVencimiento() string {
	return `SELECT codnip,fchvctocarnet FROM pers_dat_afiliac af
		JOIN personas pr ON af.nropersona=pr.nropersona
		WHERE fchvctocarnet != ''`
}

func obtenerFechaDefuncion() string {
	return `select codnip, fechadefuncion from personas where fechadefuncion != ''`
}

func obtenerCuentaBancaria() string {
	return `
			SELECT
			AR.cedula_saman, nrocuenta, tipcuentacod, instfinancod, usocuentacod
		FROM
		analisis.reducciones as AR
		JOIN pers_cta_bancarias  AS cta
		ON AR.np=cta.nropersona ORDER BY AR.cedula_saman DESC
		--LIMIT 10
		`

}

func obtenerComponenteGrado() string {
	return `
		SELECT c.componentecod, componentenombre, componentesiglas, gradocod,gradocodrangoid,gradonombrecorto,
		gradonombrelargo
		FROM ipsfa_grados AS g JOIN ipsfa_componentes AS c ON g.componentecod=c.componentecod
		ORDER BY c.componentepriorpt,g.gradocodrangoid
	`
}

func obtenerEstados() string { //MySQL
	return `SELECT estado, iso_31662, ciudad,capital FROM analisis.estados JOIN analisis.ciudades ON estados.id_estado=ciudades.id_estado`
}
func obtenerMunicipios() string { //MySQL
	return `SELECT estado, municipio FROM analisis.estados JOIN analisis.municipios ON estados.id_estado=municipios.id_estado
ORDER BY estado,municipio`
}

func obtenerMunicipiosParroquia() string { //MySQL
	return `SELECT estado, municipio, parroquia FROM analisis.estados JOIN analisis.municipios ON estados.id_estado=municipios.id_estado
			JOIN  analisis.parroquias ON analisis.municipios.id_municipio= analisis.parroquias.id_municipio
			ORDER BY estado,municipio,parroquia`
}

// func InsertarPace(militar sssifanb.Militar) string {
// 	return ""
// 	// \'' . $this->fecha_ingreso . '\',
// 	// \'' . $this->fecha_ultimo_ascenso . '\',
// 	// \'' . $this->fecha_ingreso_sistema . '\',
// 	// \'' . $this->fecha_retiro . '\',
// 	// \'' . $this->fecha_retiro_efectiva . '\',
// 	// \'' . $this->fecha_creacion . '\',
// 	// \'' . $this->fecha_ultima_modificacion . '\',
// 	// \'' . $this->fecha_reincorporacion . '\'
//
// 	// return `INSERT INTO hist_beneficiario (
// 	// 		status_id,
// 	// 		componente_id,
// 	// 		grado_id,
// 	// 		cedula,
// 	// 		nombres,
// 	// 		apellidos,
// 	// 		tiempo_servicio,
// 	// 		fecha_ingreso,
// 	// 		edo_civil,
// 	// 		n_hijos,
// 	// 		f_ult_ascenso,
// 	// 		anio_reconocido,
// 	// 		mes_reconocido,
// 	// 		dia_reconocido,
// 	// 		f_ingreso_sistema,
// 	// 		f_retiro,
// 	// 		f_retiro_efectiva,
// 	// 		st_no_ascenso,
// 	// 		numero_cuenta,
// 	// 		st_profesion,
// 	// 		sexo,
// 	// 		f_creacion,
// 	// 		usr_creacion,
// 	// 		f_ult_modificacion,
// 	// 		usr_modificacion,
// 	// 		observ_ult_modificacion,
// 	// 		motivo_paralizacion,
// 	// 		f_reincorporacion
// 	// 	) VALUES ';
// 	//
// 	// 	$sInsertar .= '(
// 	// 		\'' . $this->estatus_activo . '\',
// 	// 		` + militar.Fideicomiso.ComponenteCodigo + `,
// 	// 		` + militar.Fideicomiso.GradoCodigo + `,
// 	// 		'` + militar.Persona.DatoBasico.Cedula + `',
// 	// 		'` + militar.Persona.DatoBasico.ConcatenarNombre() + `',
// 	// 		'` + militar.Persona.DatoBasico.ConcatenarApellido() + `',
// 	// 		'` + militar.TiempoSevicio + `',
// 	//
// 	// 		'` + militar.Persona.DatoBasico.EstadoCivil + `'',
// 	// 		` + strconv.Itoa(militar.NumeroHijos()) + `,
// 	//
// 	// 		` + strconv.Itoa(militar.AnoReconocido) + ` ,
// 	// 	  ` + strconv.Itoa(militar.MesReconocido) + `,
// 	// 	 	` + strconv.Itoa(militar.DiaReconocido) + `,
// 	//
// 	//
// 	//
// 	// 		` + strconv.Itoa(militar.Fideicomiso.EstatusNoAscenso) + `,
// 	// 		'` + militar.Fideicomiso.CuentaBancaria + `',
// 	// 		` + strconv.Itoa(militar.Fideicomiso.EstatusProfesion) + `,
// 	// 		'` + militar.Persona.DatoBasico.Sexo + `',
// 	//
// 	// 		'tunel',
// 	//
// 	// 	 	'tunel',
// 	// 	 	'INSERCION POR TUNELES',
// 	// 		'` + militar.Fideicomiso.MotivoParalizacion + `',
// 	//
// 	// 	)';`
//
// 	//echo $sInsertar;
// 	//<<<<<<< HEAD
//

//
