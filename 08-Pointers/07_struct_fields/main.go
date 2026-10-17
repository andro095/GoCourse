package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Profile struct {
	Bio string
}

type User struct {
	Name    string
	Addr    Address
	Profile *Profile // Opcional
}

func main() {
	user := User{
		Name: "Ana",
		Addr: Address{
			City:  "Buenos Aires",
			State: "CABA",
		},
	}

	if user.Profile == nil {
		fmt.Println("No Profile")
	}

	user.Profile = &Profile{Bio: "Fullstack Developer"}

	fmt.Println("Bio: ", user.Profile.Bio)

	fmt.Println(user)
}
