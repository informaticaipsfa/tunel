// Administración basada en Roles
// El control de acceso basado en roles (RBAC) es una función de seguridad para
// controlar el acceso de usuarios a tareas que normalmente están restringidas al
// superusuario. Mediante la aplicación de atributos de seguridad a procesos y
// usuarios, RBAC puede dividir las capacidades de superusuario entre varios
// administradores. La gestión de derechos de procesos se implementa a través de
// privilegios. La gestión de derechos de usuarios se implementa a través de RBAC.
package seguridad

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
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
	PROOT              string = "Root"
	PPRESIDENTE        string = "Presidente"
	PADMIN             string = "Administrador"
	PGERENTE           string = "Gerente"
	PJEFE              string = "Jefe"
	GAFILIACION        string = "Afiliacion"
	ANALISTA           string = "Analista"
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
	Metodo      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Accion      string `json:"accion"`
}

// Perfil
type Perfil struct {
	Descripcion string       `json:"descripcion,omitempty"`
	Privilegios []Privilegio `json:"Privilegios,omitempty"`
}

type Rol struct {
	Descripcion string `json:"descripcion" bson:"descipcion"`
}

// Usuarios del Sistema
type Usuario struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Cedula       string        `json:"cedula"`
	Nombre       string        `json:"nombre"`
	Login        string        `json:"login"`
	Correo       string        `json:"correo,omitempty"`
	Clave        string        `json:"clave,omitempty"`
	Sucursal     string        `json:"sucursal,omitempty" bson:"sucursal"`
	Sistema      string        `json:"sistema,omitempty" bson:"sistema"`
	Rol          Rol           `json:"Roles,omitempty"`
	Token        string        `json:"token,omitempty"`
	Perfil       Perfil        `json:"Perfil,omitempty"`
	FirmaDigital FirmaDigital  `json:"FirmaDigital,omitempty"`
}

//FirmaDigital La firma permite identificar una maquina y persona autorizada por el sistema
type FirmaDigital struct {
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

//Salvar Metodo para crear usuarios del sistema
func (usr *Usuario) Salvar() error {

	// var privilegio Privilegio
	// var lst []Privilegio
	//
	// usr.Id = bson.NewObjectId()
	// usr.Nombre = "Root"
	// usr.Login = "root"
	// usr.Sucursal = "Principal"
	// usr.Clave = util.GenerarHash256([]byte(usr.Login))
	//
	// usr.Rol.ID = ROOT
	// usr.Rol.Descripcion = "Super Usuario"
	// usr.Perfil.ID = ROOT
	// usr.Perfil.Descripcion = "Super Usuario"
	//
	// privilegio.Metodo = "afiliacion.salvar"
	// privilegio.Descripcion = "Crear Usuario"
	// privilegio.Accion = "Insert()" // ES6 Metodos
	// lst = append(lst, privilegio)
	//
	// privilegio.Metodo = "afiliacion.modificar"
	// privilegio.Descripcion = "Modificar Usuario"
	// privilegio.Accion = "Update()"
	// lst = append(lst, privilegio)
	// usr.Perfil.Privilegios = lst

	var mongo sys.Mongo

	return mongo.Salvar(usr, "usuario")

}

//Validar Usuarios
func (u *Usuario) Validar(login string, clave string) (err error) {
	u.Nombre = ""
	c := sys.MGOSession.DB("ipsfa_test").C("usuario")
	err = c.Find(bson.M{"login": login, "clave": clave}).Select(bson.M{"clave": false, "firmadigital": false}).One(&u)
	return
}

//Validar Usuarios
func (u *Usuario) CambiarClave(login string, clave string, nueva string) (err error) {
	u.Nombre = ""
	c := sys.MGOSession.DB("ipsfa_test").C("usuario")
	actualizar := make(map[string]interface{})
	actualizar["clave"] = util.GenerarHash256([]byte(nueva))
	err = c.Update(bson.M{"login": login, "clave": clave}, bson.M{"$set": actualizar})
	return
}

//Consultar el sistema de usuarios
func (u *Usuario) Consultar() (v bool) {

	return
}
