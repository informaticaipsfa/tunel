package web

//Copyright Carlos Peña
//Modulo de negociación WEB
import (
	"fmt"
	"net/http"

	"github.com/gesaodin/tunel-ipsfa/sys/web/api"
	"github.com/gorilla/mux"
)

//Variables de Control
var (
	Enrutador   = mux.NewRouter()
	WsEnrutador = mux.NewRouter()
)

func Cargar() {
	CargarModulosWeb()
	CargarModulosSeguridad()
}

//CargarModulosWeb Cargador de modulos web
func CargarModulosWeb() {
	var api api.Persona
	Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/ipsfa/api/militar/crud/{id}", api.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", api.Insertar).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", api.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", api.Eliminar).Methods("DELETE")

}

func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
}

//Principal Página inicial del sistema o bienvenida
func Principal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos bienvenidos al Bus Empresarial de Datos")
}
