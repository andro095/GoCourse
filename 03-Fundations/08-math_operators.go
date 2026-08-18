package main

import "fmt"

func mathOperators() {
	a, b := 10, 3

	fmt.Printf("Sum: %d\n", a+b)
	fmt.Printf("Difference: %d\n", a-b)
	fmt.Printf("Product: %d\n", a*b)
	fmt.Printf("Division: %d\n", a/b)
	fmt.Printf("Remainder: %d\n", a%b)

	//increment and decrement
	a++
	b--
	fmt.Printf("Increment A: %d\n", a)
	fmt.Printf("Decrement B: %d\n", b)

}
