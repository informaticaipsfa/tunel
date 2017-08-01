package sssifanb

import "github.com/gesaodin/tunel-ipsfa/sys"

//Recibo de Pago
type Recibo struct {
	Numero      string  `json:"numero" bson:"numero"`
	CanalDePago string  `json:"canaldepago" bson:"canaldepago"`
	Monto       float32 `json:"monto" bson:"monto"`
	Instrumento int     `json:"instrumento" bson:"instrumento"`
}

//SalvarMGO Guardar
func (m *Militar) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB(CBASE).C(colecion)
		err = c.Insert(m)
	}

	//fmt.Println(err)

	return
}
