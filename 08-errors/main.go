package main

import (
	//"errors"
	"fmt"
	"os"
)

func ReadUsers() error {
	file, err := os.Open("Users.txt")
	if err != nil {
		return err //errors.New("CRITICAL SYSTEM FAILURE: users.txt is missing from the vault")
	} //else {
	//fmt.Println("File exists")
	//}
	fmt.Println("File exists")
	defer file.Close()
	return nil
}

func main() {
	err := ReadUsers()
	if err != nil {
		fmt.Println(err)
	}
}
