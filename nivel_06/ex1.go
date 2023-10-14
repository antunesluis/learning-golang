package main

import (
	"fmt"
)

func soma(x ...int) int {
	total := 0
	for _, val := range x {
		total += val
	}

	return total
}

func soma2(x []int) int {
	total := 0
	for _, val := range x {
		total += val
	}

	return total
}

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

	soma_total := soma(slice...)
	fmt.Printf("\nSoma total dos valores : %d\n", soma_total)

	soma_total = soma2(slice)
	fmt.Printf("\nSoma total dos valores : %d\n", soma_total)
}
