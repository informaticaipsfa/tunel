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
func (a *Archivo) LeerCA(PostPension *sql.DB, codigo string) bool {
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
		fmt.Println("lineas ", l)
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
				fmt.Println("Linea # ", i, cedula, "|", concepto)
			}

		}
		i++

	}

	fmt.Println("procesando ", i)
	_, err = PostPension.Exec(insertar)
	fmt.Println(insertar)
	if err != nil {
		fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", a.Ruta, err.Error())
		return false
	}
	return true
}
