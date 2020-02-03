package empleado

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/sys"
)

func BuscarEmpleadosSigesp() string {
	return `select per.codper,
    per.nomper,
    per.apeper,
    per.fecnacper,
    per.edocivper,
    per.sexper,
    per.rifper,
    per.fecingper,
    per.anoservpreper as anosreconocidos,
    (case     when asi.denasicar is null and nom.codnom = '0006' then 'JUBILADO'
        when asi.denasicar is null and nom.codnom = '0007' then 'PENSIONADOS'
        when asi.denasicar = 'Sin Asignación de Cargo' then car.descar
        else asi.denasicar end),
    ger.denger,
    (case     when ban.nomban is null and substr(nom.codcueban, 1,4) = '0102' then 'VENEZUELA (Solo Uso Nomina)'
        when ban.nomban is null and substr(nom.codcueban, 1,4) = '0177' then 'BANFANB (NOMINA CIVIL)'
        else ban.nomban end) as nombrebanco,
    nom.codcueban,
    (case     when nom.codnom = '0001' then 'Empleados Fijos'
        when nom.codnom = '0002' then 'Empleados Contratados'
        when nom.codnom = '0003' then 'Obreros Fijos'
        when nom.codnom = '0004' then 'Obreros Contratos'
        when nom.codnom = '0006' then 'Jubilados'
        else 'Pensionados' end) tiponomina
        from    sno_personal as per
        inner join sno_personalnomina as nom
            on per.codper = nom.codper
            and nom.codnom in ('0001','0002','0003','0006','0007')
            and nom.staper in ('1','2')
    left join srh_gerencia as ger
        on nom.codemp = ger.codemp
        and per.codger = ger.codger
    left join scb_banco as ban
        on ban.codemp = ger.codemp
        and ban.codban = nom.codban
    left join sno_asignacioncargo as asi
        on ban.codemp = asi.codemp
        and nom.codnom = asi.codnom
        and nom.codasicar = asi.codasicar
    left join sno_cargo as car
        on car.codemp = nom.codemp
        and car.codnom = nom.codnom
        and car.codcar = nom.codcar
		--where     per.codper = '015132444'
		order by nom.codnom, 1`
}

type Empleado struct {
	Persona    sssifanb.Persona `json:"Persona" bson:"persona"`
	Cargo      string
	Gerencia   string
	TipoNomina string
}

func (e *Empleado) Insertar() bool {
	sq, err := sys.PostgreSQLEMPLEADOSIGESP.Query(BuscarEmpleadosSigesp())
	if err != nil {

	}
	for sq.Next() {
		var Empl Empleado
		var DF sssifanb.DatoFinanciero
		var ced, nom, ape, fch, edo, sex, rif, fing, ano, car, ger, ban, cue, tip string
		sq.Scan(&ced, &nom, &ape, &fch, &edo, &sex, &rif, &fing, &ano,
			&car, &ger, &ban, &cue, &tip)
		//fmt.Println(ced, nom, ape, fch, edo, sex, rif, fing, ano, car, ger, ban, cue, tip)
		Empl.Persona.DatoBasico.Cedula = ced
		DF.Cuenta = cue
		DF.Institucion = ban

		Empl.Persona.DatoFinanciero = append(Empl.Persona.DatoFinanciero, DF)
		Empl.TipoNomina = tip

		coleccion := sys.MGOSession.DB("sssifanb").C("empleado")
		err = coleccion.Insert(Empl)
		if err != nil {
			fmt.Println("Error: cedula: ", ced)
			// return
		}
	}
	return true
}

func (E *Empleado) Actualizar() {
	sq, err := sys.PostgreSQLEMPLEADOSIGESP.Query(BuscarEmpleadosSigesp())
	if err != nil {

	}
	coleccion := sys.MGOSession.DB("sssifanb").C("empleado")
	for sq.Next() {
		cargo := make(map[string]interface{})

		var ced, nom, ape, fch, edo, sex, rif, fing, ano, car, ger, ban, cue, tip string
		sq.Scan(&ced, &nom, &ape, &fch, &edo, &sex, &rif, &fing, &ano,
			&car, &ger, &ban, &cue, &tip)
		cargo["cargo"] = car
		fmt.Println("Cedula", ced)
		err = coleccion.Update(bson.M{"persona.datobasico.cedula": ced}, bson.M{"$set": cargo})
		if err != nil {
			fmt.Println("Error: cedula: ", ced)
			// return
		}

	}
}
