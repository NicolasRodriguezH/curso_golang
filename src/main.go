package main

import "fmt"

type car struct {
	marca  string
	modelo int
}

func main() {
	myCar := car{marca: "Ford", modelo: 1997}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.marca = "Mustang"
	otherCar.modelo = 2022
	fmt.Println(otherCar)
}
