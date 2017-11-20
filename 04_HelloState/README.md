# Hello State

## Cosas Nuevas

Hagamos ahora que el servidor web reciba mensajes por post y se puedan
consultar por get.

### Globales

Se pueden declarar variables y constantes globales en cualquier lugar
del _package_, pero fuera de cualquier función o método,  usando los prefijos
_var_ y _const_, y pueden ser accedidos por funciones dentro del _package_,
o, de estar en mayúsculas, por cualquier _package_ que lo importe.

```golang
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
		value := r.FormPostValue("key")
		println(value)
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

Las funciones _Marshal_ y _Unmarshal_ permiten serializar y deserializar
cualquier tipo de variable. La primera devuelve un array de bytes,
que puede ser usado luego por el método Write ;), y un error en caso de
haber fallado. Los fallos son en condiciones poco comunes, como que alguno
de los valores no pueda ser representado en json como un número complejo.
Por ahora, podemos ignorar el error, en cuyo caso se le asigna a una variable
de nombre `_` (guión bajo) que el compilador entiende que no se va a usar.

Cuidado con el nombre de las propiedades y su visibilidad, porque las
propiedades que sean privadas van a ser ignoradas por _json_ dado que no las
puede ver.

```golang
type MyStruct struct {
	Number  int
	Strings []string
	privateProperty int
}

myStruct := &MyStruct{Number: 3, Strings: []string{"hello", "world"}, privateProperty: 1}
serialized, _ := json.Marshal(myStruct)
fmt.Println("serialized", string(serialized)) // serialized {"Number":3,"Strings":["hello","world"]}

var otherStruct MyStruct
err = json.Unmarshal(serialized, &otherStruct)
if err != nil {
	return
}
fmt.Println("deserialized", otherStruct) // deserialized {3 [hello world]}
```

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
