package usuario

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gesaodin/bdse/sys"
	"github.com/gesaodin/bdse/util"
)

type Usuario struct {
	Cedula   string `json:"cedula,omitempty" bson:"cedula"`
	Usuario  string `json:"usuario,omitempty" bson:"usuario"`
	Clave    string `json:"clave,omitempty" bson:"clave"`
	Nombre   string `json:"nombre,omitempty" bson:"nombre"`
	Apellido string `json:"apellido,omitempty" bson:"apellido"`
	Correo   string `json:"correo,omitempty" bson:"correo"`
	Sucursal string `json:"sucursal,omitempty" bson:"sucursal"`
	Sistema  string `json:"sistema,omitempty" bson:"sistema"`
	IP       string `json:"ip,omitempty" bson:"ip"`
}

var M util.Mensajes

func (u *Usuario) Consultar() (jSon []byte, err error) {
	sSQL := `SELECT codnip, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo FROM personas WHERE codnip='` + u.Cedula + `' LIMIT 1`
	rs, err := sys.PostgreSQLSAMAN.Query(sSQL)
	if err != nil {
		M.Msj = err.Error()
		M.Tipo = 0
		M.Fecha = time.Now()
		jSon, err = json.Marshal(M)
		return
	}
	for rs.Next() {

		var cedula, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo string
		rs.Scan(&cedula, &nombreprimero, &nombresegundo, &apellidoprimero, &apellidosegundo)
		u.Usuario = nombreprimero
		u.Clave = ""
		u.Nombre = nombreprimero + " " + nombresegundo
		u.Apellido = apellidoprimero + " " + apellidosegundo
		u.AplicarReglas()
		fmt.Println(u)
		jSon, err = json.Marshal(u)

	}

	return
}

func (u *Usuario) AplicarReglas() {
	u.Usuario = u.Cedula[len(u.Usuario)-3 : 3]
}

func (u *Usuario) Registrar() (jSon []byte, err error) {
	return
}

func (u *Usuario) Actualizar() (jSon []byte, err error) {
	return
}

func (u *Usuario) Eliminar() (jSon []byte, err error) {
	return
}

func (u *Usuario) Trazar() (jSon []byte, err error) {
	return
}
