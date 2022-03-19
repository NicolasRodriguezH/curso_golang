package main

import (
	"fmt"
	"log"
	"strconv"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("21")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	// Declaracion de constantes
	const pi float64 = 3.14
	//La otra forma
	const pi2 = 3.1416

	/* fmt.Println("pi:", pi)
	fmt.Println("pi2", pi2) */

	// Declaracion de variables enteras
	//Tipos
	/* base := 21
	var altura int = 13
	var area int

	fmt.Println(base, altura, area) */

	// Zero values
	/* var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d) */

	// Area cuadrado
	/* const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado :", areaCuadrado) */

	/* x := 10
	y := 50 */

	//Suma
	/* resul := x + y
	fmt.Println("Suma:", resul) */

	//Resta
	/* resul = y - x
	fmt.Println("Resta:", resul) */

	//Multiplicacion
	/* resul = x * y
	fmt.Println("Multiplicacion:", resul) */

	//Division
	/* resul = y / x
	fmt.Println("Division", resul) */

	//Modulo
	/* resul = y % x
	fmt.Println("Modulo", resul) */

	//Incremental
	/* x++
	fmt.Println("Incremental", x) */

	//Decremental
	/* x--
	fmt.Println("Decremental", x) */

	//Declaracion de variables
	/* helloMessage := "Hello"
	wordMessage := "Word" */

	// Println
	/* fmt.Println(helloMessage, wordMessage)
	fmt.Println(helloMessage, wordMessage) */

	//Printf
	/* nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos) */

	// Sprintf
	/* message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)
	*/
	// Tipo datos
	/* fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos) */

	//normalFunction("hola o algo asi o que")
	//tripleArgument(420, 21, "beee")

	/* value := returnValue(2)
	fmt.Println("Value: ", value)

	_, value2 := doubleReturn(2)
	fmt.Println(value2) */

	// For condicional
	/* for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n") */

	// For while
	/* counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	} */

	// For forever
	/* counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	} */

}
