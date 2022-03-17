package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters = map[int]int{}
	mut := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, i int, mtx *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mut.Lock()
				counters[i*10+j]++
				mut.Unlock()
			}
		}(counters, i, mut)
	}
	fmt.Scanln()

	mut.Lock()
	fmt.Println(counters)
	defer mut.Unlock()
}
