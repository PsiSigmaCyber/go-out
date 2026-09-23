package main

import "fmt"

type Cluster struct {
	Name        string
	ActiveNodes int
}
type User struct {
	Email string
}

/*type User struct {
	ID       int
	IsActive bool
}*/ //for if conditional

func main() {

	var MyCluster []*Cluster

	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

	MyCluster = append(MyCluster, &Cluster{Name: "US-East", ActiveNodes: 4})
	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

	MyCluster = append(MyCluster, &Cluster{Name: "Eu-Central", ActiveNodes: 10})
	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

	MyCluster = append(MyCluster, &Cluster{Name: "Ap-South", ActiveNodes: 6})
	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

	//var persons []*User

	/*persons = append(persons, &User{"cto@tech.com"})
	persons = append(persons, &User{"cfo@tech.com"})
	persons = append(persons, &User{"ceo@tech.com"})
	persons = append(persons, &User{"cso@tech.com"})*/

	persons := []User{
		User{"cto@tech.com"},
		User{"cfo@tech.com"},
		User{"ceo@tech.com"},
		User{"cso@tech.com"},
	}

	for _, persons := range persons {
		fmt.Printf("Email: %v\n", persons.Email)
		//fmt.Printf("length: %d\n", len(persons.Email))
		//fmt.Println(len(persons)) //done to genuinely crash it, variable shadoing..
	}
	/*user := []User{
		User{1, true},
		User{2, false},
		User{3, true},
	}

	for _, u := range user {
		if u.IsActive == true {
			fmt.Println(u.ID)

		}
	}*/ //if conditional..

} //Since there are 4 slices, the final capacity will be 4, and the length will also be 4
