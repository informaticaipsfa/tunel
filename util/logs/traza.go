package logs

import (
	"fmt"
	"time"

	"github.com/gesaodin/bdse/sys/seguridad"
)

type Traza struct {
	seguridad.Usuario `json:"usuario"`
	Sistema           string    `json:"sistema"`
	Accion            string    `json:"accion"`
	Metodo            string    `json:"metodo"`
	Pregunta          string    `json:"pregunta"`
	DireccionMac      string    `json:"mac"`
	DireccionIP       string    `json:"ip"`
	CreadoEnFecha     time.Time `json:"creadoenfecha"`
	HuellaDigital     string    `json:"huelladigital"`
}

func (t *Traza) Salva() error {
	fmt.Println("Err: ", t.HuellaDigital)
	return nil
}
