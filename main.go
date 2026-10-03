package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	ClearanceLevel string `json:"clearance_level"`
	Age            int    `json:"Age"`
}

func main() {
	// 1. Declare ALL flags first with unique keys
	fileInput := flag.String("file", "data.json", "Path to input JSON")
	fileOutput := flag.String("out", "output.csv", "Path to output CSV")

	// 2. Parse flags ONCE
	flag.Parse()

	// 3. Read the input file (Extract)
	rawData, err := os.ReadFile(*fileInput)
	if err != nil {
		fmt.Println("CRITICAL FAILURE reading input file:", err)
		return
	}

	// 4. Unmarshal into typed structs
	var users []User
	err = json.Unmarshal(rawData, &users)
	if err != nil {
		fmt.Println("CRITICAL FAILURE parsing JSON:", err)
		return
	}

	// 5. Hash map lookup table (Transform)
	badgeLookup := map[string]string{
		"Admin": "Tier-1 (Unrestricted)",
		"User":  "Tier-3 (Standard)",
	}

	// 6. Create the destination file (Load)
	csvFile, err := os.Create(*fileOutput)
	if err != nil {
		fmt.Println("CRITICAL FAILURE creating CSV:", err)
		return
	}
	defer csvFile.Close()

	// 7. Initialize CSV writer and ensure buffer flushes to disk
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// 8. Write header row
	writer.Write([]string{"ID", "Full Name", "Clearance Level", "Badge Status", "Age"})

	// 9. Loop over users, enrich via map, and write each row
	for _, u := range users {
		badge, ok := badgeLookup[u.ClearanceLevel]
		if u.Age >= 18 {
			if !ok {
				badge = "Unassigned"
			}

			// CSV only accepts strings, so convert integer ID to ASCII
			row := []string{
				strconv.Itoa(u.ID),
				u.FullName,
				u.ClearanceLevel,
				badge,
				strconv.Itoa(u.Age),
			}

			// if !ok {
			// 	badge = "Unassigned"
			// }

			// // CSV only accepts strings, so convert integer ID to ASCII
			// row := []string{
			// 	strconv.Itoa(u.ID),
			// 	u.FullName,
			// 	u.ClearanceLevel,
			// 	badge,
			// }

			writer.Write(row)
		}

		fmt.Printf("Success: converted %d records from %s to %s\n", len(users), *fileInput, *fileOutput)
	}
}
