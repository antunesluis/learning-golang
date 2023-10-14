package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("\nx = %d\ny = %s\nz = %t\n", x, y, z)

	fmt.Printf("\nx = %d", x)
	fmt.Printf("\ny = %s", y)
	fmt.Printf("\nz = %t", z)
}
