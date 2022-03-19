package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Marca  string
	Modelo int
}

type carPrivate struct {
	marca  string
	modelo int
}

// PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
