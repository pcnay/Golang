package main

/*

func main() {
	/*
		// Tipos de datos primitivos:
		// El código es mas eficiente se le defines el tipo de dato a utilizar, usar : var ... int...

		Números Enteros:, Por defecto GoLang asigna dependiendo de la arquitectura (X32, X64)
		int = Depende del Sistema Operativo
		int8 = 8 bits = -128 a 127
		int16 = 16 bits = -2̣^15 a 2^15 -1
		int32 = 32 bits = -2̣^31 a 2^31 -1
		int64 = 64 bits = -2̣^63 a 2^63 -1

		Número Entero Positivo:
		uint = Depende del Sistema Operativo
		uint8 = 8 bits = 0 a 2^8 -1
		uint16 = 16 bits = 0 a 2̣^16 - 1
		uint32 = 32 bits = 0 a 2^32 - 1
		uint64 = 64 bits = 0 a 2^64 -1

		Número Decimales:
		float32 = 32 bits = +/-1.18e^-38 a +/-3.4e^38
		float64 = 64 bits = +/-2.23e^-308 a +/-1.8e^308
		// Utilizando números:

		// uint
		// uint8(Alias byte) unsigned 8-bit intergers (0 to 255)
		// uint16 unsigned 16-bit intergers (0 to 65,535)
		// uint32 unsigned 32-bit intergers (0 to 4,294,967,295)
		// uint64 unsigned 64-bit intergers (0 to 184446744073709551615)

		// int
		// init8 unsigned 8-bit intergers (-128 to 127)
		// init16 unsigned 16-bit intergers (-32,768 to 32,767)
		// init32 (Alias - rune -) unsigned 32-bit intergers (2,147,483,648 to 2,147,483,647)
		// init64 unsigned 64-bit intergers (-9,223,372,036,854,775,808 to -9,223,372,036,854,775,807)

		// Si no se especifica el tipo de dato int, por defecto se compiladora dependiendo de la arquitectura que tenga la computadora: uint32 ó uint64
		// Como buenas prácticas y para optimizar el funcionamiento del programa, no utilizar mas memoria de la necesaria.

		var vocal rune = 'a'      // Representa el codigo ASCII de la vocal "a" y se coloca con un comillas simples
		var cadena2 string = "as" // Representaria una cadena y debe ir con Comillas.
		fmt.Printf("El tipo de dato : %T, valor : %d \n", vocal, vocal)
		fmt.Printf("El tipo de dato : %T, valor : %s \n", cadena2, cadena2)

		TEXTOS Y BOOLENOS
		string = ""
		bool = true ó false

		NUMERO COMPLEJOS
		Tema matematico avanzado.
		Complex64 = Real e imaginario float32
		Complex32 = Real e imaginario float64
		Ejemplo: c := 10 + 8i
	// float32, float64

	// NO se puede realizar operaciones entre uint y int, int32 y float32
	// Para realizar la operacion se tienen que convertir al mismo tipo de dato.

	var altura uint8 = 100
	var anchura uint16 = 1000
	var suma uint16 = 0
	// No permite la ejecución del programa, se tiene que cambiar al mismo tipo de dato(casting)
	//suma = altura + anchura
	// El tipo de datos "altura", no cambia
	suma = uint16(altura) + anchura
	fmt.Printf("La suma de la operacion es: %d \n", suma)

	// Identificador Blanco _ ; se asigna un valor, pero no se utiliza.
	_ = 300
	var _ string = "Cadena sin asignar a una variable."

	// Valors asignados por defecto
	// %q = Permite imprimir comillas.
	var cadena3 string
	fmt.Printf("Valor Por defecto Cadena = %q \n", cadena3)

	var entero2 int16
	var decimal float64
	// 0, 0.000
	fmt.Printf("Valor por defecto entero %d, flotante %f \n", entero2, decimal)

	var booleano bool
	// false
	fmt.Printf("Valor por defecto Booleano %v, \n", booleano)



}
*/
