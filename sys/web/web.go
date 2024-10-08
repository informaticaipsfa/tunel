package web

//Copyright Carlos Peña
//Controlador del MiddleWare
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/informaticaipsfa/tunel/sys/web/api"
)

// Variables de Control
var (
	Enrutador = mux.NewRouter()
	wUsuario  api.WUsuario
)

// Cargar los diferentes modulos del sistema
func Cargar() {
	CargarMiddleWare()
	CargarModulosNomina()
	CargarPensionados()
	CargarModulosBanco()
	CargarModulosCredito()

	CargarModulosWebSite()
	CargarModulosWebDevel()
	CargarModulosPanel()
	CargarModulosSeguridad()
	CargarBienestarSocial()
	WMAdminLTE()
	Principal()
}

// CargarModulosSeguridad Y cifrado
func CargarModulosSeguridad() {
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/app/api/wusuario/login", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Crear)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/validarphp", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/consultar/{id}/{col}", wUsuario.ValidarToken(wUsuario.Consultar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/wusuario/restablecer", wUsuario.ValidarToken(wUsuario.RestablecerClaves)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wusuario/restablecer", wUsuario.Opciones).Methods("OPTIONS")
}

// CargarMiddleWare Cargador de modulos web
func CargarMiddleWare() {
	var per api.Militar
	var comp api.APIComponente
	var esta api.APIEstado
	var wrec api.WRecibo
	var wcar api.WCarnet
	var wfam api.WFamiliar
	var wcsv api.CSV
	Enrutador.HandleFunc("/ipsfa/api/militar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.ValidarToken(per.Eliminar)).Methods("DELETE")
	Enrutador.HandleFunc("/ipsfa/api/militar/crud", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/militar/listado", wUsuario.ValidarToken(per.Listado)).Methods("POST")
	//
	Enrutador.HandleFunc("/ipsfa/api/militar/xcrud", wUsuario.ValidarToken(per.Actualizar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/xcrud", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/militar/pace/{id}", wUsuario.ValidarToken(per.ConsultarPACE)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/militar/pace/consultarbeneficiario/{id}", wUsuario.ValidarToken(per.ConsultarBeneficiario)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/militar/reportecomponente", wUsuario.ValidarToken(per.EstadisticasPorComponente)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/reportegrado", wUsuario.ValidarToken(per.EstadisticasPorGrado)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/reportefamiliar", wUsuario.ValidarToken(per.EstadisticasFamiliar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivos", wUsuario.ValidarToken(per.SubirArchivos)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivos", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivosx/{id}", wUsuario.ValidarToken(per.SubirArchivosConversion)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivosx/{id}", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivostxt", wUsuario.ValidarToken(per.SubirArchivosTXTPensiones)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivostxt", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivoscob", wUsuario.ValidarToken(per.SubirArchivosTXTCobranza)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirarchivoscob", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirtxtsisa", wUsuario.ValidarToken(per.SubirArchivosSISA)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/jwtsubirtxtsisa", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/componente/{id}", wUsuario.ValidarToken(comp.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/estado", wUsuario.ValidarToken(esta.Consultar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/familiar/crud/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.ValidarToken(wfam.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/familiar/crud", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/familiar/csvfamiliar", wUsuario.ValidarToken(wcsv.GCSVSC)).Methods("POST")

	//
	Enrutador.HandleFunc("/ipsfa/api/familiar/xcrud", wUsuario.ValidarToken(wfam.Actualizar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/familiar/xcrud", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/recibo/crud/{id}", wUsuario.ValidarToken(wrec.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wUsuario.ValidarToken(wrec.Insertar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/recibo/crud", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/carnet/listar/{id}", wUsuario.ValidarToken(wcar.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/carnet/limpiar/{estatus}/{sucursal}", wUsuario.ValidarToken(wcar.Limpiar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/carnet/apro/{estatus}/{serial}", wUsuario.ValidarToken(wcar.Aprobar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/carnet/liberar", wUsuario.ValidarToken(wcar.Liberar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/carnet/liberar", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/consultar/{id}", wUsuario.ValidarToken(per.ConsultarCedula)).Methods("GET")
}

// CargarBienestarSocial Modulos de Bienestar Social para Reembolsos Carta y apoyos
func CargarBienestarSocial() {
	var wCis api.WCis
	var wCisA api.WCisApoyo
	var wCisC api.WCisCarta
	var wfe api.WFedeVida
	var wtp api.WTratamiento
	var wfactura api.WFactura
	var wmedicina api.WMedicina

	Enrutador.HandleFunc("/ipsfa/api/wreembolso/listar/{id}/{sucursal}", wUsuario.ValidarToken(wCis.ListarReembolso)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Registrar)).Methods("POST")
	//Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wreembolsox", wUsuario.ValidarToken(wCis.Actualizar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso", wUsuario.ValidarToken(wCis.Opciones)).Methods("OPTIONS")
	//Enrutador.HandleFunc("/ipsfa/api/wreembolso/estatus", wUsuario.ValidarToken(wCis.Estatus)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso/estatusx", wUsuario.ValidarToken(wCis.Estatus)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wreembolso/estatus", wUsuario.ValidarToken(wCis.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/wreembolsoreporte", wUsuario.ValidarToken(wCis.ListarReporteFinanzas)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wapoyo/listar/{id}/{sucursal}", wUsuario.ValidarToken(wCisA.ListarApoyo)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Registrar)).Methods("POST")
	//	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wapoyox", wUsuario.ValidarToken(wCisA.Actualizar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")
	//Enrutador.HandleFunc("/ipsfa/api/wapoyo/estatus", wUsuario.ValidarToken(wCisA.Estatus)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo/estatusx", wUsuario.ValidarToken(wCisA.Estatus)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wapoyo/estatus", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wcarta/listar/{id}", wUsuario.ValidarToken(wCisC.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wcarta", wUsuario.ValidarToken(wCisC.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wcarta", wUsuario.ValidarToken(wCisA.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wtratamiento", wUsuario.ValidarToken(wtp.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wfedevida", wUsuario.ValidarToken(wfe.Registrar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wfactura", wUsuario.ValidarToken(wfactura.Consultar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wmedicina", wUsuario.ValidarToken(wmedicina.Registrar)).Methods("POST")
}

// CargarModulosNomina Nomina del personal Militar
func CargarModulosNomina() {

	var concepto api.WNomina
	var wNomina api.WNomina
	var medida api.WMedidaJudicial
	var descuentos api.WDescuentos
	var M api.Militar
	var wR api.WRechazos
	var wRetroactivo api.Retroactivo

	Enrutador.HandleFunc("/ipsfa/api/nomina/concepto", wUsuario.ValidarToken(concepto.Agregar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/concepto", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/nomina/concepto/{id}", wUsuario.ValidarToken(concepto.Consultar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/listar/concepto/", wUsuario.ValidarToken(concepto.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/conceptos/listar/", wUsuario.ValidarToken(concepto.ListarPHP)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/conceptos/contable/{id}", wUsuario.ValidarToken(concepto.ListarContable)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva", wUsuario.ValidarToken(M.ConsultarDirectiva)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/detalle/{id}", wUsuario.ValidarToken(M.ConsultarDetalleDirectiva)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/listar/{id}", wUsuario.ValidarToken(M.ListarDetalleDirectiva)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/clonar", wUsuario.ValidarToken(M.ClonarDirectiva)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/eliminar/{id}", wUsuario.ValidarToken(M.ConsultarDetalleDirectiva)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/cerrar/{id}/{estatus}", wUsuario.ValidarToken(wNomina.Gestionar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/procesar", wUsuario.ValidarToken(wNomina.Procesar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/verpartida/{id}", wUsuario.ValidarToken(wNomina.VerPartidas)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/ccpensionados", wUsuario.ValidarToken(M.ConsultarCantidadPensionados)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/listarpendientes/{mes}/{id}/{ano}", wUsuario.ValidarToken(wNomina.ListarPendientes)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/cuadrebanco/{id}/{tbl}", wUsuario.ValidarToken(wNomina.CuadreBanco)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/prima", wUsuario.ValidarToken(M.ActualizarPrima)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/directiva/actualizar", wUsuario.ValidarToken(M.ActualizarDirectiva)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/nomina/generar", wUsuario.ValidarToken(M.GenerarNomina)).Methods("POST")
	Enrutador.HandleFunc("/devel/api/nomina/generar", M.GenerarNomina).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/medidajudicial", wUsuario.ValidarToken(medida.Agregar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/medidajudicial/{id}", wUsuario.ValidarToken(medida.Actualizar)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/lstmedidalistar/{id}", wUsuario.ValidarToken(medida.ListarMedida)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/descuentos", wUsuario.ValidarToken(descuentos.Agregar)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/nomina/listarpagos", wUsuario.ValidarToken(wNomina.ListarPagos)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/listarpagosdetalles", wUsuario.ValidarToken(wNomina.ListarPagosDetalle)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/nomina/verpagosindividual/{llav}/{cedu}", wUsuario.ValidarToken(wNomina.VerPagosIndividual)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/rechazos/agregar", wUsuario.ValidarToken(wR.Agregar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/rechazos/listar/{id}", wUsuario.ValidarToken(wR.Listar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/rechazos/eliminar/{id}", wUsuario.ValidarToken(wR.Eliminar)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/eliminar/{id}", wUsuario.ValidarToken(wNomina.Eliminar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/publicar/{id}", wUsuario.ValidarToken(wNomina.Publicar)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/debaja/{id}", wUsuario.ValidarToken(wNomina.DeBaja)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/nomina/mes/activo", wUsuario.ValidarToken(wRetroactivo.MesActivo)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/nomina/mes/detalle", wUsuario.ValidarToken(wRetroactivo.MesDetalle)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/nomina/gretroctivo", wUsuario.ValidarToken(wRetroactivo.GenerarRetroactivo)).Methods("POST")
	Enrutador.HandleFunc("/devel/api/nomina/gretroactivo", wRetroactivo.GenerarRetroactivo).Methods("POST")
}

// CargarPensionados Pensionados en general
func CargarPensionados() {
	var wPensionado api.Militar

	Enrutador.HandleFunc("/ipsfa/api/pensionado/calculo/{id}", wUsuario.ValidarToken(wPensionado.Calculo)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/pensionado/consultarneto/{id}", wUsuario.ValidarToken(wPensionado.ConsultarNeto)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/consultarsobreviviente/{id}/{fam}", wUsuario.ValidarToken(wPensionado.ConsultarNetoSobreviviente)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/pensionado/derechoacrecer", wUsuario.ValidarToken(wPensionado.AplicarDerechoACrecer)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/derechoacrecer", wUsuario.ValidarToken(wPensionado.AplicarDerechoACrecerUpdate)).Methods("PUT")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/situacionpago", wUsuario.ValidarToken(wPensionado.SituacionPago)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/pensionado/calcularretroactivo", wUsuario.ValidarToken(wPensionado.CalcularRetroactivo)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/pensionado/impimirarc", wUsuario.ValidarToken(wPensionado.ImprimirARC)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/pensionado/impimirarc", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/pensionado/gitall", wUsuario.ValidarToken(wPensionado.GitAll)).Methods("POST")
}

// CargarModulosBanco Modulos de txt y reportes de banco
func CargarModulosBanco() {

	var wNom api.WNomina
	var wR api.WRechazos
	Enrutador.HandleFunc("/ipsfa/api/nomina/metodobanco/{id}/{cant}", wUsuario.ValidarToken(wNom.CrearTxt)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/patria/{id}", wUsuario.ValidarToken(wNom.Patria)).Methods("GET")
	Enrutador.HandleFunc("/dev/api/nomina/patria/{id}", wNom.Patria).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/nomina/metodobancorechazos/{id}/{cant}", wUsuario.ValidarToken(wR.CrearTxt)).Methods("GET")

}

// CargarModulosCredito Cargador de modulos web
func CargarModulosCredito() {

	var wCredito api.WCredito

	Enrutador.HandleFunc("/ipsfa/api/credito/crud", wUsuario.ValidarToken(wCredito.Guardar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/listar", wUsuario.ValidarToken(wCredito.Listar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/listar", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/credito/actualizar", wUsuario.ValidarToken(wCredito.Actualizar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/enviar", wUsuario.ValidarToken(wCredito.EnviarATesoreria)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/liquidar", wUsuario.ValidarToken(wCredito.Liquidar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/pagar", wUsuario.ValidarToken(wCredito.Pagar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/creartxt/{ano}/{mes}", wUsuario.ValidarToken(wCredito.CrearTxt)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/credito/relacionactiva", wUsuario.ValidarToken(wCredito.RelacionActiva)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/relacionpagados", wUsuario.ValidarToken(wCredito.RelacionPagados)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/credito/cuotas/{id}", wUsuario.ValidarToken(wCredito.ListarCuotas)).Methods("GET")
}

// CargarModulosWebDevel Cargador de modulos web
func CargarModulosWebDevel() {

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
	var concepto api.WNomina
	var wfactura api.WFactura
	var wmedicina api.WMedicina
	var wPensionado api.Militar
	var wcsv api.CSV

	Enrutador.HandleFunc("/devel/api/militar/crud/{id}", per.Consultar).Methods("GET")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Actualizar).Methods("PUT")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Insertar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Eliminar).Methods("DELETE")
	Enrutador.HandleFunc("/devel/api/militar/crud", per.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/militar/xcrud", per.Actualizar).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/xcrud", per.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/devel/api/militar/reportecomponente", per.EstadisticasPorComponente).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/reportegrado", per.EstadisticasPorGrado).Methods("POST")
	Enrutador.HandleFunc("/devel/api/militar/reportefamiliar", per.EstadisticasFamiliar).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/listado", per.Listado).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/militar/subirarchivos", per.SubirArchivos).Methods("POST")

	Enrutador.HandleFunc("/devel/api/militar/pace/{id}", per.ConsultarPACE).Methods("GET")
	Enrutador.HandleFunc("/devel/api/militar/pace/consultarbeneficiario/{id}", per.ConsultarBeneficiario).Methods("GET")

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
	Enrutador.HandleFunc("/devel/api/pensionado/consultarneto/{id}", wPensionado.ConsultarNeto).Methods("GET")
	Enrutador.HandleFunc("/devel/api/pensionado/consultarsobreviviente/{id}/{fam}", wPensionado.ConsultarNetoSobreviviente).Methods("GET")
	Enrutador.HandleFunc("/devel/api/nomina/conceptos/listar/", concepto.ListarPHP).Methods("GET")
	Enrutador.HandleFunc("/devel/api/nomina/conceptos/contable/{id}", concepto.ListarContable).Methods("GET")

	Enrutador.HandleFunc("/devel/api/pensionado/calculo/{id}", wPensionado.Calculo).Methods("GET")

	Enrutador.HandleFunc("/devel/api/consultar/{id}", per.ConsultarCedula).Methods("GET")

}

// CargarModulosWebSite Cargador de modulos web
func CargarModulosWebSite() {
	var wU api.WebUsuario
	var per api.Militar
	var concepto api.WNomina
	var wPensionado api.Militar

	Enrutador.HandleFunc("/ipsfa/api/web/loginWsx", wUsuario.LoginW).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/web/loginWsx", wUsuario.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipsfa/api/web/cambiarclave", wUsuario.CambiarClave).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/web/cambiarclave", wUsuario.Opciones).Methods("OPTIONS")

	//Identificación de Usuario
	Enrutador.HandleFunc("/ipsfa/api/web/identificacion", wU.Identificacion).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/web/identificacion", wU.Opciones).Methods("OPTIONS")
	//Datos Militares HTTP - DEVELOPER
	Enrutador.HandleFunc("/ipsfa/api/web/militar/{id}", per.Consultar).Methods("GET")
	//Consultar Netos de Pensionados
	Enrutador.HandleFunc("/ipsfa/api/web/nomina/conceptos/listar/", concepto.ListarPHP).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/web/pensionado/consultarneto/{id}", wPensionado.ConsultarNetoWeb).Methods("GET")
	//Consultar Netos de Pensionados Sobrevivientes
	Enrutador.HandleFunc("/ipsfa/api/web/pensionado/consultarsobreviviente/{id}/{fam}", wPensionado.ConsultarNetoSobrevivienteWeb).Methods("GET")
	//Constancia de Pensionado
	Enrutador.HandleFunc("/ipsfa/api/web/pensionado/calculo/{id}", wPensionado.Calculo).Methods("GET")
	//Fideicomiso
	Enrutador.HandleFunc("/ipsfa/api/web/pace/consultarbeneficiario/{id}", per.ConsultarBeneficiario).Methods("GET")
	//Consultar Calculadora
	Enrutador.HandleFunc("/ipsfa/api/web/pensionado/calcularretroactivo", wPensionado.CalcularRetroactivo).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/web/pensionado/calcularretroactivo", wU.Opciones).Methods("OPTIONS")
	//Consultar ARC
	Enrutador.HandleFunc("/ipsfa/api/web/pensionado/imprimirarc", wPensionado.ImprimirARC).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/web/pensionado/imprimirarc", wUsuario.Opciones).Methods("OPTIONS")

	/** 	CONTENIDO HTTPS Y SEGURO **/

	//SEGURIDAD HTTPS PRODUCTION - JWT
	Enrutador.HandleFunc("/ipsfa/loginWsx", wUsuario.LoginW).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/loginWsx", wUsuario.Opciones).Methods("OPTIONS")

	//Identificación de Usuario
	Enrutador.HandleFunc("/ipsfa/identificacion", wU.Identificacion).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/identificacion", wU.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/recuperarclave", wUsuario.RecuperarW).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/recuperarclave", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/cambiarclave", wUsuario.ValidarToken(wUsuario.CambiarClaveW)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/cambiarclave", wUsuario.Opciones).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/restablecerclave", wUsuario.ValidarToken(wUsuario.RestablecerClaves)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/restablecerclave", wUsuario.Opciones).Methods("OPTIONS")

	//Datos Militares
	Enrutador.HandleFunc("/ipsfa/militar/{id}", wUsuario.ValidarToken(per.Consultar)).Methods("GET")
	//Consultar Netos de Pensionados - JWT
	Enrutador.HandleFunc("/ipsfa/nomina/conceptos/listar/", wUsuario.ValidarToken(concepto.ListarPHP)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/pensionado/consultarneto/{id}", wUsuario.ValidarToken(wPensionado.ConsultarNetoWeb)).Methods("GET")
	//Consultar Netos de Pensionados Sobrevivientes - JWT
	Enrutador.HandleFunc("/ipsfa/pensionado/consultarsobreviviente/{id}/{fam}", wUsuario.ValidarToken(wPensionado.ConsultarNetoSobrevivienteWeb)).Methods("GET")
	//Constancia de Pensionado - JWT
	Enrutador.HandleFunc("/ipsfa/pensionado/calculo/{id}", wUsuario.ValidarToken(wPensionado.Calculo)).Methods("GET")
	//Fideicomiso - JWT
	Enrutador.HandleFunc("/ipsfa/pace/consultarbeneficiario/{id}", wUsuario.ValidarToken(per.ConsultarBeneficiario)).Methods("GET")
	//Consultar Calculadora - JWT
	Enrutador.HandleFunc("/ipsfa/pensionado/calcularretroactivo", wUsuario.ValidarToken(wPensionado.CalcularRetroactivo)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/pensionado/calcularretroactivo", wU.Opciones).Methods("OPTIONS")
	//Consultar ARC - JWT
	Enrutador.HandleFunc("/ipsfa/pensionado/imprimirarc", wUsuario.ValidarToken(wPensionado.ImprimirARC)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/pensionado/imprimirarc", wUsuario.Opciones).Methods("OPTIONS")

}

// CargarModulosPanel Panel de Contencion
func CargarModulosPanel() {

	var wpanel api.WPanel
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/vreduccion", wUsuario.ValidarToken(wpanel.ValidarReduccion)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/exreduccion", wUsuario.ValidarToken(wpanel.ExtraerReduccion)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/crearreduccion", wUsuario.ValidarToken(wpanel.CrearReduccion)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/listarcolecciones", wUsuario.ValidarToken(wpanel.ListarColecciones)).Methods("GET")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/listarpendientes", wUsuario.ValidarToken(wpanel.ListarPendientes)).Methods("GET")

	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/extraerdatosmysql", wUsuario.ValidarToken(wpanel.ExtraerDatosMySQL)).Methods("POST")

	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/gitall", wUsuario.ValidarToken(wpanel.GitAll)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/gitall", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/compilar", wUsuario.ValidarToken(wpanel.Compilar)).Methods("POST")
	Enrutador.HandleFunc("/ipsfa/api/wpanel/data/compilar", wUsuario.Opciones).Methods("OPTIONS")

}

// WMAdminLTE OpenSource tema de panel de control Tecnología Bootstrap3
func WMAdminLTE() {
	fmt.Println("Cargando Modulos de AdminLTE...")

	prefix := http.StripPrefix("/sssifanb", http.FileServer(http.Dir("public_web/SSSIFANB")))
	Enrutador.PathPrefix("/sssifanb/").Handler(prefix)
}

// Principal Página inicial del sistema o bienvenida
func Principal() {
	prefix := http.StripPrefix("/", http.FileServer(http.Dir("public_web/SSSIFANB/app.ipsfa/dist")))
	Enrutador.PathPrefix("/").Handler(prefix)
}
