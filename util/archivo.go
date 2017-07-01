package util

import "io/ioutil"

type Archivo struct {
	Ruta             string
	NombreDelArchivo string
	Codificacion     string
}

func (a *Archivo) Crear(cadena string) bool {
	return true
}

func (a *Archivo) LeerPorLinea() bool {
	return true
}

func (a *Archivo) LeerTodo() (f []byte, err error) {
	f, err = ioutil.ReadFile(a.NombreDelArchivo)
	return
}

func (a *Archivo) EscribirLinea(linea string) bool {
	return true
}

func (a *Archivo) Cerrar() bool {
	return true
}
