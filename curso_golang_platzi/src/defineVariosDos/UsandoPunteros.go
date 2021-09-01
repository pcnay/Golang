package defineVariosDos // NOmbre de la carpeta que contiene este archivo

import (
	"fmt"
)

// Se define en Mayusculas para que el modificador de acceso sea Público
// El nombre del Struct es "PC"
type PC struct {
	Marca   string
	Disco   uint8
	Memoria uint8
}

// Si se quiere asignar valores de forma generica, se debe
// No se pasa el "struct" a la funcion porque va a generar la instancia, lo que si se le debe colocar el que retorna, es es el Struct ->PC
func Datos(disco, memoria uint8, marca string) PC {
	var PartesComp PC
	PartesComp.Marca = marca
	PartesComp.Disco = disco
	PartesComp.Memoria = memoria

	return PartesComp
}

// Se tiene que pasar el struct, para poder accesar a sus atributos.
func (DatosPC PC) ImprimirDatosComp() {
	fmt.Printf("La marca = %s; Disco duro = %d; Memoria = %d \n ", DatosPC.Marca, DatosPC.Disco, DatosPC.Memoria)
}

// Obtener la informacion de la memoria.
// Se pasa a la funcion el Struct para poder accesar al dato requerido, y retornara un tipo de datos uint8, que es como se definio en el "Struct"
func (DatosPC PC) ObtenerMemoria() uint8 {
	return DatosPC.Memoria
}

// Asignar el tamaño de memoria de forma dinamica}
// Se pasa el "struct" pero de tipo puntero con el identificador "*" para accesar a su valor, se le pasara como parámetro el nuevo tamaño.
// La ventaja de esta manera es que cualquier variable que este apuntando a esta dirección tendrá siempre el valor actualizado.
func (DatosPC *PC) AsignarTamano(nuevoTamano uint8) {
	DatosPC.Memoria = nuevoTamano
}

// Ahora es duplicar el tamaño através de un método, y utilizando los punteros.
func (DatosPC *PC) DuplicarTamano() {
	DatosPC.Memoria = DatosPC.Memoria * 2
}
