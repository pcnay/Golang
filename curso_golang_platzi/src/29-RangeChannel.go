package main

/*
import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text // Recive datos
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"
	// cap = Capacidad máxima que puede almacer el Channel
	// len = Cuantas "goroutines" hay en el Channel
	fmt.Printf("Longuitud = %v y Capacidad = %v del Channel \n", len(c), cap(c))
	// Para realizar un recorrido de los channels
	close(c) // Le indica a runtime de Go que va ha cerrar el canal
	// Para que no siga esperando el "runtime" de Go valores.

	// recorrido del Channel activo
	// "range" es el adecuado para interar con cada uno de los elementos del canal.
	for message := range c {
		fmt.Printf("Imprimiendo el contenido del canal activo = %v \n ", message)
	}

	// Select :
	// Cuando se manejan multiples canales y no se sabe quien va a responde primero, se utilizan los "Select".
	// NO se definen cuantos canales manejara.
	email1 := make(chan string)
	email2 := make(chan string)

	// Hacemos que la funcion sea concurrente, se agrega gorutines
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	// Como no se sabe cual va a responder primero, se utilizan los Select.
	// 2 porque es el numero de canales que se utilizan.
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1: // Guarda la salida de este canal.
			fmt.Printf("Email recibido de * email1 * %v \n ", m1)
		case m2 := <-email2:
			fmt.Printf("Email recibido de * email2 * %v \n ", m2)
		}

	}

}
*/
