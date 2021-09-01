package main

import (
	"edteam/src/greet" // Se tiene que importar para que se puede usar
	"fmt"
)

func main() {
	// Paquetes = Es una carpeta que contiene programas que proveen funcionalidad.
	// En los nomnbres de los paquetes deben ser concisos y claros.

	fmt.Printf("Hello \n")
	fmt.Printf("Saludo en Ingles = %s \n ", greet.English())
	fmt.Printf("Saludo en Espanol = %s \n", greet.Spanish())
	fmt.Printf("Saludo en Italiano = %s \n", greet.Italian())

}
