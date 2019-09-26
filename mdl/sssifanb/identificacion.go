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
		fmt.Println("Buscando error ", err.Error())
	}
	wU.Cedula = militar.ID
	wU.Nombre = militar.Persona.DatoBasico.NombrePrimero
	wU.Apellido = militar.Persona.DatoBasico.ApellidoPrimero
	wU.Titular = true
	wU.Clave = util.GenerarHash256([]byte(clv))
	wU.Componente = comp
	wU.Grado = militar.Grado.Abreviatura
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
func (IU *IdentificacionUsuario) BuscarSobreviviente(id string, tipo string) (err error, wU seguridad.WUsuario) {

	s := `SELECT titular, nombres, apellidos FROM familiar WHERE cedula = ` + id

	sq, err := sys.PostgreSQLPENSION.Query(s)
	util.Error(err)

	for sq.Next() {
		var tit, nomb, apel string
		var familiar seguridad.WFamiliar
		sq.Scan(&tit, &nomb, &apel)
		familiar.Cedula = tit
		familiar.Nombre = nomb
		familiar.Apellido = apel
		wU.Familiar = append(wU.Familiar, familiar)
	}

	fmt.Println("Controlando los datos...")

	return
}
