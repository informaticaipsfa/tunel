package investigacion

import (
	"time"
)

const (
	BUENO                int32 = 1
	MALO                 int32 = 2
	DISPONIBLE           int32 = 1
	PRESTADO             int32 = 2
	MULETAS              int32 = 1
	BASTONES             int32 = 2
	ANDADERAS            int32 = 3
	SILLADERUEDAS        int32 = 4
	TRICICLOSMOTORIZADOS int32 = 5
)

type MovilidadAsistida struct {
	Serial    int32
	Condicion int32 // 1 bueno 2 malo
	Estado    int32 //1 disponible 2 prestado
	Tipo      int32 // 1 muletas 2 bastones 3 andaderas 4 silla de ruedas 5 triciclos motorizados
}

type Inventario struct {
	MovilidadAsistida MovilidadAsistida
	Cantidad          int32
	FechaIngreso      time.Time
}

type Prestamo struct {
	DatoPersona       DatoPersonal
	MovilidadAsistida MovilidadAsistida
	FechaPrestamo     time.Time
	FechaDevolucion   time.Time
}

type DatoPersonal struct {
	Cedula         string
	Situacion      string
	NombreCompleto string
	Direccion      Direccion
	Telefono       string
	Parentesco     string
}

//Direccion ruta y secciones
type Direccion struct {
	Tipo         int    `json:"tipo,omitempty" bson:"tipo"` //domiciliaria, trabajo, emergencia
	Ciudad       string `json:"ciudad,omitempty" bson:"ciudad"`
	Estado       string `json:"estado,omitempty" bson:"estado"`
	Municipio    string `json:"municipio,omitempty" bson:"municipio"`
	Parroquia    string `json:"parroquia,omitempty" bson:"parroquia"`
	CalleAvenida string `json:"calleavenida" bson:"calleavenida"`
	Casa         string `json:"casa" bson:"casa"`
	Apartamento  string `json:"apartamento" bson:"apartamento"`
	Numero       int    `json:"numero,omitempty" bson:"numero"`
}
