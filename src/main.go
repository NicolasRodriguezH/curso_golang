package main

import "fmt"

func main() {
	//Slices
	//Tipo{elementos} //COMPOSITE LITERAL
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	s = append(s, 6, 7, 8)
	fmt.Println(s)

	a := []int{9, 10}
	s = append(s, a...)
	fmt.Println(s)

	//Eliminar elementeos del slice
	s = append(s[:5], s[7:]...)
	fmt.Println(s)
}
