package main

import "fmt"

func main() {
	//var MySlice []string

	/*MySlice = append(MySlice, "auth-server")
	MySlice = append(MySlice, "Database-server")
	MySlice = append(MySlice, "Cache-server")
	MySlice = append(MySlice, "Load-server")*/ //line by line mutation

	/*MySlice := []string{
		"Auth-Server",
		"Database-Server",
		"Cache-Server",
	}*/ //compsoite literal/upfront initiation

	/*MySlice[0] = "auth-server"
	MySlice[1] = "Database-Server"
	MySlice[2] = "cache-server"
	MySlice[3] = "Load-balancer"*/

	/*Servers := []string{
		"Auth",
		"Databse",
		"Cache",
	}

	fmt.Println("Servers", len(Servers), cap(Servers), Servers) */ //composite literal feels just like array
	var servers []string

	// We print the Length and Capacity of the empty slice
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	// Add ONE server
	servers = append(servers, "Auth")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	// Add a SECOND server
	servers = append(servers, "Database")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	// Add a THIRD server
	servers = append(servers, "Cache")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	servers = append(servers, "Load")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	servers = append(servers, "Max")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	servers = append(servers, "ch")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	servers = append(servers, "csha")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	servers = append(servers, "jha")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

	servers = append(servers, "anga")
	fmt.Printf("Length: %d | Capacity: %d | Data: %v\n", len(servers), cap(servers), servers)

} //doubling stratagy

/*package main

import "fmt"

func main() {
	//create a standard for lopp
	for i := 0; i < 3; i++ {
		fmt.Println("Spin attempt: ", i)
	}
}
*/
/*package main

import "fmt"

func main() {
	//statusCode := 500

	// FILL THIS IN:
	// Write a switch statement that evaluates 'statusCode'
	switch statusCode := 7800; statusCode {
	// case 200: Print "Connection Successful"
	case 200:
		fmt.Println("Connection Successful")
	// case 404: Print "Not Found"
	case 404:
		fmt.Println("Not Found")
	// case 500: Print "Server Error"
	case 500:
		fmt.Println("Server Error")
	// default: Print "Unknown Code"
	default:
		fmt.Println("Unknown Code")

	}

}
*/
