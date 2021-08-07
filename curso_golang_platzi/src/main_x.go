package main

/*
import (
	// para agregarle un Alias con el nombe "pk"
	//pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

//Se esta definiendo el método de la clase.
// MiPC = Es la instancia de la clase "pc" para ser utilizada como método.
func (MiPC pc) ping() {
	fmt.Println(MiPC.brand, "Pong")
}

// Para cambiar el contendio de la variable utilizando la direccion de la misma
// *pc = Para accesar al valor de la dirección de memoria de "pc", es decir el contenido de la varible
// Es la forma mas optima de modificar el contenido de las variables, además es la forma mas común de utilizarla en "Golang"
func (MiPC *pc) duplicateRAM() {
	MiPC.ram = MiPC.ram * 2
}

// Función Principal.
func main2() {
	// Punteros
	// Esta es la forma que funciona para asignar valores, de la otra forma de asignacion no funciona.
	a := 50
	b := &a // & Es la direccion de memoria de "a"
	//fmt.Printf("Valor de la variable 'a = ' %d, el valor de la direccion de memoria 'b = ' %d ", a, b)
	fmt.Println(a)
	fmt.Println(b)

	// Ahora imprimir el contenido de la dirección de memoria.
	fmt.Println(b)  // Desplegara la dirección de memoria
	fmt.Println(*b) // El contenido de la dirección de memoria.

	// En los punteros cuando se cambia el valor, todas las demas variables que apunten a esa dirección de mmemoria van a cambiar.
	fmt.Println("Cambiando el valor de la variable, 'a' que apunta a esa direccion cambiara ")
	*b = 100
	fmt.Println("valor nuevo de 'a'", a)
	//  // En los punteros cuando se cambia el valor, todas las demas variables que apunten a esa dirección de memoria van a cambiar.
	// Instanciando la clase
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	// Utilizando el método definido anteriormente.
	myPC.ping()
	fmt.Println(myPC)
	myPC.duplicateRAM()

	// Imprime los valores duplicados
	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

	/*
		// Accesando a la "Struct" (clase) que se definie en la carpeta "/mypackages"
		//var myCar mypackage.CarPublico
		var myCar pk.CarPublico
		myCar.Brand = "Ferrari"
		myCar.Year = 2020
		fmt.Println("Contenido de Struct, que fue defino en otro archivo ", myCar)

		pk.PrintMessage("Hola Platzi")

		// Este codigo no funciona, porque la Struct (clase) es de caracter privado, solo es accible dentro de "mypackage"

			var myCar2 pk.carPrivada
			myCar2.brand = "Ford"
			myCar2.year = "2021"
			fmt.Println("Imprimiendo Struct privado ", myCar2)
*/

/*
		// Se debe definir antes de la funcion "main()"
		// Defieniendo la clases, sus atributos.
		type car struct {
			brand string
			year  int
	}

		// Instanciar una "Struct"
		myCar := car{brand: "Ford", year: 2015}
		fmt.Println(myCar)

		// Otra forma de instanciar la "Struct"
		var otherCar car
		otherCar.brand = "Ferrari"
		// Como no se define el atributo "year", tomara por defecto el Valor Zero del tipo de datos, para este caso es "Int"
		fmt.Println(otherCar)


			// Los maps son mas eficientes que los "Slices", es la mejor opcion para manajer mas de dos listas,
			//De forma nativa utilizan la concurrencia.
			m := make(map[string]int)
			m["Jose"] = 14
			m["pedro"] = 30
			fmt.Println("Contenido del Map ", m)

			// Recorriendo un "Map"
			// El recorrido es de forma aleatoria.
			// Si se requiere que sean en el mismo orden, utilizar los "Slice"
			for i, v := range m {
				fmt.Println(i, v)
			}

			// Obtener un valor del Map
			contenido := m["Jose"]
			fmt.Println("Contenido de un elemento del Map['Jose'] ", contenido)

			// Cuando no exista la llave, retorna el valor Zero dependiendo del tipo de datos que se definio, para este caso es "int", por lo que es 0.
			contenido2 := m["Joses"]
			fmt.Println("Contenido de un elemento del Map['Joses'] NO existe ", contenido2)

			contenido3, valor_retorno := m["Joses"]
			fmt.Println("Contenido de un elemento del Map['Joses'] NO existe y retorna el error ", contenido3, valor_retorno)


				// Se debe definir antes de la Func main()
				func isPalindromo(texto string) {
					var textReverse string
					for i := len(texto) - 1; i >= 0; i-- {
						//textReverse = textReverse+string(texto[i])
						// texto = Es una cadena de caracteres que es como arreglo
						textReverse += string(texto[i])
					}
					// strings.ToLower = Convierte todo a minusculas
					// https://pkg.go.dev/strings = Manual en linea de Go para mas funciones de "Strings".

					if strings.ToLower(texto) == strings.ToLower(textReverse) {
						fmt.Println("Es Palindromo")
					} else {
						fmt.Println("NO Palindromo")
					}

				}

				// Se tiene que definir de esta manera, la otra defincion no funciona
				slice := []string{"hola", "que", "haces"}
				for i, valor := range slice {
					// fmt.Println(i, valor)
					fmt.Printf("Imprimiendo el contenido del arreglo indice %d, valor del indice ' %s ' \n", i, valor)
				}

				// Imprimiendo son el índice.
				for _, valor := range slice {
					// fmt.Println(i, valor)
					fmt.Printf("Imprimiendo el contenido del arreglo sin el indice, el valor del indice ' %s ' \n", valor)
				}

				// Imprimiendo solo el indice.
				for i := range slice {
					// fmt.Println(i, valor)
					fmt.Printf("Imprimiendo solo el indice %d  \n", i)
				}

				// hacer una funcion si una palabra es un palindromo (que se lea igual de izq. a drecha como derecha a izq.)
				isPalindromo("Amor a Roma")



					// Arrays.
					// Son imutable, no se puede agregar otro tipo de dato.
					var arreglo [4]int
					arreglo[0] = 1
					arreglo[2] = 2

					fmt.Println(arreglo)
					// len = Cuantos elementos tiene un arreglo.
					// cap = Capacidad máxima del arreglo
					fmt.Printf("Cuantos elementos tiene el arreglo %d \n", len(arreglo))
					fmt.Printf("Capacidad máxima del arreglo %d \n", cap(arreglo))

					// Slice
					// No se le indica el tamaño.
					// Son mutables, se pueden agregar elementos.
					// Se tiene que asignar de esta forma, de lo contrario muestre error en el editor.
					slice := []int{0, 1, 2, 3, 4, 5, 6}
					fmt.Printf("Mostrano un slice %d, Cuantos elmentos tiene %d, La capacidad %d \n", slice, len(slice), cap(slice))

					// Métodos del "slice" y "Array"
					fmt.Printf("Imprimir el elemento del indice 0 del slice %d \n", slice[0])
					fmt.Printf("Imprimir el elemento desde el indice 0 hasta el 3 del slice %d \n", slice[:3])
					// :3 = Es Excluyente es solo para el rango que se indica
					fmt.Printf("Imprimir el elemento desde el indice 2 hasta el 4 del slice %d \n", slice[2:4])
					// :2 = Es inlcuyente para el rango que se indica
					// :4 = Es Excluyente es solo para el rango que se indica
					fmt.Printf("Imprimir el elemento desde el indice 4 hasta el último %d \n", slice[4:])
					// :4 = Es incluyendte es solo para el rango que se indica

					// Agregando elementos al "Slice"
					slice = append(slice, 7)
					fmt.Printf("Imprimiendo el elemento agregado en el arreglo %d \n", slice)

					// Agregando una lista en el "slice"
					nuevaSlice := []int{8, 9, 10, 11, 12}
					slice = append(slice, nuevaSlice...)
					fmt.Printf("Imprimiendo la lista agregada en el arreglo %d \n", slice)



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

	/* Se deben definir antes de la funcion "Main"
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

}
*/
