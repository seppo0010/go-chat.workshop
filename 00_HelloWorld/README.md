# Hello World

## Cosas nuevas

### Instalar golang

El primer paso para usar golang es instalarlo :)

En OS X, se puede hacer con [homebrew](https://brew.sh/), `brew install golang`.

En Ubuntu, `apt install golang`.

Para otros sistemas operativos buscar la distribución oficial en
https://golang.org/

### Punto de entrada

Todos los archivos de go pertenecen a un _package_. Cada directorio sólo
puede tener un _package_ y todos sus archivos tienen que declararse
parte del mismo.
Al inicio de cada archivo de _go_ se tiene que declarar a qué package
pertenece el mismo.

```golang
package mypackage
```

El _package_ ejecutable se tiene que llamar _main_ y la función principal
también _main_.

### Imprimir

_println_ es una función especial del lenguaje que
imprime un string seguido de un salto de línea en el _standard output_.

### _The Go Playground_

Para hacer pequeñas pruebas de _go_, se puede usar el _Playground_, un editor
online que ejecuta el código que introduzcamos.

https://play.golang.org/

## Instrucciones

Crear un archivo _main.go_ que al ejecutarlo diga "hello world". Para
ejecutarlo, correr `go run main.go`

### Solución

```golang
package main

func main() {
    println("hello world")
}
```

[Siguiente](../01_HelloFormatting)
