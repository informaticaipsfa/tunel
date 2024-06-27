package sssifanb

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/investigacion"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
)

// Familiar Busquedas
type Familiar struct {
	ID                     string                 `json:"id" bson:"id"`
	Persona                Persona                `json:"Persona" bson:"persona"`
	FechaAfiliacion        time.Time              `json:"fechaafiliacion" bson:"fechaafiliacion"`
	Parentesco             string                 `json:"parentesco" bson:"parentesco"` //0:Mama, 1:papa, 2: Esposa  3: hijo
	EsMilitar              bool                   `json:"esmilitar" bson:"esmilitar"`
	Condicion              int                    `json:"condicion" bson:"condicion"` //Sano o Condicion especial
	Estudia                int                    `json:"estudia" bson:"estudia"`
	Benficio               bool                   `json:"beneficio" bson:"beneficio"` //
	Documento              int                    `json:"documento" bson:"documento"`
	Adoptado               bool                   `json:"adoptado" bson:"adoptado"`
	DocumentoPadre         string                 `json:"documentopadre" bson:"documentopadre"`
	HistoriaMedica         string                 `json:"historiamedica" bson:"historiamedica"`
	Donante                string                 `json:"donante" bson:"donante"`
	SituacionPago          string                 `json:"situacionpago" bson:"situacionpago"` //Este estatus permite habilitar un pago en pensiones 201 | 202 o paralizar
	RazonPago              string                 `json:"razonpago" bson:"razonpago"`         //Razon de Pago
	EstatusCarnet          int                    `json:"estatuscarnet" bson:"estatuscarnet"`
	GrupoSanguineo         string                 `json:"gruposanguineo" bson:"gruposanguineo"`
	PorcentajePrestaciones float64                `json:"pprestaciones,omitempty" bson:"pprestaciones"`
	TIF                    Carnet                 `json:"Tif" bson:"tif"`
	FeDeVida               investigacion.FeDeVida `json:"FeDeVida" bson:"fedevida"`
	EstatusDePension       bool                   `json:"estatuspension" bson:"estatuspension"`
	CondicionPago          string                 `json:"condicionpago" bson:"condicionpago"` //Para ver si es Cheque o Banco
	// EstatusAfiliacion string `json:"estatus" bson:"adoptado"`
	// RazonAfiliacion   string `json:"adoptado" bson:"adoptado"`
}

// FamiliarEstadistica Busquedas
type FamiliarEstadistica struct {
	ID              string    `json:"id" bson:"id"`
	FechaNacimiento time.Time `json:"fechanacimiento" bson:"fechanacimiento"`
	Componente      string    `json:"componente,omitempty" bson:"componente"`
	Grado           string    `json:"grado,omitempty" bson:"grado"`         //grado
	Categoria       string    `json:"categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion       string    `json:"situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	IDF             string    `json:"idf" bson:"idf"`
	Parentesco      string    `json:"parentesco" bson:"parentesco"` //0:Mama, 1:papa, 2: Esposa  3: hijo
	EsMilitar       bool      `json:"esmilitar" bson:"esmilitar"`
	Benficio        bool      `json:"beneficio" bson:"beneficio"` //
	Condicion       int       `json:"condicion" bson:"condicion"` //Sano o Condicion especial
}

// AplicarReglasBeneficio OJO SEGUROS HORIZONTES
func (f *Familiar) AplicarReglasBeneficio() {
	edad, _, _ := util.CalcularTiempo(f.Persona.DatoBasico.FechaNacimiento)
	if f.Parentesco == "HJ" {
		f.Benficio = false

		if f.Condicion == 1 {
			f.Benficio = true
		} else {
			if edad < 18 {
				f.Benficio = true
			} else if f.Estudia == 1 && edad < 27 {
				f.Benficio = true
			}
		}
	} else { // ESPOSA Y PADRES
		f.Benficio = true
	}

}

// AplicarReglasParentesco Reglas
func (f *Familiar) AplicarReglasParentesco() {

}

// ConvertirFechaHumano Validacion
func (f *Familiar) ConvertirFechaHumano() {

}

// MGOActualizar Vida Militar
func (f *Familiar) MGOActualizar() (jSon []byte, err error) {
	var msj Mensaje
	msj.Mensaje = "Su data ha sido actualizada."
	msj.Tipo = 2
	jSon, err = json.Marshal(msj)
	f.MGOActualizar()
	return
}

// Actualizar Actualizando en MONGO
func (f *Familiar) Actualizar(usuario string) (jSon []byte, err error) {

	id := f.Persona.DatoBasico.Cedula
	familiar := make(map[string]interface{})
	familiar["familiar.$.persona"] = f.Persona
	//fmt.Println(f.Persona.DatoFinanciero)
	a, _, _ := f.Persona.DatoBasico.FechaDefuncion.Date()
	if a > 1900 {
		f.Benficio = false
	}
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	if f.ID != id {
		fmt.Println("Cambio de Cedula de un familiar: ", id)
		_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": f.ID}, bson.M{"$set": familiar})
	} else {
		_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": id}, bson.M{"$set": familiar})
	}
	if err != nil {
		log.Println("Cedula: " + id + " -> " + err.Error())
		return
	}

	fechaafiliacion := make(map[string]interface{})
	fechaafiliacion["familiar.$.fechaafiliacion"] = time.Now()
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": fechaafiliacion})
	if err != nil {
		fmt.Println("Incluyendo parentesco eRR Cedula: " + id + " -> " + err.Error())
	}

	//
	parentesco := make(map[string]interface{})
	parentesco["familiar.$.parentesco"] = f.Parentesco
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": parentesco})
	if err != nil {
		fmt.Println("Incluyendo parentesco eRR Cedula: " + id + " -> " + err.Error())
	}

	beneficio := make(map[string]interface{})
	beneficio["familiar.$.beneficio"] = f.Benficio
	_, err = c.UpdateAll(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": beneficio})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
	}
	historia := make(map[string]interface{})
	historia["familiar.$.historiamedica"] = f.HistoriaMedica
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": historia})
	if err != nil {
		fmt.Println("eRR Cedula: " + id + " -> " + err.Error())
	}
	donante := make(map[string]interface{})
	donante["familiar.$.donante"] = f.Donante
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": donante})
	if err != nil {
		fmt.Println("Donante (Cedula): " + id + " -> " + err.Error())
	}

	documentopadre := make(map[string]interface{})
	documentopadre["familiar.$.documentopadre"] = f.DocumentoPadre
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": documentopadre})
	if err != nil {
		fmt.Println("Parentesco (Cedula): " + id + " -> " + err.Error())
	}

	condicion := make(map[string]interface{})
	condicion["familiar.$.condicion"] = f.Condicion
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": condicion})
	if err != nil {
		fmt.Println("Condicion (Cedula): " + id + " -> " + err.Error())
	}

	derecho := make(map[string]interface{})
	derecho["familiar.$.pprestaciones"] = f.PorcentajePrestaciones
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": derecho})
	if err != nil {
		fmt.Println("Condicion (Cedula): " + id + " -> " + err.Error())
	}
	estudia := make(map[string]interface{})
	estudia["familiar.$.estudia"] = f.Estudia
	err = c.Update(bson.M{"familiar.persona.datobasico.cedula": id, "id": f.DocumentoPadre}, bson.M{"$set": estudia})
	if err != nil {
		fmt.Println("Estudia (Cedula): " + id + " -> " + err.Error())
	}

	var mOriginal Militar
	mOriginal, _ = consultarMongo(f.DocumentoPadre)
	go f.ActualizarPorReduccion(mOriginal.Grado.Abreviatura, mOriginal.Componente.Abreviatura)
	go f.ActualizarCuentaBancaria(usuario)
	return
}

// ActualizarPorReduccion Control de Reduccion de datos
func (f *Familiar) ActualizarPorReduccion(grado string, componente string) {

	//Reducción
	reduc := make(map[string]interface{})
	cred := sys.MGOSession.DB(sys.CBASE).C(sys.CREDUCCION)
	reduc["cedula"] = f.Persona.DatoBasico.Cedula
	reduc["fechanacimiento"] = f.Persona.DatoBasico.FechaNacimiento
	reduc["nombre"] = f.Persona.DatoBasico.ConcatenarNombreApellido()
	reduc["grado"] = grado
	reduc["componente"] = componente
	err := cred.Update(bson.M{"cedula": f.Persona.DatoBasico.Cedula}, bson.M{"$set": reduc})
	if err != nil {
		fmt.Println("Err", err.Error())
	}
}

// ActualizarCuentaBancaria Cuenta Bancaria
func (f *Familiar) ActualizarCuentaBancaria(usuario string) {
	var cabecera, cuerpo, autorizado, tipo, banco, cuenta, titular string
	if f.PorcentajePrestaciones > 0 {

		cabecera = `UPDATE familiar SET `

		if len(f.Persona.DatoFinanciero) > 0 {
			autorizado = f.Persona.DatoFinanciero[0].Autorizado
			tipo = f.Persona.DatoFinanciero[0].Tipo
			banco = f.Persona.DatoFinanciero[0].Institucion
			cuenta = f.Persona.DatoFinanciero[0].Cuenta
			titular = f.Persona.DatoFinanciero[0].Titular
		}

		cuerpo += `autorizado ='` + autorizado + `',
			edo_civil='` + f.Persona.DatoBasico.Nacionalidad + `',
			tipo='` + tipo + `',
			banco='` + banco + `',
			numero='` + cuenta + `',
			f_ult_modificacion=Now(),
			usr_modificacion='` + usuario + `',
			nombre_autorizado='` + titular + `' WHERE cedula='` + f.Persona.DatoBasico.Cedula + `';`
	}

	query := cabecera + cuerpo
	_, err := sys.PostgreSQLPENSION.Exec(query)
	if err != nil {
		fmt.Println("familiar.go ( ActualizarCuentaBancaria ) : ", err.Error())
	} else {
		fmt.Println("familiar.go ( ActualizarCuentaBancaria ) : OK!")
	}
}

// AplicarReglasCarnetHijos Reglas
func (f *Familiar) AplicarReglasCarnetHijos() (TIM Carnet) {
	var mes, dia string
	carnet := make(map[string]interface{})
	fechaActual := time.Now()
	Anio, Mes, Dia := fechaActual.Date()
	edad, _, _ := util.CalcularTiempo(f.Persona.DatoBasico.FechaNacimiento)
	layOut := "2006-01-02"

	switch {
	case edad < 13:
		Anio += 5
		break
	case edad >= 13 && edad <= 15:
		Anio += 18 - edad
		break
	case edad >= 15 && edad <= 17:
		Anio += 18 - edad
		break
	case edad > 17 && edad <= 27 && f.Estudia == 1:
		Anio += 26 - edad
		break
	case edad > 18 && f.Condicion == 1:
		Anio += 5
		break
	}

	mes = strconv.Itoa(int(Mes))
	if int(Mes) < 10 {
		mes = "0" + strconv.Itoa(int(Mes))
	}
	dia = strconv.Itoa(Dia)
	if Dia < 10 {
		dia = "0" + strconv.Itoa(Dia)
	}
	fecha := strconv.Itoa(Anio) + "-" + mes + "-" + dia

	fechaVencimientoCarnet, _ := time.Parse(layOut, fecha)
	TIM.Serial = TIM.GenerarSerial()
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimientoCarnet
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.Tipo = 0
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hijos")
	}
	return
}

// AplicarReglasCarnetHijos Reglas
func (f *Familiar) AplicarReglasCarnetHermanos() (TIM Carnet) {
	var mes, dia string
	carnet := make(map[string]interface{})
	fechaActual := time.Now()
	Anio, Mes, Dia := fechaActual.Date()
	layOut := "2006-01-02"
	Anio += 1
	mes = strconv.Itoa(int(Mes))
	if int(Mes) < 10 {
		mes = "0" + strconv.Itoa(int(Mes))
	}
	dia = strconv.Itoa(Dia)
	if Dia < 10 {
		dia = "0" + strconv.Itoa(Dia)
	}
	fecha := strconv.Itoa(Anio) + "-" + mes + "-" + dia
	fechaVencimientoCarnet, _ := time.Parse(layOut, fecha)
	TIM.Serial = TIM.GenerarSerial()
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimientoCarnet
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.Tipo = 0
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hermanos")
	}
	return
}

// IncluirFamiliar Agregar
func (f *Familiar) IncluirFamiliar(usuario string) (err error) {
	familiar := make(map[string]interface{})
	familiar["familiar"] = f
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	// var fam map[string]interface{}
	var parametro string
	//
	// fmt.Println("cedula", f.DocumentoPadre, " familiar ", f.Persona.DatoBasico.Cedula)
	// donde := bson.M{"id": f.DocumentoPadre, "familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula}
	// e := c.Find(donde).Select(bson.M{"familiar.persona.datobasico.cedula": true, "_id": 0}).One(&fam)
	// if e != nil {
	// 	fmt.Println("Insertando...", e.Error())
	// 	// return
	// }

	// m := fam.(map[string]map[string]map[string]map[string]string)
	// fmt.Println(fam)
	// fmt.Println("Actualizando...")
	// fami := fam["familiar"]
	// for fami {
	// 	fmt.Println(c, v)
	//
	// }

	parametro = "$push"
	// if fam.Persona.DatoBasico.Cedula != "" {
	// 	parametro = "$set"
	// }
	err = c.Update(bson.M{"id": f.DocumentoPadre}, bson.M{parametro: familiar})
	if err != nil {
		fmt.Println(" " + err.Error())
		return
	}
	return
}

// ContarFamiliar Contando Familiares
func (f *Familiar) ContarFamiliar() {

}

// AplicarReglasCarnetPadres Reglas
func (f *Familiar) AplicarReglasCarnetPadres() (TIM Carnet) {
	carnet := make(map[string]interface{})
	var mes, dia string
	var fechaVencimiento time.Time
	fechaActual := time.Now()
	AnnoA, MesA, DiaA := fechaActual.Date()
	layout := "2006-01-02"

	if f.Parentesco == "PD" {
		AnnoA += 10
		mes = strconv.Itoa(int(MesA))
		if int(MesA) < 10 {
			mes = "0" + strconv.Itoa(int(MesA))
		}
		dia = strconv.Itoa(DiaA)
		if DiaA < 10 {
			dia = "0" + strconv.Itoa(DiaA)
		}
		fvenc := strconv.Itoa(AnnoA) + "-" + mes + "-" + dia
		fechaVencimiento, _ = time.Parse(layout, fvenc)
	}

	TIM.Serial = TIM.GenerarSerial()
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimiento
	// TIM.CodigoComponente = m.Componente.Abreviatura
	// TIM.Grado.Abreviatura = m.Grado.Abreviatura
	TIM.Responsable = f.Persona.DatoBasico.Cedula
	TIM.Tipo = 1
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hijos")
	}
	return

}

// AplicarReglasCarnetEsposa Reglas
func (f *Familiar) AplicarReglasCarnetEsposa() (TIM Carnet) {
	carnet := make(map[string]interface{})
	var mes, dia string
	var fechaVencimiento time.Time
	fechaActual := time.Now()
	AnnoA, MesA, DiaA := fechaActual.Date()
	layout := "2006-01-02"

	if f.Parentesco == "EA" {
		AnnoA += 10
		mes = strconv.Itoa(int(MesA))
		if int(MesA) < 10 {
			mes = "0" + strconv.Itoa(int(MesA))
		}
		dia = strconv.Itoa(DiaA)
		if DiaA < 10 {
			dia = "0" + strconv.Itoa(DiaA)
		}
		fvenc := strconv.Itoa(AnnoA) + "-" + mes + "-" + dia
		fechaVencimiento, _ = time.Parse(layout, fvenc)
	}

	TIM.Serial = TIM.GenerarSerial()
	TIM.Serial = TIM.Usuario + TIM.Serial
	TIM.Nombre = f.Persona.DatoBasico.NombrePrimero
	TIM.Apellido = f.Persona.DatoBasico.ApellidoPrimero
	TIM.FechaCreacion = fechaActual
	TIM.FechaVencimiento = fechaVencimiento
	// TIM.CodigoComponente = m.Componente.Abreviatura
	// TIM.Grado.Abreviatura = m.Grado.Abreviatura
	TIM.Responsable = f.Persona.DatoBasico.Cedula
	TIM.Tipo = 1
	TIM.Estatus = 0
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	carnet["familiar.$.estatuscarnet"] = 1
	err := c.Update(bson.M{"familiar.persona.datobasico.cedula": f.Persona.DatoBasico.Cedula, "id": f.DocumentoPadre}, bson.M{"$set": carnet})
	if err != nil {
		fmt.Println("Err. Creando Estatus de Carnet para hijos")
	}
	return

}

// Estadisticas Actualizando
func (f *Familiar) Estadisticas() {
	var militar []Militar
	c := sys.MGOSession.DB(sys.CBASE).C(sys.CMILITAR)
	buscar := bson.M{"familiar": bson.M{"$elemMatch": bson.M{"esmilitar": false}}}
	seleccion := bson.M{
		"id":                          true,
		"situacion":                   true,
		"familiar.beneficio":          true,
		"familiar.persona.datobasico": true}

	err := c.Find(buscar).Select(seleccion).All(&militar)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	i := 0
	for _, v := range militar {
		// var cant int
		i++
		for _, val := range v.Familiar {

			// if val.Benficio == true {
			// 	cant++
			// }
			var fam FamiliarEstadistica
			fam.ID = v.ID
			fam.Situacion = v.Situacion
			fam.Benficio = val.Benficio
			fam.Parentesco = val.Parentesco
			fam.IDF = val.Persona.DatoBasico.Cedula

			col := sys.MGOSession.DB(sys.CBASE).C(sys.CFAMILIAR)
			err = col.Insert(fam)
		}

	}
}
