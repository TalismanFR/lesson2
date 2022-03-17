package main

import "fmt"

func main() {
	//first()
	//two()
	//tree()
	ext4()
}

func first() {
	values := []string{"1", "2", "3"}

	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}

	fmt.Scanln()
}

func two() {
	for i := 0; i < 3; i++ {
		go foobyval(&i)
	}

	fmt.Scanln()
}

func foobyval(n *int) {
	fmt.Println(*n)
}

func tree() {
	ms := []MyInt{1, 2, 3}
	for _, m := range ms {
		go m.Show()
	}

	fmt.Scanln()
}

type MyInt int

func (mi *MyInt) Show() {
	fmt.Println(*mi)
}

func ext4() {
	for i := 0; i < 3; i++ {
		v := i
		go func() {
			foobyval1(v)
		}()
	}

	{
		v := 2
		_ = v
	}
	fmt.Scanln()
}

func foobyval1(n int) {
	fmt.Println(n)
}
