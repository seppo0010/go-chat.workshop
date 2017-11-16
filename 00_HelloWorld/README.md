Hello world

# Cosas nuevas

## Instalar golang

El primer paso para usar golang es instalarlo :)

En OS X, se puede hacer con [homebrew](https://brew.sh/), `brew install golang`.

En Ubuntu, `apt install golang`.

Para otros sistemas operativos buscar la distribución oficial en
https://golang.org/

## Punto de entrada

Todos los archivos de go pertenecen a un _package_. Cada directorio sólo
puede tener un _package_ y todos sus archivos tienen que declararse
parte del mismo.
El _package_ ejecutable se tiene que llamar _main_ y la función principal
también _main_.

## Imprimir

_println_ es una función especial del lenguaje que
imprime un string seguido de un salto de línea en el _standard output_.


# Hello World

Crear un archivo _main.go_ que al ejecutarlo diga "hello world".

## Solución

```golang
package main

func main() {
    println("hello world")
}
```
