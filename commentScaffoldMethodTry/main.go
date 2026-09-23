package main

import "fmt"

//Create a struct for numbers, NO, i don't need to create a struct, i can make a sice named num and add numbers into it
func main() {
	//Create a slaice with numbers in the slice

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// create a loop to go through the slice
	for _, u := range num {
		// if the numver is divisible by 2 then print
		if u%2 == 0 {
			fmt.Println(u, " is even")

		} else {
			fmt.Println(u, " is odd")
		}

	}

}
