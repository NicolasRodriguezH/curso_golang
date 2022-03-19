package main

import "fmt"

func main() {
	// Array
	var arrAy [4]int
	arrAy[0] = 1
	arrAy[1] = 2
	fmt.Println(arrAy, len(arrAy), cap(arrAy))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
