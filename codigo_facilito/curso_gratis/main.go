package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strconv"
)

func main() {
	// edad := 22
	// edad2 := "22"

	// String to int
	// edad_str := strconv.Itoa(edad)
	// fmt.Println("Mi edad es: ", edad_str)

	// Convirtiendo e imprimiendo
	// fmt.Println("De int to string", strconv.Itoa(edad))

	// Convert string to int
	// edad_int, _ := strconv.Atoi(edad2)
	// fmt.Println("De string to int", edad_int)

	// Imprimiendo los diferentes datos:
	// variableInt := 22
	// fmt.Printf("Imprimiendo el int %d:\n ", variableInt) 

	// variableString := "22.000"
	// fmt.Printf("Imprimiendo el string %v:\n ", variableString) 

	// variableBool := false
	// fmt.Printf("Imprimiendo el boolean %t:\n ", variableBool) 

	// var v int
	// fmt.Println("Ingresa la edad: ")
	// fmt.Scanf("%v\n", &v)
	// fmt.Printf("Edad es:%d", v)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Ingresa el nombre:")
	// text, e := reader.ReadString('\n') // Las "" declaran un string y '' un caracter
	// if e != nil {
	// 	fmt.Println(e)
	// }else{
	// 	fmt.Println("Se salvo: "+ text)
	// }

	i := 0
	for i<10 {
		fmt.Println(i)
		i++
	}
}

// Golang Pringf verbs.
// http://xahlee.info/golang/golang_printf_verbs.html