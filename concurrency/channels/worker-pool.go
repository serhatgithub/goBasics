package main

import (
	"fmt"
	"sync"
)

func main() {
	Jobs := make(chan int, 100)
	Results := make(chan int, 100)
	wg := sync.WaitGroup{}

	for w := 1; w <= 100; w++ {
		wg.Add(1)
		go worker(w, Jobs, Results, &wg)
	}

	for j := 1; j <= 100; j++ {
		Jobs <- j
	}
	close(Jobs)

	wg.Wait()

	close(Results)

	for r := range Results {
		fmt.Println("Sonuç:", r)
	}

}

func worker(w int, Jobs <-chan int, Results chan<- int, wg *sync.WaitGroup) {
	for j := range Jobs {
		fmt.Printf("işci: %d, Sayi %d\n", w, j)
		wg.Done()
		Results <- j * j

	}
}
