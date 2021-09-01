package main

/*
import (
	"errors"
	"fmt"
	"io/ioutil"
)

// Desde la creacion del lenguaje "Go" se penso en manejar los errores desde el momento que se presenten.
// En otros lenguajes manejan excepciones pero son controladas.

// Creando una función para controlar el error.
func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No puede dividor por Cero ")
	}
	return dividendo / divisor, nil
}

/*
Este otra manera de manejar las funciones de multiples retornos.
No se recomienda porque es mas confuso cuando se tienen funciones muy grandes.

func division(dividendo, divisor int) (resultado int, err error ) {
	if divisor == 0 {
		err = errors.New("No puede dividor por Cero ")
		// Retornara los valores Zeros del tipo de datos para este caso es 0.
		return
	}
	resultado = dividendo / divisor
	return
}



func main() {
	// se utiliza una de las funciones integradas que se llama "ioutil"
	// Esta funcion retorna dos valores : el contenido del archivo y el "error"
	// Desde la version 1.16 , esta funcion se llama : "os.ReadFile"
	// Se coloca esta ruta porque se ejecuta desde la carpeta "edteam", pero el archivo fuente esta en "src/5.5-Controlar-Errores.go", igual manera el archivo "texto.txt"
	content, err := ioutil.ReadFile("./src/texto.txt")
	/*
		if err != nil {
			fmt.Printf("Ocurrio un error %v \n", err)
		} else {
			// Se realiza el castig de la variable, ya que la funcion retorna el "bytes" que le corresponde
			fmt.Printf("Imprimiendo el contenido del archivo %v\n", string(content))
		}


	// La otra modificacion del "IF"
	if err != nil {
		fmt.Printf("Ocurrio un error %v \n", err)
		// Para la ejecución del programa
		return
	}
	fmt.Printf("Imprimiendo el contenido del archivo %v\n", string(content))

	//Llamando a la funcion "division"
	// Si se cambia el valor de 5 por 0, se genera el error.
	resultado, error := division(15, 5)
	// Ahora se controla el error
	if error != nil {
		fmt.Printf("Ocurrio un error %v \n", error)
		// Para la ejecución del programa
		return
	}

	fmt.Printf("El resultado de la division %d \n", resultado)

}
*/
