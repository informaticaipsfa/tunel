package seguridad

import jwt "github.com/dgrijalva/jwt-go"

type Reclamaciones struct {
	Usuario Usuario `json:"Usuario" bson:"usuario"`
	Rol     Rol     `json:"Rol" bson:"Rol"`
	jwt.StandardClaims
}
