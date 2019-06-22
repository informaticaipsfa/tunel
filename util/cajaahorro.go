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
	var coma string
	a.iniciarVariable()
	insertar := a.Cabecera
	archivo, err := os.Open(a.Ruta)
	Error(err)
	scan := bufio.NewScanner(archivo)
	i := 0
	for scan.Scan() {
		linea := strings.Split(scan.Text(), "|")
		l := len(linea)
		a.CantidadLineas++
		tipo := "2"
		if l > 4 {
			if a.CantidadLineas > 1 {
				coma = ","
			} else {
				coma = ""
			}

			if l > 3 {
				tipo = linea[4]
			}
			concepto := linea[0]
			cedula, _ := strconv.Atoi(strings.Split(linea[1], ".")[0])
			familiar, _ := strconv.Atoi(strings.Split(linea[2], ".")[0])
			monto := linea[3]
			insertar += coma + "('" + strconv.Itoa(cedula) + "','" + strconv.Itoa(familiar) + "','" + codigo + "','" + concepto + "'," + monto + ", " + tipo + ", Now() )"
		} else { //DE LO CONTRARIO ARCHIVO DE SOBREVIVIENTES
			tipo := "2"
			if l > 3 {
				tipo = linea[3]
			}
			if l > 2 {
				if a.CantidadLineas > 1 {
					coma = ","
				} else {
					coma = ""
				}
				concepto := linea[0]
				cedula, _ := strconv.Atoi(strings.Split(linea[1], ".")[0])
				monto := linea[2]
				insertar += coma + "('" + strconv.Itoa(cedula) + "','','" + codigo + "','" + concepto + "'," + monto + ", " + tipo + ", Now() )"
				//fmt.Println("Linea # ", i, cedula, "|", concepto)
			}
		}

		i++
	}

	fmt.Println("procesando ", i)
	_, err = PostPension.Exec(insertar)
	fmt.Println("Control")
	if err != nil {
		fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", a.Ruta, err.Error())
		return false
	}
	return true
}
