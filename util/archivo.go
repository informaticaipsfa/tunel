package util

<<<<<<< HEAD
import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/tealeg/xlsx"
)
=======
import "io/ioutil"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406

type Archivo struct {
	Ruta             string
	NombreDelArchivo string
	Codificacion     string
}

func (a *Archivo) Crear(cadena string) bool {
	return true
}

<<<<<<< HEAD
func (a *Archivo) LeerPorLinea(excelFileName string, PostgreSQLPENSIONSIGESP *sql.DB) bool {
	var iconstante, iconcepto string
	var codconcepto string
	xlFile, err := xlsx.OpenFile(excelFileName)

	switch excelFileName[4:7] {
	case "inv":
		codconcepto = "0000000027"
		break
	case "rcp":
		codconcepto = "0000000061"
		break
	case "sob":
		codconcepto = "0000000063"
		break
	}
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	coma := ""
	iconstante = `INSERT INTO sno_constantepersonal (codemp,codnom,codper,codcons,moncon,montopcon) VALUES `
	iconcepto = `INSERT INTO sno_conceptopersonal (codemp, codnom, codper, codconc, aplcon, valcon, acuemp,
		acuiniemp, acupat, acuinipat, acuinipataux, acupataux, acuiniempaux, acuempaux, valconaux) VALUES `
	fmt.Println("Preparando indices para el insert")
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if i > 0 {
				coma = ","
			}
			cedula := CompletarCeros(row.Cells[0].String(), 0, 10)
			monto := row.Cells[1].String()
			iconstante += coma + `('0001','0001','` + cedula + `','` + codconcepto + `',` + monto + `,0)`
			iconcepto += coma + `('0001','0001','` + cedula + `','` + codconcepto + `',1, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL)`
			i++
		}
	}
	fmt.Println("Insertando...")
	_, err = PostgreSQLPENSIONSIGESP.Exec(iconstante)
	if err != nil {
		fmt.Println("Error en la inserción: ", err.Error())
	}
	_, err = PostgreSQLPENSIONSIGESP.Exec(iconcepto)
	if err != nil {
		fmt.Println("Error en la inserción ", err.Error())
	}

	fmt.Println("Proceso exitoso...")
	fmt.Println(excelFileName[4:7])
	return true
}

// func (a *Archivo) LeerPorLinea(excelFileName string) bool {
// 	var cedula string
// 	xlFile, err := xlsx.OpenFile(excelFileName)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	Leer Archivos de XLS del Tratamiento Prolongado
// 	for _, sheet := range xlFile.Sheets {
//
// 		for _, row := range sheet.Rows {
// 			fmt.Println(row)
// 			cedula = row.Cells
// 		}
// 	}
// 	return true
// }

=======
func (a *Archivo) LeerPorLinea() bool {
	return true
}

>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
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
