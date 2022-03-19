package main

import (
	"curso_golang/src/reto"
)

func main() {
	var newPc reto.Pc
	newPc.Ram = 16
	newPc.Disk = 221
	newPc.Marca = "Razer"

	//fmt.Println(newPc)
	newPc.DuplicateRam()
}
