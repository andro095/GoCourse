package main

import "fmt"

const taxRate float32 = 0.13

func miniproject() {
	var name, email string
	var subtotal float32

	fmt.Print("Insert your name: ")
	fmt.Scan(&name)

	fmt.Print("Insert your email: ")
	fmt.Scan(&email)

	fmt.Print("Insert subtotal: ")
	fmt.Scan(&subtotal)

	taxes := subtotal * taxRate
	total := subtotal + taxes

	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Taxes: %.2f\n", taxes)
	fmt.Printf("Total: %.2f\n", total)

}
