package logs

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
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

func RegistrarLog() {
	//create your file with desired read/write permissions
	f, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//defer to close when you're done with it, not because you think it's idiomatic!
	defer f.Close()
	//set output of logs to f
	log.SetOutput(f)
	//test case
	// log.Println("check to make sure it works")
}
