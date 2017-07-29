package sssifanb

type Fideicomiso struct {
	GradoCodigo        string `json:"grado" bson:"grado"`
	ComponenteCodigo   string `json:"Componente" bson:"componente"`
	NumeroHijos        int    `json:"numerohijos" bson:"numerohijos"`
	AnoReconocido      int    `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido      int    `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido      int    `json:"dreconocido,omitempty" bson:"dreconocido"`
	CuentaBancaria     string `json:"cuentabancaria" bson:"cuentabancaria"`
	EstatusNoAscenso   int    `json:"estatusnoascenso,omitempty" bson:"estatusnoascenso"`
	EstatusProfesion   int    `json:"estatusprofesion,omitempty" bson:"estatusprofesion"`
	MotivoParalizacion string `json:"motivoparalizacion" bson:"motivoparalizacion"`
}
