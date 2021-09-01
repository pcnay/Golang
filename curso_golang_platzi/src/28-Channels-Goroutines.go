package main

/*
import (
	"fmt"
)

// Channels = Es la comunicacion entre las gorutienes.
// Solo manejan un tipo de dato.
// Permite compartir datos en las Gorutines de forma nativa.
// Si en la aplicacion que manejan, en todo momento se tiene alto desempeño, es utiliza las instrucciones primitivas de "WaitGroup"

// Los WaitGroup son mas optimos y mejor rendimiento, pero son dificiles de controlar.
// Para eficientizar el codigo y como buena práctica se debe especificar si el canal es de entrada o salida.
// Para el caso de solo Entrada, se coloca en la lado Derecho : c chan<- string)
// Para el caso de solo Salida, se coloca en la lado Izquierdo: c <-chan string)
func say(text string, c chan<- string) {
	c <- text
}

// Para la salida de informacion del canal.
//func say2(texto2 string, c <-chan string) {
//	texto2 = <-c
//}

func main() {
	//Definiendo el "channel" de tipo "string - Cadena"
	// como buena practica es recomendable indicar cuantas goroutines manejara, para este caso solo 1 gorutine
	c := make(chan string, 1)
	fmt.Println("hello")

	// NO despliega "Bye", hasta que se modifique la funcion, pasandole como argumento el canal
	// go say("Bye")
	go say("Bye", c)

	// Para que muestre el datos, es decir que la goroutine se ejecute a tiempo, se debe llamar el contenido de la funcion.
	fmt.Printf("%v \n", <-c)

}

*/
