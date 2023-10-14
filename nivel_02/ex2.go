package main

import (
	"fmt"
)

var val_1, val_2 int

func main() {
	fmt.Print("Digite um valor: ")
	fmt.Scanln(&val_1)

	fmt.Print("Digite um valor: ")
	fmt.Scanln(&val_2)

	fmt.Println("")
	fmt.Printf("%d == %d : %t\n", val_1, val_2, val_1 == val_2)
	fmt.Printf("%d != %d : %t\n", val_1, val_2, val_1 != val_2)
	fmt.Printf("%d <= %d : %t\n", val_1, val_2, val_1 <= val_2)
	fmt.Printf("%d < %d : %t\n", val_1, val_2, val_1 < val_2)
	fmt.Printf("%d >= %d : %t\n", val_1, val_2, val_1 >= val_2)
	fmt.Printf("%d > %d : %t\n", val_1, val_2, val_1 > val_2)
}
