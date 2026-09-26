/*
The Problem: Valid Palindrome
Write a function that takes a string and returns true if it reads the same forwards and backwards, and false if it does not.

Your Arsenal (The Mechanics)
Before you write the logic, you only need to know how Go physically handles strings:

Strings are just Lists: If text := "cat", it has 3 slots.

Zero-Indexed: The first letter is at text[0].

The Length Tool: len(text) will return the number 3.
(This means the last letter is always at len(text) - 1).

Equality Check: You can compare two things using == (is equal) or != (is not equal).

The Challenge
Write a completely blank Go file.
You need a main() function to test your words.
You need an isPalindrome function that takes a string and returns a boolean.

How you build the machine inside isPalindrome is entirely up to you. You have to figure out how to compare the front of the word to the back of the word, and how to stop when you know the answer.
*/
//===============================================IGNORE THE COMMENTS HERE!!!!!!!!!!!!!!!!!!!!===================================================================
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

//===========================================================================================================================================
// package main

// import "fmt"

// // The LeetCode Function
// func isPalindrome(word string) bool {
// 	// Pointer 1 starts at index 0 (the beginning)
// 	left := 0

// 	// Pointer 2 starts at the very end of the word
// 	// (len(word) gets the total length. We subtract 1 because indexes start at 0)
// 	right := len(word) - 1

// 	for left < right {
// 		// FILL THIS IN:
// 		// 1. Grab the letter at the 'left' index from the 'word' string.
// 		// Syntax: leftLetter := word[left]
// 		leftLetter := word[left]

// 		// FILL THIS IN:
// 		// 2. Grab the letter at the 'right' index from the 'word' string.
// 		rightLetter := word[right]

// 		// FILL THIS IN:
// 		// 3. If they do NOT match (!=), it's not a palindrome. Return false.
// 		// Syntax: if leftLetter != rightLetter { return false }
// 		if leftLetter != rightLetter {
// 			return false
// 		}

// 		// 4. Move the pointers inward (I wrote this part)
// 		left++
// 		right--
// 	}

// 	// If the loop finishes and didn't return false, it must be a palindrome!
// 	return true
// }

// func main() {
// 	// Let's test it locally before we go to LeetCode
// 	fmt.Println("Is racecar a palindrome?", isPalindrome("racecar"))
// 	fmt.Println("Is hello a palindrome?", isPalindrome("hello"))
// }
