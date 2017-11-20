# Hello Errors

## Cosas nuevas

### _if_

Los _if_ en _go_ son muy parecidos a los del resto de los lenguajes, salvo
quizás porque no requiere paréntesis para las condiciones.

```golang
if 1 < 2 && 3 >= 0 {
    println("el compilador sabe comparar números")
}
```

### Estructuras

Las estructuras de datos pueden contener multiples valores e implementar
sus propios métodos. Como siempre, métodos y propiedades en mayúscula
son públicas.

```golang
type MyStruct struct {
	number1 int
	number2 int
}

// declaramos el método sum en la estructura MyStruct
// `me` es el nombre que va a tener la estructura dentro del método
// cualquier instancia de la estructura va a tener este método disponible
func (me MyStruct) sum(number3 int) int {
	return me.number1 + me.number2 + number3
}

func main() {
	myStruct := MyStruct{number1: 1, number2: 3}
	value := myStruct.sum(5)
	println(value) // 9
}
```

#### Debuggear

Si se quiere ver el contenido de un struct, se puede usar el _package_ _fmt_
que tiene una opción para ver el contenido de cualquier variable, algo parecido
al _console.log_ de javascript.

```golang
import "fmt"

type MyStruct struct {
	number1 int
	number2 int
}

func main() {
	myStruct := MyStruct{number1: 1, number2: 3}
	fmt.Printf("%#v\n", myStruct)
}
```

### Interfaces

Las interfaces son tipos de datos que implementan un conjunto de
métodos. Se consideran implementadas automáticamente por cualquier tipo de
dato que tenga *todos* los métodos listados.

```golang
type SumAndSubstract interface {
	sum(number int) int
	substract(number int) int
}

type OneNumber struct {
	number int
}

func (me OneNumber) sum(number2 int) int {
	return me.number + number2
}

func (me OneNumber) substract(number2 int) int {
	return me.number - number2
}

type TwoNumbers struct {
	number1 int
	number2 int
}

func (me TwoNumbers) sum(number3 int) int {
	return me.number1 + me.number2 + number3
}

func (me TwoNumbers) substract(number3 int) int {
	return me.number1 + me.number2 - number3
}

func main() {
	var v SumAndSubstract
	v = TwoNumbers{number1: 1, number2: 3}
	println(v.sum(5)) // 9
	v = OneNumber{number: 1}
	println(v.sum(5)) // 6
}
```

#### Interfaces, ¿para qué?

En lenguajes de tipo dinámicos como javascript, ruby, python, PHP,
las interfaces no tienen tanto sentido como en un lenguaje con tipos estáticos.

Las interfaces permiten declarar que un parámetro o variable es de un tipo
desconocido, pero del cual se conoce cierto comportamiento. Por ejemplo,
siguiendo el caso anterior, se puede declarar una función que reciba dos
_SumAndSubstract_ y opere con ellos sin saber exactamente de qué tipo son.

```golang
func operate(v1 SumAndSubstract, v2 SumAndSubstract, v3 int) int {
	return v1.substract(v2.sum(1))
}

func main() {
	v1 := TwoNumbers{number1: 10, number2: 100}
	v2 := OneNumber{number: 1000}
	v3 := 10000
	println(operate(v1, v2, v3)) // -891
	println(operate(v2, v1, v3)) // 889
}
```

### El mundo sin excepciones

La función
[_ListenAndServe_](https://golang.org/pkg/net/http/#Server.ListenAndServe)
que usamos para empezar nuestro servidor web devuelve _error_. _error_
es una interfaz que sólo implementa el método `Error() string`.
Cualquier operación que pueda fallar suele devolver un _error_ que es
_nil_ en caso de éxito y no nulo en caso de fallo. Cada vez que se
invoca a una función que puede fallar _habría que_ verificar la
respuesta, comparándola con _nil_ y manejar el fallo o hacer que la función
que invoca devuelva el error.

### Cómo leer la documentación de _go_

En https://golang.org/pkg/ está la documentación de los _packages_ que vienen
en _go_. Cada uno de ellos empieza diciendo cómo importarlo, para qué sirve,
y luego una lista de funciones y estructuras que contiene.

Por ejemplo, el _package_ _os_  se importa con `import "os"` y tiene una
función _Exit_, según se puede ver acá: https://golang.org/pkg/os/#Exit

```golang
package main

import "os"

func main() {
    os.Exit(2)
}
```

## Instrucciones

Recibir el error de _ListenAndServe_ y si no es nulo imprimir la
descripción del error y salir con un estado distinto de cero.

### Solución

A partir de acá la idea es que intenten solucionarlo sin ver la respuesta.
Si se rinden, la respuesta está disponible en el subdirectorio _solucion_.

Para probar cómo hacer que falle se puede ejecutar el programa dos veces.
La primera tendría que tener éxito, pero la segunda falla porque el primer
proceso ya está usando el puerto.

Se debería ver algo así

```bash
$ go run main.go &
[1] 86196
$ go run main.go
listen tcp :8080: bind: address already in use
exit status 1
$
```

[Siguiente](../04_HelloState)
