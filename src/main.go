package main

import (
	pk "curso_golang/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Marca = "Corsa"
	myCar.Modelo = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Nnasjkbasjkab")
}
