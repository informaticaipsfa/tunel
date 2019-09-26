package seguridad

import jwt "github.com/dgrijalva/jwt-go"

type Reclamaciones struct {
	Usuario Usuario `json:"Usuario" bson:"usuario"`
	Rol     Rol     `json:"Rol" bson:"Rol"`
	jwt.StandardClaims
}

type WReclamaciones struct {
	WUsuario WUsuario `json:"WUsuario" bson:"wusuario"`
	Rol      Rol      `json:"Rol" bson:"Rol"`
	jwt.StandardClaims
}
