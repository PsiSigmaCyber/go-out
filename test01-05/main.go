package main

import (
	"fmt"
	"os"
)

type Job struct {
	ID     string
	Status string
}

func (u *Job) Process() {
	if u.Status == "pending" {
		u.Status = "Completed"
	}

}
func main() {
	Joblist := []*Job{

		&Job{ID: "a1", Status: "pending"},
		&Job{ID: "b2", Status: "pending"},
		&Job{ID: "c3", Status: "Failed"},
	}
	for _, u := range Joblist {
		u.Process()

	}
	file, err := os.Create("final_state.txt")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}
	for _, u := range Joblist {
		fmt.Fprintln(file, "ID:", u.ID, ", Status:", u.Status)
	}
	file.Close()

}
