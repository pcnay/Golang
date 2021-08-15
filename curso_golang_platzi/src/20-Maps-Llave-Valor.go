package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pedro"] = 20

	// se utiliza el verbo "%v" para desplegar el valor del Diccionario en Go Lang
	fmt.Printf("Para Desplegar el contenido de un diccionario = %v \n", m)

	// Recorrer el Map
	// Utiliza la concurrencia, por lo que el orden de los elementos son de forma aleatoria
	// Si se requiere que lleven un orden, utilizar los "slice"

	for indice, valor := range m {
		fmt.Printf("Recorriendo el 'map' indice = %v, valor = %v \n", indice, valor)
	}

	// Encontrar un valor.
	valor := m["Jose"]
	fmt.Printf("El valor del indice %s, con el valor %v \n", "jose", valor)

}
