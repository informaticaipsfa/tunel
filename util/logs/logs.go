package logs

import (
	"crypto/md5"
	"encoding/hex"
)

type Log struct {
	Id int
	Traza
}

func (l *Log) Salvar() {
	return
}

func (l *Log) Consultar(id string) bool {
	return true
}

func (l *Log) listar(id string) (lo []Log) {
	return
}

func (l *Log) Firmar() {
	valor := md5.Sum([]byte(l.Usuario.Nombre + l.DireccionIP + l.DireccionMac + l.CreadoEnFecha.String()))
	l.HuellaDigital = hex.EncodeToString(valor[:])
}

func (l *Log) ValidarFirma() (b bool) {

	valor := md5.Sum([]byte(l.Usuario.Nombre + l.DireccionIP + l.DireccionMac + l.CreadoEnFecha.String()))
	s := hex.EncodeToString(valor[:])
	if s == l.HuellaDigital {
		b = true
	}
	return b
}
