package main

import "fmt"

func zero() {
	var i int
	var s string
	var ok bool
	var nums []int

	fmt.Printf("i=%d (%T)\n", i, i)
	fmt.Printf("s=%q (%T)\n", s, s)
	fmt.Printf("ok=%v (%T)\n", ok, ok)
	fmt.Printf("nums=%v (%T)\n", nums, nums)
}
