package main

import (
	"fmt"
)

func main() {
	array := [5]int{4, 2, 3, 4, 5}

	for i := 0; i < len(array); i++ {
		fmt.Printf("array [%d] = %d, tipo: %T \n", i, array[i], i)
	}
}
