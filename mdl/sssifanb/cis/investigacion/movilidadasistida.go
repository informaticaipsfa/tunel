package investigacion

import (
	"time"

	"github.com/gesaodin/tunel-ipsfa/mdl/sssifanb"
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
	Direccion      sssifanb.Direccion
	Telefono       string
	Parentesco     string
}
