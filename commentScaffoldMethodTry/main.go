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

/*package main

import "fmt"

// ---------------------------------------------------------
// WORKER 1: THE STANDARD FUNCTION (Stands outside main)
// It takes NO inputs: ()
// It returns TWO outputs: (string, string)
// ---------------------------------------------------------
func CheckSystem() (string, string) {

	// FILL THIS IN:
	// Type the word 'return' followed by "Online" and "US-East" separated by a comma
	return "Online", "US-East"

}

// ---------------------------------------------------------
// WORKER 2: THE MAIN ENGINE
// ---------------------------------------------------------
func main() {

	// FILL THIS IN:
	// Create two variables (status, region) and assign them the execution of CheckSystem()
	// It should look like: variable1, variable2 := FunctionName()
	status, region := CheckSystem()

	// PROVE IT WORKED (I wrote this for you)
	fmt.Println("System Status:", status)
	fmt.Println("System Region:", region)
}
*/
