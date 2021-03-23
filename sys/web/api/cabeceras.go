package api

import "net/http"

//Cabecera Conexion de archivos locales
func Cabecera(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, TRACE, HEAD, PATCH, CONNECT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

//CabeceraW Conexion de archivos Para app
func CabeceraW(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, TRACE, HEAD, PATCH, CONNECT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

//CabeceraRechazada Archivos de rechazos
func CabeceraRechazada(w http.ResponseWriter, estatus int, m string) {
	w.WriteHeader(estatus)
	msj := []byte(m)
	w.Write(msj)
}
