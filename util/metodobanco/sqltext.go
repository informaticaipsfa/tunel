package metodobanco

func SQL_QUERY_PATRIA(firma string) string {
	return `
	select
		cc.cedula,
		numero, 
		monto,
		cc.nombre,
		NULLIF('V', sc.tipo_cedula ) as tipo_cedula
from
	(
select REGEXP_REPLACE(cedu, E'[\\s\\t]+', '', 'g') as cedula, numero, SUM(monto) as monto, nombre from (
select
	rcuentas.cedu, cdep as cedula, TBL1.nume as numero, neto_monto as monto, nomb as nombre
from
	(select
		distinct(cdep),
		nume,
		COUNT(cdep) as cant,
		SUM(neto) as neto_monto
	from
		(
		select
			pg.nomi, pg.cedu,
			regexp_replace(pg.nomb,
			'[^a-zA-Y0-9 ]',
			'','g') as nomb, pg.nume,pg.tipo,pg.banc,pg.neto,
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
			pg.cedu desc
) as DR
	group by
		cdep,
		nume
	order by
		cant desc
) as TBL1
join ( select distinct on (nume) * from 
(select * from space.cuentas_temp ct 
order by cedu ASC) as tbl ) as rcuentas on
	TBL1.nume = rcuentas.nume
where rcuentas.nume !='0102'

order by
	rcuentas.nume) as mayor
	
group by cedu, numero, nombre
order by cedu
 ) as cc
left join 
	( select 
	distinct (cedula), tipo_cedula, sigla1, nombre1, nombre2, apellido1, apellido2, sexo 
from saime.saime_cedula ) AS sc on sc.cedula = cc.cedula `
}
