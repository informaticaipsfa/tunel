package metodobanco

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/informaticaipsfa/tunel/util"
)

// Archivos Generalizando normas del archivo
type Archivos struct{}

// ComprimirTxt Comprimir la carpeta generada para los archivos bancarios
func (m *Archivos) ComprimirTxt(llave string) bool {

	zip := "zip -r " + llave + ".zip " + llave
	cmd := "cd " + URLBancoZIP + ";" + zip
	fmt.Println(cmd)
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
		return false
	}
	//fmt.Printf("%s", out)
	return true
}

// Borrar Permite eliminar el directorio de los archivos asosiados a un hash
func (m *Archivos) Borrar(llave string) bool {
	cmd := "rm -rf " + llave + "*"
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
		return false
	}
	//fmt.Printf("%s", out)
	return true

}

// generarCedula Permite generar el campo de cedula completando ceros
// al mismo tiempo limpia la cadena de espacios en blanco y elimna puntos
// orientacion establece el inicio si es de derecha a izquierda 0 BDV
// cantidad estable el formato de autollenado para completar ceros segun su orientacion
func generarCedula(cedula sql.NullString, orientacion int, cantidad int) string {
	sCedula := util.ValidarNullString(cedula)
	return util.CompletarCeros(util.EliminarEspacioBlanco(util.EliminarPuntoDecimal(sCedula)), orientacion, cantidad)[:cantidad]
}

// generarNombre basado en el autocompletar con espacio BDV
func generarNombre(nombre sql.NullString, orientacion int, cantidad int) string {
	return util.CompletarEspacios(util.ValidarNullString(nombre), orientacion, cantidad)[:cantidad]
}

// generarMonto Permite evaluar valores de postgres y convertirlos en cadena con float64
func generarMonto(neto sql.NullFloat64, orientacion int, cantidad int) (monto float64, smonto string) {
	monto = util.ValidarNullFloat64(neto)
	smonto = util.EliminarPuntoDecimal(strconv.FormatFloat(util.ValidarNullFloat64(neto), 'f', 2, 64))
	smonto = util.CompletarCeros(smonto, orientacion, cantidad)
	return
}

// generarCuentaBancaria Crear formato de cuetnas bancarias
func generarCuentaBancaria(cuenta sql.NullString) string {
	sCuenta := util.ValidarNullString(cuenta)
	return util.CompletarCeros(util.EliminarUnderScore(sCuenta), 0, 20)[:20] //20 Numero de cuenta)
}

// genearFecha Archivos del Banco Venezuala
// formato dd/mm/aa
func generarFecha() string {
	fecha := time.Now()
	dd := fecha.String()[8:10]
	mm := fecha.String()[5:7]
	aa := fecha.String()[2:4]
	return dd + "/" + mm + "/" + aa
}

// crearDirectorio permite iniciar la carpeta donde se crearan los documentos
func crearDirectorio(Dir string, desactivar bool, firma string, tabla string) string {
	if !desactivar {
		if Dir == "" {
			Dir = URLBanco + firma + definirArchivo(tabla)
		}

		err := os.Mkdir(Dir, 0777)
		util.Error(err)
	}
	return Dir
}

// definirArchivo para su asignacion y creacion en los documentos
func definirArchivo(tabla string) (valor string) {

	valor = ""
	if tabla == "rechazos" {
		valor = "-XR"
	}
	return
}
