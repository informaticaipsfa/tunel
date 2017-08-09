/*
Copyright 2017 Carlos Peña.Todos los derechos reservados.
En informática un Bus de Servicio Empresarial (ESB por sus siglas en inglés)
es un modelo de arquitectura de software que gestiona la comunicación entre
servicios web. Es un componente fundamental de la Arquitectura Orientada a
Servicios.
Un ESB generalmente proporciona una capa de abstracción construida sobre
una implementación de un sistema de mensajes de empresa que permita a los
expertos en integración explotar el valor del envío de mensajes sin tener que
escribir código. Al contrario que sucede con la clásica integración de
aplicaciones de empresa (IAE) que se basa en una pila monolítica sobre una
arquitectura hub and spoke, un bus de servicio de empresa se construye sobre
unas funciones base que se dividen en sus partes constituyentes, con una
implantación distribuida cuando se hace necesario, de modo que trabajen
armoniosamente según la demanda.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/sys/web"
	"github.com/gorilla/context"
)

func init() {
	fmt.Println("")
	fmt.Println("Versión del Panel ", sys.Version)
	fmt.Println("")
	if sys.MongoDB {
		//fmt.Println("Metodo de Encriptamiento ", seguridad.Encriptamiento, "...")
		// sys.MongoDBConexion()
		//sys.PostgresDBConexion()
		fmt.Println("")
		fmt.Println("..........................................................")
		fmt.Println("... Iniciando Carga de Elemento Para el servidor WEB   ...")
		fmt.Println("..........................................................")
		fmt.Println("")

	}
}

func main() {

	// var Militar sssifanb.Militar
	// Militar.ConsultarMGO("10107698")
	// fmt.Println(Militar.Familiar[1].Persona.DatoBasico.NombrePrimero)
	// j, _ := Militar.Familiar[0].AplicarReglasCarnetPadres()
	// fmt.Printf("%s\n", j)
	// var carnet sssifanb.Carnet
	// j, _ := carnet.Listar(0)
	// fmt.Println(j)
	//fmt.Println(seguridad.GetMacAddr())

	// j, _ := Militar.GenerarCarnet()
	// fmt.Printf("%s \n", j)

	// var Semillero fanb.Semillero
	// i, _ := Semillero.Maximo()
	// fmt.Println(util.CompletarCeros(strconv.Itoa(i), 1, 10))

	// var es estadistica.Estructura
	//
	// es.Reduccion()
	// var usuario seguridad.Usuario
	// usuario.Generico()

	//var priorizador estadistica.Priorizador00
	// var familiares estadistica.Estructura
	// var migrado estadistica.Estructura
	// migrado.CrearUsuarios()

	// migrado.CargarEstados()
	// migrado.CargarMunicipio()
	// migrado.CargarParroquia()
	// migrado.CargarComponenteGrado()
	//go priorizador.Capturador()
	// migrado.Migracion()
	// migrado.CargarFamiliar()
	// migrado.CargarCtaBancaria()
	// migrado.CargarMilitar()
	// migrado.CargarPensiones()
	// familiares.ActualizarFamiliar()

	// fmt.Println("Inciando la carga del sistema")
	// web.Cargar()
	// srv := &http.Server{
	// 	Handler:      context.ClearHandler(web.Enrutador),
	// 	Addr:         ":8090",
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// fmt.Println("Servidor Escuchando en el puerto: ", sys.PUERTO)
	// srv.ListenAndServe()
	//
	fmt.Println("Inciando la carga del sistema")
	web.Cargar()
	srv := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto: ", sys.PUERTO)
	go srv.ListenAndServe()
	//
	//https://dominio.com/* Protocolo de capa de seguridad
	server := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO_SSL,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto: ", sys.PUERTO_SSL)
	log.Fatal(server.ListenAndServeTLS("sys/seguridad/https/cert.pem", "sys/seguridad/https/key.pem"))

}
