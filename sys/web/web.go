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
	WMAdminLTE()
	CargarModulosWebDevel()
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
	Enrutador.HandleFunc("/ipsfa/api/militar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Eliminar)).Methods("DELETE")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/componente/{id}", wUsuario.ValidarToken(comp.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/estado", wUsuario.ValidarToken(esta.Consultar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/familiar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/recibo/crud/{id}", wUsuario.ValidarToken(wrec.Consultar)).Methods("GET")
	//Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wrec.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wUsuario.ValidarToken(wrec.Insertar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/carnet/listar/{id}", wUsuario.ValidarToken(wcar.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/carnet/apro/{estatus}/{serial}", wUsuario.ValidarToken(wcar.Aprobar)).Methods("GET")
}

func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	// Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")

}

//Principal Página inicial del sistema o bienvenida
func Principal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos bienvenidos al Bus Empresarial de Datos")
}

//WMAdminLTE OpenSource tema de panel de control
//Tecnología Bootstrap3
func WMAdminLTE() {
	fmt.Println("Cargando Modulos de AdminLTE...")
	// var GP = GPanel{}
	// Enrutador.HandleFunc("/sssifanb/{id}", GP.IrA)
	prefix := http.StripPrefix("/sssifanb", http.FileServer(http.Dir("public_web/SSSIFANB")))
	Enrutador.PathPrefix("/sssifanb/").Handler(prefix)
	// prefixx := http.StripPrefix("/bdse-admin/public/temp", http.FileServer(http.Dir("public/temp")))
	// Enrutador.PathPrefix("/bdse-admin/public/temp/").Handler(prefixx)
}

//CargarModulosWebDevel Cargador de modulos web
func CargarModulosWebDevel() {
	var wUsuario api.WUsuario
	var wCis api.WCis
	var per api.Militar
	var comp api.APIComponente
	var esta api.APIEstado
	var wrec api.WRecibo
	var wcar api.WCarnet
	var wfam api.WFamiliar

	Enrutador.HandleFunc("/devel/api/militar/crud/{id}", per.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Insertar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Eliminar).Methods("DELETE")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/componente/{id}", comp.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/estado", esta.Consultar).Methods("GET")

	Enrutador.HandleFunc("/devel/api/familiar/crud/{id}", per.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/familiar/crud", wfam.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/familiar/crud", wfam.Insertar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/familiar/crud", wfam.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/recibo/crud/{id}", wrec.Consultar).Methods("GET")
	//Enrutador.HandleFunc("/devel/api/recibo/crud", wrec.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/recibo/crud", wrec.Insertar).Methods("POST")

	Enrutador.HandleFunc("/devel/api/carnet/listar/{id}", wcar.Listar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/carnet/apro/{estatus}/{serial}", wcar.Aprobar).Methods("GET")

	Enrutador.HandleFunc("/devel/api/wusuario", wUsuario.CambiarClave).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wusuario", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/wreembolso/listar/{id}", wCis.ListarReembolso).Methods("GET")
	Enrutador.HandleFunc("/devel/api/wreembolso", wCis.Registrar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wreembolso", wCis.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wreembolso", wCis.Opciones).Methods("OPTIONS")

}
