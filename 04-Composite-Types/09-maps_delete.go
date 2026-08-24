package main

import "fmt"

func maps_delete() {
	my_map := map[string]int{
		"Hello": 1,
		"World": 2,
	}

	fmt.Println(my_map)

	delete(my_map, "Hello")
	clear(my_map)

	fmt.Println(my_map)
	fmt.Println(len(my_map))
}
