package seguridad

import (
	"time"

	"github.com/informaticaipsfa/tunel/sys"
	"gopkg.in/mgo.v2/bson"
)

//WFamiliar Control de familiares para asignaciones
type WCausante struct {
	Cedula     string `json:"cedula" bson:"cedula"`
	Nombre     string `json:"nombre" bson:"nombre"`
	Apellido   string `json:"apellido" bson:"apellido"`
	Componente string `json:"componente" bson:"componente"`
	Grado      string `json:"grado" bson:"grado"`
}

//WUsuario del Sistema
type WUsuario struct {
	ID            string       `json:"id,omitempty" bson:"id"`
	Cedula        string       `json:"cedula" bson:"cedula"`
	Nombre        string       `json:"nombre" bson:"nombre"`
	Apellido      string       `json:"apellido" bson:"apellido"`
	Causante      []WCausante  `json:"causante,omitempty" bson:"causante"`
	Login         string       `json:"usuario" bson:"login"`
	Clave         string       `json:"clave,omitempty" bson:"clave"`
	Correo        string       `json:"correo" bson:"correo"`
	FechaCreacion time.Time    `json:"fechacreacion,omitempty" bson:"fechacreacion"`
	Estatus       int          `json:"estatus" bson:"estatus"`
	Situacion     string       `json:"situacion,omitempty" bson:"situacion"` //PM - PC
	Componente    string       `json:"componente" bson:"componente"`
	Grado         string       `json:"grado" bson:"grado"`
	Rol           Rol          `json:"Roles,omitempty" bson:"roles"`
	Token         string       `json:"token,omitempty" bson:"token"`
	Perfil        Perfil       `json:"Perfil,omitempty" bson:"perfil"`
	FirmaDigital  FirmaDigital `json:"FirmaDigital,omitempty" bson:"firmadigital"`
	Telefono      string       `json:"telefono,omitempty" bson:"telefono"`
	Titular       bool         `json:"titular,omitempty" bson:"titular"`
	Sobreviviente bool         `json:"sobreviviente,omitempty" bson:"sobreviviente"`
	Empleado      bool         `json:"empleado,omitempty" bson:"empleado"`
}

//WVwalidar Validacion de Usuarios
func (u *WUsuario) Validar(login string, clave string) (err error) {
	u.Nombre = ""
	c := sys.MGOSession.DB(sys.CBASE).C(sys.WUSUARIO)
	err = c.Find(bson.M{"cedula": login, "clave": clave}).Select(bson.M{"clave": false}).One(&u)

	return
}

//WVwalidar Validacion de Usuarios
func (u *WUsuario) Existe(login string) (err error) {
	u.Nombre = ""
	c := sys.MGOSession.DB(sys.CBASE).C(sys.WUSUARIO)
	err = c.Find(bson.M{"cedula": login}).Select(bson.M{"clave": false}).One(&u)

	return
}
