package main

import "fmt"

func main() {

LOOP:
	for i := 0; i < 4; i++ {
		for k := 0; k < 4; k++ {
			if k == 2 {
				continue
			}
			fmt.Println("i=", i, " k=", k)
			if k == 3 {
				break LOOP
			}
		}
	}
}
