package estadistica

//HistoriaPension Historico de pensiones
func HistoriaPension() string { //
	return `
	SELECT p.codnip,pensionvigente, direcsalcod, fchinicpension,  sueldobasico, primatransporte,primadescenc,primaannoserv,
			primanoascenso,porcprimanoascenso,primaespecial,primaprofesional,porcprimaprof,subtotal, porcprestmonto,pensionasignada,
			bonovac,bonovacaguinaldo
FROM pension_calc pc JOIN personas p ON pc.nropersona=p.nropersona
-- WHERE p.codnip='9150043'
ORDER BY p.codnip,pc.auditfechacambio ASC;`
}
