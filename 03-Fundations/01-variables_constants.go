package main

import "fmt"

const App = "Go Course 2.0" // app = internal, App = exportable

const MaxUsers = 100

func constants() {
	var name string = "Andre"

	lastname := "Rodriguez"

	fmt.Printf("Name: %s %s\n", name, lastname)
	fmt.Println(App)
	fmt.Println(MaxUsers)
}
