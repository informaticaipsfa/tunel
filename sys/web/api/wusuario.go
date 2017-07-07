package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gesaodin/tunel-ipsfa/mdl/usuario"
	"github.com/gesaodin/tunel-ipsfa/sys/seguridad"
	"github.com/gesaodin/tunel-ipsfa/util"
	"github.com/gorilla/mux"
)

type WUsuario struct {
}

//Consultar
func (u *WUsuario) Consultar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var usr usuario.Usuario
	var variable = mux.Vars(r)
	usr.Cedula = variable["id"]
	j, _ := usr.Consultar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//Login conexion para solicitud de token
func (u *WUsuario) Login(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.SUsuario

	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&usuario)
	util.Error(e)

	if usuario.Nombre == "carlos" && usuario.Clave == "za63qj2p" {
		usuario.Nombre = "Carlos"
		usuario.Clave = ""
		usuario.ID = 0
		token := seguridad.GenerarJWT(usuario)
		result := seguridad.RespuestaToken{Token: token}
		j, e := json.Marshal(result)
		util.Error(e)

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		w.Header().Set("Content-Type", "application/text")
		fmt.Println("Error en la conexion del usuario")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Usuario y clave no validas")
	}
}

//ValidarToken Validacion de usuario
func (u *WUsuario) ValidarToken(fn http.HandlerFunc) http.HandlerFunc {
	var mensaje util.Mensajes

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Conexion establecida desde: ", r.Header.Get("Origin"))

		token, e := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &seguridad.Reclamaciones{}, func(token *jwt.Token) (interface{}, error) {

			return seguridad.LlavePublica, nil
		})

		if e != nil {
			switch e.(type) {
			case *jwt.ValidationError:
				vErr := e.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:

					Cabecera(w, r)
					w.WriteHeader(http.StatusUnauthorized)
					mensaje.Tipo = 2
					mensaje.Msj = "El token ha expirado"
					j, _ := json.Marshal(mensaje)
					w.Write(j)
					return
				case jwt.ValidationErrorSignatureInvalid:
					Cabecera(w, r)
					w.WriteHeader(http.StatusForbidden)
					mensaje.Tipo = 3
					mensaje.Msj = "La firma del token no coincide"
					j, _ := json.Marshal(mensaje)
					w.Write(j)
					return
				default:
					Cabecera(w, r)
					w.WriteHeader(http.StatusForbidden)
					mensaje.Tipo = 4
					mensaje.Msj = "Acceso denegado"
					j, _ := json.Marshal(mensaje)
					w.Write(j)
					return
				}
			default:
				fmt.Fprintln(w, "El token no es valido")
				return
			}
		}

		if token.Valid {
			// session.Values["ok"] = true
			// session.Values["name"] = ""
			// session.Save(r, w)

			fn(w, r)
		} else {
			CabeceraRechazada(w, http.StatusForbidden, "El token no es valido")
			return
		}
	})
}

func (u *WUsuario) Autorizado(w http.ResponseWriter, r *http.Request) {
	var mensaje util.Mensajes
	Cabecera(w, r)
	w.WriteHeader(http.StatusOK)
	mensaje.Tipo = 1
	mensaje.Msj = "Acceso Autorizado"
	j, _ := json.Marshal(mensaje)
	w.Write(j)
}
