// seguridad (del latín securitas) cotidianamente se puede
// referir a la ausencia de riesgo o a la confianza en algo
// o en alguien. Sin embargo, el término puede tomar diversos
// sentidos según el área o campo a la que haga referencia en la
// seguridad. En términos generales, la seguridad se define como "el
// estado de bienestar que percibe y disfruta el ser humano".
package seguridad

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

//Session Seccion de acceso
type Session struct {
	Nombre string
	Acceso string
	Nivel  int
}

//Stores resultados
var Stores = sessions.NewCookieStore([]byte("#za63qj2p-6pt33pSUz#"))

func init() {
	Stores.Options = &sessions.Options{
		Domain:   "192.168.43.174",
		Path:     "/",
		MaxAge:   1800, //Media Hora en segundos
		HttpOnly: true,
	}
	//ObtnerIP()
}

//Crear conexion y variables
func (S *Session) Crear(w http.ResponseWriter, r *http.Request) {

}

//ObtnerIP Direccion Fisica de la maquina
func ObtnerIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		//fmt.Println(a)
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}
