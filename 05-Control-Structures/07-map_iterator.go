package main

import "fmt"

func map_iterator() {
	my_map := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 20; i++ {
		fmt.Println("Loop:", i)
		for key, value := range my_map {
			fmt.Println(key, value)
		}
	}

}
