package main

import "fmt"

type User struct {
	ID    int
	Email string
}

func main() {
	user := User{
		ID:    1,
		Email: "monk@mode.com",
	}
	fmt.Printf("user")
}
