// se refiere a la confianza en algo.
package seguridad

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	//jwt "github.com/dgrijalva/jwt-go"
	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/informaticaipsfa/tunel/util"
)

// Constantes Generales
const (
	ENCRIPTAMIENTO             = "md5"
	ACTIVARLIMITECONEXIONES    = true
	DESACTIVARLIMITECONEXIONES = false
)

// Variables de Seguridad
var (
	LlavePrivada *rsa.PrivateKey
	LlavePublica *rsa.PublicKey
	LlaveJWT     string
)

// init Función inicial del sistema
func init() {
	bytePrivados, err := ioutil.ReadFile("./sys/seguridad/private.rsa")
	util.Fatal(err)
	LlavePrivada, err = jwt.ParseRSAPrivateKeyFromPEM(bytePrivados)
	bytePublicos, err := ioutil.ReadFile("./sys/seguridad/public.rsa.pub")
	util.Fatal(err)
	LlavePublica, err = jwt.ParseRSAPublicKeyFromPEM(bytePublicos)
}

// GenerarJWT Json Web Token
func GenerarJWT(u Usuario) string {

	peticion := Reclamaciones{
		Usuario: u,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 259200)),
			Issuer:    "Conexion Bus Empresarial",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, peticion)
	rs, e := token.SignedString(LlavePrivada)
	util.Fatal(e)
	return rs
}

// WGenerarJWT Json Web Token
func WGenerarJWT(u WUsuario) string {
	peticion := WReclamaciones{
		WUsuario: u,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 360)),
			Issuer:    "Conexion Bus Empresarial",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, peticion)
	rs, e := token.SignedString(LlavePrivada)
	util.Fatal(e)
	return rs
}
