// Learn the basic of os.Args
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args reads command line arguments  
	fmt.Printf("%#v\n", os.Args[1:])
}
