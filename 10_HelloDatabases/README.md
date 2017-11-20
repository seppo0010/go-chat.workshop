# Hello Databases

## Cosas nuevas

### _os.Args_

Los parámetros con los que se invocan los programas en _go_ están
disponibles como un _[]string_ en
[os.Args](https://golang.org/pkg/os/#pkg-variables), siendo el primer elemento
el nombre del programa.

```golang
package main
import "os"
func main() {
	println(os.Args)
}
```

```bash
$ go run *.go args1
$ go build *.go -o myprog && ./myprog args1
```

### Cerrar canales

Los canales que usamos para intercambiar mensajes tienen un estado
abierto por defecto y se pueden cerrar. Una vez cerrado, no puede
enviarse más mensajes. Esto sirve para poder usar _range_ con un canal y
detectar cuando terminar.

```golang
c := make(chan int)
go func() {
	c <- 3
	c <- 2
	c <- 1
	close(c)
}()

for i := range c {
	fmt.Println(i)
}
```

### _sync.Map_

En _go_ existen mapas de clave-valor, pero no son _multi-thread safe_.
Sin embargo existe la versión segura que implementa los métodos necesarios.

```golang
m := &sync.Map{}
val, found := m.Load("key")
fmt.Println(val, found) // <nil> false

m.Store("key", "value")
val, found = m.Load("key")
fmt.Println(val, found) // value true
```

### _Cast_

Una variable puede estar declarada con su tipo escalar o estructura, pero
alternativamente puede ser una interfaz, en cuyo caso se puede querer
saber si es de un tipo específico o convertirlo si se sabe, por ejemplo
al recuperar un valor de un _sync.Map_ su tipo es _interface{}_ que
requiere _castearlo_ para poder usar sus métodos o propiedades.

```golang
var a interface{}
a = 1
if astr, ok := a.(string); ok {
        fmt.Println("es el string", astr)
} else if aint, ok := a.(int); ok {
        fmt.Println("es el entero", aint)
} else {
        fmt.Println("no es string ni entero", aint)
}
```

### Redis

Redis es una base de datos que se adecua al caso de uso que tenemos
porque permite guardar listas de datos y publicar mensajes en tiempo
real.

En OS X se puede instalar usando `brew install redis` mientras que en
Ubuntu `apt install redis`. En Windows se puede usar
[este fork](https://github.com/MicrosoftArchive/redis/releases/tag/win-3.0.504)
Para otras versiones ver [acá](https://redis.io/download).

Para lo que vamos a hacer sólo vamos a necesitar conocer los siguientes
comandos:

* [RPUSH](https://redis.io/commands/rpush)
* [LRANGE](https://redis.io/commands/lrange)
* [PUBLISH](https://redis.io/commands/publish)
* [SUBSCRIBE](https://redis.io/commands/subscribe)

La librería para acceder a Redis está en
[github.com/go-redis/redis](https://github.com/go-redis/redis) y su
documentación [acá](https://godoc.org/github.com/go-redis/redis).

## Hello Databases

Hacer que _Server_ sea ahora una interfaz con los métodos _addMessage_,
_getMessages_, _subscribe_ y _unsubscribe_, cada uno recibiendo los
mismos parámetros que venía recibiendo, pero cualquiera de ellos puede
devolver, además de lo que devolvía, un error.

Cuando se produce un error en un pedido HTTP, devolver un error 500;
cuando es dentro de un websocket, desconectar al cliente.

La implementación de Server existente ahora será ServerMemory.
Implementar ServerRedis usando la base de datos correspondiente.
Atención con cómo se maneja la desuscripción de los clientes que puede
permitir errores de concurrencia.

Recibir un parámetro al iniciarse el servidor que sea _--redis_ o
_--memory_ para elegir qué implementación de Server usar.

[Siguiente](../11_HelloTests)
