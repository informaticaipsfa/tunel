package web

import "net/http"

func Cabecera(w http.ResponseWriter, origen string) {
	//w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", origen)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

}

func CabeceraRechazada(w http.ResponseWriter, estatus int, m string) {
	w.WriteHeader(estatus)
	msj := []byte(m)
	w.Write(msj)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// var u seguridad.Usuario
	//
	// Cabecera(w, r.Header.Get("Origin"))
	// e := json.NewDecoder(r.Body).Decode(&u)
	// util.Error(e)
	//
	// //fmt.Println("Pasando el Decode", u)
	// if u.Nombre == "carlos" && u.Clave == "za63qj2p" {
	// 	u.Nombre = "Carlos"
	// 	u.Clave = ""
	// 	u.Id = 0
	// 	token := seguridad.GenerarJWT(u)
	// 	result := seguridad.RespuestaToken{Token: token}
	// 	j, e := json.Marshal(result)
	// 	util.Error(e)
	// 	Mensajeria.Usuario["gpanel"].ch <- []byte("Iniciando Sesión")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(j)
	// } else {
	// 	w.Header().Set("Content-Type", "application/text")
	// 	fmt.Println("Error en la conexion del usuario")
	// 	w.WriteHeader(http.StatusForbidden)
	// 	fmt.Fprintln(w, "Usuario y clave no validas")
	// }
}
