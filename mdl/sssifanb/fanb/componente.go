package fanb
import (
  "github.com/gesaodin/tunel-ipsfa/sys"
)
type Componente struct {
  Codigo string
  Nombre string
  Siglas string
  Grado []Grado
}

type Grado struct {
  Codigo string
  Rango string
  Nombre string
  Descripcion string
}

//SalvarMGO Guardar
func (comp *Componente) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB("ipsfa_test").C(colecion)
		err = c.Insert(comp)
	} else {
		c := sys.MGOSession.DB("ipsfa_test").C("componente")
		err = c.Insert(comp)
	}

	return
}
