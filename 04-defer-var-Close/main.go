package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. OPEN THE FILE
	file, _ := os.Create("armour.txt")
	fmt.Println("File pipeline opened.")

	// FILL THIS IN:
	// 2. USE DEFER TO GUARANTEE IT CLOSES
	// Type the word 'defer' followed by the command to close the file
	defer file.Close()

	// 3. DO SOMETHING DANGEROUS
	fmt.Println("Executing dangerous process...")

	// FILL THIS IN:
	// 4. CRASH THE SYSTEM ON PURPOSE
	// Type 'panic("System Meltdown")' to force a fatal crash
	panic("System Meltdown")

}
