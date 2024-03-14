package metodobanco

func SQL_QUERY_PATRIA(firma string) string {
	return `
select
	cdep as cedula, TBL1.nume as numero, neto_monto as monto, nomb as nombre
from
	(
	select
		distinct(cdep),
		nume,
		COUNT(cdep) as cant,
		SUM(neto) as neto_monto
	from
		(
		select
			pg.cedu,
			regexp_replace(pg.nomb,
			'[^a-zA-Y0-9 ]',
			'',
			'g') as nomb,
			pg.nume,
			pg.tipo,
			pg.banc,
			pg.neto,
			case
				when pg.situ != 'FCP' then pg.cedu
				else 
			case
					when pg.caut = ''
					or pg.caut = '0' then pg.cfam
					else pg.caut
				end
			end as cdep,
			pg.cfam,
			pg.caut,
			regexp_replace(pg.naut,
			'[^a-zA-Y0-9 ]',
			'',
			'g') as autor,
			pg.situ
		from
			space.nomina nom
		join space.pagos as pg on
			nom.oid = pg.nomi
		where
			llav = '` + firma + `'
			and nume != ''
			and nume != 'N/A'
			and nume != 'N/'
		order by
			banc,
			pg.cedu desc
) as DR
	group by
		cdep,
		nume
	order by
		cant desc

) as TBL1
join space.cuentas_temp as rcuentas on
	TBL1.nume = rcuentas.nume
where rcuentas.nume !='0102'
order by
	rcuentas.nume

	`
}
