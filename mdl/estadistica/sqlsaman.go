package estadistica

func historiaPension() string { //
	return `
    SELECT pensionvigente, direcsalcod, fchinicpension,  sueldobasico, primatransporte,primadescenc,primaannoserv,
    primanoascenso,porcprimanoascenso,primaespecial,primaprofesional,porcprimaprof,subtotal, porcprestmonto,pensionasignada,
    bonovac,bonovacaguinaldo
    FROM pension_calc where nropersona=9609 order by auditfechacambio ASC;`
}
