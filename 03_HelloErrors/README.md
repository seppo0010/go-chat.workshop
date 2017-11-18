# Hello Errors

## Cosas nuevas

### Estructuras

Las estructuras de datos pueden contener multiples valores e implementar
sus propios métodos. Como siempre, métodos y propiedades en mayúscula
son públicas.

```golang
type MyStruct struct {
	number1 int
	number2 int
}

func (me MyStruct) sum(number3 int) int {
	return me.number1 + me.number2 + number3
}

func main() {
	fmt.Println(MyStruct{number1: 1, number2: 3}.sum(5)) // 9
}
```

### Interfaces

Las interfaces son tipos de datos que implementan un conjunto de
métodos. En _go_ no pueden tener implementaciones por defecto y se
consideran implementadas automáticamente por cualquier tipo de dato que
tenga los métodos listados.

```golang
type MyStruct struct {
	number1 int
	number2 int
}

func (me MyStruct) sum(number3 int) int {
	return me.number1 + me.number2 + number3
}

type MyInterface interface {
	sum(number int) int
}

func main() {
	var v MyInterface
	v = MyStruct{number1: 1, number2: 3}
	fmt.Println(v.sum(5)) // 9
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
respuesta, comparandola con _nil_ y manejar el fallo o hacer que la función
que invoca devuelva el error.

## Instrucciones

Recibir el error de _ListenAndServe_ y si no es nulo imprimir la
descripción del error y salir con un estado distinto de cero
([Pista](https://golang.org/pkg/os/#Exit)).

### Solución

A partir de acá la idea es que intenten solucionarlo sin ver la respuesta.
Si se rinden, la respuesta está disponible en el subdirectorio _solucion_.

[Siguiente](../04_HelloState)
