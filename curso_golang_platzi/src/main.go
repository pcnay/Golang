package main

import (
	"fmt"
)

// Definiendo la funcion
func usandoFuncion() {
	fmt.Println("Cadena de Texto ")
}

func usandoFuncParametros(cadena string) {
	fmt.Println("Usando la funcion con parámetros")
	fmt.Println(cadena)
}

func funcionVariosParam(entero int32, decimal float32, cadena string) {
	fmt.Println("Desplegando los tres parametros de la Funcion")
	fmt.Println("Entero : ", entero, "Decimal : ", decimal, "Cadena :", cadena)

}

func funcionDevuelveValor(a uint8) uint8 {
	return a * 2
}
func funcionDevuelvValores(entero uint) (valor1, valor2 uint) {
	return entero, entero * 100
}

// Función Principal.
func main() {
	fmt.Println("Hola")
	fmt.Println("Mundo")
	// "defer" = Imprime antes que se termine de ejecutar la funcion principal
	// Se puede utilizar por ejemplo cuando abres una conexion a la base de datos y al final que cierre la conexion
	// Buena practica es solo un "defer" por funcion.

	defer fmt.Println("Se imprimira hasta el final HOLA")
	fmt.Println("Mundo")

	// Continue y Brak se utilizan en el ciclo "for"
	for i := 0; i < 10; i++ {
		fmt.Printf("Ciclo For contador = %d \n", i)

		// "Continue", se utiliza para cuando se tiene controlado un error, pero que se sigue continuando
		if i == 2 {
			fmt.Println("Es el 2")
			continue
		}
		if i == 8 {
			fmt.Println("Se ejecuta un Break en el número 8, saldra del ciclo")
			break
		}

	}

	/*
		// Condicional Switch
		var modulo float32
		modulo = 5 % 2
		switch modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
		}
		//swtich sin condiciones
		// Se utiliza cuando la condicion no se conoce
		var valor_1 uint = 300
		switch {
		case valor_1 > 300:
			fmt.Printf("El valor_1 > %d \n ", valor_1)
		case valor_1 < 400:
			fmt.Printf("El valor_1 menor a %d \n ", 400)
		// Se ejecuta cuando no se cumple ninguna condicion.
		default:
			fmt.Println("NO se cumple la condicion")
		}

			// Otra manera
			switch modulo := 4 % 2; modulo {
			case 0:
				fmt.Println("Es par")
			default:
				fmt.Println("Es impar")
			}
	*/

	/*
		var valor1 int = 1
		var valor2 int = 2
		if valor1 == 1 {
			fmt.Println("Es 1")
		} else {
			fmt.Println("NO es 1")
		}
		if valor1 == 1 && valor2 == 2 {
			fmt.Println("Se cumple la condicion, usadndo el operador AND")
		}
		if valor1 == 0 || valor2 == 2 {
			fmt.Println("Se cumple, usando el operador OR")
		}
		// Conversion de tipos de datos.
		var num_entero int = 10
		var texto_Cadena string

		// https://pkg.go.dev/strconv
		fmt.Println("Conviertiendo de Entero a Cadena :")
		texto_Cadena = strconv.Itoa(num_entero)
		fmt.Println(texto_Cadena)
		fmt.Println()

		fmt.Println("Convirtiendo de Cadena a Entero")
		var cadena string = "50"
		//var err string
		//var cadena string = "50"
		//var cadena_entero int
		cadena_entero, err := strconv.Atoi(cadena)
		fmt.Println("Convertido de Cadena a Entero ", cadena_entero, err)



			// Ciclo For
			fmt.Println("Utilizano Ciclo For ")
			for i := 0; i <= 10; i++ {
				fmt.Printf("Numeros del ciclo %d \n", i)
			}

			fmt.Println()
			fmt.Println("Utilizando el For While ")
			// For While
			var counter uint8 = 0
			for counter < 10 {
				fmt.Printf("Numeros a Desplegar %d \n", counter)
				counter++
			}

			// For forever
			// Para romper el ciclo se oprime al tecla CTRL-C en la ventana.
			fmt.Println()
			fmt.Println("Utilizando el For Forever")
			var counterForever uint8 = 0
			for {
				fmt.Printf("Contador Forever %d ", counterForever)
				counterForever++
			}


				usandoFuncion()
				usandoFuncParametros("Cadena Como Parametro")
				funcionVariosParam(20, 350.34, "Cadena Texto")
				var obtenerValor uint8
				obtenerValor = funcionDevuelveValor(3)
				fmt.Printf("%d Valor a retornar \n  ", obtenerValor)

				var variosValores1, variosValores2 uint
				// Retorna dos valores, lo coloca en el orden de como se declara la variable
				variosValores1, variosValores2 = funcionDevuelvValores(5)
				fmt.Printf("%d variosValores1 %d variosValores2 \n", variosValores1, variosValores2)

				// Para el caso de que solo se requiere un solo parámetro.
				//var variosValores3, variosValores4 uint
				// Se coloca comentario la linea anterior porque no funciona para el caso que se quiera omitir una variable.
				variosValores3, _ := funcionDevuelvValores(10)
				fmt.Printf("%d variosValores3 \n", variosValores3)

				// En lugar de escribirlo de forma repetida.
				//fmt.Println("Cadena de Texto 1")
				//fmt.Println("Cadena de Texto 2")
				//fmt.Println("Cadena de Texto 3")


				/*
						// Usando el paquete : FMT
						// Manual en línea de este paquete : https://pkg.go.dev/fmt
						var helloMessage string = "Hellow"
						var worldMessage string = "Mundo"

						// Imprimiendo el mensaje pero con retorno automatico.
						fmt.Println("Imprimir el mensaje : ", helloMessage)
						fmt.Println("Imprimir el mensaje : ", worldMessage)

						// Printf = Le da formato adicional a la cadena que imprime, es mas utilizado
						var nombre string = "Platzi"
						var cursos uint16 = 500
						// %s = Imprimira una cadena
						// %d = Que recibe un entero
						// %v = Cuando no se sepa que tipo de dato se desplegara.
						// \n = Para retorno de carro (enter)
						fmt.Printf("%s tiene mas de %d cursos \n ", nombre, cursos)
						fmt.Printf("%v tiene mas de %v cursos \n ", nombre, cursos)

						//Sprintf = Genera una cadena, no lo despliega en consola.
						var mensajeAsignado string
						mensajeAsignado = fmt.Sprintf("%s tiene mas de %d cursos \n ", nombre, cursos)
						fmt.Println(mensajeAsignado)

						// Para determinar el tipo de datos de una variable
						fmt.Println("Imprimiendo el tipo de dato de una variable :")
						fmt.Printf("helloMessage : %T \n", helloMessage)
						fmt.Printf("cursos : %T \n", cursos)

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

						TEXTOS Y BOOLENOS
						string = ""
						bool = true ó false

						NUMERO COMPLEJOS
						Tema matematico avanzado.
						Complex64 = Real e imaginario float32
						Complex32 = Real e imaginario float64
						Ejemplo: c := 10 + 8i
	*/

	/*
		// Operadores aritmeticos
		x := 10
		y := 50

		// suma
		resultado := x + y
		fmt.Println("Resultado Suma = ", resultado)
		// Resta
		resultado2 := x - y
		fmt.Println("Resultado Resta = ", resultado2)

		// Multiplecacion
		resultado3 := x * y
		fmt.Println("Resultado Multiplicacion = ", resultado3)
		// Division
		resultado5 := y / x
		fmt.Println("Resultado Divisor = ", resultado5)

		// Modulo, el remante de la division, es utilizado para determinar si un número es Par o Impar.
		resultado6 := y % x
		fmt.Println("Resultado Remante Division = ", resultado6)

		// Incremental
		x++
		fmt.Println("Resultado Incremntal = ", x)

		// Decremental
		x--
		fmt.Println("Resultado Incremntal = ", x)

		// Ejercicio :
		// Calcular el área de: Rectangulo, Trapecio y Circulo.
		// Rectangulo = Largo X Ancho
		// Trapecio = ((BaseLaraga + BaseCorte)/2)XAltura
		// Circulo =Pi X r2

		fmt.Println("Area de Rectangulo LargoxAncho")
		largo := 5
		ancho := 20
		areaRect := largo * ancho
		fmt.Println("Area Rectangulo = ", areaRect)

		fmt.Println("Area Trapecio = ((BaseLarga + BaseCorta)/2)XAltura")
		var BaseLarga int = 5
		var BaseCorta int = 20
		var Altura int = 20
		areaTrapecio := ((BaseLarga + BaseCorta) / 2) * Altura
		fmt.Println("Area Trapecio = ", areaTrapecio)

		fmt.Println("Area Circulo = PiXRadio Al cuadrado")
		const pi float64 = 3.141598383
		var radio float64 = 9
		// math.Pow(Numero-Elevar, AquePotencia)
		areaCirculo := pi * (math.Pow(radio, 2))
		fmt.Println("9 al cuadrado = ", math.Pow(radio, 2))
		fmt.Println("Area Circulo = ", areaCirculo)


			// Declaracion de constantes
			const pi float64 = 3.1416
			const pi2 = 3.141689
			//fmt.Println("Hola Mundo")
			fmt.Println("pi : ", pi)
			fmt.Println("pi2 : ", pi2)

			//Declaración de variable enteras
			base := 12
			var altura int = 14
			var area int
			// Mientras no se usen no se prodra ejecutar el programa, es una práctica de programacion (TypeScrit lo usa en Angular, en "Golang" ya lo tiene por defecto). Para evitar el error, puedes utilizarla con "Println" o colocarla en comentario
			fmt.Println("Valor numero base = ", base)
			fmt.Println("Valor numero Altura =", altura)
			fmt.Println("Valor numero Area =", area)
			fmt.Println() // Imprime una línea en Blanco
			fmt.Println("Mostrando los valores por defcto ")
			// Valores Zero
			var a int
			var b float64
			var c string
			var d bool

			fmt.Println("Valor Entero 'a' = ", a)
			fmt.Println("Valor Decimales 'b' = ", b)
			fmt.Println("Valor Cadena 'c' = ", c)
			fmt.Println("Valor Booleano 'd' = ", d)

			fmt.Println() // Imprime una línea en Blanco
			// Otro ejemplo usando las variables.
			const baseCuadrado = 10
			areaCuadrado := baseCuadrado * baseCuadrado
			fmt.Println("Area del cuadrado = ", areaCuadrado)
	*/

}
