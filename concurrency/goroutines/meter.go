package main

import (
	"fmt"
	"sync"
)

func main() {

	map2 := make(map[string]int)
	map2["ziyaretci"] = 0

	wg := sync.WaitGroup{}
	wg.Add(100)
	lock := sync.Mutex{}

	for i := 0; i < 100; i++ {
		go visiters(map2, &lock, &wg)
	}
	wg.Wait()

	fmt.Println(map2)
}

func visiters(map2 map[string]int, lock *sync.Mutex, wg *sync.WaitGroup) {
	lock.Lock()
	map2["ziyaretci"]++
	lock.Unlock()
	wg.Done()

}
