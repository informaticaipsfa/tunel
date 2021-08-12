package metodobanco_test

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/informaticaipsfa/tunel/util/metodobanco"
	_ "github.com/lib/pq"
)

const (
	_CADENA   string = "user=postgres dbname=pensiones password=za63qj2p host=localhost sslmode=disable"
	_FIRMA    string = "7af394c54166bb519624a976c6acd97d" //firma del archivo nomina a generar
	_DIR      string = "tmp/" + _FIRMA
	_CANTIDAD int    = 5000
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
			vzla.Cantidad = _CANTIDAD
			vzla.DesactivarArchivo = false
			vzla.Directorio = _DIR

			err = vzla.Generar(psql, "CA")

			if err == nil {

				t.Log("Proceso finalizado Venezuela CA Registro: ", vzla.Registros, " Total: ", strconv.FormatFloat(vzla.Total, 'f', 2, 64))
			} else {
				t.Error("El proceso se ejecuto pero no se encontro el archivo")
				t.Fail()
			}

			err = vzla.Generar(psql, "CC")

			if err == nil {

				t.Log("Proceso finalizado Venezuela CC Registro: ", vzla.Registros, " Total: ", strconv.FormatFloat(vzla.Total, 'f', 2, 64))
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
			banfan.CodigoEmpresa = "0026"
			banfan.NumeroEmpresa = "01770001421100683232"
			banfan.Firma = _FIRMA
			banfan.Cantidad = _CANTIDAD
			banfan.DesactivarArchivo = false
			banfan.Directorio = _DIR

			err = banfan.Generar(psql)

			if err == nil {

				t.Log("Proceso finalizado Banfanb Registro: ", banfan.Registros, " Total: ", strconv.FormatFloat(banfan.Total, 'f', 2, 64))
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
			bic.CodigoEmpresa = "0651"
			bic.NumeroEmpresa = "01750484310076626369"
			bic.Firma = _FIRMA
			bic.Cantidad = _CANTIDAD
			bic.DesactivarArchivo = false
			bic.Directorio = _DIR

			err = bic.Generar(psql)

			if err == nil {

				t.Log("Proceso finalizado Bicentenario Registro: ", bic.Registros, " Total: ", strconv.FormatFloat(bic.Total, 'f', 2, 64))
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
