package main

import "fmt"

func incrementValue(number int) {
	number++
}

func mutateFirst(my_slice []int) {
	my_slice[0] = 999
}

func mutateSecond(my_slice []int) {
	my_slice[1] = 999
}

func push(s []int) []int {
	s = append(s, 42)
	return s
}

func main() {
	x := 10

	incrementValue(x)

	fmt.Println("Value of x:", x)

	a := []int{1, 2, 3, 4, 5}

	mutateFirst(a)

	fmt.Println("Value of a:", a)

	b := push(a)

	fmt.Println("Value of b:", b)
	fmt.Println("Value of a:", a)

	mutateSecond(b)

	fmt.Println("Value of b:", b)
	fmt.Println("Value of a:", a)
}
