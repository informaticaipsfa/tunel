package util

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//LeerCobranza Leer archivos de la caja de ahorro
func (a *Archivo) LeerCobranza(PostPension *sql.DB, codigo string, doc string) bool {
	var coma string

	//codigomd5 := strings.Split(codigo, "|")

	insertar := "INSERT INTO space.credito_pagos (cedu,llav,cred,mont,esta,fech) VALUES "
	archivo, err := os.Open(a.Ruta)
	Error(err)
	scan := bufio.NewScanner(archivo)
	i := 0
	for scan.Scan() {
		columna := strings.Split(scan.Text(), "|")
		//l := len(columna)

		coma = ""
		if i > 2 {
			coma = ","
		}
		cedula, _ := strconv.Atoi(strings.Split(columna[0], ".")[0])
		// estatus := columna[7]
		// monto := columna[2]
		// cuota := columna[3]
		// credito := columna[5]

		insertar += coma + "('" + strconv.Itoa(cedula)
		// insertar += "','" + codigomd5[0] + "','" + concepto + "'," + monto + ", " + tipo + ", Now(), '" + codigomd5[1] + "' )"

		i++

	}

	fmt.Println("procesando ", i, doc)
	_, err = PostPension.Exec(insertar)
	//fmt.Println(insertar)
	if err != nil {
		fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", a.Ruta, err.Error())
		return false
	}
	return true
}
