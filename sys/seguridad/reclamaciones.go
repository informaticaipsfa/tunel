package seguridad

import jwt "github.com/golang-jwt/jwt/v5"

type Reclamaciones struct {
	Usuario Usuario `json:"Usuario" bson:"usuario"`
	Rol     Rol     `json:"Rol" bson:"Rol"`
	//jwt.StandardClaims
	jwt.RegisteredClaims
}

type WReclamaciones struct {
	WUsuario WUsuario `json:"WUsuario" bson:"wusuario"`
	Rol      Rol      `json:"Rol" bson:"Rol"`
	jwt.RegisteredClaims
}
