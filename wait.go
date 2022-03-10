package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go gor(i, wg)
	}

	wg.Wait()
}

func gor(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 5; j++ {
		fmt.Printf(printMessages(in, j))
		runtime.Gosched()
	}
}

func printMessages(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "▮",
		strings.Repeat(" ", 10-in),
		"Thread", in, "Iteration", j, strings.Repeat("<", j))
}
