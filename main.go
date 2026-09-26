package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(char string) bool {
	var cleanWord string
	char = strings.ToLower(char)
	for i := 0; i < len(char); i++ {
		val := char[i]
		let := val >= 'a' && val <= 'z'
		num := val >= '0' && val <= '9'

		if let || num {
			cleanWord = cleanWord + string(val)
		}
	}

	//declared char is a string and that this function returns a bollean answer
	//letter checking
	left := 0                   //charter in the left is 0
	right := len(cleanWord) - 1 //characters in the right is -1
	for left < right {          //when value of left is greater than right (makes no sense to me.. we are checking for equality..)
		if cleanWord[left] != cleanWord[right] { // are you starting that one is greater and other is less
			return false //and then we are saying they if that is true
			//left++ //then make them equal?
			//right--
			//case false: // if theyre not equal then it isn't palindrome?
			//fmt.Println(char, "isn't a plaindrome")

		}
		left++
		right--

	}
	return true //the bollean answer it has to return on the basis of running the for loop
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the word: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Printf("You typed: %s", input)
	//===================================FOR SHITTY INPUT=============================================
	//var word string                //decalre a string variable
	//fmt.Println("Enter the word:") //ask for the string
	// pulling char out of isPlaindrome to here

	//fmt.Scan(&word) //store the string
	//===================================SHITTY INPUT TOOL ENDS========================================
	if isPalindrome(input) == true {
		fmt.Println(input, "is a palindrome")
	} else {
		fmt.Println(input, "isn't a palindrome")
	}
}
