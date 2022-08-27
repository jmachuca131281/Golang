package main

import (
	"fmt"
)

func cambiarValor( a int ) { // Crea la copia de x y se lapasamos a a
	fmt.Println("\nDirección de a: ", &a)
	fmt.Println("Valor de a: ", a)
	a = 36
	fmt.Println("Ultimo valor de a: ", a)
}

func main() {
	x :=  25
	fmt.Println("Valor variable x: ", x)
	fmt.Println("Direccíón de x: ", &x)
	y := &x  // Se asigna el valor o referencia de x
	fmt.Println("\nValor de y que se le pasa dirección de x: ", y)
	fmt.Println("Valor de y: ", *y)
	cambiarValor(x) // Pasando el argumento por referencia. Se envia el 25
}