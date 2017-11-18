# Hello Organizing

## Cosas nuevas

### Multiples archivos

Podemos tener más de un archivo en el _package_, en cuyo caso todos los
identificadores que declaran se comparten y no es necesario hacer ningún
_import_ entre ellos, aunque sí cada uno tiene que importar cada módulo
externo que vaya a usar.s

`main.go`
```golang
package main

func main() {
	println(myVar)
}
```

`myvar.go`
```golang
package main

const myVar = "hello world"
```

Al momento de compilar o ejecutar, hay que pasar todos los archivos del
_package_

```bash
$ go run main.go myvar.go
```

## Instrucciones

Ahora tenemos bien diferenciada el código del servidor de mensajes y la
que se encarga de manejar pedidos HTTP. Podemos extraer el servidor a su
propio archivo y así _main.go_ sólo se encarga de conectar al servidor
HTTP con el servidor de datos.

[Siguiente](../08_Dependencies)
