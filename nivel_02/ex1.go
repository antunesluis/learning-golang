package main

import (
	"fmt"
)

var val int

func main() {

	fmt.Print("Digite um valor: ")
	fmt.Scanln(&val)

	fmt.Printf("octal = %o\n", val)
	fmt.Printf("decimal = %d\n", val)
	fmt.Printf("hexadecimal = %#x\n", val)

}
