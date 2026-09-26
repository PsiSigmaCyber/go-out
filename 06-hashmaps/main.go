/*I will give you a sentence: "the dog chased the cat and the cat ran"
You need to write a Go program that uses a map[string]int to count how many times each word appears, and print it out. (e.g., the: 3, cat: 2, dog: 1).*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "the dog chased the cat and the cat ran"
	words := strings.Fields(text) // This chops the text into a list of words

	// FILL THIS IN:
	// 1. Create a map named 'freq' where the Key is a string, and Value is an int.
	// Syntax: freq := make(map[string]int)
	freq := make(map[string]int)

	// 2. Loop over the words and count them
	for _, word := range words {
		// FILL THIS IN:
		// Add 1 to the count for the current 'word'
		// Syntax: freq[word] = freq[word] + 1
		freq[word] = freq[word] + 1
	}

	fmt.Println(freq)
}

// package main

// import "fmt"

// //create a struct for the sentence
// type sentence struct{
// 	words int //why creating a struct for a single thing? could just do with a declared var i think..
// }
// func main(){
// 	char := make(map[string]sentence) // make a map for the sentence
// 	char["read"]=sentence{words:the dog chased the cat and the cat ran} // put the value to be lost in the RAM
// 	// nope, dunno

//}
//-----------------------------------------------------SAME THING DIFFERENT WAY MAYBE-----------------------------------------------------
// package main

// import "fmt"

// type words struct{
// 	//call appendwords func to make integer variables using a for loop

// }

// func main(){
// 	//print question to input the sentence
// 	//scan the sentence
// 	// make a map for the struct
// 	// print the map?
// }
// func appendWords(){ //create function to append the sentence into words and put them in a struct
// }
//-------------------------------------------------------------------------------------------------------------------------------------
/*Define a struct named Session that tracks two things: Username (string) and IsAdmin (boolean).

In main(), create a Map that links a Session ID (a string, e.g., "Session_44") to a Session struct.

Load one user into the Map: Session ID "Session_44" belongs to "Alice", and her IsAdmin status is currently false.

Write a completely separate helper function called escalatePrivilege. It needs to accept your Map and a Session ID as arguments.

Inside that helper function, change the user's IsAdmin status to true. (Remember the Photocopy Trap).

Back in main(), print the Map to prove Alice is now an admin.*/

// package main

// import "fmt"

// //define struct
// type Session struct {
// 	Username string
// 	IsAdmin  bool
// }

// func main() {
// 	//create map that links a Session ID (a string, e.g., "Session_44") to a Session struct.
// 	Session_ID := make(map[string]Session)
// 	Session_ID["Session_44"] = Session{Username: "Alice", IsAdmin: false}
// 	escalatePrivilege(Session_ID, "Session_44")
// 	//print the Map to prove Alice is now an admin
// 	fmt.Println(Session_ID["Session_44"].Username, "is an admin")

// }
// func escalatePrivilege(Session_ID map[string]Session, Session_44 string) {
// 	//needs to accept your Map and a Session ID as arguments.
// 	target := Session_ID[Session_44]

// 	//change the user's IsAdmin status to true
// 	target.IsAdmin = true
// 	Session_ID[Session_44] = target
// }

//------------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// // 1. The Helper Function
// // It takes a map as an argument.
// func nukeJob(targetMap map[string]string, jobID string) {
// 	delete(targetMap, jobID)
// 	fmt.Println("[NETWORK] Target destroyed:", jobID)
// }

// func main() {
// 	// 2. The Main Map
// 	activeJobs := make(map[string]string)
// 	activeJobs["Job_777"] = "Running"
// 	activeJobs["Job_999"] = "Running"

// 	fmt.Println("Before Nuke:", activeJobs)

// 	// FILL THIS IN:
// 	// 3. Call the helper function.
// 	// Pass it your 'activeJobs' map, and tell it to destroy "Job_777".
// 	// Syntax: nukeJob(yourMapName, "TargetString")
// 	nukeJob(activeJobs, "Job_777")

// 	// 4. Print the map again to see if the original was permanently altered
// 	fmt.Println("After Nuke:", activeJobs)
// }

//---------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type Job struct {
// 	Status  string
// 	Retries int
//}

//func main() {
//registry := make(map[string]Job)
//registry["Job_101"] = Job{Status: "Failed", Retries: 3}

// 1. THE TRAP (Try uncommenting this line first, run it, and read the error)
//registry["Job_101"].Retries = 4

// 2. THE FIX
// FILL THIS IN:
// Step A: Pull out the photocopy: target := registry["Job_101"]
//target := registry["Job_101"]
// Step B: Update the photocopy: target.Retries = 4
//target.Retries = 4
// Step C: Overwrite the map: registry["Job_101"] = target
//registry["Job_101"] = target

//fmt.Printf("Final Retries for Job 101: %d\n", registry["Job_101"].Retries)
//}

//-----------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// // 1. The Blueprint
// type Job struct {
// 	Status  string
// 	Retries int
// }

// func main() {
// 	// FILL THIS IN:
// 	// 2. Ignite the Map
// 	registry := make(map[string]Job)
// 	// The Key is a string. The Value is the 'Job' struct.
// 	// Syntax: registry := make(map[string]Job)

// 	// FILL THIS IN:
// 	// 3. Load a full Struct into the Map
// 	registry["Job_101"] = Job{Status: "Failed", Retries: 3}
// 	// Assign a new Job to "Job_101" with Status "Failed" and Retries 3
// 	// Syntax: registry["Job_101"] = Job{Status: "Failed", Retries: 3}

// 	// 4. Retrieve and Print
// 	target := registry["Job_101"]
// 	fmt.Printf("Job 101 Status: %s | Retries: %d\n", target.Status, target.Retries)
// }

//-----------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	registry := make(map[string]string)
// 	registry["Job_101"] = "Completed"

// 	// FILL THIS IN:
// 	// 1. Permanently delete the job from the map
// 	// Syntax: delete(mapName, "Key")
// 	delete(registry, "Job_101")

// 	// FILL THIS IN:
// 	// 2. Use the Comma-Ok Radar to check if it's still there
// 	// Syntax: status, exists := registry["Job_101"]
// 	status, exists := registry["Job_101"]

// 	// 3. Print the results (I wrote this)
// 	fmt.Println("Status:", status)
// 	fmt.Println("Does it exist?", exists)
//}

//-------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	registry := make(map[string]string)

// 	// 1. OVERFLOW THE BUCKET
// 	// Loop 15 times to force Go to create multiple memory buckets
// 	for i := 1; i <= 15; i++ {
// 		key := fmt.Sprintf("Job_%d", i) // Formats the string like "Job_1", "Job_2", etc.
// 		registry[key] = "Active"
// 	}

// 	// 2. WATCH THE CHAOS
// 	fmt.Println("Reading the Map:")
// 	for key := range registry {
// 		fmt.Printf("%s ", key) // Print on one line with a space
// 	}
// 	fmt.Println() // Add a final new line
//}
//-------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	registry := make(map[string]string)

// 	// Load the web
// 	registry["Job_1"] = "Completed"
// 	registry["Job_2"] = "Active"
// 	registry["Job_3"] = "Failed"

// 	// FILL THIS IN:
// 	// Write a for-range loop that pulls the 'key' and 'val' from registry.
// 	for key, val := range registry {
// 		// Syntax: for key, val := range registry { ... }
// 		// Inside the loop, print: key, "->", val
// 		fmt.Println(key, "->", val)
// 	}

//}
//---------------------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	// FILL THIS IN:
// 	// 1. Ignite the Map (The Warp Tunnel)
// 	registry := make(map[string]string)
// 	// Create a map named 'registry' where the Key is a string, and the Value is a string.
// 	// Syntax: registry := make(map[string]string)

// 	// FILL THIS IN:
// 	// 2. Load the data.
// 	// Assign the string "Completed" to the key "Job_101".
// 	registry["Job_101"] = "Completed"
// 	// Assign the string "Active" to the key "Job_8492".
// 	registry["Job_8492"] = "Active"
// 	// Syntax: registry["Job_101"] = "Completed"

//		// 3. Teleport directly to Job_8492 and print its status (I wrote this for you)
//		fmt.Println("Status of Job 8492:", registry["Job_8492"])
//	}
//----------------------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	registry := make(map[string]string)
// 	registry["Job_101"] = "Completed"

// 	// 1. PROVE WE ARE USING RAW MEMORY (Print the pointer address)
// 	fmt.Printf("The Map is physically located at RAM address: %p\n", registry)

// 	// 2. THE RADAR PING (val, ok)
// 	// We are asking for Job_999.
// 	// 'status' gets the data. 'exists' gets a true/false boolean.
// 	status, exists := registry["Job_999"]

// 	fmt.Println("Status of 999:", status)
// 	fmt.Println("Did it actually exist in memory?:", exists)
// }
