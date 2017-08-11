package cis

type ApoyoEconomico struct {
	CasoParticular    []CasoParticular
	CovenioSeguro     []ConvenioSeguro
	FondoContingencia []FondoContingencia
}

func (ae *ApoyoEconomico) Listar() []ApoyoEconomico {

}
