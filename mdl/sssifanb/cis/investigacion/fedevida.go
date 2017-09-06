package investigacion

import (
	"time"
)

const (
	RETIRO    int32 = 1
	INVALIDEZ int32 = 2
	GRACIA    int32 = 3
)

type FeDeVida struct {
	fechadeemision time.Time
	datospersona   DatoPersonal
	tipopension    int32 //1 retiro, 2 invalidez, 3 gracia
	estado         bool
}
