// El consumidor de recursos establece la necesidad de uso e iteraciones
// a las cuales el usuario se enfrenta a diario. Recurrencia
package logs

import (
	"fmt"
	"time"
)

const (
	Irregular = 0
	Regular   = 1
)

type Comportamiento struct {
	tipo   int
	Accion string
}

type Anomalia struct {
	Comportamiento
	funcion string
	tiempo  time.Time
}

func (r *Anomalia) Agregar() {

}

func (r *Anomalia) Notificar(ValorEsperado interface{}) {
	switch v := ValorEsperado.(type) {
	case string:
		fmt.Println(v)
	case int32, int64:
		fmt.Println(v)
	default:
		fmt.Println("Valor Inesperado")
	}
}
