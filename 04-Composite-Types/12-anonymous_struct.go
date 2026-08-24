package main

import "fmt"

func anonymousStruct() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "Andre"
	person.age = 19
	person.pet = "Dog"

	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Rex",
		kind: "Dog",
	}

	fmt.Println(pet)

}
