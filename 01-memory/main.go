package main

import "fmt"

// DAY 1: The Blueprint (Strict Typing)
type User struct {
	ID       int
	Email    string
	IsActive bool
}

// DAY 2 & 3: Pass-by-Value (The Photocopy)
func (u User) ProvePhotocopy() {
	fmt.Printf("Photocopy Address (Inside): %p\n", &u)
}

// DAY 2 & 3: Pass-by-Reference (The Pointer)
func (u *User) Deactivate() {
	u.IsActive = false
	fmt.Printf("Pointer Address (Inside):   %p\n", u)
}

func main() {
	user := User{
		ID:       1,
		Email:    "monk@mode.com",
		IsActive: true,
	}

	fmt.Println("--- MEMORY ADDRESS TEST ---")
	fmt.Printf("Original Address (Main):    %p\n", &user)

	user.ProvePhotocopy()
	user.Deactivate()

	fmt.Println("\n--- MUTATION TEST ---")
	fmt.Printf("Final State: %+v\n", user)
}
