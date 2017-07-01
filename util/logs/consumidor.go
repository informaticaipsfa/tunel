// El consumidor de recursos establece la necesidad de uso e iteraciones
// a las cuales el usuario se enfrenta a diario. Recurrencia
package logs

import (
	"time"

	"github.com/gesaodin/bdse/sys/seguridad"
)

type Recurso struct {
	seguridad.Usuario
	funcion string
	tiempo  time.Time
}

func (r *Recurso) Agregar() {

}

func (r *Recurso) Calcular() int {
	cant := 0
	return cant
}

func (r *Recurso) Notificar() {

}
