package main

import (
	"fmt"
	"strings"
)

type User struct {
	name  string
	age   int
	email string
}

func updateProfile(user *User){
	user.name = strings.ToUpper(user.name)
	user.age += 1
}

func main() {
	user := User{
		name:  "Adarsh",
		age:   21,
		email: "adarsh@xyz.com",
	}
	updateProfile(&user)
	fmt.Println("Updated User Profile:", user)
}