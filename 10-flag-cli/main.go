package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("file", "data.json", "Path to json")
	flag.Parse()

	data, err := os.ReadFile(*filePath)

	// THE TRAP: Never silence this.
	if err != nil {
		fmt.Println("CRITICAL FAILURE:", err)
		return
	}

	fmt.Println(string(data))
}

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	//filePath := flag.String("file", "data.json", "Path to json file")
// 	//flag.Parse()
// 	//fmt.Printf("getting file location %s\n", *filePath)
// 	filePath := flag.String("file", "data.json", "Path to json")
// 	flag.Parse()

// 	data, _ := os.ReadFile(*filePath)
// 	fmt.Println(string(data))
// }

// // data, err := os.ReadFile("data.json" /**filePath*/)
// // if err != nil {
// // 	fmt.Println("Crirical failure: ", err)
// // 	return
// // }
// // fmt.Println("\nfile content: ")
// // fmt.Println(string(data))

// //}
