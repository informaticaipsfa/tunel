#!/bin/sh

echo "Bajando el servicio esb.ipsfa"
echo "Por favor espere..."
pkill esb.ipsfa

echo "Eliminando la versión actual del servicio esb.ipsfa"
echo "Por favor	espere..."
rm -rf esb.ipsfa

echo "Compilando la nueva versión del servicio esb.ipsfa"
echo "Por favor	espere..."
go build -o esb.ipsfa servicio.go && ./esb.ipsfa & 

