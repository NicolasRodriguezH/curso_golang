package main

import (
	"fmt"
)

func main() {
	parImpar(4)

	if verification("Nicolas", "12345") {
		fmt.Println("Acceso permitido")
	} else {
		fmt.Println("Acceso denegado")
	}
}

func parImpar(value int) {
	if value%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
}

func verification(userName, pass string) bool {
	return userName == "Nicolas" && pass == "123"
}
