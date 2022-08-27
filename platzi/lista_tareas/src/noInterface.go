package main

import (
	"fmt"
)

type animal interface {
	mover() string
}

type perro  struct{}
type pez    struct{}
type pajaro struct{}


func (perro) mover() string {
	return "Soy un perro y camino"
}

func (pez) mover() string {
	return "Soy un pez y estoy nadando"
}

func (pajaro) mover() string {
	return "Soy pajaro y estoy volando"
}

// func moverPerro (p perro){
// 	fmt.Println(p.mover())
// }

// func moverPez (p pez){
// 	fmt.Println(p.nadar())
// }

// func moverPajaro (p pajaro){
// 	fmt.Println(p.volar())
// }

func moverAnimal(a animal){
	fmt.Println(a.mover())
}

func main() {
	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaro{}
	moverAnimal(pa)

	fmt.Println("")

}