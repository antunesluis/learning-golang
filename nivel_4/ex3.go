package main

import (
	"fmt"
)

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice = append(slice, 52)
	slice = append(slice, 53, 54, 55)

	aux_slice := []int{56, 57, 58, 59, 60}
	slice = append(slice, aux_slice...)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice [%d] = %d, tipo: %T \n", i, slice[i], i)
	}
}
