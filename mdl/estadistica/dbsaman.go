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
		SELECT	codnip AS cedula, prs.nropersona, fecha_ingreso
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
			anio_reconocido,mes_reconocido,dia_reconocido,f_retiro,n_hijos,numero_cuenta
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
							bnf.numero_cuenta
						FROM pers_dat_militares AS pm
						LEFT JOIN pension ON pension.nropersona=pm.nropersona
						JOIN personas AS p ON pm.nropersona=p.nropersona
						LEFT JOIN beneficiario AS bnf ON p.codnip=bnf.cedula
						JOIN ipsfa_componentes AS icom ON pm.componentecod=icom.componentecod
						JOIN ipsfa_grados AS igra ON pm.gradocod=igra.gradocod AND pm.componentecod=igra.componentecod
					) AS B ON B.nropersona = TBL.np  -- limit 1000 -- WHERE cedula_saman='16872776' --  WHERE B.perssituaccod = 'ACT' --`
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
			ORDER BY AR.cedula_saman,p.nropersona, gradofchrecipsfa`
	//WHERE p.nropersona IN (1393199,79227)

}

func obtenerCuentaBancaria() string {
	return `
		SELECT
			AR.cedula_saman, nrocuenta, tipcuentacod, instfinancod
		FROM
		analisis.reducciones as AR
		JOIN (SELECT DISTINCT nropersona, nrocuenta, tipcuentacod, instfinancod FROM pers_cta_bancarias where usocuentacod='PRI') AS cta
		ON AR.np=cta.nropersona
		--WHERE AR.cedula_saman='16872776'
		`
}

func obtenerComponenteGrado() string{
	return `
		SELECT c.componentecod, componentenombre, componentesiglas, gradocod,gradocodrangoid,gradonombrecorto,
		gradonombrelargo
		FROM ipsfa_grados AS g JOIN ipsfa_componentes AS c ON g.componentecod=c.componentecod
		ORDER BY c.componentepriorpt,g.gradocodrangoid
	`
}

func obtenerEstados() string { //MySQL
	return `SELECT estado, iso_31662, ciudad,capital FROM estados JOIN ciudades ON estados.id_estado=ciudades.id_estado`
}
func obtenerMunicipios() string { //MySQL
	return `SELECT estado, municipio FROM estados JOIN municipios ON estados.id_estado=municipios.id_estado`
}
func obtenerParroquia() string { //MySQL
	return `SELECT * FROM estados JOIN municipios ON estados.id_estado=municipios.id_estado`
}
