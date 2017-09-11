package util

import (
	"io/ioutil"
)

type Archivo struct {
	Ruta             string
	NombreDelArchivo string
	Codificacion     string
}

func (a *Archivo) Crear(cadena string) bool {
	return true
}

func (a *Archivo) LeerPorLinea(excelFileName string) bool {
	// var cedula string
	// xlFile, err := xlsx.OpenFile(excelFileName)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//Leer Archivos de XLS del Tratamiento Prolongado
	// for _, sheet := range xlFile.Sheets {
	//
	// 	for _, row := range sheet.Rows {
	// 		fmt.Println(row)
	// 		cedula = row.Cells
	// 	}
	// }
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
