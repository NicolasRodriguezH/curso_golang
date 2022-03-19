package reto

import "fmt"

// Pc Pc para actualizar Ram
type Pc struct {
	Ram   int
	Disk  int
	Marca string
}

func (myCompu *Pc) DuplicateRam() {
	myCompu.Ram = myCompu.Ram * 2

	fmt.Println(myCompu)
}
