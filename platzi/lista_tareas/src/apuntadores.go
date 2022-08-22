package main

import (
	"fmt"
)

func cambiarValor(a int) {
	fmt.Println("Dirección de a: ", &a)
	a = 36
}

func main() {
	x :=  25
	fmt.Println(x)
	fmt.Println("Direccíón de b: ", &x)
	// y := &x
	// fmt.Println("Direccion y: ", *y)
}