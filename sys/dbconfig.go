package sys

type BaseDeDatosPermisos struct {
	CrearBaseDeDatos     bool
	CrearTablas          bool
	CrearFunciones       bool
	CrearDisparadores    bool
	EliminarBaseDeDatos  bool
	EliminarTablas       bool
	EliminarFunciones    bool
	EliminarDisparadores bool
}

//pg_restore --host localhost --port 5432 --username "postgres" --dbname "banca" --no-password  --verbose "/home/crash/bdse/public/temp/banca20032017.backup"
