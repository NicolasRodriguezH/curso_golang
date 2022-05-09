package main

import "fmt"

func main() {
	//Imprime los años vividos con for infinito y break
	bi := 1997
	for {
		fmt.Println(bi)
		bi++
		if bi == 2022 {
			break
		}
	}

	//Imprime el modulo de 4 en rango de 10 a 100
	for n := 10; n <= 100; n++ {
		fmt.Printf("Si dividimos %v ebtre 4, el resto (tambien modulo) es %v\n", n, n%4)
	}

	//Imprime la validacion de un if
	x := 21
	if x != 22 {
		fmt.Println("El if funciona correctamente")
		fmt.Println(x)
	}

	//Un switch sin expresion especificada
	switch {
	case false:
		fmt.Println("No deberia imprimir")
	case true:
		fmt.Println("Deberia imprimir")
	}

	//Un switch con expresion, le estaba agregando var en la asignacion a la variable y no es necesario ya que esta en scope a nivel de paquete
	deporteFav := "Fitness"
	switch deporteFav {
	case "Baseball":
		fmt.Println("Ve al estadio")
	case "Workout":
		fmt.Println("En la casa")
	case "Fitness":
		fmt.Println("Go to the fitness")
	}
}
