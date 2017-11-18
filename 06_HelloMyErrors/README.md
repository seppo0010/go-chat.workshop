# Hello My Errors

## Errores globales

Es una buena práctica definir los errores predefinidos en variables
globales, así después pueden compararse e identificar el tipo de error.
Por ejemplo, en [io](https://golang.org/pkg/io/#pkg-variables) se define
la variable _ErrShortBuffer_ para indicar que el tamaño del _buffer_ es
insuficiente, así quien lo invoca puede manejarlo como le parezca
adecuado, por ejemplo extendiendo el _buffer_ y reintentando.

## _early return_

Es convención en _go_ usar _return_ tempranos en vez de _else_ cuando es
posible, especialmente en el manejo de errores.

```golang
err := doSomething()
if err != nil {
    return err
}
doSomethingElse()
// ...
```

# Instrucciones

Definir dos errores como variables globales usando
[_errors.New_](https://golang.org/pkg/errors/#New), una para contenido
vacío y otro para autor vacío.

Hacer que el método `addMessage` falle si los datos son inválidos, y que el
servidor HTTP debería devolver un código de estado 400.
([Pista](https://golang.org/pkg/net/http/#ResponseWriter)).

Después de escribir el código de estado, podemos escribir un JSON que en
la clave `error` indique el mensaje correspondiente.
