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
	Root                 string = "0xRO" //Todos los privilegios del sistema
	Consulta             string = "0xCO"
	Administrador        string = "0xAD"
	AdministradorDeGrupo string = "0xAA"
	Invitado             string = "0xIN"
	Produccion           string = "0xPR"
	Desarrollador        string = "0xDE"
	Pasante              string = "0xPA"
	Operador             string = "0xOP"
	Prueba               string = "0xPR"
	Hack                 string = "0xHA"
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
	Id          string `json:"id"`
	Descripcion string `json:"descripcion"`
}

// Usuarios del Sistema
type Usuario struct {
	Id           int         `json:"id"`
	Nombre       string      `json:"nombre"`
	Correo       string      `json:"correo,omitempty"`
	Clave        string      `json:"clave,omitempty"`
	Token        string      `json:"token,omitempty"`
	Perfil       interface{} `json:"perfil,omitempty"`
	Rol          string      `json:"rol,omitempty"`
	FirmaDigital interface{} `json:"firma,omitempty"`
	//Rol          interface{} `json:"rol,omitempty"`
}

// La firma permite identificar una maquina y persona autorizada por el sistema
type FirmaDigital struct {
	Id           int
	Usuario      Usuario
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

func (u *Usuario) Salvar(us Usuario) {

}

func (u *Usuario) Consultar(usuario string, clave string) (v bool) {

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
		u.Id = oid
		u.Nombre = util.ValidarNullString(nomb)
		u.Correo = util.ValidarNullString(corr)
		u.Rol = util.ValidarNullString(rol)
		v = true
	}

	return
}
