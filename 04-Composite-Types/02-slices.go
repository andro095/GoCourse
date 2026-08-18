package main

import (
	"fmt"
	"slices"
)

func slice() {
	// Literal Slice
	var my_slice = []int{10, 20, 30}

	fmt.Println(my_slice)

	my_slice[2] = 40

	fmt.Println(my_slice)

	// Zero Slice (no elements at all)
	var zero_slice []int

	fmt.Println(zero_slice == nil)

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}

	// s := []string{"Hello", "World"}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	// fmt.Println(slices.Equal(x, s)) // Error because different types

}
