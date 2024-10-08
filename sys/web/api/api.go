package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb"
	"github.com/informaticaipsfa/tunel/mdl/sssifanb/fanb"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/sys/seguridad"
	"github.com/informaticaipsfa/tunel/util"
)

var UsuarioConectado seguridad.Usuario

// Militar militares
type Militar struct {
	Frase string
	Tipo  int
}

// Componente Control Militar
type Componente struct {
	Componente string
	Grado      string
	Situacion  string
}

// Consultar Militares
func (p *Militar) Consultar(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	Cabecera(w, r)
	var dataJSON sssifanb.Militar
	var cedula = mux.Vars(r)
	dataJSON.Persona.DatoBasico.Cedula = cedula["id"]
	j, e := dataJSON.Consultar()
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = cedula["id"]
	traza.Documento = "Consultando Militar"
	traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Actualizar Datos Generales
func (p *Militar) Actualizar(w http.ResponseWriter, r *http.Request) {

	Cabecera(w, r)
	ip := strings.Split(r.RemoteAddr, ":")
	var dataJSON sssifanb.Militar
	err := json.NewDecoder(r.Body).Decode(&dataJSON)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error ", err.Error())
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}

	j, _ := dataJSON.Actualizar(UsuarioConectado.Login, ip[0])
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Insertar Militar
func (p *Militar) Insertar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje
	var militar sssifanb.Militar

	ip := strings.Split(r.RemoteAddr, ":")

	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	err := json.NewDecoder(r.Body).Decode(&militar)
	M.Tipo = 1
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error al insertar", err.Error())
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	//e := militar.SalvarMGOI("militares", objeto)

	if UsuarioConectado.Login[:3] != "act" {
		e := militar.SalvarMGO()
		if e != nil {
			M.Mensaje = e.Error()
			M.Tipo = 0
		}

		traza.Log = militar.ID
		traza.Documento = "Agregando: " + militar.Grado.Abreviatura + "|" + militar.Situacion +
			"|" + militar.FechaIngresoComponente.String() + "|" + militar.FechaAscenso.String()
		traza.CrearHistoricoConsulta("hmilitar")

	} else {
		M.Mensaje = "Su cuenta no poseé acceso para ingresar nuevos militares"
		M.Tipo = 2
	}
	j, _ := militar.Consultar()
	//j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Eliminar Militar
func (p *Militar) Eliminar(w http.ResponseWriter, r *http.Request) {

}

// EstadisticasPorComponente
func (p *Militar) EstadisticasPorComponente(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var militar sssifanb.Militar
	j, _ := militar.EstadisticasPorComponente()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ComponenteDecode Decodificando
type ComponenteDecode struct {
	Grado string
}

// EstadisticasPorGrado EstadisticasPorGrado
func (p *Militar) EstadisticasPorGrado(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var militar sssifanb.Militar
	var componente ComponenteDecode
	err := json.NewDecoder(r.Body).Decode(&componente)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Estoy en un error ", err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}
	j, _ := militar.EstadisticasPorGrado(componente.Grado)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (p *Militar) EstadisticasFamiliar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	// ip := strings.Split(r.RemoteAddr, ":")
	var militar sssifanb.Militar

	j, _ := militar.EstadisticasFamiliar()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Listado Militares
func (p *Militar) Listado(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	var M sssifanb.Mensaje
	var mil Militar
	var dataJSON sssifanb.Militar

	Cabecera(w, r)
	ip := strings.Split(r.RemoteAddr, ":")

	err := json.NewDecoder(r.Body).Decode(&mil)
	if err != nil {
		fmt.Println(err.Error())
		M.Mensaje = err.Error()
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	}
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = mil.Frase
	traza.Documento = "Consultando Militar"
	traza.CrearHistoricoConsulta("historicoconsultas")
	j, _ := dataJSON.BusquedaFullText(mil.Frase, mil.Tipo)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Opciones Militar
func (p *Militar) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	fmt.Println("OPTIONS...")

}

func (p *Militar) SubirArchivos(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje
	var militarF sssifanb.Militar

	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}

	cedula := r.FormValue("txtFileID")

	m := r.MultipartForm
	files := m.File["archivo"]
	directorio := "./public_web/SSSIFANB/afiliacion/temp/" + cedula + "/"
	if cedula == "" {
		M.Mensaje = "Carga fallida"
		M.Tipo = -1
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	} else if cedula == "DESERTOR" {
		directorio = "./tmp/"

	}

	errr := os.Mkdir(directorio, 0775)
	if errr != nil {
		fmt.Println("la carpeta ya existe.")
	}
	cadena := ""
	for i, _ := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println("No se pudo escribir el archivo verifique los privilegios.")
			return
		}
		// fmt.Println("Subiendo Archivo")
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println("Entrando en un Erro...", err)
			return
		}
		// fmt.Println("Instertando....")
		cadena += files[i].Filename + ";"
		if cedula == "DESERTOR" {
			var desertor sssifanb.Desertor
			go desertor.LeerDesertores(files[i].Filename)
		}

	}

	if UsuarioConectado.Login[:3] != "act" {
		traza.Documento = "Agregando Historial Digital ( " + cedula + " )"
		traza.Log = cadena
		traza.CrearHistoricoConsulta("hmilitar")
		M.Mensaje = "Carga exitosa"
		M.Tipo = 2
		militarF.ActualizarFoto(cedula)
	} else {
		M.Mensaje = "Carga fallida"
		M.Tipo = -1
	}

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// SubirArchivosTXTPensiones Pensionados
func (p *Militar) SubirArchivosTXTPensiones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje

	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}
	m := r.MultipartForm
	files := m.File["input-folder-2"]
	codigo := r.FormValue("txtFileID")
	directorio := "./public_web/SSSIFANB/pensiones/temp/nomina/"
	errr := os.Mkdir(directorio, 0777)
	if errr != nil {
		fmt.Println("El directorio ya existe!")
	}
	cadena := ""
	for i, _ := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println(er.Error())
			return
		}
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println(err)
			return
		}
		cadena += files[i].Filename + ";"
		ProcesarTxt(files[i].Filename, codigo)

	} // Fin de archivos
	M.Mensaje = cadena
	M.Tipo = 2

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ProcesarTxt Proceso de archivo
func ProcesarTxt(doc string, codigo string) {
	var a util.Archivo
	a.Ruta = "./public_web/SSSIFANB/pensiones/temp/nomina/" + doc
	a.LeerCA(sys.PostgreSQLPENSION, codigo, doc)
}

// SubirArchivosTXTCobranza Pensionados
func (p *Militar) SubirArchivosTXTCobranza(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje

	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}
	m := r.MultipartForm
	files := m.File["input-folder-2"]
	codigo := r.FormValue("txtFileID")
	directorio := "./public_web/SSSIFANB/tmp/cobranza/"
	errr := os.Mkdir(directorio, 0777)
	if errr != nil {
		fmt.Println("El directorio ya existe!")
	}
	cadena := ""
	for i, _ := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println(er.Error())
			return
		}
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println(err)
			return
		}
		cadena += files[i].Filename + ";"
		ProcesarCob(files[i].Filename, codigo)

	} // Fin de archivos
	M.Mensaje = cadena
	M.Tipo = 2

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// ProcesarCob Proceso de archivo
func ProcesarCob(doc string, codigo string) {
	var a util.Archivo
	a.Ruta = "./public_web/SSSIFANB/tmp/cobranza/" + doc
	a.LeerCA(sys.PostgreSQLPENSION, codigo, doc)
}

// ConsultarCedula Obtener cedulas de familiar o militar
func (p *Militar) ConsultarCedula(w http.ResponseWriter, r *http.Request) {
	var traza fanb.Traza
	Cabecera(w, r)
	var dataJSON sssifanb.Militar
	var cedula = mux.Vars(r)

	j, e := dataJSON.ConsultarCedula(cedula["id"])
	if e != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error al consultar los datos"))
		return
	}
	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login
	traza.Log = cedula["id"]
	traza.Documento = "C::SistemaTurnos"
	traza.CrearHistoricoConsulta("historicoconsultas")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (p *Militar) SubirArchivosSISA(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje

	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}
	m := r.MultipartForm
	files := m.File["input-folder-2"]
	codigo := r.FormValue("txtFileID")
	directorio := "./public_web/SSSIFANB/pensiones/temp/sisa/"
	errr := os.Mkdir(directorio, 0777)
	if errr != nil {
		fmt.Println("El directorio ya existe!")
	}
	cadena := ""
	for i := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println(er.Error())
			return
		}
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println(err)
			return
		}
		cadena += files[i].Filename
		ProcesarTxtSISA(files[i].Filename, codigo)

	} // Fin de archivos
	M.Mensaje = cadena
	M.Tipo = 2

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ProcesarTxt Proceso de archivo
func ProcesarTxtSISA(doc string, codigonomina string) {
	var a util.Archivo
	valores := strings.Split(codigonomina, "|")
	a.Ruta = "./public_web/SSSIFANB/pensiones/temp/sisa/" + doc
	a.LeerSisa(sys.PostgreSQLPENSION, valores[0], doc, valores[1], valores[2])
}

func (p *Militar) SubirArchivosConversion(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje
	var militarF sssifanb.Militar

	var componente = mux.Vars(r)
	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}

	cedula := r.FormValue("txtFileID")

	m := r.MultipartForm
	files := m.File["archivo"]
	directorio := "./public_web/SSSIFANB/afiliacion/temp/" + cedula + "/"
	if cedula == "" {
		M.Mensaje = "Carga fallida"
		M.Tipo = -1
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return

	}

	errr := os.Mkdir(directorio, 0775)
	if errr != nil {
		fmt.Println("la carpeta ya existe.")
	}
	cadena := ""
	for i, _ := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println("No se pudo escribir el archivo verifique los privilegios.")
			return
		}
		// fmt.Println("Subiendo Archivo")
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println("Entrando en un Erro...", err)
			return
		}

		if files[i].Filename == "ifoto.jpg" {
			go util.ConverPhoto(directorio, componente["id"])
		}
	}

	if UsuarioConectado.Login[:3] != "act" {
		traza.Documento = "Agregando Historial Digital ( " + cedula + " )"
		traza.Log = cadena
		traza.CrearHistoricoConsulta("hmilitar")
		M.Mensaje = "Carga exitosa"
		M.Tipo = 2
		militarF.ActualizarFoto(cedula)
	} else {
		M.Mensaje = "Carga fallida"
		M.Tipo = -1
	}

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
