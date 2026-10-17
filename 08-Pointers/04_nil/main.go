package main

import "fmt"

type User struct {
	Name string
}

type Config struct {
	Host string
	Port int
}

func printName(user *User) {

	if user == nil {
		fmt.Println("User is nil")
		return
	}
	fmt.Println("User name: ", user.Name)
}

func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,
	}
}

func main() {
	var u *User
	printName(u)

	u = &User{Name: "Andre"}
	printName(u)

	config := NewConfig()
	fmt.Println("configuration: ", config.Host, config.Port)

}
