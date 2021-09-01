package greet // Deben indicar el nombre de la carpeta donde se crean los archivos.

// Cuando se declaran fuera de las funciones, esta variable puede ser accedida en todas las funciones.
// NO se puede utilizar el asignador de variables :=
// Solo se puede declarar con "var"
var saluda = "Mundo"

// Para que pueda exportarse a otros paquetes tanto las variables, funciones,
//estructuras, es solo colocar una letra Mayusculas al inicio.
// Para este caso GrretEnglish , Italian.
// Pero la variable "saluda" es solo visible para este paquete "greet.go"

// Los nombre de los paquetes deben ser Claros y concisos.
// Asignando un letra inicial en Mayuscula, "Go" la exportara a otros archivos, ya que lo hace pública.

// English retorna saludo en ingles.
func English() string {
	return "Hi " + saluda
}

func Italian() string {
	return "Ciao " + saluda
}
