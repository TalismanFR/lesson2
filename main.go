package main

import "fmt"

func main() {
	var x [10]int

	x[0] = 2

	xx := x

	xx[0] = 3
	fmt.Println(x, cap(x), xx, cap(xx))

	arr1 := x[0:3]
	arr1 = append(arr1, 3)
	fmt.Printf("arr1 type= %T %v  cap=%v len=%v\n", arr1, arr1, cap(arr1), len(arr1))
	fmt.Println("arr ", x)

	var slice1 = make([]int, 5)
	fmt.Printf("slice1 type= %T %v  cap=%v len=%v\n", slice1, slice1, cap(slice1), len(slice1))
	//slice1 = append(slice1, 4)
	fmt.Printf("slice1 after append type= %T %v  cap=%v len=%v\n", slice1, slice1, cap(slice1), len(slice1))

	slice2 := slice1

	slice2[1] = 4
	slice2 = append(slice2, 5)
	slice2[1] = 3
	fmt.Printf("slice1 and slice2. Slice1: %v  cap=%v len=%v \n Slice2: %v  cap=%v len=%v\n",
		slice1, cap(slice1), len(slice1), slice2, cap(slice2), len(slice2))

	slice3 := make([]int, 0, 10)

	var slice4 []int
	copy(slice4, slice3)

	slice4 = append(slice4, 1)

	fmt.Printf("slice3 %v len==%v cap=%v", slice3, len(slice3), cap(slice3))

}
