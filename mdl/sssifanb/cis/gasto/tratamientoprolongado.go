package cis

type TratamientoProlongado struct {
    Medicamento       string
    TiempoTratamiento int
    Farmacia          string
}

func (tp *TratamientoProlongado) Listar() []TratamientoProlongado {

}
