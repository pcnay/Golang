package main

/*
import (
	"fmt"
)

// Es un método que puede compartir otros métodos
// Por ejemplo se tiene un método para varios -struct-, por lo que se busca una sola entrada para los métodos, es mediante por una "Interfaz"

type figuras2D interface {
	area() float64 // Es como se definio en los metodos de las figuras geometricas
}
type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

// Calcular el área
func (c cuadrado) area() float64 {
	return c.base * c.base
}
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Printf("Area : %f \n", f.area())
}

func main() {
	// myCuadrado := cuadrado{base:2}
	// Instanciando los "struct"
	var myCuadrado cuadrado
	var myRectangulo rectangulo
	myCuadrado.base = 2
	myRectangulo.base = 2
	myRectangulo.altura = 4

	// Como se nota se tienen métodos comunes en los struct "area" para cada figura Geometrica.
	// La forma tradicional es ejecutar el método "area" de cada figura Geometrica.
	myCuadrado.area()
	myRectangulo.area()

	// Utilizando las interfaces:
	// Toma el metodo de "area", dependiendo que "struct" manden como parámetro.
	// Cuando se tiene programas muy grandes, es mas eficiente realizado de esta manera.
	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista de Interfaces:
	// Se puede agregar varios tipos de datos.

	// Lista de intefaces
	// En "Go" no se puede agregar la siguiente lista: [1,"bo",true,23.45]
	// Ya que en "Go" en los arreglos se tiene que definir que tipo de datos contendra.
	// Como solucionarlo.
	myInterface := []interface{}{1, "bo", true, 23.45}
	fmt.Println(myInterface...)
	fmt.Printf("Imprimiendo el -slide- con multiples tipos de datos %v \n", myInterface)

}
*/
