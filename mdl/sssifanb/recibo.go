package sssifanb

import (
	"fmt"
	"time"

	"github.com/gesaodin/tunel-ipsfa/sys"
)

//Recibo de Pago
type Recibo struct {
	ID          string    `json:"id" bson:"id"`
	IDF         string    `json:"idf" bson:"idf"` //ID FAMILIAR
	Numero      string    `json:"numero" bson:"numero"`
	CanalDePago string    `json:"canal" bson:"canal"`
	Fecha       time.Time `json:"fecha" bson:"fecha"`
	Monto       float32   `json:"monto" bson:"monto"`
	Motivo      string    `json:"motivo" bson:"motivo"`
	IP          string    `json:"ip" bson:"ip"`
	Usuario     string    `json:"usuario" bson:"usuario"`
}

//Salvar Guardar
func (r *Recibo) Salvar() (err error) {
	var TIM Carnet
	TIM.ID = r.ID
	TIM.IDF = r.IDF
	TIM.IP = r.IP
	TIM.Motivo = r.Motivo
	TIM.Usuario = r.Usuario
	fmt.Println(r.Usuario)
	TIM.Salvar()
	c := sys.MGOSession.DB(CBASE).C(CRECIBO)
	err = c.Insert(r)
	return
}

//Consultar Recibos
func (r *Recibo) Consultar(id string) (err error) {

	return
}

//Listar Recibos
func (r *Recibo) Listar(id string) (err error) {

	return
}
