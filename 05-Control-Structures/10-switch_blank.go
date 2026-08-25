package main

import "fmt"

func switchBlank() {
	a := 4

	switch {
	case a < 2:
		fmt.Println("a is less than 2")
	case a < 5:
		fmt.Println("a is less than 5")
	default:
		fmt.Println("a is greater than or equal to 5")
	}
}
