package main
/*
import (
	"fmt"
)

type pc struct {
	ram   uint
	disk  uint
	brand string
}

// Se pasa como parámetro el "struct" para poder utilizar como método
// MyPC = es el nombre que recibe el struct que se envio, puede tener cualquier nombre.
func (MyPC pc) ping() {
	fmt.Printf("Imprimiendo % v, evento : %s \n", MyPC.brand, "Pong")
}

// Se pasa como argumento el "Struct" como puntero ahora que se accedera al valor de la direccion de memoria, el "*"
// Recordando que "MyPC" puede ser cualquier nombre.
// "duplicateRAM()... = Es lo que retorna la funcion.
func (MyPC *pc) duplicateRAM() {
	MyPC.ram = MyPC.ram * 2
}

func main() {
	// Los punteros se utilizan para :
	// Personalizar funciones.
	// Llevar funciones a otro paquete de forma fácil y eficiente.

	var a int = 50
	b := &a // Esta asignando la direccion de memoria a la variable "b"

	// *b = Para acceder a la direcciópn de memoria.

	fmt.Printf("Imprimiendo el valor de [a] = %d y la direccion de memoria =  %v \n", a, b)

	// Para imprimir el contenido de la variable "b" que apunta a la direccion de memoria.
	fmt.Printf("Imprimiendo el contenido [b] de la direccion de memoria %v \n", *b)

	// Si se cambia el valor de la direccion de memoria, las variables que apunten a esa dirección de memoria tambien cambiaran su valor.
	// Ejemplo:
	*b = 100
	// lo cambia, ya que estan apuntando a la misma direccion de memoria.
	fmt.Printf("Imprimiendo el valor de [a] = %v cuando -b- se modifico \n", a)

	// Ejemplo:
	// Se creara un "struct"
	// Instanciando la clase struct
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Printf("Imprimiendo el contenido de la clase = %v \n", myPC)

	// Para desplegar el contenido de la funcion que se paso como parámtro.
	myPC.ping()

	// Ahora ejemplo con punteros.
	// Se utilizaran para modificar los struct en lo referentes al valor de las variables (atributos, métodos, funciones, etc).

	// Se actualizara la RAM
	// Lo común es hacer una asignacion directa
	myPC.ram = 20
	fmt.Printf("Valor nuevo de RAM = %v \n", myPC.ram)

	// Ahora se utilizaran funciones que utiliza punteros para alterar los valores de los atributos de "StrucT"

	// Ahora utilizar la funcion "duplicateRAM" antes defenida.
	fmt.Printf("Imprimiendo el estado actual del -Struct- %v \n", myPC)
	myPC.duplicateRAM()
	fmt.Printf("Imprimiendo el estado despues haber ejecutado la funcion -duplicateRAM- de -Struct- %v \n", myPC)
}
*/