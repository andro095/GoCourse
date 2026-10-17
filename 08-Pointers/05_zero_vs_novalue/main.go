package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

type UpdateUser struct {
	Name     *string
	Age      *int
	IsActive *bool
}

func Apply(user *User, update UpdateUser) {
	if update.Name != nil {
		user.Name = *update.Name
	}

	if update.Age != nil {
		user.Age = *update.Age
	}

	if update.IsActive != nil {
		user.IsActive = *update.IsActive
	}
}

func main() {
	user := User{
		Name:     "Andre",
		Age:      20,
		IsActive: true,
	}

	fmt.Println("User: ", user)

	// newName := "John"
	// zeroAge := 0
	inactive := false

	Apply(&user, UpdateUser{
		// Name:     &newName,
		// Age:      &zeroAge,
		IsActive: &inactive,
	})

	fmt.Println("Updated user: ", user)
}
