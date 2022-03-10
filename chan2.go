package main

import (
	"fmt"
)

func main() {
	in := make(chan int, 3)
	go func(out chan<- int) {
		for i := 0; i < 10; i++ {
			fmt.Println("Write channel ", i)
			out <- i
			fmt.Println("Write channel completed", i)
		}
		close(out)
	}(in)

	val, ok := <-in
	fmt.Println(val, ok)

	for i := range in {
		fmt.Println("\t val ", i)
	}

}
