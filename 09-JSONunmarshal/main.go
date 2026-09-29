package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	ClearanceLevel string `json:"clearance_level"`
}

func main() {
	rawBytes, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("couldn't read file", err)
		return
	}
	var users []User
	err = json.Unmarshal(rawBytes, &users)
	if err != nil {
		fmt.Println("Can't parse Json", err)
		return
	}
	fmt.Printf("%+v\n", users)
}

/* trying a stupid Marshall with write file.. only to realize you never actually save data from .json or .txt.. it's from a dbms..
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//create a nested struct for user
type User struct{
	ID int `json:"id"`
	FullName string `json:"full_name"`
	Security Security `json:"security"`
}
type Security struct{
	Clearance string `json:"clearance"`
	RetinaScan bool `json:"retina_scan"`
}

func main(){
	//create a .json file
	file, err:= os.Create("User.txt")
	//take user input for the fields in struct use for loop and bufio(os.stdin)
	for file, u range User{
		reader := bufio.NewReader(os.Stdin)
	}

	//close file
	defer file.Close()
	//read file to show user
	file.Read(User.txt)
}

*/
