package main

/*
import (
	"fmt"
)

// llamada_porDetras = Es la funcion que filtrara, que retornara un valor Booleano.
// La funcion filter tiene parámetros: 1ro = Numeros enteros, 2do = Es una funcion (con paramentros Enteros) retorna valor Booleano; que retornara un slice de enteros.
func filter(nums []int, llamada_porDetras func(int) bool) []int {
	result := []int{}
	// Recorres el slice "nums"
	// Llaman la función que se definen en el "main"
	for _, valor := range nums {
		if llamada_porDetras(valor) {
			result = append(result, valor) // Si es verdadero se agrega el número al "slice"
		}
	}
	return result
}

// Esta retornara una funcion de tipo string
func hello(name string) func() string {
	return func() string {
		return "Hello " + name
	}

}

// Ahora mandando un parámetro a la funcion que se envia como argumento
func hello2(name string) func(string) string {
	return func(text string) string {
		return "Hello " + name + " " + text
	}
}

func main() {
	// Las funciones son tipos de datos
	nums := []int{1, 4, 70, 5, 67, 90, 2}
	// Se pasa como parámetro la funcion y se define
	// a su vez la funcion "filter" retorna los numeros que cumplen la condicion
	result := filter(nums, func(num int) bool {
		if num <= 10 {
			return true
		} else {
			return false
		}
	})

	fmt.Printf("Se imprime el valor %d \n ", result)

	// Una funcion que retorna un funcion de tipo "string", para ejecutarlo se tiene que colocar los parentesis
	//hello("pedro") = Se ejecuta la primera función; () = Ejecuta la segunda funcion.
	x := hello("Pedro")()
	fmt.Printf("Imprimiendo el valor de X = %v \n", x)

	// Otra forma:
	y := hello("SorJuana")
	fmt.Printf("Imprimiendo el valor de x = %v \n", y())

	z := hello2("Ruben")("Orozco")
	fmt.Printf("Imprimiendo el valor Hello2 de x = %v \n", z)

	// Ahora mandando un parámetro a la funcion que se envia como argumento
}
*/
