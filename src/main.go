package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 42)
	d := (42 != 42)
	e := (42 > 42)
	f := (42 < 42)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
