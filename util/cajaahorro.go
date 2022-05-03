package util

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//LeerCA Leer archivos de la caja de ahorro
func (a *Archivo) LeerCA(PostPension *sql.DB, codigo string, doc string) bool {
	var coma, concepto, tipo string

	codigomd5 := strings.Split(codigo, "|")

	a.iniciarVariable()
	insertar := a.Cabecera
	archivo, err := os.Open(a.Ruta)
	Error(err)
	scan := bufio.NewScanner(archivo)
	i := 0
	for scan.Scan() {
		linea := strings.Split(scan.Text(), ";")
		l := len(linea)
		if i == 0 {
			concepto = linea[0]
			tipo = linea[4]
		} else { //Leyendo la primera linea
			a.CantidadLineas++
			if l > 2 {
				if a.CantidadLineas > 1 {
					coma = ","
				} else {
					coma = ""
				}

				cedula, _ := strconv.Atoi(strings.Split(linea[0], ".")[0])
				familiar, _ := strconv.Atoi(strings.Split(linea[1], ".")[0])
				monto := linea[2]
				insertar += coma + "('" + strconv.Itoa(cedula) + "','" + strconv.Itoa(familiar)
				insertar += "','" + codigomd5[0] + "','" + concepto + "'," + monto + ", " + tipo + ", Now(), '" + codigomd5[1] + "' )"

			} else { //DE LO CONTRARIO

				if a.CantidadLineas > 1 {
					coma = ","
				} else {
					coma = ""
				}

				cedula, _ := strconv.Atoi(strings.Split(linea[0], ".")[0])
				monto := linea[1]
				insertar += coma + "('" + strconv.Itoa(cedula) + "','','" + codigomd5[0]
				insertar += "','" + concepto + "'," + monto + ", " + tipo + ", Now(), '" + codigomd5[1] + "')"
				//fmt.Println("Linea # ", i, cedula, "|", concepto)
			}

		}
		i++

	}

	fmt.Println("procesando ", i, doc)
	_, err = PostPension.Exec(insertar)
	//fmt.Println(insertar)
	if err != nil {
		fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", a.Ruta)
		a.EscribirLinea([]byte(err.Error()))
		return false
	}

	return true
}

//LeerArchivos del SIS
func (a *Archivo) LeerSisa(PostPension *sql.DB, codigo string, doc string, concepto string, mes string) bool {
	var coma string

	a.iniciarVariableSisa()
	insertar := a.Cabecera
	archivo, err := os.Open(a.Ruta)
	Error(err)
	scan := bufio.NewScanner(archivo)
	i := 0
	fmt.Println(codigo, concepto)
	for scan.Scan() {
		linea := strings.Split(scan.Text(), ";")
		l := len(linea)

		if i > 0 {
			if l > 1 {

				coma = ""
				if i > 1 {
					coma = ","

				}
				//fmt.Println(linea[0])
				ced := linea[0][1:]
				cedula, _ := strconv.Atoi(strings.Split(ced, ".")[0])
				monto := ReemplazarComaPorPunto(EliminarEspacioBlanco(linea[1]))
				insertar += coma + "('" + strconv.Itoa(cedula) + "','','" + mes
				insertar += "','" + concepto + "'," + monto + ", 3, Now(), '" + codigo + "')"
				//fmt.Println("Linea # ", i, cedula, "|", concepto)
			}
		}
		i++

	}

	fmt.Println("procesando ", i-1, doc)
	_, err = PostPension.Exec(insertar)
	//fmt.Println(insertar)
	if err != nil {
		fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", a.Ruta)
		a.EscribirLinea([]byte(err.Error()))
		fmt.Println(err.Error())
		return false
	}

	return true
}
