package util

import (
	"database/sql"
	"io/ioutil"
)

type Archivo struct {
	Responsable      int
	Ruta             string
	NombreDelArchivo string
	Codificacion     string
	Cabecera         string
	Leer             bool
	Salvar           bool
	Fecha            string
	CantidadLineas   int
	Registros        int
	PostgreSQL       *sql.DB
	Canal            chan []byte
}

//Directorio operaciones sobre un directorio
type Directorio struct {
	Ruta    string
	Listado []string `json:"listado"`
}

func (a *Archivo) iniciarVariable() {
	a.Cabecera = "INSERT INTO space.nomina_archivo (cedu,fami, llav, conc,mont,tipo,fech, proc) VALUES "
	a.CantidadLineas = 0
	a.Leer = false
	a.Salvar = false
}

func (a *Archivo) Crear(cadena string) bool {
	return true
}

//LeerTodo Cargar todo tipo de archivos al sistema
func (a *Archivo) LeerTodo() (f []byte, err error) {
	f, err = ioutil.ReadFile(a.NombreDelArchivo)
	return
}

func (a *Archivo) EscribirLinea(linea string) bool {
	return true
}

//Cerrar Archivos
func (a *Archivo) Cerrar() bool {
	return true
}
