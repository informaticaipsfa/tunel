package web

//Copyright Carlos Peña
//Modulo de negociación WEB
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/sys/web/api"
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

	var wCis api.WCis
	var wCisA api.WCisApoyo
	var wCisC api.WCisCarta
	var wfe api.WFedeVida
	var wtp api.WTratamiento

	var wfactura api.WFactura
	var wmedicina api.WMedicina

	Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/ipsfa/api/militar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Eliminar)).Methods("DELETE")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/militar/reportecomponente", wUsuario.ValidarToken(per.EstadisticasPorComponente)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/reportegrado", wUsuario.ValidarToken(per.EstadisticasPorGrado)).Methods("POST")

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

	Enrutador.HandleFunc("/ipsfa/api/wreembolso/listar/{id}", wUsuario.ValidarToken(wCis.ListarReembolso)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso/estatus", wUsuario.ValidarToken(wCis.Estatus)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso/estatus", wUsuario.ValidarToken(wCis.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wreembolsoreporte", wUsuario.ValidarToken(wCis.ListarReporteFinanzas)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wapoyo/listar/{id}", wUsuario.ValidarToken(wCisA.ListarApoyo)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo/estatus", wUsuario.ValidarToken(wCisA.Estatus)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo/estatus", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wcarta/listar/{id}", wUsuario.ValidarToken(wCisC.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wcarta", wUsuario.ValidarToken(wCisC.Registrar)).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/wcarta", wCisA.Actualizar).Methods("PUT")
	// Enrutador.HandleFunc("/ipsfa/api/wcarta/estatus", wCisA.Estatus).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wcarta", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wtratamiento", wUsuario.ValidarToken(wtp.Registrar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wfedevida", wUsuario.ValidarToken(wfe.Registrar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wfactura", wUsuario.ValidarToken(wfactura.Consultar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wmedicina", wUsuario.ValidarToken(wmedicina.Registrar)).Methods("POST")
}

func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	// Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/wusuario", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wusuario", wUsuario.CambiarClave).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wusuario", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/devel/api/wusuario/listar", wUsuario.Listar).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/wusuario/validarphp", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("GET")
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
	var wCisA api.WCisApoyo
	var wCisC api.WCisCarta
	var wfe api.WFedeVida
	var per api.Militar
	var comp api.APIComponente
	var esta api.APIEstado
	var wrec api.WRecibo
	var wcar api.WCarnet
	var wfam api.WFamiliar
	var wtp api.WTratamiento

	var wfactura api.WFactura
	var wmedicina api.WMedicina

	Enrutador.HandleFunc("/devel/api/militar/crud/{id}", per.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Insertar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Eliminar).Methods("DELETE")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/devel/api/militar/reportecomponente", per.EstadisticasPorComponente).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/reportegrado", per.EstadisticasPorGrado).Methods("POST")

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

	Enrutador.HandleFunc("/devel/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/wusuario/listar", wUsuario.Listar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/wusuario/crud", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wusuario/crud", wUsuario.CambiarClave).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wusuario/crud", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/wreembolso/listar/{id}", wCis.ListarReembolso).Methods("GET")
	Enrutador.HandleFunc("/devel/api/wreembolso", wCis.Registrar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wreembolso", wCis.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wreembolso", wCis.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/devel/api/wreembolso/estatus", wCis.Estatus).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wreembolso/estatus", wCis.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/devel/api/wreembolsoreporte", wCis.ListarReporteFinanzas).Methods("POST")

	Enrutador.HandleFunc("/devel/api/wapoyo/listar/{id}", wCisA.ListarApoyo).Methods("GET")
	Enrutador.HandleFunc("/devel/api/wapoyo", wCisA.Registrar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wapoyo", wCisA.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wapoyo", wCisA.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/devel/api/wapoyo/estatus", wCisA.Estatus).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wapoyo/estatus", wCisA.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/wcarta/listar/{id}", wCisC.Listar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/wcarta", wCisC.Registrar).Methods("POST")
	// Enrutador.HandleFunc("/devel/api/wcarta", wCisA.Actualizar).Methods("PUT")
	// Enrutador.HandleFunc("/devel/api/wcarta/estatus", wCisA.Estatus).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/wcarta", wCisA.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/wtratamiento", wtp.Registrar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wfedevida", wfe.Registrar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wfactura", wfactura.Consultar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/wmedicina", wmedicina.Registrar).Methods("POST")
}
