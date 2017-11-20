# Hello Web

Ya teníamos un "Hello world" en el _stdout_, ahora vamos a levantar un
servidor http, escuchar en un puerto y responder peticiones.

## Cosas nuevas

### import

Para importar un _package_ alcanza con poner `import "packagename"` si
está en la [stdlib](https://golang.org/pkg/). Si se quieren hacer
múltiples _imports_ en el mismo archivo se usa una vez la palabra
_import_ y el resto se pone entre paréntesis, una por línea. Por
convención se hace alfabéticamente. Por ejemplo:

```golang
import (
        "fmt"
        "net/http"
)
```

Una vez importado se puede usar el nombre del _package_ directamente. En
caso de tener barras de separación, el nombre es el último componente.

### Funciones

Para declarar una función se usa el _keyword_ `func` seguido de los
argumentos y opcionalmente un valor de retorno. Por ejemplo

```golang
func sum(a float64, b float64) float64 {
        return a + b
}
```


### Privacidad

Todos los identificadores de un _package_ son públicos o privados
automáticamente de acuerdo a si su primera letra está en mayúscula o
minúscula, respectivamente. Entonces cada vez que se use un método
o estructura importada necesariamente va a tener su identificador en
mayúscula.

```golang
func iAmPrivate() {}
func IAmPublic() {}
```

### Declarar una variable

Para introducir una variable nueva hay que asignarle un valor con el operador
`:=`, mientras que para cambiar el valor de una variable existente se le
debe asignar sólo con `=`.
También se puede declarar una variable sin inicializar, en cuyo caso es
automáticamente inicializada al valor vacío correspondiente al tipo (0
para números, _string_ vacío para _string_, falso para booleanos, etc)

```golang
var a int
b := false
b = true
```

### Conversión de tipo

Los valores en _go_ tienen un tipo y no se convierten automáticamente a
otro sino que tienen que ser explícitas. Hay tipos que se pueden
convertir entre ellos como los numéricos, pero no se puede convertir
directamente un _string_ a número o viceversa.

Sí se puede convertir un _string_ a una lista de runas (_[]rune_) o una lista
de _bytes_ (_[]byte_). La primera es una lista de los caracteres y la
segunda de números del 0 al 255. La diferencia aparece con caracteres
que no son ASCII, en donde las runas toman como unidad al caracter que
se imprime mientras que los _bytes_ no.

```golang
println(len([]byte("🍑"))) // 4
println(len([]rune("🍑"))) // 1
```

La forma de invocar una conversión de tipo es llamar al tipo como si
fuese una función.

```golang
[]byte("hello world")
int(2.3)
```

### nil

Hay valores en _go_ que pueden ser _nil_ y otros que no. Los números,
string, booleanos y estructuras no pueden serlo y son automáticamente
inicializados, mientras que listas y punteros sí pueden serlo.

## Instrucciones

Con todo lo visto podemos escribir las seis líneas de código
necesarias para responder pedidos HTTP.

Para eso vamos a importar
[net/http](https://golang.org/pkg/net/http/), definir una función que
reciba un
[ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter) y un
[Request](https://golang.org/pkg/net/http/#Request) y escriba "Hello
world" como respuesta siempre.

Luego en nuestro _main_ vamos a usar
[HandleFunc](https://golang.org/pkg/net/http/#ServeMux.HandleFunc) para
designar esa función como encargada de manejar las peticiones a "/" y
[ListenAndServe](https://golang.org/pkg/net/http/#Server.ListenAndServe)
para empezar a escuchar en un puerto. Estas funciones globales de
_net/http_ son utilitarios para hacer un servidor web sencillo.

Una vez que esté corriendo, abrir un explorador e ir a http://localhost:8080/
en donde deberíamos leer `Hello world`.

Una vez terminada la prueba, con ctrl+C en la terminal donde están ejecutando
el programa se puede salir.

### Solución

```golang
import "net/http"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// La función Write recibe []byte, pero cuando creamos con comillas
	// el valor es string; como go no hace la conversión automáticamente,
	// tenemos que castear para poder pasarle este valor.
	w.Write([]byte("Hello world"))
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
	// la función ListenAndServe, cuando tiene éxito, se queda esperando
	// por nuevas conexiones y nunca devuelve ningún valor, por lo que la
	// siguiente línea no va a ser ejecutada
	println("nunca voy a aparecer")
}
```

[Siguiente](../03_HelloErrors)
