package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/gorilla/mux"

	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys/seguridad"
	"github.com/informaticaipsfa/tunel/util"
)

type WUsuario struct{}

// Crear Usuario del sistema
func (u *WUsuario) Crear(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.Usuario
	var m util.Mensajes
	Cabecera(w, r)
	ip := strings.Split(r.RemoteAddr, ":")
	usuario.FirmaDigital.DireccionIP = ip[0]
	usuario.FirmaDigital.Tiempo = time.Now()
	usuario.FechaCreacion = time.Now()

	e := json.NewDecoder(r.Body).Decode(&usuario)
	util.Error(e)
	// if ip[0] == "192.168.6.45" {

	e = usuario.Salvar()
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		m.Msj = e.Error()
		m.Tipo = 1
		fmt.Println("Err: ", e.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		m.Msj = "Usuario creado"
		m.Tipo = 0
	}
	// } else {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	m.Msj = "El equipo donde no esta autorizado"
	// 	m.Tipo = 2
	// }

	m.Fecha = time.Now()
	j, _ := json.Marshal(m)
	w.Write(j)
}

// Clave de tipos
type Clave struct {
	ID        string `json:"id"`
	Cedula    string `json:"cedula"`
	Login     string `json:"login"`
	Clave     string `json:"clave"`
	Nueva     string `json:"nueva"`
	Repetir   string `json:"repetir"`
	Correo    string `json:"correo"`
	Coleccion string `json:"coleccion"`
}

// CambiarClave ID
func (u *WUsuario) CambiarClave(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M util.Mensajes
	var usr seguridad.Usuario
	var datos Clave

	e := json.NewDecoder(r.Body).Decode(&datos)
	util.Error(e)
	if datos.Coleccion == "" {
		datos.Coleccion = "usuario"
	}
	ok := usr.CambiarClave(datos.Login, datos.Clave, datos.Nueva, datos.Coleccion)
	M.Tipo = 1
	if ok != nil {
		M.Msj = ok.Error()
		M.Tipo = 0
	}
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Login conexion para solicitud de token
func (u *WUsuario) Login(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.Usuario
	var traza fanb.Traza
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&usuario)
	util.Error(e)
	fmt.Println(usuario)

	usuario.Validar(usuario.Nombre, util.GenerarHash256([]byte(usuario.Clave)))

	if usuario.Nombre != "" {
		//usuario.Nombre = "Carlos"
		usuario.Clave = ""

		token := seguridad.GenerarJWT(usuario)
		result := seguridad.RespuestaToken{Token: token}
		j, e := json.Marshal(result)
		util.Error(e)
		ip := strings.Split(r.RemoteAddr, ":")
		traza.Usuario = usuario.Login
		traza.Time = time.Now()
		traza.Log = "Inicio de sesión"
		traza.IP = ip[0]
		traza.Documento = "Usuario"
		traza.Crear()
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		w.Header().Set("Content-Type", "application/text")
		fmt.Println("Error en la conexion del usuario")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Usuario y clave no validas")
	}
}

// LoginW conexion para solicitud de token
func (u *WUsuario) LoginW(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.WUsuario
	// var traza fanb.Traza
	CabeceraW(w, r)
	e := json.NewDecoder(r.Body).Decode(&usuario)
	fmt.Println(usuario)
	util.Error(e)
	//fmt.Println("login... ", usuario.Cedula, usuario.Nombre, util.GenerarHash256([]byte(usuario.Clave)))
	err := usuario.Validar(usuario.Cedula, util.GenerarHash256([]byte(usuario.Clave)))
	if err != nil {

		w.Header().Set("Content-Type", "application/text")
		fmt.Println("Error en la conexion del usuario")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Usuario y clave no validas")

	} else {

		if usuario.Cedula != "" && usuario.Componente != "" {
			usuario.Clave = ""
			token := seguridad.WGenerarJWT(usuario)
			result := seguridad.RespuestaToken{Token: token}
			j, e := json.Marshal(result)
			util.Error(e)
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	}

}

// ValidarToken Validacion de usuario
func (u *WUsuario) ValidarToken(fn http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entrando... ")
		Cabecera(w, r)

		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return seguridad.LlavePublica, nil
		}, request.WithClaims(&seguridad.Reclamaciones{}))

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid token:", err)
			return
		}
		fmt.Println(token.Claims.(*seguridad.Reclamaciones).Usuario)

		/**
		token, e := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &seguridad.Reclamaciones{}, func(token *jwt.Token) (interface{}, error) {
			//var claims jwt.Claims
			reclamacion := token.Claims.(*seguridad.Reclamaciones)
			UsuarioConectado = reclamacion.Usuario
			//fmt.Println("Se esta conectando: ", reclamacion.Usuario.Nombre, " Desde: ", reclamacion.Usuario.FirmaDigital.DireccionIP)
			return seguridad.LlavePublica, nil
		})

		if e != nil {
			switch e.(type) {
			case *jwt.ValidationError:
				vErr := e.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					w.WriteHeader(http.StatusUnauthorized)
					mensaje.Tipo = 2
					mensaje.Msj = "El token ha expirado"
					j, _ := json.Marshal(mensaje)
					w.Write(j)
					return
				case jwt.ValidationErrorSignatureInvalid:
					w.WriteHeader(http.StatusForbidden)
					mensaje.Tipo = 3
					mensaje.Msj = "La firma del token no coincide"
					j, _ := json.Marshal(mensaje)
					w.Write(j)
					return
				default:
					w.WriteHeader(http.StatusForbidden)
					mensaje.Tipo = 4
					mensaje.Msj = "Acceso denegado"
					j, _ := json.Marshal(mensaje)
					w.Write(j)
					return
				}
			default:
				w.WriteHeader(http.StatusForbidden)
				mensaje.Tipo = 5
				mensaje.Msj = "El token no es valido"
				j, _ := json.Marshal(mensaje)
				w.Write(j)
				return
			}
		}*/

		if token.Valid {
			fn(w, r)
		} else {
			CabeceraRechazada(w, http.StatusForbidden, "Err. token no es valido")
			return
		}
	})
}

// Autorizado Formando archivos
func (u *WUsuario) Autorizado(w http.ResponseWriter, r *http.Request) {
	var mensaje util.Mensajes
	Cabecera(w, r)
	w.WriteHeader(http.StatusOK)
	mensaje.Tipo = 1
	mensaje.Msj = "Acceso Autorizado"
	j, _ := json.Marshal(mensaje)
	w.Write(j)
}

// Consultar conexion para solicitud de token
func (u *WUsuario) Consultar(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.Usuario
	var traza fanb.Traza
	Cabecera(w, r)
	var cedula = mux.Vars(r)
	ced := cedula["id"]
	col := cedula["col"]

	j, _ := usuario.Consultar(ced, col)

	ip := strings.Split(r.RemoteAddr, ":")
	traza.Usuario = usuario.Login
	traza.Time = time.Now()
	traza.Log = "Consultar Usuario"
	traza.IP = ip[0]
	traza.Documento = "Usuario"
	traza.CrearHistoricoConsulta("husuario")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// Listar Usuario del sistema
func (u *WUsuario) Listar(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.Usuario
	Cabecera(w, r)
	j, _ := usuario.Listar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Opciones Militar
func (u *WUsuario) Opciones(w http.ResponseWriter, r *http.Request) {
	CabeceraW(w, r)
	fmt.Println("Conectandose usuario vía Extranet...")
	//fmt.Fprintf(w, "Saludos")

}

// CambiarClaveW Control de Cambio de Clave
func (u *WUsuario) CambiarClaveW(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M util.Mensajes
	var usr seguridad.WUsuario
	var datos Clave

	e := json.NewDecoder(r.Body).Decode(&datos)
	util.Error(e)

	ok := usr.CambiarClave(datos.Correo, datos.Clave, datos.Nueva)
	M.Tipo = 1
	M.Msj = "Felicitaciones"
	if ok != nil {
		M.Msj = ok.Error()
		M.Tipo = 0
	}
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// RestablecerClaves ID
func (u *WUsuario) RestablecerClaves(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M util.Mensajes
	var usr seguridad.Usuario
	var datos Clave

	e := json.NewDecoder(r.Body).Decode(&datos)
	util.Error(e)
	if datos.Coleccion == "" {
		datos.Coleccion = "usuario"
	}
	ok := usr.RestablecerClaves(datos.Cedula, datos.Correo, datos.Clave, "wusuario")
	M.Tipo = 1
	M.Fecha = time.Now()
	M.Msj = "Felicitaciones su clave ha sido actualizada"
	if ok != nil {
		M.Msj = ok.Error()
		M.Tipo = 0
	}
	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// RecuperarW conexion para solicitud de token
func (u *WUsuario) RecuperarW(w http.ResponseWriter, r *http.Request) {
	var usuario seguridad.WUsuario
	CabeceraW(w, r)
	e := json.NewDecoder(r.Body).Decode(&usuario)
	util.Error(e)

	if usuario.Correo != "" {

		err := usuario.Recuperar(usuario.Correo)
		if err != nil {
			w.Header().Set("Content-Type", "application/text")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Correo no validas")

		} else {

			if usuario.Cedula != "" && usuario.Componente != "" {
				usuario.Clave = ""
				token := seguridad.WGenerarJWT(usuario)
				result := seguridad.RespuestaToken{Token: token}
				j, e := json.Marshal(result)
				util.Error(e)
				w.WriteHeader(http.StatusOK)
				w.Write(j)
			}
		}
	} else {
		w.Header().Set("Content-Type", "application/text")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Correo no validas")

	}
}
