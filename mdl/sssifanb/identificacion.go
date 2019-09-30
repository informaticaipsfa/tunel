package sssifanb

import (
	"fmt"
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/sys/seguridad"
	"github.com/informaticaipsfa/tunel/util"
	"gopkg.in/mgo.v2/bson"
)

//IdentificacionUsuario Analizar para determinar que usario es
type IdentificacionUsuario struct {
	ID     string
	Nombre string
}

//BuscarTitular e identificar un cédula, Como lo es un titulas o sobreviviente
func (IU *IdentificacionUsuario) BuscarTitular(id string, tipo string, clv string, comp string, corr string) (err error, wU seguridad.WUsuario) {
	var militar Militar
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	err = c.Find(bson.M{"id": id}).One(&militar)

	if err != nil {
		fmt.Println("Buscando error de login de usuarios... ", err.Error())
		return
	}
	fmt.Println(comp, " ", militar.Componente.Abreviatura)
	if comp != militar.Componente.Abreviatura {
		return
	}
	wU.Cedula = militar.ID
	wU.Nombre = militar.Persona.DatoBasico.NombrePrimero
	wU.Apellido = militar.Persona.DatoBasico.ApellidoPrimero
	wU.Titular = true
	wU.Clave = util.GenerarHash256([]byte(clv))
	wU.Componente = comp
	wU.Grado = militar.Grado.Abreviatura
	wU.Situacion = militar.Situacion

	wU.FechaCreacion = time.Now()
	wU.Correo = corr

	var usuario seguridad.WUsuario
	usuario = wU

	cu := sys.MGOSession.DB(sys.CBASE).C(sys.WUSUARIO)
	err = cu.Insert(usuario)

	if err != nil {
		fmt.Println("insertando error ", err.Error())
	}

	// reduc := make(map[string]interface{})
	// cred := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	// reduc["cedula"] = militar.ID
	// reduc["persona.correo.principal"] = corr
	// err = cred.Update(bson.M{"cedula": militar.ID}, bson.M{"$set": reduc})
	// if err != nil {
	// 	fmt.Println("Err", err.Error())
	// }

	return
}

//BuscarSobreviviente e identificar un cédula, Como lo es un titulas o sobreviviente
func (IU *IdentificacionUsuario) BuscarSobreviviente(id string, tipo string, corr string, clv string) (err error, wU seguridad.WUsuario) {

	s := `SELECT fm.cedula, fm.nombres, fm.apellidos, fm.titular,
					bf.nombres, bf.apellidos, cp.descripcion
				FROM familiar fm
				JOIN beneficiario bf ON bf.cedula=fm.titular
				JOIN componente cp ON bf.componente_id=cp.id
				WHERE fm.cedula='` + id + `'

				`
	sq, err := sys.PostgreSQLPENSION.Query(s)
	util.Error(err)

	for sq.Next() {
		var cedu, tit, nomb, apel, nombbf, apelbf, descri string
		var causante seguridad.WCausante
		sq.Scan(&cedu, &nomb, &apel, &tit, &nombbf, &apelbf, &descri)
		causante.Cedula = tit
		causante.Nombre = nombbf
		causante.Apellido = apelbf
		causante.Componente = descri

		wU.Cedula = cedu
		wU.Nombre = nomb
		wU.Apellido = apel
		wU.Sobreviviente = true
		wU.Clave = util.GenerarHash256([]byte(clv))
		wU.Correo = corr
		wU.Componente = descri
		wU.Causante = append(wU.Causante, causante)
	}
	cu := sys.MGOSession.DB(sys.CBASE).C(sys.WUSUARIO)
	err = cu.Insert(wU)
	if err != nil {
		fmt.Println("insertando error ", err.Error())
	}
	return
}
