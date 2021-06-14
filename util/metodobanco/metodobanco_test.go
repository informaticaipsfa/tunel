package metodobanco_test

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/informaticaipsfa/tunel/util/metodobanco"
	_ "github.com/lib/pq"
)

const (
	_CADENA string = "user=postgres dbname=pensiones password=za63qj2p host=localhost sslmode=disable"
	_FIRMA  string = "8a3afb8f96ec8a0c032fc829dc586d42" //firma del archivo nomina a generar
	_DIR    string = "tmp/8a3afb8f96ec8a0c032fc829dc586d42"
)

var test = []func(t *testing.T){
	func(t *testing.T) {
		psql, err := sql.Open("postgres", _CADENA)
		if err != nil {
			t.Error("No se logro establecer la conexion verifique el usurio y la clave")
			t.Fail()
		} else {

			var vzla metodobanco.Venezuela
			vzla.Tabla = "pagos"
			vzla.CodigoEmpresa = "0"
			vzla.NumeroEmpresa = "01020488720000002147"
			vzla.Firma = _FIRMA
			vzla.Cantidad = 1100
			vzla.DesactivarArchivo = false
			vzla.Directorio = _DIR

			err = vzla.Generar(psql, "CA")

			if err == nil {

				t.Log("Proceso finalizado Registro: ", vzla.Registros, " Total: ", strconv.FormatFloat(vzla.Total, 'f', 2, 64))
			} else {
				t.Error("El proceso se ejecuto pero no se encontro el archivo")
				t.Fail()
			}
		}
	}, //Testing Banfanb
	func(t *testing.T) {
		psql, err := sql.Open("postgres", _CADENA)
		if err != nil {
			t.Error("No se logro establecer la conexion verifique el usurio y la clave")
			t.Fail()
		} else {

			var banfan metodobanco.Banfanb
			banfan.Tabla = "pagos"
			banfan.CodigoEmpresa = "0"
			banfan.NumeroEmpresa = "01020488720000002147"
			banfan.Firma = _FIRMA
			banfan.Cantidad = 1100
			banfan.DesactivarArchivo = false
			banfan.Directorio = _DIR

			err = banfan.Generar(psql)

			if err == nil {

				t.Log("Proceso finalizado Registro: ", banfan.Registros, " Total: ", strconv.FormatFloat(banfan.Total, 'f', 2, 64))
			} else {
				t.Error("El proceso se ejecuto pero no se encontro el archivo")
				t.Fail()
			}
		}

	}, //Testing Bicentenario
	func(t *testing.T) {
		psql, err := sql.Open("postgres", _CADENA)
		if err != nil {
			t.Error("No se logro establecer la conexion verifique el usurio y la clave")
			t.Fail()
		} else {

			var bic metodobanco.Bicentenario
			bic.Tabla = "pagos"
			bic.CodigoEmpresa = "0"
			bic.NumeroEmpresa = "01020488720000002147"
			bic.Firma = _FIRMA
			bic.Cantidad = 1100
			bic.DesactivarArchivo = false
			bic.Directorio = _DIR

			err = bic.Generar(psql)

			if err == nil {

				t.Log("Proceso finalizado Registro: ", bic.Registros, " Total: ", strconv.FormatFloat(bic.Total, 'f', 2, 64))
			} else {
				t.Error("El proceso se ejecuto pero no se encontro el archivo")
				t.Fail()
			}
		}

	},
}

func TestGenerarBanco(t *testing.T) {

	for i, fn := range test {
		t.Run(strconv.Itoa(i), fn)
	}

}
