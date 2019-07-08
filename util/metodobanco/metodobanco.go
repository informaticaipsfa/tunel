package metodobanco

import (
	"fmt"
	"os/exec"
)

type Archivos struct{}

//ComprimirTxt Comprimir la carpeta generada para los archivos bancarios
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

//Borrar Permite eliminar el directorio de los archivos asosiados a un hash
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
