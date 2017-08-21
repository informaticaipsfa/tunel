package estadistica

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb/cis/tramitacion"
	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
)

func HistorialReembolso() (jSon []byte, err error) {
	var msj Mensaje
	sq, err := sys.PostgreSQLSAMAN.Query(HistoriaReembolsos())
	if err != nil {
		msj.Mensaje = "Err: " + err.Error()
		msj.Tipo = 1
		jSon, err = json.Marshal(msj)
	}
	i := 0
	c := sys.MGOSession.DB("sssifanb").C("militar")

	for sq.Next() {
		var reembolso tramitacion.Reembolso
		var ced, nro, npmil, nppag, tipocod, pagomonto, montoaprobado,
			instfinannombre, cuenta, concnombre, canal, codnip string
		var componente, grado, clase, situac string
		var tipocuenta, nombre, nombrea, parentesco sql.NullString
		var solicitud, aprobacion sql.NullString
		var Concepto tramitacion.Concepto

		sq.Scan(&ced, &nro, &npmil, &nppag, &tipocod,
			&solicitud, &aprobacion, &pagomonto, &montoaprobado,
			&instfinannombre, &cuenta, &tipocuenta, &concnombre, &canal,
			&componente, &grado, &clase,
			&situac, &nombre, &codnip, &nombrea, &parentesco)
		i++
		layOut := "2006-01-02"
		fechasolicitud := util.ValidarNullString(solicitud)
		if fechasolicitud != VNULL {
			dateString := strings.Replace(fechasolicitud, "/", "-", -1)
			dateStamp, er := time.Parse(layOut, dateString)
			if er == nil {
				reembolso.FechaCreacion = dateStamp
			}
		}
		estatus := 0
		fechaaprobado := util.ValidarNullString(aprobacion)
		if fechaaprobado != VNULL {
			dateString := strings.Replace(fechaaprobado, "/", "-", -1)
			dateStamp, er := time.Parse(layOut, dateString)
			if er == nil {
				reembolso.FechaAprobado = dateStamp
				estatus = 4
			}
		}

		reembolso.Numero = nro
		reembolso.CuentaBancaria.Cedula = ced
		reembolso.CuentaBancaria.Tipo = util.ValidarNullString(tipocuenta)
		reembolso.CuentaBancaria.Institucion = instfinannombre
		reembolso.CuentaBancaria.Cuenta = cuenta
		paren := util.ValidarNullString(parentesco)
		if paren == "null" {
			paren = "MILITAR"
		}
		afiliado := ced + "|" + util.ValidarNullString(nombre) + "(" + paren + ")"
		if ced == codnip {
			afiliado = codnip + "|" + util.ValidarNullString(nombrea) + "(" + paren + ")"
		}
		Concepto.Afiliado = afiliado
		Concepto.Descripcion = concnombre
		reembolso.Requisitos = []int{1, 2, 3}
		reembolso.Concepto = append(reembolso.Concepto, Concepto)
		reembolso.Clase = clase
		reembolso.Componente = componente
		reembolso.Grado = grado
		reembolso.Situacion = situac
		reembolso.Estatus = estatus
		reemb := make(map[string]interface{})
		reemb["cis.serviciomedico.programa.reembolso"] = reembolso
		e := c.Update(bson.M{"id": ced}, bson.M{"$push": reemb})
		if e != nil {
			fmt.Println("Erro: cedula: ", ced)
			return
		}
		fmt.Println(i, "Cedula: ", ced, ": ", reembolso)
	}

	return
}
