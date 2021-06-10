package metodobanco_test

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/informaticaipsfa/tunel/util/metodobanco"
	_ "github.com/lib/pq"
)

func TestGenerarBDV(t *testing.T) {
	cadena := "user=postgres dbname=pensiones password=za63qj2p host=localhost sslmode=disable"
	psql, err := sql.Open("postgres", cadena)
	if err != nil {
		t.Error("No se logro establecer la conexion verifique el usurio y la clave")
		t.Fail()
	} else {

		var vzla metodobanco.Venezuela
		vzla.Tabla = "pagos"
		vzla.CodigoEmpresa = "0"
		vzla.NumeroEmpresa = "01020488720000002147"
		vzla.Firma = "8a3afb8f96ec8a0c032fc829dc586d42" //firma del archivo nomina a generar
		vzla.Cantidad = 1100
		vzla.DesactivarArchivo = false
		vzla.Directorio = "tmp/8a3afb8f96ec8a0c032fc829dc586d42"

		err = vzla.Generar(psql, "CA")

		if err == nil {

			t.Log("Proceso finalizado Registro: ", vzla.Registros, " Total: ", strconv.FormatFloat(vzla.Total, 'f', 2, 64))
		} else {
			t.Error("El proceso se ejecuto pero no se encontro el archivo")
			t.Fail()
		}
	}
}
