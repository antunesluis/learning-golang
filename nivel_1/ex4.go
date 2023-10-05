package main

import (
	"fmt"
)

type meutipo int

var l meutipo

func main() {
	fmt.Printf("valor = %v\n", l)
	fmt.Printf("tipo = %T\n", l)

	l = 42
	fmt.Printf("novo valor = %v\n", l)
}
