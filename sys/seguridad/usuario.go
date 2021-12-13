// Administración basada en Roles
// El control de acceso basado en roles (RBAC) es una función de seguridad para
// controlar el acceso de usuarios a tareas que normalmente están restringidas al
// superusuario. Mediante la aplicación de atributos de seguridad a procesos y
// usuarios, RBAC puede dividir las capacidades de superusuario entre varios
// administradores. La gestión de derechos de procesos se implementa a través de
// privilegios. La gestión de derechos de usuarios se implementa a través de RBAC.
package seguridad

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
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
	Menu        []Menu       `json:"Menu,omitempty"`
}

type Menu struct {
	Url    string `json:"url,omitempty"`
	Js     string `json:"js,omitempty"`
	Icono  string `json:"icono,omitempty"`
	Nombre string `json:"nombre,omitempty"`
	Accion string `json:"accion,omitempty"`
	Color  string `json:"color,omitempty"`
}

type Rol struct {
	Descripcion string `json:"descripcion" bson:"descipcion"`
}

// Usuarios del Sistema
type Usuario struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	Cedula        string        `json:"cedula" bson:"cedula"`
	Nombre        string        `json:"nombre" bson:"nombre"`
	Login         string        `json:"usuario" bson:"login"`
	Correo        string        `json:"correo" bson:"correo"`
	FechaCreacion time.Time     `json:"fechacreacion,omitempty" bson:"fechacreacion"`
	Estatus       int           `json:"estatus" bson:"estatus"`
	Clave         string        `json:"clave,omitempty" bson:"clave"`
	Situacion     string        `json:"situacion,omitempty" bson:"situacion"` //PM - PC
	Sucursal      string        `json:"sucursal,omitempty" bson:"sucursal" bson:"sucursal"`
	Departamento  string        `json:"departamento,omitempty" bson:"departamento"`
	Sistema       string        `json:"sistema,omitempty" bson:"sistema"`
	Rol           Rol           `json:"Roles,omitempty" bson:"roles"`
	Token         string        `json:"token,omitempty" bson:"token"`
	Perfil        Perfil        `json:"Perfil,omitempty" bson:"perfil"`
	FirmaDigital  FirmaDigital  `json:"FirmaDigital,omitempty" bson:"firmadigital"`
	Direccion     string        `json:"direccion,omitempty" bson:"direccion"`
	Telefono      string        `json:"telefono,omitempty" bson:"telefono"`
	Cargo         string        `json:"cargo,omitempty" bson:"cargo"`
	Modulo        []Modulo      `json:"modulo,omitempty" bson:"modulo"`
}

//FirmaDigital La firma permite identificar una maquina y persona autorizada por el sistema
type FirmaDigital struct {
	DireccionMac string    `json:"direccionmac,omitempty" bson:"direccionmac"`
	DireccionIP  string    `json:"direccionip,omitempty" bson:"direccionip"`
	Tiempo       time.Time `json:"tiempo,omitempty" bson:"tiempo"`
}

type RespuestaToken struct {
	Token string `json:"token"`
}

type Modulo struct {
	Id         string `json:"id"`
	Nombre     string `json:"nombre"`
	URL        string `json:"url"`
	Comentario string `json:"comentario"`
	Version    string `json:"version"`
	Autor      string `json:"autor"`
}

func (f *FirmaDigital) Registrar() bool {

	return true
}

//Salvar Metodo para crear usuarios del sistema
func (usr *Usuario) Salvar() error {
	usr.ID = bson.NewObjectId()
	usr.Clave = util.GenerarHash256([]byte(usr.Clave))
	usr.FechaCreacion = time.Now()
	fmt.Println("Creando Usuario")

	c := sys.MGOSession.DB(sys.CBASE).C(sys.CUSUARIO)
	return c.Insert(usr)

}

//Validar Usuarios
func (u *Usuario) Validar(login string, clave string) (err error) {
	u.Nombre = ""
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CUSUARIO)
	err = c.Find(bson.M{"login": login, "clave": clave}).Select(bson.M{"clave": false}).One(&u)

	return
}

func CrearClaveTodos() {
	var usuario []Usuario
	// var lst []interface{}
	c := sys.MGOSession.DB(sys.CBASE).C("usuario")
	err := c.Find(nil).All(&usuario)
	if err != nil {
		return
	}
	// usuario = lst
	for _, v := range usuario {
		clave := util.GenerarHash256([]byte(v.Cedula))
		fmt.Println(v.Cedula, " -> ", v.Clave, " -> ", clave)
		err = c.Update(bson.M{"cedula": v.Cedula}, bson.M{"$set": bson.M{"clave": clave}})
		if err != nil {
			fmt.Println("Err.", err.Error())
			return
		}
	}
	return
}

//CambiarClave Usuarios
func (u *Usuario) CambiarClave(login string, clave string, nueva string, coleccion string) (err error) {
	u.Nombre = ""
	c := sys.MGOSession.DB(sys.CBASE).C(coleccion)
	actualizar := make(map[string]interface{})
	actualizar["clave"] = util.GenerarHash256([]byte(nueva))
	antigua := util.GenerarHash256([]byte(clave))
	err = c.Update(bson.M{"login": login, "clave": antigua}, bson.M{"$set": actualizar})
	return
}

//Consultar el sistema de usuarios
func (u *Usuario) Consultar(cedula string, coleccion string) (j []byte, err error) {
	u.Nombre = ""
	campo := "login"
	var itObjeto interface{}

	c := sys.MGOSession.DB(sys.CBASE).C(coleccion)
	if coleccion == "wusuario" {
		campo = "cedula"
	}
	err = c.Find(bson.M{campo: cedula}).Select(bson.M{"clave": false}).One(&itObjeto)
	j, _ = json.Marshal(itObjeto)
	return
}

//RestablecerClaves  de usuarios Internos y WEB
func (u *Usuario) RestablecerClaves(cedula string, correo string, clave string, coleccion string) (err error) {

	c := sys.MGOSession.DB(sys.CBASE).C(coleccion)
	actualizar := make(map[string]interface{})
	actualizar["clave"] = util.GenerarHash256([]byte(clave))
	//	err = c.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": actualizar})
	err = c.Update(bson.M{"cedula": cedula, "correo": correo}, bson.M{"$set": actualizar})
	return
}

//Listar el sistema de usuarios
func (u *Usuario) Listar() (j []byte, err error) {
	var lstUsuario []Usuario
	c := sys.MGOSession.DB(sys.CBASE).C("usuario")
	err = c.Find(bson.M{}).Select(bson.M{"clave": false}).All(&lstUsuario)
	j, _ = json.Marshal(lstUsuario)
	return
}

func (u *Usuario) Generico() error {
	var privilegio Privilegio
	var lst []Privilegio
	var usr Usuario
	usr.ID = bson.NewObjectId()
	usr.Nombre = "Informatica - Consulta"
	usr.Login = "usuario"
	usr.Sucursal = "Principal"
	usr.Clave = util.GenerarHash256([]byte("123"))

	// usr.Rol.ID = ROOT
	usr.Rol.Descripcion = "Super Usuario"
	// usr.Perfil.ID = ROOT
	usr.Perfil.Descripcion = "Super Usuario"

	privilegio.Metodo = "afiliacion.salvar"
	privilegio.Descripcion = "Crear Usuario"
	privilegio.Accion = "Insert()" // ES6 Metodos
	lst = append(lst, privilegio)

	privilegio.Metodo = "afiliacion.modificar"
	privilegio.Descripcion = "Modificar Usuario"
	privilegio.Accion = "Update()"
	lst = append(lst, privilegio)
	usr.Perfil.Privilegios = lst

	var mongo sys.Mongo

	return mongo.Salvar(usr, "usuario")
}
