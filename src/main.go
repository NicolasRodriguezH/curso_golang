package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := task{
		nombre:      "Completar cursoGO",
		descripcion: "Voy a culminar curso GOYOUTUBR",
	}
	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actualizarNombre("Finalizar el curso de golanggg")
	t.actualizarDescripcion("Completar ese curso cuanto antes")
	fmt.Printf("%+v\n", t)
}
