package main

import "fmt"

type Cluster struct {
	Name        string
	ActiveNodes int
}

func main() {

	var MyCluster []*Cluster

	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

	MyCluster = append(MyCluster, &Cluster{Name: "US-East", ActiveNodes: 4})
	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

	MyCluster = append(MyCluster, &Cluster{Name: "Eu-Central", ActiveNodes: 10})
	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

	MyCluster = append(MyCluster, &Cluster{Name: "Ap-South", ActiveNodes: 6})
	fmt.Printf("Length: %d| Capacity: %d| Data: %v\n ", len(MyCluster), cap(MyCluster), MyCluster)

} //Since there are 4 slices, the final capacity will be 4, and the length will also be 4
