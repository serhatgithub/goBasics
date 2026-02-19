package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	table := make(chan int)

	go player("Ping", table, &wg)
	go player("Pong", table, &wg)

	table <- 1
	wg.Wait()
	println("Oyun bitti.")

}

func player(name string, table chan int, wg *sync.WaitGroup) {

	for {
		ball, ok := <-table
		if !ok {
			fmt.Println(name, "Masadan ayrıldı.")
			wg.Done()
			return
		}

		if ball >= 10 {
			fmt.Println("Kazanan:", name)
			close(table)
			wg.Done()
			return
		}

		fmt.Printf("%s vurdu: %d\n", name, ball)

		table <- ball + 1

	}
}
