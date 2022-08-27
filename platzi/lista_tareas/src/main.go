package main

import (
	"fmt"
)

type task struct {
	nombre string
	descripcion string
	completado bool
}

type taskList struct {
	tasks []*task  // Definimos el Slice
}

func (t *taskList) agregarALista(tl *task) { // Receiver
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(description string) {
	t.descripcion = description
}

func (t *task) actualizarNombre(n string) {
	t.nombre = n
}

func (v *taskList) imprimirLista() {
	for _, v := range v.tasks {
		fmt.Println("Nombre: ", v.nombre)
		fmt.Println("Descripción: ", v.descripcion)
	}
}

func (v *taskList) imprimirListaCompleta() {
	for _, v := range v.tasks {
		if v.completado {
			fmt.Println("Nombre: ", v.nombre)
			fmt.Println("Descripción: ", v.descripcion)
		}
	}
}


func main() {
	// t := &task {
	// 	nombre: "task 1\n",
	// 	descripcion: "task 1\n",
	// 	completado: false,
	// }

	// fmt.Printf("Imprimiendo el valor t: %+v\n", t)
	// t.marcarCompleta()
	// t.actualizarNombre("Actulizar nombre1\n")
	// t.actualizarDescripcion("Actualizar descripción\n")
	// fmt.Printf("Imprimiendo el valor de t: %+v\n", t)


	t1 := &task {
		nombre: "task 1 de t1\n",
		descripcion: "Descripcion del task 1\n",
	}

	t2 := &task {
		nombre: "task 2 de t2\n",
		descripcion: "Descripcion del task 2\n",
	}

	t3 := &task {
		nombre: "task 3 de t3\n",
		descripcion: "Descripcion del task 3\n",
	}

	lista := &taskList{ // Slice lista
		tasks: []*task { t1, t2, },
	}

	// fmt.Println("Estoy por aquí => ", lista.tasks[1])
	lista.agregarALista(t3)
	// fmt.Println("Elementos en el slice: ", len(lista.tasks))
	// lista.eliminarDeLista(1)
	// fmt.Println("Elementos en el slice: ", len(lista.tasks))

	// for i := 0; i < len (lista.tasks); i++ {
	// 	fmt.Println("Recorriendo toda la lista:\n\n", lista.tasks[i] )
	// 	fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
	// }
	
	// // Forma mas utilizada.
	// for index, tarea := range lista.tasks {
	// 	fmt.Println("Index: ", index, "Tarea: ", tarea.nombre)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)	
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		fmt.Println("Soy el 5")
	// 		continue
	// 	}
	// 	fmt.Println(i)	
	// }

	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("Tareas completadas")
	lista.imprimirListaCompleta()


	mapaTareas := make(map[string]*taskList)
	mapaTareas["Nestor"] = lista
	mapaTareas["Nestor"].imprimirLista()
	fmt.Println("Imprimiendo con los mapas: ", mapaTareas["Nestor"])
}




/**
Los tres puntos al final de t.task[index + 1:]… (operador ellipsis) es
porque el segundo parámetro del append no es un slice y la función append
recibe un item, con este operador lo que hacemos es decirle a go que tome
ese slice y lo “desempaquete” para que sean muchos parámetros de 1 solo 
item y no un slice.
Ejemplo:

t.tasks = append(t.tasks, task1, task2, task3)
Lo que haría el operador de ellipsis seria pasar un slice a este formato 
de muchos parámetros.
*/