package metodobanco

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/informaticaipsfa/tunel/util"
)

//Cobranza Control de Banco
type Cobranza struct {
	Firma         string
	Cantidad      int
	CodigoEmpresa string
	NumeroEmpresa string
	Fecha         string
	Tabla         string
}

//CabeceraSQL Creando consulta para archivos
func (CB *Cobranza) CabeceraSQL(desde string, hasta string, componente string) string {
	return `SELECT cedula, nomb
  FROM space.credito crd WHERE comp = '` + componente + `' AND crea BETWEEN '` + desde + `' AND '` + hasta + `'`

}

//GenerarCobranza Creando consulta para archivos
func (CB *Cobranza) GenerarCobranza(PostgreSQLPENSIONSIGESP *sql.DB) {
	fmt.Println(CB.CabeceraSQL("", "", "GN"))
	sq, err := PostgreSQLPENSIONSIGESP.Query(CB.CabeceraSQL("", "", "GN"))
	util.Error(err)

	i := 0

	directorio := URLCobranza + "cobranza/" + CB.Firma

	errr := os.Mkdir(directorio, 0777)
	util.Error(errr)
	for sq.Next() {
		i++
		var cedula, nombre sql.NullString

		e := sq.Scan(&cedula, &nombre)
		util.Error(e)
	}

}
