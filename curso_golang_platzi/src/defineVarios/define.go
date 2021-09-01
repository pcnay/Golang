package defineVarios // Es el nombre de la carpeta.
import (
	"fmt"
)

type Pcs struct {
	Brand string
	Disk  uint
	Ram   uint
}

// Asigna_Valores Publicos.
// Se pasan los parametros para hacerlo general, retornara de tipo Struct
func Asigna_Valores(brand string, disk, ram uint) Pcs {
	// MyPC := Pcs{Brand: "Lenovo", Disk: 50, Ram: 8}
	//Asignando valores de otra manera
	var MyPC Pcs
	MyPC.Brand = brand
	MyPC.Disk = disk
	MyPC.Ram = ram
	return MyPC
}

// Para imprimirlo de forma mas general
// Reciendo como parametro de la función el Struct "Pcs", con el nombre de la variable que lo recibe "DatosPc"
func (DatosPC Pcs) ImprimeDatos() {
	fmt.Printf("Imprimiendo los datos del equpo, Marca = %s, Disco Duro = %d, emoria RAM = %d \n", DatosPC.Brand, DatosPC.Disk, DatosPC.Ram)

}

// Asignado la memoria usando punteros
func (DatosPC *Pcs) AsignarMemoria(valorMemoria uint) {
	DatosPC.Ram = valorMemoria

}

// Retorna un valor de tipo uint
func (DatosPc Pcs) ObtenerMemoria() uint {
	var datoMemoria uint
	datoMemoria = DatosPc.Ram
	return datoMemoria
}

// Duplicar la memoria a través de los punteros.
// Se utiliza "*" ya que va a modificar el valor y todas las variables que este apuntando a este direccion de memoria se modificaran.
func (DatosPC *Pcs) DuplicarMem() {
	DatosPC.Ram = DatosPC.Ram * 2
}
