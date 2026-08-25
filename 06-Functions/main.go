package main

import "fmt"

func main() {
	ok := PrintOK("Everything's fine!")
	PrintOK("All systems operational")
	PrintOK("Everything's fine!")

	fmt.Println(ok)
}

func PrintOK(msg string) string {
	fmt.Printf("OK - %s\n", msg)
	return msg
}
