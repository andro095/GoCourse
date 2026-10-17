package main

import "fmt"

func main() {
	x := 10

	p := &x

	fmt.Println("Value of x:", x)
	fmt.Println("Value of p:", p)
	fmt.Println("Value of *p:", *p)

	*p = 20

	fmt.Println("Value of x:", x)
	fmt.Println("Value of p:", p)
	fmt.Println("Value of *p:", *p)
}
