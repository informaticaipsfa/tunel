// El consumidor de recursos establece la necesidad de uso e iteraciones
// a las cuales el usuario se enfrenta a diario. Recurrencia
package logs

import (
	"time"

<<<<<<< HEAD
	"github.com/informaticaipsfa/tunel/sys/seguridad"
=======
	"github.com/gesaodin/tunel-ipsfa/sys/seguridad"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
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
