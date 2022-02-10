package main

import (
	"fmt"
)

func main() {
	ter("123", "456")
	names := []string{"333", "4444"}

	ter(names...)
}

func ter(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}
