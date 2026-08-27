package main

import "fmt"

func main() {
	// ok := PrintOK("Everything's fine!")
	// PrintOK("All systems operational")
	// PrintOK("Everything's fine!")

	// fmt.Println(ok)

	// demoDeferLIFO()

	fmt.Println(factorial(5))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
}

func PrintOK(msg string) string {
	fmt.Printf("OK - %s\n", msg)
	return msg
}

func demoDeferLIFO() {
	x := 10

	defer fmt.Println(x)

	x = 99

	defer PrintOK("Defer #1")
	defer PrintOK("Defer #2")

	PrintOK("Not deffered!")

}

func factorial(n int) int {
	if n < 0 {
		panic("Negative number not allowed")
	}
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + sum(nums[1:])
}
