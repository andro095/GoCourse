package main

import "fmt"

func structs() {
	type Person struct {
		name string
		age  int
		pet  string
	}

	// var andre Person

	// Struct literal
	// ricardo := Person{}

	// fmt.Println(andre)
	// fmt.Println(ricardo)

	andre := Person{
		"Andre",
		19,
		"Dog",
	}

	ricardo := Person{
		pet:  "Cat",
		age:  30,
		name: "Ricardo",
	}

	andre.name = "André Sebastian"

	fmt.Println(andre.name, ricardo.name)
}
