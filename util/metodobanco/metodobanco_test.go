package metodobanco_test

import (
	"fmt"
	"testing"

	"github.com/informaticaipsfa/tunel/util/metodobanco"
)

func TestGenerarBDV(t *testing.T) {

	var archvo metodobanco.Archivos

	resp := archvo.Borrar("asdf")
	fmt.Println(resp)
	if !resp {
		t.Log("Proceso finalizado correctamente")
	} else {
		t.Error("El proceso se ejecuto pero no se encontro el archivo")
		t.Fail()
	}
}

//func TesGenerarBBI(t *testing.T){
//)
