package main

import "fmt"

// Methods agreement
type Speaker interface {
	Speak() string
}

type Person struct {
}

func (Person) Speak() string {
	return "Hello"
}

// type error interface {
// 	Error() string
// }

type MyErr struct{}

func (e *MyErr) Error() string {
	return "Crash!"
}

func returnsTypedNil() error {
	var e *MyErr = nil
	return e
}

func returnRealNil() error {
	return nil
}

func main() {
	err1 := returnsTypedNil()

	fmt.Println(err1 == nil)
	fmt.Printf("%T\n", err1)

	err2 := returnRealNil()
	fmt.Println(err2 == nil)
	fmt.Printf("%T\n", err2)
}
