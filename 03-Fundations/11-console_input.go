package main

import "fmt"

func consoleInput() {
	var name string
	var age int

	fmt.Println("Insert your name: ")
	fmt.Scan(&name)
	fmt.Println("Insert your age: ")
	fmt.Scan(&age)

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
}
