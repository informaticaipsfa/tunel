// se refiere a la confianza en algo.
package seguridad

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gesaodin/tunel-ipsfa/util"
)

//Constantes Generales
const (
	ENCRIPTAMIENTO             = "md5"
	ACTIVARLIMITECONEXIONES    = true
	DESACTIVARLIMITECONEXIONES = false
)

//Variables de Seguridad
var (
	LlavePrivada *rsa.PrivateKey
	LlavePublica *rsa.PublicKey
	LlaveJWT     string
)

//init Función inicial del sistema
func init() {
	bytePrivados, err := ioutil.ReadFile("./sys/seguridad/private.rsa")
	util.Fatal(err)
	LlavePrivada, err = jwt.ParseRSAPrivateKeyFromPEM(bytePrivados)
	bytePublicos, err := ioutil.ReadFile("./sys/seguridad/public.rsa.pub")
	util.Fatal(err)
	LlavePublica, err = jwt.ParseRSAPublicKeyFromPEM(bytePublicos)
}

//GenerarJWT Json Web Token
func GenerarJWT(u SUsuario) string {
	rol := Rol{ID: "01", Descripcion: "Development"}
	peticion := Reclamaciones{
		SUsuario: u,
		Rol:      rol,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "Conexion Bus Empresarial",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, peticion)
	rs, e := token.SignedString(LlavePrivada)
	util.Fatal(e)
	return rs
}
