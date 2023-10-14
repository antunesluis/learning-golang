package main

import (
	"fmt"
)

var val int

func main() {

	fmt.Print("Digite um valor: ")
	fmt.Scanln(&val)

	fmt.Printf("\nRepresentacoes de %d:", val)
	fmt.Printf("\n\tbinário = %b", val)
	fmt.Printf("\n\tdecimal = %d", val)
	fmt.Printf("\n\thexadecimal = %#x\n", val)

	alt := val << 1

	fmt.Printf("\nRepresentacoes de %d (%d << 1):", alt, val)
	fmt.Printf("\n\tbinário = %b", alt)
	fmt.Printf("\n\tdecimal = %d", alt)
	fmt.Printf("\n\thexadecimal = %#x\n", alt)

}
