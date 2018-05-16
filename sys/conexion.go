//configuraciones del sistema
package sys

import (
	"database/sql"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
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

//ConexionSAMAN Funcion de Conexion a Postgres
func ConexionSAMANWEB(mapa map[string]CadenaDeConexion) {
	c := mapa["samanweb"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PsqlWEB, _ = sql.Open("postgres", cadena)
	if PsqlWEB.Ping() != nil {
		fmt.Println("[SamanWEB:   Error...] ", PsqlWEB.Ping())
	} else {
		fmt.Println("[SamanWEB: ", c.Host, "  OK...]")
	}

}

//ConexionEMPLEADO Funcion de Conexion a Postgres
func ConexionEMPLEADO(mapa map[string]CadenaDeConexion) {
	c := mapa["empleado"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLEMPLEADOSIGESP, _ = sql.Open("postgres", cadena)
	if PostgreSQLEMPLEADOSIGESP.Ping() != nil {
		fmt.Println("[Empleado:   Error...] ", PostgreSQLEMPLEADOSIGESP.Ping())
	} else {
		fmt.Println("[Empleado: ", c.Host, "  OK...]")
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

//ConexionMYSQL
func ConexionMYSQL(mapa map[string]CadenaDeConexion) {
	c := mapa["mysql"]
	cadena := c.Usuario + ":" + c.Clave + "@tcp(" + c.Host + ":3306)/sssifanb"
	MysqlFullText, _ = sql.Open("mysql", cadena)
	if MysqlFullText.Ping() != nil {
		fmt.Println("[mysql FULLTEXT: Error...] ", MysqlFullText.Ping())
	} else {
		fmt.Println("[mysql FULLTEXT: OK...]")
	}
	return
}
