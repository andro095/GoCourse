package main

import "fmt"

func forRange() {
	evenNumbers := []int{2, 4, 6, 8, 10, 12}

	for _, v := range evenNumbers {
		fmt.Printf("Value: %d\n", v)
	}

	for i := range evenNumbers {
		fmt.Println(evenNumbers[i])
	}
}
