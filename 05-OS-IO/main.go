package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Shoot data down Pipe 1 (Standard Output)
	fmt.Println("[NORMAL] System boot complete. All engines online.")

	// 2. Shoot data down Pipe 2 (Standard Error)
	fmt.Fprintln(os.Stderr, "[CRITICAL] Memory leak detected in Sector 7!")
}

// 1. go run main.go
// Pipe 1 (Normal) and Pipe 2 (Error) both dump directly onto your monitor.

// 2. go run main.go | clip
// Intercepts ONLY Pipe 1 (Normal) and puts it on the clipboard.
// Pipe 2 (Error) bypasses the clipboard and still hits the monitor.

// 3. go run main.go 2>&1 | clip
// The '2>&1' bends Pipe 2 into Pipe 1.
// Then '| clip' catches that combined stream and puts EVERYTHING on the clipboard.
