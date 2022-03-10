package main

import (
	"fmt"
	"runtime"
	"strings"
)

var gorutineCount = 5

func main() {
	for i := 0; i < gorutineCount; i++ {
		go func(th int) {
			for j := 0; j < 5; j++ {
				fmt.Printf(printMessage(th, j))
				runtime.Gosched()
			}
		}(i)
	}

	fmt.Scanln()
}

func printMessage(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "▮",
		strings.Repeat(" ", gorutineCount-in),
		"Thread", in, "Iteration", j, strings.Repeat("<", j))
}
