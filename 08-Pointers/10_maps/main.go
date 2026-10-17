package main

import "fmt"

func put(my_map map[string]int, key string, value int) {
	my_map[key] = value
}

func main() {
	var my_map map[string]int

	fmt.Println("leer nil map:", my_map["x"])

	my_map = make(map[string]int)

	fmt.Println("Empty map:", my_map == nil)

	put(my_map, "a", 1)
	put(my_map, "b", 2)

	fmt.Println(my_map)
}
