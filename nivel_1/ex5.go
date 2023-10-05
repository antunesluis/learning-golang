package main

import (
	"fmt"
)

type meutipo int

var x meutipo
var y int

func main() {
	fmt.Printf("\nvalor de x = %v", x)
	fmt.Printf("\ntipo de x = %T", x)

	x = 42
	fmt.Printf("\nnovo valor de x= %v\n", x)

	y = int(x)
	fmt.Printf("\ny = %v", y)
	fmt.Printf("\ntipo de y = %T", y)
}
