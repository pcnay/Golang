package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// Instanciar echo
	e := echo.New() // Se tiene que importar libnrería "echo"
	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start(":3033"))

}
