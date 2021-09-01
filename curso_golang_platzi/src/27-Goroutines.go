package main

/*
import (
	"fmt"
	"sync"
	"time"
)

// Se agrega el segundo parametro para hacer uso de la WaitGroup
func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // Se ejecuta hasta el último este método del WaitGroup
	fmt.Printf("Imprimiendo el texto que enviaron como parámetro %s \n", text)
}

// La funcion "main()" esta ejecutandose en una Goroutine, una vez ejecutada muere
func main() {
	// Genera un conjunto de gorutine y las libera poco a poco.
	var wg sync.WaitGroup

	// say("Hello")
	// Se cambio por la modificacion de la funcion para el uso de "WaitGroup"
	fmt.Println("Hello")

	// Esta función no se ejecuta, dado que como la funcion "main()" como es una gorutine, se termina al ejecuta - say("Hello") -, no alcanza a ejecutarse la gorutine -say ("world")- porque no esta en el hilo de ejecucion de la gorutine "main()"
	// Lo que se tiene que hacer :
	// 1 - Basica pero no recomendable, tampoco eficiente, dado que se agrega tiempo de espera.
	// Para solucionarlo se utilizan "channel", mas adelante se vera.

	//time.Sleep(time.Second * 1) // Dara mas tiempo para que se ejecute el hilo de la gorutine "say("World")"

	// 2 - Es utilizar "WaitGroup", se comentara el * time * para utilizar esta instrucción
	//sync = Permite interactuar de forma primitiva con las "gorutine", es mas eficiente, pero es mas complejo mantenerlo.
	wg.Add(1)            // Agrega una gorutine al WaitGroup para que espera la gorutine main() para dar tiempo a la gorutine "say("World")
	go say("World", &wg) // Para que la función sea concurrente se antepone la palabra "go"

	wg.Wait() // Se le indica a la gorutine main() que espere, hasta que se finalize todas las gorutine de WaitGroup

	// Uitlizando las gorutine con funciones anonimas.
	go func(text string) {
		// fmt.Printf("Adios \n")
		// Pasandole como argumento a la funcion.
		fmt.Printf("%s \n", text)
	}("Adios")
	time.Sleep(time.Second * 1) // Dara mas tiempo para que se ejecute el hilo de la gorutine "say("World")"

}
*/
