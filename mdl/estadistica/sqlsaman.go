package estadistica

//HistoriaPension Historico de pensiones
func HistoriaPension() string { //
	return `
	SELECT p.codnip,pensionvigente, direcsalcod, fchinicpension,  sueldobasico, primatransporte,primadescenc,primaannoserv,
				primanoascenso,porcprimanoascenso,primaespecial,primaprofesional,porcprimaprof,subtotal, porcprestmonto,pensionasignada,
				bonovac,bonovacaguinaldo
	FROM pension_calc pc JOIN personas p ON pc.nropersona=p.nropersona
	-- WHERE p.codnip='9150043'
	ORDER BY p.codnip,pc.auditfechacambio ASC --limit 10;`
}

func HistorialUsuario() string {
	return `
	select usuariocodigo,codnip, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo  from seg_usuarios
		join personas on seg_usuarios.nropersona= personas.nropersona
		where usuariocodigo LIKE '%afi%' AND estatususrcod ='ACT'
	`
}

func HistoriaReembolsos() string {
	return `
			SELECT
				codnip,
				rbs.reembsolicnro AS oid,
				rbs.nropersonaafilmil,
				rbs.nropersonapago,rbs.reembtipocod,
				rbs.reembfchsolicitud AS fechaSolicitud,
				rbs.reembfchaprobacion AS fechaAprobacion,
				ordenpagomonto AS montoSolicitado,
				reembconcmontoapr AS montoAprobado,
				inf.instfinannombre,
				cuenta,
				rdc.reembconcnombre,canalliquidnombre,
				componentecod,
				gradocod,
				perscategcod,
				perssituaccod
			FROM personas
				INNER JOIN ci_reembolso_solic rbs ON personas.nropersona=rbs.nropersonaafilmil
				INNER JOIN canal_liquidacion cnl ON rbs.canalliquidcod=cnl.canalliquidcod
				INNER JOIN ci_reembolso_det rdt ON rbs.reembsolicnro=rdt.reembsolicnro
				INNER JOIN ci_reembolso_concep rdc on rdt.reembconccod=rdc.reembconccod
				INNER JOIN ci_reembolso_tipo ON (rbs.reembtipocod=ci_reembolso_tipo.reembtipocod)
				LEFT JOIN inst_financieras inf ON inf.instfinancod=rbs.instfinancod
			WHERE rbs.reembtipocod = 'DAF'
			ORDER BY codnip, rbs.reembfchsolicitud`
}
