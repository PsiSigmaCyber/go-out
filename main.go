package main

import "fmt"

type User struct {
	ID       int
	IsActive bool
}

func main() {

	user := []User{
		User{1, true},
		User{2, false},
		User{3, true},
	}

	for _, u := range user {
		if u.IsActive == true {
			fmt.Println(u.ID)

		}
	}

}
