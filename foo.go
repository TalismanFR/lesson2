package main

import (
	"fmt"
	"os"
)

func main() {
	err := foo()
	if err == nil {
		fmt.Println("not error")
	} else {
		fmt.Printf("Error %v\n", err)
	}
}

func foo() error {
	var err *os.PathError

	return err
}
