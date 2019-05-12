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

//Cargar los diferentes modulos del sistema
func Cargar() {
	CargarModulosWeb()
	CargarModulosNomina()
	CargarPensionados()
	CargarModulosBanco()
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
	var wpanel api.WPanel
	var wcsv api.CSV

	Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/ipsfa/api/militar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Eliminar)).Methods("DELETE")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/militar/listado", wUsuario.ValidarToken(per.Listado)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/militar/pace/{id}", wUsuario.ValidarToken(per.ConsultarPACE)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/militar/reportecomponente", wUsuario.ValidarToken(per.EstadisticasPorComponente)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/reportegrado", wUsuario.ValidarToken(per.EstadisticasPorGrado)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/reportefamiliar", wUsuario.ValidarToken(per.EstadisticasFamiliar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivos", wUsuario.ValidarToken(per.SubirArchivos)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivostxt", wUsuario.ValidarToken(per.SubirArchivosTXTPensiones)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/componente/{id}", wUsuario.ValidarToken(comp.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/estado", wUsuario.ValidarToken(esta.Consultar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/familiar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/familiar/csvfamiliar", wUsuario.ValidarToken(wcsv.GCSVSC)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/recibo/crud/{id}", wUsuario.ValidarToken(wrec.Consultar)).Methods("GET")
	//Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wrec.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wUsuario.ValidarToken(wrec.Insertar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/carnet/listar/{id}", wUsuario.ValidarToken(wcar.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/carnet/limpiar/{estatus}/{sucursal}", wUsuario.ValidarToken(wcar.Limpiar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/carnet/apro/{estatus}/{serial}", wUsuario.ValidarToken(wcar.Aprobar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/wreembolso/listar/{id}/{sucursal}", wUsuario.ValidarToken(wCis.ListarReembolso)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso/estatus", wUsuario.ValidarToken(wCis.Estatus)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso/estatus", wUsuario.ValidarToken(wCis.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wreembolsoreporte", wUsuario.ValidarToken(wCis.ListarReporteFinanzas)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wapoyo/listar/{id}/{sucursal}", wUsuario.ValidarToken(wCisA.ListarApoyo)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo/estatus", wUsuario.ValidarToken(wCisA.Estatus)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo/estatus", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wcarta/listar/{id}", wUsuario.ValidarToken(wCisC.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wcarta", wUsuario.ValidarToken(wCisC.Registrar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wcarta", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wtratamiento", wUsuario.ValidarToken(wtp.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wfedevida", wUsuario.ValidarToken(wfe.Registrar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wfactura", wUsuario.ValidarToken(wfactura.Consultar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wmedicina", wUsuario.ValidarToken(wmedicina.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/vreduccion", wUsuario.ValidarToken(wpanel.ValidarReduccion)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/exreduccion", wUsuario.ValidarToken(wpanel.ExtraerReduccion)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/crearreduccion", wUsuario.ValidarToken(wpanel.CrearReduccion)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/listarcolecciones", wUsuario.ValidarToken(wpanel.ListarColecciones)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/listarpendientes", wUsuario.ValidarToken(wpanel.ListarPendientes)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/extraerdatosmysql", wUsuario.ValidarToken(wpanel.ExtraerDatosMySQL)).Methods("POST")
}

//CargarModulosNomina Nomina del personal Militar
func CargarModulosNomina() {
	var wUsuario api.WUsuario
	var concepto api.WNomina
	var wNomina api.WNomina
	var medida api.WMedidaJudicial
	var descuentos api.WDescuentos
	var M api.Militar

	Enrutador.HandleFunc("/ipsfa/api/nomina/concepto", wUsuario.ValidarToken(concepto.Agregar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/concepto/{id}", wUsuario.ValidarToken(concepto.Consultar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/listar/concepto/", wUsuario.ValidarToken(concepto.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva", wUsuario.ValidarToken(M.ConsultarDirectiva)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/detalle/{id}", wUsuario.ValidarToken(M.ConsultarDetalleDirectiva)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/listar/{id}", wUsuario.ValidarToken(M.ListarDetalleDirectiva)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/clonar", wUsuario.ValidarToken(M.ClonarDirectiva)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/eliminar/{id}", wUsuario.ValidarToken(M.ConsultarDetalleDirectiva)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/cerrar/{id}/{estatus}", wUsuario.ValidarToken(wNomina.Gestionar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/procesar", wUsuario.ValidarToken(wNomina.Procesar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/verpartida/{id}", wUsuario.ValidarToken(wNomina.VerPartidas)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/ccpensionados", wUsuario.ValidarToken(M.ConsultarCantidadPensionados)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/listarpendientes/{id}", wUsuario.ValidarToken(wNomina.ListarPendientes)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/listarpagos", wUsuario.ValidarToken(wNomina.ListarPagos)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/cuadrebanco/{id}", wUsuario.ValidarToken(wNomina.CuadreBanco)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/prima", wUsuario.ValidarToken(M.ActualizarPrima)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/actualizar", wUsuario.ValidarToken(M.ActualizarDirectiva)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/nomina/generar", wUsuario.ValidarToken(M.GenerarNomina)).Methods("POST")
	Enrutador.HandleFunc("/devel/api/nomina/generar", M.GenerarNomina).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/medidajudicial", wUsuario.ValidarToken(medida.Agregar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/medidajudicial/{id}", wUsuario.ValidarToken(medida.Actualizar)).Methods("PUT")

	Enrutador.HandleFunc("/ipsfa/api/descuentos", wUsuario.ValidarToken(descuentos.Agregar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/nomina/listarpagosdetalles", wUsuario.ValidarToken(wNomina.ListarPagosDetalle)).Methods("POST")
}

//CargarPensionados Pensionados en general
func CargarPensionados() {
	var wUsuario api.WUsuario
	var wPensionado api.Militar
	Enrutador.HandleFunc("/ipsfa/api/pensionado/calculo/{id}", wUsuario.ValidarToken(wPensionado.Calculo)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/consultarneto/{id}", wUsuario.ValidarToken(wPensionado.ConsultarNeto)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/consultarsobreviviente/{id}", wUsuario.ValidarToken(wPensionado.ConsultarNetoSobreviviente)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/derechoacrecer", wUsuario.ValidarToken(wPensionado.AplicarDerechoACrecer)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/situacionpago", wUsuario.ValidarToken(wPensionado.SituacionPago)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/calcularretroactivo", wUsuario.ValidarToken(wPensionado.CalcularRetroactivo)).Methods("POST")

}

//CargarModulosBanco Modulos de txt y reportes de banco
func CargarModulosBanco() {
	var wUsuario api.WUsuario
	var wNom api.WNomina
	Enrutador.HandleFunc("/ipsfa/api/nomina/metodobanco/{id}/{cant}", wUsuario.ValidarToken(wNom.CrearTxt)).Methods("GET")

}

//CargarModulosSeguridad Y cifrado
func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	// Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Crear)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

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
	var wcsv api.CSV

	Enrutador.HandleFunc("/devel/api/militar/crud/{id}", per.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Insertar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Eliminar).Methods("DELETE")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/devel/api/militar/reportecomponente", per.EstadisticasPorComponente).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/reportegrado", per.EstadisticasPorGrado).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/reportefamiliar", per.EstadisticasFamiliar).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/listado", per.Listado).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/subirarchivos", per.SubirArchivos).Methods("POST")

	Enrutador.HandleFunc("/devel/api/militar/pace/{id}", per.ConsultarPACE).Methods("GET")

	Enrutador.HandleFunc("/devel/api/componente/{id}", comp.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/estado", esta.Consultar).Methods("GET")

	Enrutador.HandleFunc("/devel/api/familiar/crud/{id}", per.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/familiar/crud", wfam.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/familiar/crud", wfam.Insertar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/familiar/crud", wfam.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/devel/api/familiar/csvfamiliar", wcsv.GCSVSC).Methods("POST")

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
