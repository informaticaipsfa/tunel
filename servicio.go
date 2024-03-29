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
implantación distribuida cuando se hace necesario, de modo que trabajen
arquitectura hub and spoke, un bus de servicio de empresa se construye sobre
unas funciones base que se dividen en sus partes constituyentes, con una
armoniosamente según la demanda.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/informaticaipsfa/tunel/sys"
	"github.com/informaticaipsfa/tunel/sys/web"
)

func init() {
	fmt.Println("")
	fmt.Println("Versión del Panel ", sys.Version)
	fmt.Println("")
	if sys.MongoDB {
		fmt.Println("")
		fmt.Println("..........................................................")
		fmt.Println("... Iniciando Carga de Elemento Para el servidor WEB   ...")
		fmt.Println("..........................................................")
		fmt.Println("")

	}
}

func main() {

	fmt.Println("Inciando la carga del sistema")
	web.Cargar()

	srv := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto: ", sys.PUERTO)
	go srv.ListenAndServe()

	srvs := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO_STANDAR,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto: ", sys.PUERTO_STANDAR)
	go srvs.ListenAndServe()
	//
	//https://dominio.com/* Protocolo de capa de seguridad
	server := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO_SSL,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto: ", sys.PUERTO_SSL)
	go server.ListenAndServeTLS("sys/seguridad/https/cert.pem", "sys/seguridad/https/key.pem")

	serverx := &http.Server{
		Handler:      context.ClearHandler(web.Enrutador),
		Addr:         ":" + sys.PUERTO_SSL_STANDAR,
		WriteTimeout: 280 * time.Second,
		ReadTimeout:  280 * time.Second,
	}
	fmt.Println("Servidor Escuchando en el puerto: ", sys.PUERTO_SSL_STANDAR)
	log.Fatal(serverx.ListenAndServeTLS("sys/seguridad/https/cert.pem", "sys/seguridad/https/key.pem"))
}
