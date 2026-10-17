package main

import "fmt"

func increment(number *int) {
	*number++
}

type User struct {
	Name string
	Age  int
}

func birthday(u *User) {
	u.Age++
}

func main() {
	x := 10

	increment(&x)
	fmt.Println("Value of x:", x)

	user := User{
		Name: "Ana",
		Age:  20,
	}

	birthday(&user)
	fmt.Println(user)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(num1, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}
