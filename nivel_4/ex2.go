package main

import (
	"fmt"
)

func main() {
	slice := []int{}

	fmt.Println("Digite um valor a ser adiconado (0 para parar): ")
	for {
		var val int

		fmt.Scanln(&val)
		if val == 0 {
			break
		}
		slice = append(slice, val)
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice [%d] = %d, tipo: %T \n", i, slice[i], i)
	}
}
