package main

import "fmt"

func conversion() {
	var number1 int = 10

	var number2 float64 = 10.5

	total := float64(number1) + number2

	fmt.Println(total)
}
