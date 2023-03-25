## Iniciando con Golang

Ir al sitio oficial de [go.dev](https://go.dev/) sección descargas y elegir la opción del sistema operativo que se desee descargar, por ejemplo para linux descargaremos el archivo `go1.20.2.linux-amd64.tar.gz`

1. Ejecutamos en la terminal de la carpeta de descargas lo siguiente:
```sh
# Solamente si ya se tiene go instalado para remover la instalación actual
sudo rm -rf /usr/local/go

# Despues ejecutamos en descargas el siguiente comando
sudo tar -C /usr/local -xzvf go1.20.2.linux-amd64.tar.gz

tar: descomprime
-C: ruta donde extraera el archivo -> /usr/local
-x: extrae el archivo
-z: indica que el archivo tiene extensión .gz
-v: indica que es verbose es decir imprime en terminal el proceso de extracción
-f: indicamos el archivo con el que trabaja -> go1.20.2.linux-amd64.tar.gz
```
2. Agregamos /usr/local/go/bin al PATH de las variables de ambiente

```sh
sudo nano ~/.bashrc
export PATH=$PATH:/usr/local/go/bin
source ~/.bashrc
```

## Estructura de un archivo de Go

```go
package main
 
import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```