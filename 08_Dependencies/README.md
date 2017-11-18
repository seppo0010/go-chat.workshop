# Hello Dependencies

## Cosas Nuevas

### Examinar y modificar una lista

Ya habíamos visto como agregar elementos a una lista. Se puede acceder a
la longitud usando la función _len_ y a cada elemento usando su índice
entre corchetes. Para recorrer una lista se usa la palabra _range_ que
devuelve por cada elemento su índice y su valor.

Se puede crear una nueva lista usando un rango de índices en la lista
existente entre corchetes usando dos puntos de separador. Hay más
patrones comunes disponibles para copiar y pegar acá
https://github.com/golang/go/wiki/SliceTricks

```golang
myList := []int{1, 2, 3, 4, 5, 6}
fmt.Println(len(myList)) // 6
fmt.Println(myList[3]) // 4
for index, value := range myList {
	fmt.Println(index, value) // 0 1, 1 2...
}
fmt.Println(myList[3:5]) // [4 5]
fmt.Println(myList[:2]) // [1 2]
fmt.Println(myList[4:]) // [5 6]
```

### _variadics_

_variadics_ son funciones que aceptan una cantidad variable de
parámetros. Ya usamos un par de ellas como _fmt.Println_ y _append_. Se
puede definir una función como variadic simplemente usando `...type`
como último tipo de parámetro recibido, y dicha variable será un
`[]type` dentro de la función.

```golang
func myFunc(args ...interface{}) {
	fmt.Println(args)
}
```

También es interesante es que se puede invocar una función expandiendo
una lista en los argumentos, es decir que si se tiene un _[]interface{}_,
eso puede usarse como parámetro para _myFunc_ agregando tres puntos
al final de la variable.

```golang
myVar := []interface{}{1, 2, "3"}
myFunc(myVar...)
```

Esto nos permite, usando la función _append_, crear una nueva lista
quitando uno de los elementos de una lista que se tiene.

```golang
myList := []int{1, 2, 3, 4}
myList = append(myList[:1], myList[2:]...)
fmt.Println(myList) // [1 3 4]
```

### Servir archivos estáticos

El _package_ _http_ trae una función para manejar archivos estáticos
automáticamente. Con
[_ServeFile_](https://golang.org/pkg/net/http/#ServeFile) sólo
necesitamos que nuestra respuesta elija el archivo a usar como
respuesta.

```golang
func handleRequest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
```

### Usar dependencias externas

Cada _package_ de _go_ de la _stdlib_ como venimos usando se importa con
su nombre, pero este formato es extensible y URLs pueden ser usadas
como dependencias.

Por ejemplo, la librería más popular para hacer servidores de _websockets_
se llama _gorilla/websockets_

```golang
import "github.com/gorilla/websocket"
```

La dependencia no se instala sola sino que hay que ejecutar `go get`
desde la terminal

```bash
$ go get github.com/gorilla/websocket
```

Una vez hecho estos dos pasos, se puede usar _websocket_ como veníamos
usando los _packages_.

### Gorilla

_websockets_ son conexiones abiertas con el explorador que sirven para
recibir y mandar mensajes una vez cargada la página.

[_gorilla_](https://github.com/gorilla/websocket) es una librería de
_go_ para manejar estas conexiones fácilmente. En su
[documentación](https://godoc.org/github.com/gorilla/websocket)
podemos ver un ejemplo de cómo usarla.

```golang
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// ... Use conn to send and receive messages.
}
```

```golang
for {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		return
	}
	if err := conn.WriteMessage(messageType, p); err != nil {
		return
	}
}
```

Esto es suficiente como para empezar a usar la librería. Leyendo la
documentación también podemos ver que la conexión tiene,
además de _WriteMessage_, _WriteJSON(v interface{})_ que nos puede ser útil :)

## Instrucciones

Renombrar el endpoint que devuelve un _JSON_ de mensajes a `/messages`,
y hacer que `/` sirva el index.html provisto.

En _Server_, crear métodos _subscribe_ y _unsubscribe_. El primero
debería devolver un _chan \*Message_ donde luego notificará cada mensaje que
se publica, mientras que el segundo tiene que recibir un canal y
eliminar su subscripción.

Agregar _gorilla_ como dependencia, y en `/ws` se cree un websocket
para el cliente. Cuando se crea, nos subscribimos al servidor y cada
mensaje que recibimos se lo enviamos a la conexión como _JSON_. Si falla
la escritura, desuscribimos el canal y cerramos la conexión.

[Siguiente](../09_ConcurrencyII)
