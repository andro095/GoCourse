package main

import "fmt"

func mainSwitch() {
loop: // Label
	for i := 1; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6, 8:
			fmt.Println(i, " is even")
		case 3:
			fmt.Println(i, " is special (3)")
		case 7:
			fmt.Println(i, " is finished")
			break loop
		default:
			fmt.Println(i, " is odd")
		}
	}
}
