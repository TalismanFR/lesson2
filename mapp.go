package main

import "fmt"

func main() {
	names := make(map[string]string)

	names["student1"] = "Денис"
	names["student2"] = "Кирил"

	fmt.Printf("name %T %v\n", names, names)

	names1 := names

	names1["student3"] = "Тимур"
	fmt.Printf("name %T %v\n", names, names)

	stud := names1["student4"]
	fmt.Printf("%T %s\n", stud, stud)
	fmt.Println(names1)

	if _, ok := names1["student5"]; !ok {
		fmt.Println("student5 отсутствует в мапе ", stud)
	}

}
