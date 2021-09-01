package greet // Debe corresponder con la carpeta que lo contiene.

// Saludo en español
// Se declaran las funciones a nivel de paquete.
// Se puede utilizar la variable "saluda" porque se definio fuera de la funcion greet (greet.go)
func Spanish() string {
	return "Hola " + saluda
}
