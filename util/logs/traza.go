package logs

import (
	"fmt"
	"time"

<<<<<<< HEAD
	"github.com/informaticaipsfa/tunel/sys/seguridad"
=======
	"github.com/gesaodin/tunel-ipsfa/sys/seguridad"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
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
