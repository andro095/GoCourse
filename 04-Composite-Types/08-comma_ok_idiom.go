package main

import "fmt"

func comma_ok_idiom() {
	my_map := map[string]int{
		"Hello": 1,
		"World": 2,
	}

	value, ok := my_map["Hello"]

	fmt.Println(value, ok)
}
