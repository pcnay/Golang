package main

/*
import (
	"fmt"
)

func main() {
	arreglo := [7]string{"Arriba", "Abajo", "Izquierda", "Derecha", "Verde", "Rojo", "Violeta"}
	var colors []string
	colors = arreglo[4:5] // Apunta al array "Arreglo" desde los elemntos 4, al 7
	fmt.Printf("Imprimiendo todo el arreglo 'arreglo' %v \n", arreglo[:])
	fmt.Printf("Imprimiendo solo los colores %v \n ", colors[:])

	// Para obtener el numero de elementos que contiene el "Arreglo"
	fmt.Printf("\n")
	fmt.Printf("Imprimiendo el numero de elementos del array 'arreglo' = %d \n ", len(arreglo))

	fmt.Printf("\n")
	fmt.Printf("Imprimiendo el número de elementos del slide 'colors' = %d \n ", len(colors))

	// cap = No. de elementos del array donde apunta el "slice", a partir del índice de donde se creo el slice.
	fmt.Printf("\n")
	fmt.Printf("Impriendo la capacidad del array 'arreglo' %d \n ", cap(colors))

	// Agregando elementos al "slice", se genera un nuevo Arreglo, se agrega pero no se refleja en en "array"
	colors = append(colors, "Gris", "Negro", "Blanco ")
	fmt.Printf("\n")
	fmt.Printf("Imprimiendo el elemento agregado al 'slice' %v \n", colors[:])

	// Se modifica el array 'arreglo', ya que la capacidad se toma desde indice donde se agrego el "slice"
	// Si se tiene como limite un numero mas (excluyente), no agregara elementos al slice, tomarlo encuenta.
	fmt.Printf("\n")
	fmt.Printf("Imprimiendo el numero de elementos del array 'arreglo' = %v \n ", arreglo[:])

	// Creando un "slice"
	// slice_nuevo := []string
	// slice_nuevo := []string {"Campo1","Campo2"}
	// 0 = Tamaño
	// 3 = Capacidad

	slice_nuevo := make([]string, 0, 3)

	fmt.Printf("Nuevo 'slide' %v \n ", slice_nuevo)
	fmt.Printf("La capacidad del 'slide' = %v \n", cap(slice_nuevo))

	slice_nuevo = append(slice_nuevo, "Piña", "Fresa", "Mandarina", "Ciruela")
	fmt.Printf("\n")
	fmt.Printf("Contenido del arreglo 'slice_nuevo' %v \n", slice_nuevo)
	fmt.Printf("Agregando dos elementos, mostrando la capacidad %v \n", cap(slice_nuevo))
	fmt.Printf("Contenido completo %v \n", len(slice_nuevo))

	// Valores Zero
	var arreglo2 []string
	//arreglo2 := []string

	fmt.Printf("Contenido de 'arreglo2' con valores Zero = %v \n ", arreglo2)
	fmt.Printf("Tamaño 'arreglo2' con valores Zero = %d \n", len(arreglo2))
	fmt.Printf("Capacidad 'arreglo2' con valores Zero = %d \n", cap(arreglo2))
}
*/
