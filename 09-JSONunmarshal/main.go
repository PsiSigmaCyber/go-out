package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	ClearanceLevel string `json:"clearance_level"`
}

func main() {
	rawBytes, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("couldn't read file", err)
		return
	}
	var users []User
	err = json.Unmarshal(rawBytes, &users)
	if err != nil {
		fmt.Println("Can't parse Json", err)
		return
	}
	fmt.Printf("%+v\n", users)
}
