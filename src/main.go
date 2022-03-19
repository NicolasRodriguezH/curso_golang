package main

import "fmt"

// pc pc para definir computadores
type pc struct {
	ram   int
	disk  int
	marca string
}

func (myCp pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB DISK y es una %s", myCp.ram, myCp.disk, myCp.marca)
}

func main() {
	myPc := pc{ram: 16, disk: 121, marca: "acme"}
	fmt.Println(myPc)
}
