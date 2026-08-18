package main

import "fmt"

func slicesFuncs() {
	// var my_slice []int
	// fmt.Println(my_slice == nil)

	// my_slice := []int{1, 2, 3}

	// fmt.Println(len(my_slice))

	// my_slice = append(my_slice, 10, 12, 30, 15, 40, 50, 12, 13)

	// fmt.Println(len(my_slice))
	// fmt.Println(my_slice)

	// fmt.Println(my_slice, len(my_slice), cap(my_slice))

	// make_slice := make([]int, 5)
	make_slice := make([]int, 0, 10)

	fmt.Println(make_slice)

	make_slice = append(make_slice, 10, 20, 30)

	fmt.Println(make_slice)
	fmt.Println(len(make_slice))
	fmt.Println(cap(make_slice))

	clear(make_slice)

	fmt.Println(make_slice)
	fmt.Println(len(make_slice))
	fmt.Println(cap(make_slice))

}
