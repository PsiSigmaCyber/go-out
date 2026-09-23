package main

import (
	"fmt"
	"os"
)

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
	file, err := os.Create("state.txt")
	fmt.Fprintln(file, "Writing this to hard drive")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}

	for _, u := range user {
		if u.IsActive == true {
			fmt.Fprintln(file, u.ID)

		}

	}
	file.Close()

}
