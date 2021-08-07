package mypackage // El nombre de la carpeta
import "fmt"

// CarPublic (Es obligatorio desde // CarPublic, el resto es opcional) Haciendo pública, acceible para otros paquetes
type CarPublico struct {
	Brand string
	Year  int
	//year int // Acceso privado
}

// Definiendo un Struct (Clase) privada.
type carPrivada struct {
	brand string
	year  int
}

// PrintMessage para poder
func PrintMessage(mensaje string) {
	fmt.Println(mensaje)

}
