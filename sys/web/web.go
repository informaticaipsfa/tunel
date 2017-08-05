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
	var wUsuario api.WUsuario
	var per api.Militar
	var comp api.APIComponente
	var esta api.APIEstado
	var wrec api.WRecibo
	var wcar api.WCarnet
	var wfam api.WFamiliar

	Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/ipsfa/api/militar/crud/{id}", per.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", per.Eliminar).Methods("DELETE")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/componente/{id}", comp.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/estado", esta.Consultar).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/familiar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Insertar)).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/recibo/crud/{id}", wUsuario.ValidarToken(wrec.Consultar)).Methods("GET")
	//Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wrec.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wUsuario.ValidarToken(wrec.Insertar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/carnet/listar/{id}", wcar.Listar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/carnet/apro/{estatus}/{serial}", wcar.Aprobar).Methods("GET")
}

func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	// Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crear", wUsuario.Crear).Methods("POST")

}

//Principal Página inicial del sistema o bienvenida
func Principal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos bienvenidos al Bus Empresarial de Datos")
}
