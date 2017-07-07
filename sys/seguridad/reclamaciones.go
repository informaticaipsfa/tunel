package seguridad

import jwt "github.com/dgrijalva/jwt-go"

type Reclamaciones struct {
	SUsuario SUsuario
	Rol
	jwt.StandardClaims
}
