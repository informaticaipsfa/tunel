package metodobanco

import (
	"fmt"
	"os/exec"
)

type Archivos struct{}

//ComprimirTxt Comprimir la carpeta generada para los archivos bancarios
func (m *Archivos) ComprimirTxt(llave string) bool {
	zip := "zip -r " + llave + ".zip " + llave
	cmd := "cd public_web/SSSIFANB/afiliacion/temp/banco/;" + zip
	//fmt.Println(cmd)
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
		return false
	}
	//fmt.Printf("%s", out)
	return true
}
