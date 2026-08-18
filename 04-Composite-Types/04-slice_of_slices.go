package main

import "fmt"

func slice_of_slices() {
	x := []string{"a", "b", "c", "d", "e"}

	y := x[0:2]

	z := x[1:3]

	d := x[2:4]

	// e := x[:]

	e := make([]string, 5)
	copy_e := copy(e, x)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(e, copy_e)

	e[0] = "x"

	fmt.Println(x)
	fmt.Println(e)

}
