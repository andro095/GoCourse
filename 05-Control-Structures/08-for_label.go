package main

import "fmt"

func forLabel() {
	examples := []string{"Hollo", "World", "Go", "Lang"}

outer:
	for _, example := range examples {
		for i, value := range example {
			fmt.Println(i, value, string(value))
			if value == 'o' {
				continue outer
			}
		}
	}

}
