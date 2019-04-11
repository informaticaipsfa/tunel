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
func (a *Archivo) LeerCA(PostPension *sql.DB) bool {
	var coma string
	a.iniciarVariable()
	insertar := a.Cabecera
	archivo, err := os.Open(a.Ruta)
	Error(err)
	scan := bufio.NewScanner(archivo)

	for scan.Scan() {
		linea := strings.Split(scan.Text(), "|")
		l := len(linea)
		a.CantidadLineas++
		//fmt.Println("Linea -> ", linea)
		if l > 3 {

			if a.CantidadLineas > 2 {
				coma = ","
			} else {
				coma = ""
			}
			insertar += coma
			if a.CantidadLineas > 1 {

				//re := regexp.MustCompile(`[-().]`)
				concepto := strings.Split(linea[0], ".")
				cedula, _ := strconv.Atoi(strings.Split(linea[1], ".")[0])
				monto := linea[2]
				// premio, comision := RComaXPunto(linea[l-1]), RComaXPunto(linea[l-2])
				insertar += "('" + strconv.Itoa(cedula) + "','" + concepto[0] + "'," + monto + ", 2, Now() )"

			}

		}
	}
	//fmt.Println(insertar)
	_, err = PostPension.Exec(insertar)
	if err != nil {
		fmt.Println("ERR. CAJA DE AHORRO ", err.Error())
		return false
	}
	return true
}
