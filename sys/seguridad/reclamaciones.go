package seguridad

import jwt "github.com/dgrijalva/jwt-go"

type Reclamaciones struct {
	Usuario
	Rol string
	jwt.StandardClaims
}
