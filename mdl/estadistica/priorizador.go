//estadistica modelada
package estadistica

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
	"github.com/gesaodin/tunel-ipsfa/sys"
)

type Priorizador struct {
}

func filtrar() {

}

func (p *Priorizador) Capturador() {
	for {
		p.Evaluador()
		time.Sleep(100 * time.Millisecond)
	}
}

//Evaluador
func (p *Priorizador) Evaluador() (jSon []byte, err error) {
	var msj Mensaje
	s := `SELECT cedula,nombreprimero,nombresegundo,apellidoprimero,apellidosegundo,esta FROM seguimiento.persona_esta AS pe
  JOIN public.personas AS p ON p.codnip=pe.cedula
  WHERE esta != 0`
	sq, err := sys.PostgreSQLSAMAN.Query(s)
	if err != nil {
		msj.Mensaje = "Error: Consulta ya existe."
		msj.Tipo = 2
		msj.Pgsql = err.Error()
		jSon, err = json.Marshal(msj)
		fmt.Println(msj)
		return
	}
	for sq.Next() {
		var militar sssifanb.Militar
		var cedula, nombreprimero, nombresegundo, apellidoprimero, apellidosegundo string
		var estatus int
		sq.Scan(&cedula, &nombreprimero, &nombresegundo, &apellidoprimero, &apellidosegundo, &estatus)
		militar.Persona.DatoBasico.Cedula = cedula
		militar.Persona.DatoBasico.NombrePrimero = nombreprimero
		militar.Persona.DatoBasico.NombreSegundo = nombresegundo
		militar.SalvarMGO()
		msj.Mensaje = "OK: Sincronización (SAMAN -> MONGODB: " + militar.Persona.DatoBasico.Cedula + ") "
		fmt.Println(msj.Mensaje)
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
		sSQL := `UPDATE seguimiento.persona_esta SET esta=0 WHERE cedula='` + militar.Persona.DatoBasico.Cedula + `';`
		errorG(sSQL)
	}
	return
}
