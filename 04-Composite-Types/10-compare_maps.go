package main

import (
	"fmt"
	"maps"
)

func compare_maps() {
	my_map_a := map[string]int{
		"Hello": 1,
		"World": 2,
	}

	my_map_b := map[string]int{
		"Hello": 1,
		"World": 2,
	}

	fmt.Println(maps.Equal(my_map_a, my_map_b))
}
