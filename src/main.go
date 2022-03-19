package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue & break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es dos")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Hasta aca")
			break
		}
	}
}
