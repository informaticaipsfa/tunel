//configuraciones del sistema
package sys

import (
	"database/sql"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	"github.com/informaticaipsfa/tunel/util"
	_ "github.com/lib/pq"
)

//MongoDBConexion Conexion a Mongo DB
func MongoDBConexion(mapa map[string]CadenaDeConexion) {
	c := mapa["mongodb"]
	MGOSession, Error = mgo.Dial(c.Host + ":27017")
	fmt.Println("Cargando Conexión Con MongoDB...")
	util.Error(Error)
}

//ConexionSAMAN Funcion de Conexion a Postgres
func ConexionSAMAN(mapa map[string]CadenaDeConexion) {
	c := mapa["saman"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLSAMAN, _ = sql.Open("postgres", cadena)
	if PostgreSQLSAMAN.Ping() != nil {
		fmt.Println("[Saman:   Error...] ", PostgreSQLSAMAN.Ping())
	} else {
		fmt.Println("[Saman: ", c.Host, "  OK...]")
	}
}

//ConexionPACE
func ConexionPACE(mapa map[string]CadenaDeConexion) {
	c := mapa["pace"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLPACE, _ = sql.Open("postgres", cadena)
	if PostgreSQLPACE.Ping() != nil {
		fmt.Println("[Pace: ", c.Host, " Error...] ", PostgreSQLPACE.Ping())
	} else {
		fmt.Println("[Pace: ", c.Host, " OK...]")
	}
}

//ConexionTARJETA
func ConexionPENSION(mapa map[string]CadenaDeConexion) {
	c := mapa["pension"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLPENSION, _ = sql.Open("postgres", cadena)
	if PostgreSQLPENSION.Ping() != nil {
		fmt.Println("[Pensiones: Error...] ", PostgreSQLPENSION.Ping())
	} else {
		fmt.Println("[Pensiones: OK...]")
	}
	return
}

//ConexionTARJETA
func ConexionPENSIONSIGESP(mapa map[string]CadenaDeConexion) {
	c := mapa["pensiones"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLPENSIONSIGESP, _ = sql.Open("postgres", cadena)
	if PostgreSQLPENSIONSIGESP.Ping() != nil {
		fmt.Println("[Pensiones SIGESP: Error...] ", PostgreSQLPENSIONSIGESP.Ping())
	} else {
		fmt.Println("[Pensiones SIGESP: OK...]")
	}
	return
}
