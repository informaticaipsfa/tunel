// Administración basada en Roles
// El control de acceso basado en roles (RBAC) es una función de seguridad para
// controlar el acceso de usuarios a tareas que normalmente están restringidas al
// superusuario. Mediante la aplicación de atributos de seguridad a procesos y
// usuarios, RBAC puede dividir las capacidades de superusuario entre varios
// administradores. La gestión de derechos de procesos se implementa a través de
// privilegios. La gestión de derechos de usuarios se implementa a través de RBAC.
package seguridad

import "time"

const (
	ROOT               string = "0xRO"
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
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Controlador string `json:"controlador"`
	Metodo      string `json:"metodo"`
	Accion      string `json:"accion"`
}

// Perfil
type Perfil struct {
	ID          string       `json:"id, omitempty"`
	Descripcion string       `json:"descripcion,omitempty"`
	Privilegios []Privilegio `json:"Privilegios,omitempty"`
}

type Rol struct {
	ID          string `json:"id"`
	Descripcion string `json:"descripcion"`
}

// Usuarios del Sistema
type Usuario struct {
	ID           int          `json:"id"`
	Nombre       string       `json:"nombre"`
	Correo       string       `json:"correo,omitempty"`
	Clave        string       `json:"clave,omitempty"`
	Sucursal     string       `json:"sucursal,omitempty" bson:"sucursal"`
	Sistema      string       `json:"sistema,omitempty" bson:"sistema"`
	Token        string       `json:"token,omitempty"`
	Perfil       Perfil       `json:"Perfil,omitempty"`
	FirmaDigital FirmaDigital `json:"FirmaDigital,omitempty"`
	IP           string       `json:"ip,omitempty" bson:"ip"`
	Rol          `json:"Rol,omitempty"`
}

//FirmaDigital permite identificar una maquina y persona autorizada por el sistema
type FirmaDigital struct {
	ID           int
	DireccionMac string
	DireccionIP  string
	Tiempo       time.Time
}

func (f *FirmaDigital) Registrar() bool {

	return true
}

//Salvar Metodo para crear usuarios del sistema
func (u *Usuario) Salvar() {
	var usr Usuario
	var privilegio Privilegio
	var lst []Privilegio

	usr.ID = 0
	usr.Nombre = "Super Usuario"

	usr.Rol.ID = ROOT
	usr.Perfil.ID = ROOT
	usr.Perfil.Descripcion = "Super Usuario"

	privilegio.Nombre = "Salvar"
	privilegio.Descripcion = "Crear Usuario"
	privilegio.Accion = "Insert()" // ES6 Metodos
	lst = append(lst, privilegio)

	privilegio.Nombre = "Modificar"
	privilegio.Descripcion = "Modificar Usuario"
	privilegio.Accion = "Update()"
	lst = append(lst, privilegio)
	usr.Perfil.Privilegios = lst

}

//Consultar el sistema de usuarios
func (u *Usuario) Consultar() (v bool) {

	return
}
