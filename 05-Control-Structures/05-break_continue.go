package main

import "fmt"

func breakContinue() {
	n := 0

	for {
		n++
		fmt.Println("Number:", n)

		if n == 100 {
			break
		}

		if n == 50 {
			fmt.Println("Number is 50")
			continue
		}
	}
}
