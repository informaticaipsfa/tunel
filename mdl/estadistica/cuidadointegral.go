package estadistica

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

<<<<<<< HEAD
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/cis/tramitacion"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/util"
=======
	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/tramitacion"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
>>>>>>> ea581ffe0c74c05e26fc1e8f862f22c48b479406
)

func HistorialReembolso() (jSon []byte, err error) {
	var msj Mensaje
	var cedula, codigo, fechasolicitud, fechaaprobado, paren, afiliado, titular string
	var fcreado, faprobado time.Time
	var estatus int
	sq, err := sys.PostgreSQLSAMAN.Query(HistoriaReembolsos())
	if err != nil {
		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	c := sys.MGOSession.DB("sssifanb").C("militar")
	estatus = 0

	var LstConcepto []tramitacion.Concepto
	var reembolso tramitacion.Reembolso
	for sq.Next() {
		var montoaprobado, pagomonto sql.NullFloat64
		var ced, nro, npmil, nppag, tipocod,
			instfinannombre, cuenta, concnombre, canal, codnip string
		var componente, grado, clase, situac string
		var tipocuenta, nombre, nombrea, parentesco, beneficiado, cedulabenef sql.NullString
		var solicitud, aprobacion sql.NullString
		var Concepto tramitacion.Concepto

		sq.Scan(&ced, &nro, &npmil, &nppag, &tipocod,
			&solicitud, &aprobacion, &pagomonto, &montoaprobado,
			&instfinannombre, &cuenta, &tipocuenta, &concnombre, &canal,
			&componente, &grado, &clase,
			&situac, &nombre, &codnip, &nombrea, &parentesco, &beneficiado, &cedulabenef)

		if i == 0 {
			cedula = ced
			codigo = nro
			titular = util.ValidarNullString(nombre)
			estatus = 0

			layOut := "2006-01-02"
			fechasolicitud = util.ValidarNullString(solicitud)
			if fechasolicitud != VNULL {
				dateString := strings.Replace(fechasolicitud, "/", "-", -1)
				dateStamp, er := time.Parse(layOut, dateString)
				if er == nil {
					fcreado = dateStamp
				}
			}

			fechaaprobado = util.ValidarNullString(aprobacion)
			if fechaaprobado != VNULL {
				dateString := strings.Replace(fechaaprobado, "/", "-", -1)
				dateStamp, er := time.Parse(layOut, dateString)
				if er == nil {
					faprobado = dateStamp
					estatus = 99
				}
			}
			reembolso.Numero = codigo
			reembolso.FechaCreacion = fcreado
			reembolso.FechaAprobado = faprobado
			reembolso.MontoSolicitado = util.ValidarNullFloat64(pagomonto)
			reembolso.MontoAprobado = util.ValidarNullFloat64(montoaprobado)
			reembolso.CuentaBancaria.Cedula = cedula
			reembolso.CuentaBancaria.Tipo = util.ValidarNullString(tipocuenta)
			reembolso.CuentaBancaria.Institucion = instfinannombre
			reembolso.CuentaBancaria.Cuenta = cuenta

			reembolso.Requisitos = []int{1, 2, 3}
			reembolso.Concepto = LstConcepto
			reembolso.Clase = clase
			reembolso.Componente = componente
			reembolso.Grado = grado
			reembolso.Situacion = situac
			reembolso.Estatus = estatus

		}

		if cedula != ced || codigo != nro {
			reembolso.Concepto = LstConcepto
			reemb := make(map[string]interface{})
			reemb["cis.serviciomedico.programa.reembolso"] = reembolso
			e := c.Update(bson.M{"id": cedula}, bson.M{"$push": reemb})
			if e != nil {
				fmt.Println("Error: cedula: ", cedula)
				// return
			}

			if estatus == 0 {
				//Listado de Reportes
				var creembolso tramitacion.ColeccionReembolso
				creembolso.ID = cedula
				creembolso.Numero = reembolso.Numero
				creembolso.Nombre = titular
				creembolso.Usuario = "sssifanb"
				creembolso.Estatus = 0
				creembolso.Reembolso = reembolso
				creembolso.FechaCreacion = reembolso.FechaCreacion
				creembolso.MontoSolicitado = reembolso.MontoSolicitado
				creembolso.FechaAprobado = reembolso.FechaAprobado
				creembolso.MontoAprobado = reembolso.MontoAprobado

				coleccion := sys.MGOSession.DB("sssifanb").C("reembolso")
				err = coleccion.Insert(creembolso)
				if err != nil {
					fmt.Println("Error: cedula: ", ced)
					// return
				}
			}

			fmt.Println("Insertando ", reembolso.Numero)

			estatus = 0
			LstConcepto = nil
			var copiarReembolso tramitacion.Reembolso
			reembolso = copiarReembolso
			cedula = ced
			codigo = nro
			titular = util.ValidarNullString(nombre)

			layOut := "2006-01-02"
			fechasolicitud = util.ValidarNullString(solicitud)
			if fechasolicitud != VNULL {
				dateString := strings.Replace(fechasolicitud, "/", "-", -1)
				dateStamp, er := time.Parse(layOut, dateString)
				if er == nil {
					reembolso.FechaCreacion = dateStamp
				}
			}

			fechaaprobado = util.ValidarNullString(aprobacion)
			if fechaaprobado != VNULL {
				dateString := strings.Replace(fechaaprobado, "/", "-", -1)
				dateStamp, er := time.Parse(layOut, dateString)
				if er == nil {
					reembolso.FechaAprobado = dateStamp
					estatus = 99
				}
			}
			reembolso.Numero = codigo
			reembolso.MontoSolicitado = util.ValidarNullFloat64(pagomonto)
			reembolso.MontoAprobado = util.ValidarNullFloat64(montoaprobado)
			reembolso.CuentaBancaria.Cedula = ced
			reembolso.CuentaBancaria.Titular = util.ValidarNullString(nombrea)
			reembolso.CuentaBancaria.Tipo = util.ValidarNullString(tipocuenta)
			reembolso.CuentaBancaria.Institucion = instfinannombre
			reembolso.CuentaBancaria.Cuenta = cuenta

			reembolso.Requisitos = []int{1, 2, 3}

			reembolso.Clase = clase
			reembolso.Componente = componente
			reembolso.Grado = grado
			reembolso.Situacion = situac
			reembolso.Estatus = estatus

		}

		paren = util.ValidarNullString(parentesco)
		if paren == "null" {
			paren = "MILITAR"
		}
		afiliado = ced + "-" + util.ValidarNullString(nombre) + "(" + paren + ")"
		xced := util.ValidarNullString(cedulabenef)
		if ced != xced {
			afiliado = xced + "-" + util.ValidarNullString(beneficiado) + "(" + paren + ")"
		}

		Concepto.Afiliado = afiliado
		Concepto.Descripcion = concnombre
		LstConcepto = append(LstConcepto, Concepto)
		i++
		fmt.Println("Pos : ", i)
	} // FIN DEL REPITA

	reembolso.Concepto = LstConcepto
	reemb := make(map[string]interface{})
	reemb["cis.serviciomedico.programa.reembolso"] = reembolso
	e := c.Update(bson.M{"id": cedula}, bson.M{"$push": reemb})
	if e != nil {
		fmt.Println("Error: cedula: ", cedula)
		// return
	}

	if estatus == 0 {
		//Listado de Reportes
		var creembolso tramitacion.ColeccionReembolso
		creembolso.ID = cedula
		creembolso.Numero = codigo
		creembolso.Nombre = titular
		creembolso.Usuario = "sssifanb"
		creembolso.Estatus = 0
		creembolso.Reembolso = reembolso
		creembolso.FechaCreacion = reembolso.FechaCreacion
		creembolso.MontoSolicitado = reembolso.MontoSolicitado
		creembolso.FechaAprobado = reembolso.FechaAprobado
		creembolso.MontoAprobado = reembolso.MontoAprobado

		coleccion := sys.MGOSession.DB("sssifanb").C("reembolso")
		err = coleccion.Insert(creembolso)
		if err != nil {
			fmt.Println("Error: cedula: ", cedula)
			// return
		}
	}

	return
}
