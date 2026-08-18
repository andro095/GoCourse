package main

import "fmt"

func arrays() {
	// var number_list [3]int
	// Literal array
	// var number_list = [3]int{10, 20, 30}
	// Simplified literal array
	var number_list = [...]int{10, 20, 30}

	fmt.Println(number_list)

	number_list[0] = 50

	fmt.Println(number_list[0])

	fmt.Println(len(number_list))

}
