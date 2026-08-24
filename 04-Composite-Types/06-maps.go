package main

import "fmt"

func maps_example() {
	// var nilMap map[string]int

	// nilMap["Hello"] = 1

	totalWins := map[string]int{}

	// fmt.Println(totalWins == nil)

	totalWins["Blue Jays"] = 1
	totalWins["Red Sox"] = 2
	totalWins["Yankees"] = 5

	fmt.Println(totalWins["Yankees"])
}
