package main

import (
	"fmt"
	"math/rand"
)

func if_else() {
	// should_buy := false
	// debts := false

	// if should_buy {
	// 	if debts {
	// 		fmt.Println("We shouldn't buy.")
	// 	} else {
	// 		fmt.Println("We should buy.")
	// 	}
	// } else {
	// 	fmt.Println("We shouldn't buy.")
	// }

	if n := rand.Intn(10); n == 0 {
		fmt.Println("Zero")
	} else if n > 5 {
		fmt.Println("Bigger than 5")
	} else {
		fmt.Println("Between 1 and 5")
	}

	// switch
}
