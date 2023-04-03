package main

import "fmt"

func main() {
	// Mix of data type and variables and expressions
	orgApple := 2.50        // buy package of 10pcs
	godivaIceScream := 6.70 // 2 pcs
	cheeseCake := 10.99     // 1 pc

	var numOfItem int
	numOrgApple := 10
	numGodivaIceScream := 2
	numCheeseCake := 1

	numOfItem = numOrgApple + numGodivaIceScream + numCheeseCake

	var total = (orgApple * 10) + (godivaIceScream * 2) + cheeseCake

	fmt.Println("Total number of item in shopping cart is", numOfItem)
	fmt.Println("Total shopping $", total)
}
