// se refiere a la confianza en algo.
package seguridad

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gesaodin/bdse/util"
)

//Constantes Generales
const (
	Encriptamiento             = "md5"
	ActivarLimiteDeConexion    = true
	DesactivarLimiteDeConexion = false
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
func GenerarJWT(u Usuario) string {
	peticion := Reclamaciones{
		Usuario: u,
		Rol:     "Development",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Conexion Bus Empresarial",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, peticion)
	rs, e := token.SignedString(LlavePrivada)
	util.Fatal(e)
	return rs
}
