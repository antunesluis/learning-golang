package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- i * j
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
}
