package main

import "fmt"

func main() {
	type age int

	var ageOfTeacher age

	ageOfTeacher = 32

	fmt.Printf("%T %v\n", ageOfTeacher, ageOfTeacher)

	agenew := 5
	//str := "de"
	if ageOfTeacher == age(agenew) {

	}
}
