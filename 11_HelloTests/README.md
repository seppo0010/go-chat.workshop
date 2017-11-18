# Hello Tests

## Cosas nuevas

### _Tests_

#### Archivos

Los archivos con el sufijo "\_test.go" se pueden usar solamente para
pruebas. En general acompañan al archivo que prueban, por lo que
_myfeature.go_ va a ser testeado por _myfeature\_test.go_.

#### Funciones

Cada test está encapsulado en una función que debe comenzar con "Test" y
recibir sólo un parámetro, de tipo _\*testing.T_. Cualquier otra función
en este archivo estará disponible sólo cuando se corren las pruebas y no
se va a ejecutar como _test_.

El parámetro recibido, por convención nombrado _t_, tiene varias
funciones para indicar el éxito o fracaso del test. _t.Error_ indica
que falló y recibe parámetros para mostrar el fallo, mientras que
_t.Errorf_ permite proveer un string para darle un formato al error a
mostrar.

```golang
func TestTimeIsEven(t *testing.T) {
	timestamp := time.Now().Unix()
	if timestamp % 2 != 0 {
		t.Errorf("expected timestamp to be odd, %v received instead", timestamp)
	}
}
```

#### Ejecutar

Para correr los tests, sólo necesitamos ejecutar `go test` que va a
buscar todas las funciones de prueba declaradas y las va a correr y
reportar el error
```
$ go test
PASS
ok
```

#### Paralelizar

Las pruebas que puedan ser corridas en paralelo se pueden indicar
llamando al método _Parallel_ de _testing.T_ y el ejecutor va encargarse
de hacerlo. Por defecto las pruebas corren de a una por vez.

### Más Redis

Redis provee la función _PUBSUB NUMSUB_ para indicar cuantos clientes
están conectados a un canal.

## Hello tests

Agregar un método a _Server_ (y a sus dos implementaciones) donde se
indique la cantidad de clientes conectados.

Agregar tres tests a cada implementación: uno para que la cantidad de clientes
conectados al _chat_ sea correcta después de subscribirse y de desuscribirse,
otro para ver que al subscribirse a un canal se reciben los nuevos mensajes
publicados y finalmente uno que cuando se pida la lista de mensajes, los
mensajes publicados se reciban.

Cuidado que a veces algunas tareas pueden estar corriendo en otra _go
routine_ y hacer que algún efecto sea inmediato.
