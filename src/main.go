package main

import "fmt"

var a int

type num int

var p num

func main() {

	s := `Este seria un 
	arrow 
	string 
	literal`

	fmt.Println(s)
	fmt.Printf("%T\n\n", s)

	a = 42
	fmt.Println(a)
	fmt.Printf("%d\t%b\t%#x\n\n", a, a, a)

	p = 34
	fmt.Println(p)
	a = int(p)
	fmt.Printf("%T\n", a)
	fmt.Println(a)
}
