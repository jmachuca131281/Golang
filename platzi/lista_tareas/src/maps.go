package main

import (
	"fmt"
)

func main(){
	m1 := make(map[string]int) // Respetar la configuración incial
	m1["a"] = 8
	fmt.Println(m1)      // Imprime el mapa.
	fmt.Println(m1["a"]) // Imprime el valor.
}