package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type calc struct{}

// Receiver func
func (c calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, "No esta soportado")
		return 0
	}
}

func parsear(entrada string) int {
	 operador, _ := strconv.Atoi(entrada)
	 return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada, operador))
}









/**
Struct: tipo de dato que tu mismo creas, el cual puedes agregar métodos y propiedades.

type nombre struct{}

Funcionan como formularios en papel que podría usar.
Recogen diferentes datos y los organizan con diferentes nombres de campos. Cuando inicia
una variable con una nueva struct, es como si fotocopiase un formulario y lo dejase listo
para completarse.

Las Receiver Function, le da la propiedad al struct de tener el método dentro de el, 
es decir cuando instanciemos a la struct vamos a poder llamar a los métodos dentro
de la Receiver Function.

func (nombre struct) nombrefunción (variable string) int{}

referencia : https://www.digitalocean.com/community/tutorials/defining-structs-in-go-es

*/