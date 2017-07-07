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
			gradocod,gradocodrangoid,gradonombrecorto,gradonombrelargo
	 FROM (
						SELECT DISTINCT ON (A.cedula_saman) cedula_saman, cedula_pace, np FROM (
						SELECT * FROM analisis.reducciones WHERE militar > 0 ORDER BY militar DESC ) AS A
					) AS TBL JOIN (
						SELECT p.codnip,p.nropersona,
							p.tipnip,p.nombreprimero,nombresegundo,apellidoprimero,apellidosegundo,fechanacimiento,
							p.sexocod,p.edocivilcod,
							perscategcod, perssituaccod,persclasecod,
							fchingcomponente,fchultimoascenso,fchegreso,
							annoreconocido,mesreconocido,diareconocido,
							icom.componentecod,icom.componentenombre,icom.componentesiglas,
							igra.gradocod,igra.gradocodrangoid,igra.gradonombrecorto,igra.gradonombrelargo
						FROM pers_dat_militares AS pm
						JOIN personas AS p ON pm.nropersona=p.nropersona
						JOIN ipsfa_componentes AS icom ON pm.componentecod=icom.componentecod
						JOIN ipsfa_grados AS igra ON pm.gradocod=igra.gradocod AND pm.componentecod=igra.componentecod
					) AS B ON B.nropersona = TBL.np -- WHERE B.perssituaccod = 'ACT' --limit 1000`
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
	return `SELECT
						ipg.componentecod, ipg.gradocod, ipg.perscategcod,
						ipg.perssituaccod, ipg.gradofchobten, ipg.gradoresuelto,
						ipg.persclasecod,
						ipg.gradonroenresuelto,ipg.gradofchrecipsfa,
						ipg.auditfechacambio,ipg.audithoracambio,
						ipg.auditfechacreacion,ipg.audithoracreacion,ipg.razonhistcod
					FROM
						analisis.reducciones as AR
					JOIN personas p ON AR.np=p.nropersona
					JOIN ipsfa_grado_x_pers ipg ON ipg.nropersona = AR.np
					ORDER BY AR.cedula_saman;
`
	//WHERE p.nropersona IN (1393199,79227)

}
