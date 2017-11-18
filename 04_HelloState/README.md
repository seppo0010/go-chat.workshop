# Hello State

## Cosas Nuevas

Hagamos ahora que el servidor web reciba mensajes por post y se puedan
consultar por get

### Globales

Se pueden declarar variables y constantes globales en cualquier lugar
del _package_ usando los prefijos _var_ y _const_, y pueden ser accedidos
por funciones dentro del _package_, o, de estar en mayúsculas, por
cualquier _package_ que lo importe.

```
var number = 1
const StringValue = ""
func MyFunction() {
	println(number, StringValue)
}
```

### Listas

Podemos guardar una colección de elementos del mismo tipo en una lista.
La sintaxis para declararla es _[]type_ y para inicializarla se puede
hacer _[]type{}_ o _make([]type, len, capacity)_ siendo _type_ el tipo de
elementos, _len_ la cantidad de elementos en la lista y _capacity_ el
espacio que le queremos reservar a la lista. La lista se va a ir
expandiendo automáticamente de ser necesario.

Para agregar elementos a la lista (y como no hay _generics_) se usa la
función del lenguaje _append_, que recibe como primer elemento la lista
a la que se le quiere agregar algo y luego cada elemento que se quiera
agregar. La lista no es modificada sino que se devuelve una nueva lista.

```golang
myList := []int{} // []
myList = append(myList, 1) // [1]
myList = append(myList, 2, 3) // [1, 2, 3]
```

### POST y GET

La misma función se está invocando para los pedidos HTTP que llegan.
Para saber el método, el _Request_ lleva un _Method_ que puede
compararse con
[_MethodPost_, _MethodGet_, etc](https://golang.org/pkg/net/http/#pkg-constants).

Los datos enviados como formulario de HTTP están expuestos en una
función
[_PostFormValue_](https://golang.org/pkg/net/http/#Request.PostFormValue)
del _Request_.

```golang
func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// manejar post
	} else if r.Method == http.MethodGet {
		// manejar get
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
```

### interface{}

Dijimos ya que las interfaces son implicitamente implementadas cuando
todos sus métodos están implementados, de lo que se deduce que una
interfaz vacía es implementada por todos los valores. Por eso,
`interface{}` viene a cumplir las veces de variable que puede ser de
cualquier tipo.

```golang
var a interface{}
a = 1
a = "hello"
```
### JSON

_go_ viene con soporte de [JSON](https://golang.org/pkg/encoding/json/).
Para obtener la representación en JSON de cualquier variable, hay que
crear un _Encoder_ con
[_NewEncoder_](https://golang.org/pkg/encoding/json/#NewEncoder) con un
[io.Writer](https://golang.org/pkg/io/#Writer) y después llamar
[_Encode_](https://golang.org/pkg/encoding/json/#Encoder.Encode) con el
valor que queremos. El parámetro es de tipo `interface{}` así que
cualquier enviar cualquier cosa. Sobre el _Writer_ se va a escribir la
serialización del valor. Hay distintas implementaciones de _Writer_, se puede
usar para escribir [en memoria](https://golang.org/pkg/bytes/#Buffer),
[en archivos](https://golang.org/pkg/os/#File),
[en red](https://golang.org/pkg/net/#TCPConn), etc.

## Instrucciones

Crear una estructura que represente un mensaje con autor y contenido, ambos de
tipo texto.
Crear una variable global que sea una lista de mensajes.
Cuando se recibe una petición por _POST_, agregar un mensaje con el
autor y contenido correspondiente.
Cuando se recibe una petición por _GET_, devolver la lista de mensajes.

Debería poder responder de esta forma:

```bash
$ curl localhost:8080
[]
$ curl localhost:8080 --data author=me\&content=hello
$ curl localhost:8080
[{"Author":"me","Content":"hello"}]
$ curl localhost:8080 --data author=not-me\&content=how\ are\ you\?
$ curl localhost:8080
[{"Author":"me","Content":"hello"},{"Author":"not-me","Content":"how are you?"}]
```

[Siguiente](../05_HelloConcurrency)
