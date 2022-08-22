package main

import (
	"fmt"
)

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}

type task struct {
	nombre string
	descripcion string
	completado bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(description string) {
	t.descripcion = description
}

func (t *task) actualizarNombre(name string) {
	t.nombre = name
}

func main() {
	t1 := &task {
		nombre: "Esto es lo que tengo en task 1\n",
		descripcion: "Descripcion del task 1\n",
	}

	t2 := &task {
		nombre: "Esto es lo que tengo en task 2\n",
		descripcion: "Descripcion del task 2\n",
	}

	t3 := &task {
		nombre: "Esto es lo que tengo en task 3\n",
		descripcion: "Descripcion del task 3\n",
	}
	// fmt.Printf("%+v\n", t)
	// t.marcarCompleta()
	// t.actualizarNombre("Esto es lo que tengo en actulizar nombre1\n")
	// t.actualizarDescripcion("Esta es la descripción\n")
	// fmt.Printf("%+v\n", t)

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	fmt.Println(lista.tasks[1])
	lista.agregarALista(t3)
	fmt.Println(len(lista.tasks))
	fmt.Println(cap(lista.tasks))
	lista.eliminarDeLista(1)
	fmt.Println(len(lista.tasks))
}