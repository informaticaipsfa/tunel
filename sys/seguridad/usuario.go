// Administración basada en Roles
// El control de acceso basado en roles (RBAC) es una función de seguridad para
// controlar el acceso de usuarios a tareas que normalmente están restringidas al
// superusuario. Mediante la aplicación de atributos de seguridad a procesos y
// usuarios, RBAC puede dividir las capacidades de superusuario entre varios
// administradores. La gestión de derechos de procesos se implementa a través de
// privilegios. La gestión de derechos de usuarios se implementa a través de RBAC.
package seguridad

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"

	"gopkg.in/mgo.v2/bson"
)

const (
	ROOT               string = "0xRO" //Todos los privilegios del sistema
	CONSULTA           string = "0xCO"
	ADMINISTRADOR      string = "0xAD"
	ADMINISTRADORGRUPO string = "0xAA"
	INVITADO           string = "0xIN"
	PRODUCCION         string = "0xPR"
	DESARROLLADOR      string = "0xDE"
	PASANTE            string = "0xPA"
	OPERADOR           string = "0xOP"
	TEST               string = "0xPR"
	HACK               string = "0xHA"
)

type MetodoSeguro struct {
	Consultar  bool `json:"consultar" bson:"consultar" `
	Insertar   bool `json:"insertar"`
	Actualizar bool `json:"actualizar"`
	Eliminar   bool `json:"eliminar"`
	Crud       bool `json:"crud"`
	CrearSQL   bool `json:"crearsql"`
	Todo       bool `json:"todo"`
	Funcion    bool `json:"funcion"`
}

// Privilegio
type Privilegio struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Controlador string        `json:"controlador"`
	Metodo      string        `json:"metodo"`
	Accion      string        `json:"accion"`
}

// Perfil
type Perfil struct {
	Id          string        `json:"id, omitempty"`
	Descripcion string        `json:"descripcion,omitempty"`
	Privilegios []interface{} `json:"privilegios,omitempty"`
}

type Rol struct {
	ID          string `json:"id"`
	Descripcion string `json:"descripcion"`
}

// Usuarios del Sistema
type SUsuario struct {
	ID           int         `json:"id"`
	Nombre       string      `json:"nombre"`
	Correo       string      `json:"correo,omitempty"`
	Clave        string      `json:"clave,omitempty"`
	Token        string      `json:"token,omitempty"`
	Perfil       interface{} `json:"perfil,omitempty"`
	FirmaDigital interface{} `json:"firma,omitempty"`
	Rol          `json:"rol,omitempty"`
}

// La firma permite identificar una maquina y persona autorizada por el sistema
type FirmaDigital struct {
	Id int
	SUsuario
	DireccionMac string
	DireccionIP  string
	Tiempo       time.Time
}

type RespuestaToken struct {
	Token string `json:"token"`
}

func (f *FirmaDigital) Registrar() bool {

	return true
}

func (u *SUsuario) Salvar(us SUsuario) {

}

func (u *SUsuario) Consultar(usuario string, clave string) (v bool) {

	data := []byte(usuario + clave)
	b := md5.Sum(data)
	v = false
	var encry string = hex.EncodeToString(b[:])
	sC := "SELECT oid, nomb, ncom, corr, fech,esta,rol FROM usuario WHERE toke = '" + encry + "';"

	row, e := sys.PostgreSQLSAMAN.Query(sC)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer row.Close()

	for row.Next() {
		var oid int
		var nomb, ncom, corr, fech, esta, rol sql.NullString

		row.Scan(&oid, &nomb, &ncom, &corr, &fech, &esta, &rol)
		u.ID = oid
		u.Nombre = util.ValidarNullString(nomb)
		u.Correo = util.ValidarNullString(corr)
		//u.Rol = util.ValidarNullString(rol)
		v = true
	}

	return
}
