package main

import (
	"fmt"
)

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice_fatiada := append(slice[:3], slice[len(slice)-4:]...)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice [%d] = %d, tipo: %T \n", i, slice_fatiada[i], i)
	}
}
