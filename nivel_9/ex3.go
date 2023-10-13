package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var nGoRoutines int
	fmt.Println("Digite quantas goroutines deseja criar:")
	fmt.Scanln(&nGoRoutines)

	wg.Add(nGoRoutines)
	go criaGoRoutines(nGoRoutines)
	wg.Wait()
}

func criaGoRoutines(n int) {
	for i := 0; i < n; i++ {
		j := i
		go func(j int) {
			fmt.Println("Sou a go routine numero:", j)
			wg.Done()
		}(j)
	}
}
