package util

import (
	"bufio"
	"os"
	"strings"
)

//LeerCA Archivos de Loteria
func (a *Archivo) LeerCA() bool {

	archivo, err := os.Open(a.Ruta)
	Error(err)
	scan := bufio.NewScanner(archivo)

	for scan.Scan() {
		linea := strings.Fields(scan.Text())
		l := len(linea)
		if l > 3 {
			if "c" == strings.Trim(linea[0], " ") {
				a.Leer = true
				a.CantidadLineas++
			}
			if a.Leer {
				if a.CantidadLineas > 2 && strings.Trim(linea[0], " ") != "TOTALES:" && strings.Trim(linea[0], " ") != "" {
					// coma = ","
				} else {
					// coma = ""
				}
				// insertar += coma
				if a.CantidadLineas > 1 && "TOTALES:" != strings.Trim(linea[0], " ") && strings.Trim(linea[0], " ") != "" {

					// re := regexp.MustCompile(`[-()]`)
					// agen := re.Split(linea[0], -1)
					// agencia, venta := agen[1], RComaXPunto(linea[l-3])
					// premio, comision := RComaXPunto(linea[l-1]), RComaXPunto(linea[l-2])
					// insertar += "('" + agencia + "'," + venta + "," + premio + ","
					// insertar += comision + ",1,'" + a.Fecha + "',Now(),"
					// insertar += strconv.Itoa(posicionarchivo) + "," + strconv.Itoa(oid) + ")"
					// a.Salvar = true
				}
				a.CantidadLineas++
			}
		}
	}
	return true
}
