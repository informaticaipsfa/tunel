package sssifanb

type Pension struct {
	GradoCodigo            string                   `json:"grado" bson:"grado"`
	ComponenteCodigo       string                   `json:"componente" bson:"componente"`
	Clase                  string                   `json:"clase" bson:"clase"`
	Categoria              string                   `json:"categoria" bson:"categoria"`
	Situacion              string                   `json:"situacion" bson:"situacion"`
	FechaPromocion         string                   `json:"fpromocion" bson:"fpromocion"`
	FechaUltimoAscenso     string                   `json:"fultimoascenso" bson:"fultimoascenso"`
	AnoServicio            int                      `json:"aservicio" bson:"aservicio"`
	MesServicio            int                      `json:"mservicio" bson:"mservicio"`
	DiaServicio            int                      `json:"dservicio" bson:"dservicio"`
	NumeroHijos            int                      `json:"numerohijos" bson:"numerohijos"`
	DatoFinanciero         DatoFinanciero           `json:"DatoFinanciero" bson:"datofinanciero"`
	PensionAsignada        float64                  `json:"pensionasignada" bson:"pensionasignada"`
	HistorialSueldo        []HistorialPensionSueldo `json:"HistorialSueldo" bson:"historialsueldo"`
	PorcentajePrestaciones float64                  `json:"pprestaciones" bson:"pprestaciones"`
}

type HistorialPensionSueldo struct {
	Directiva       string  `json:"directiva" bson:"directiva"`
	Sueldo          float64 `json:"sueldo" bson:"sueldo"`
	Prima           Prima   `json:"Prima" bson:"prima"`
	PensionAsignada float64 `json:"pensionasignada" bson:"pensionasignada"`
	BonoVacacional  float64 `json:"bonovacacional" bson:"bonovacacional"`
	BonoAguinaldo   float64 `json:"bonoaguinaldo" bson:"bonoaguinaldo"`
}

type Prima struct {
	Transporte          float64 `json:"transporte" bson:"transporte"`
	Descendencia        float64 `json:"descendencia" bson:"descendencia"`
	NoAscenso           float64 `json:"noascenso" bson:"noascenso"`
	PorcentajeNoAscenso float64 `json:"pnoascenso" bson:"pnoascenso"`
	Especial            float64 `json:"especial" bson:"especial"`
	SubTotal            float64 `json:"subtotal" bson:"subtotal"`
}
