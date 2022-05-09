package main

import "fmt"

func main() {

	//Arrays no recomended
	var x [5]int
	x[3] = 21
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))

	//Slices
	//Tipo{elementos} //COMPOSITE LITERAL
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(s[0])

	fmt.Println(s[1:3])
	fmt.Println(s[2:4])
	fmt.Println(len(s))
	fmt.Println(cap(s))

	//Range slice
	for i, v := range s {
		fmt.Println(i, " ", v)
	}

	//Clause for slice
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}
