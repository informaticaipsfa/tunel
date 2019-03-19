package estadistica

//HistoriaPension Historico de pensiones
func HistoriaPension() string { //
	return `
	SELECT p.codnip,pensionvigente, direcsalcod, fchinicpension,  sueldobasico, primatransporte,primadescenc,primaannoserv,
				primanoascenso,porcprimanoascenso,primaespecial,primaprofesional,porcprimaprof,subtotal, porcprestmonto,pensionasignada,
				bonovac,bonovacaguinaldo, pm.fchegreso
	FROM pension_calc pc
	JOIN personas p ON pc.nropersona=p.nropersona
	JOIN pers_dat_militares pm ON p.nropersona=pm.nropersona
	WHERE p.codnip='9150043'
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
				prs.codnip,
				rbs.reembsolicnro,
				rbs.nropersonaafilmil,
				rbs.nropersonapago,
				rbs.reembtipocod,
				rbs.reembfchsolicitud,
				rbs.reembfchaprobacion,
				ordenpagomonto,
				reembconcmontoapr,
				inf.instfinannombre,
				cuenta,
				tpc.tipcuentanombre,
				rdc.reembconcnombre,
				canalliquidnombre,
				componentecod,
				gradocod,
				perscategcod,
				perssituaccod,
				prs.nombrecompleto,
				prc.codnip,
				prc.nombrecompleto,
				pdb.persrelstipcod,
				prdet.nombrecompleto AS benef,
				prdet.codnip AS cedbeneficiado
			FROM personas prs
			INNER JOIN ci_reembolso_solic rbs ON prs.nropersona=rbs.nropersonaafilmil
			INNER JOIN tipos_cuenta tpc ON tpc.tipcuentacod=rbs.tipcuentacod
			INNER JOIN canal_liquidacion cnl ON rbs.canalliquidcod=cnl.canalliquidcod
			INNER JOIN ci_reembolso_det rdt ON rbs.reembsolicnro=rdt.reembsolicnro
			INNER JOIN ci_reembolso_concep rdc on rdt.reembconccod=rdc.reembconccod
			INNER JOIN ci_reembolso_tipo ON (rbs.reembtipocod=ci_reembolso_tipo.reembtipocod)
			LEFT JOIN inst_financieras inf ON inf.instfinancod=rbs.instfinancod
			INNER JOIN personas prc ON rbs.nropersonapago=prc.nropersona
			INNER JOIN personas prdet ON rdt.nropersafilfam=prdet.nropersona
			LEFT JOIN pers_relaciones pdb ON prdet.nropersona = pdb.nropersona
			WHERE rbs.reembtipocod = 'DAF'
			AND prs.codnip IN ( '10010499', '16210806')
			ORDER BY prs.codnip, rbs.reembsolicnro, rbs.reembfchsolicitud`
}
