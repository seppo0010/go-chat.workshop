# Hello Concurrency

En el paso anterior teníamos un error sutil pero crítico: no estabamos
manejando correctamente las conexiones concurrentes. Las listas (o
_slices_) en _go_ no pueden ser escrita desde dos _threads_ porque la
función _append_ no es atómica, por lo que si dos usuarios estuviesen
escribiendo mensajes a la vez nuestro código fallaría.

Los _threads_ en este caso son creados automáticamente por el servidor
HTTP, cada petición crea su propio hilo. Vale aclarar que estos
_threads_ no son _threads_ del sistema operativo sino abstracciones de
_go_ más livianos.

Hay varias formas de corregir este problema, por ejemplo usando un
semáforo para que los distintos hilos frenen y se esperen, pero la forma
idiomática de hacerlo en _go_ es crear un canal donde publicar los
mensajes y un _hilo_ que se encargue de ejecutarlos.

## Cosas nuevas

### _goroutines_

Los hilos se crean invocando una función usando la palabra _go_ antes de
la llamada. Esto crea un nuevo _thread_ para la ejecución de la función
y hace que el actual siga a la siguiente línea.

```golang
go func() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Slept")
}()
fmt.Println("Main")
time.Sleep(100 * time.Millisecond)
```

### Canales

Los canales son la forma que se usa para comunicarse entre distintos
hilos. Cada canal tiene un tipo asociado que es el tipo de datos que
puede publicarse. Para crear un canal hay que llamar
_make(chan myChannelType)_, para publicar un mensaje se llama al
operador `<-` con el valor, y para recibir uno `= <-`

```golang
messages := make(chan string)
go func() { messages <- "ping" }()
msg := <-messages
fmt.Println(msg)
```

### Constructores

En _go_ no existen constructores. Por convención se suele crear un
método llamado _new_ o _New_ seguido por el nombre del _struct_ cuando
se necesita inicializar alguna de las propiedades.

```golang
type MyStruct struct{}

func NewMyStruct() MyStruct {
	return MyStruct{}
}
```

### Punteros

Hasta ahora veníamos pasando variables _por valor_ pero también se puede
hacer _por referencia_, tanto para los parámetros como para los
_struct_s que implementan los métodos. Esto hace que no se copien los datos
cada vez que una función es invocada, y que al modificarse su contenido
dentro de la función, el cambio se vea reflejado afuer.
Si en vez de usar una referencia, recibimos una variable por valor y lo
modificamos lo que cambia es la copia local y el cambio no se propaga.

Para representar un tipo de puntero usar el prefijo _\*_ en la
declaración, y para obtener el puntero a una variable el prefijo _&_.

``golang
type MyStruct struct {
	number int
}

func (s *MyStruct) Swap(val *int) {
	oldValue := s.number
	s.number = *val
	*val = oldValue
}

func main() {
	s := MyStruct{number: 1}
	val := 2
	s.Swap(&val)
	fmt.Println(val, s.number) // 1 2
}
```

### _for_

_for_ en _go_ tiene tres formas: con cero, uno y tres argumentos. El
primero corre para siempre, el segundo corre mientras la condición sea
verdad (como un _while_) y el tercero es el tradicional de estado
inicial, condición de continuidad y sentencia a ejecutar después de cada
iteración.

```golang
for {
	println("corro por siempre")
}

for rand.Float32() / math.MaxFloat32 < 0.9 {
	println("corro una cantidad aleatoria de veces")
}

for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```

## Instrucciones

Crear una estructura servidor que tenga una lista de mensajes y exponga
los métodos _addMessage_ y _getMessages_. El primero recibe un puntero a
un mensaje y el segundo devuelve una lista de punteros a un mensaje.

En el _constructor_ de la estructura, crear un canal donde publicar los
mensajes y una _goroutine_ que vaya agregandolos a la lista.

Para agregar un mensaje, publicarlo en el canal. Para leer los mensajes,
simplemente devolver la lista ya que no es un problema el acceso
concurrente cuando es sólo lectura.

[Siguiente](../06_HelloMyErrors)
