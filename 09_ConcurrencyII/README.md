# Hello Concurrency II

## Cosas nuevas

### _select_

Ya vimos que podíamos recibir mensajes de un _chan_, y que el _thread_
queda esperando hasta que eso pase. También podemos recibir de
cualquiera de múltiples canales usando _select_ y un _case_ por cada
canal que queremos esperar. El primero que esté disponible va a ser
llamado.

```golang
c1 := make(chan string)
c2 := make(chan string)

go func() {
        time.Sleep(time.Second * 1)
        c1 <- "uno"
}()
go func() {
        time.Sleep(time.Second * 2)
        c2 <- "dos"
}()

for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
                fmt.Println("received", msg1)
        case msg2 := <-c2:
                fmt.Println("received", msg2)
        }
}
```

## Instrucciones

Cuando agregamos y quitamos canales de la lista de subscripciones
estamos nuevamente modificando una lista desde múltiples hilos, lo que
genera problemas de concurrencia.

Crear un canal para agregar y uno para quitar subscripciones a los
mensajes y hacer ambas modificaciones desde la misma _goroutine_ donde
se publican los mensajes.

Si en algún momento escriben _chan chan_ no se asusten.

[Siguiente](../10_HelloDatabases)
