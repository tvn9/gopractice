// Path separator example
package main

import (
	"fmt"
	"path"
)

func main() {
	// declare dir and file variables
	var dir, file string

	// path.Split function returns two strings value dir, and file name
	dir, file = path.Split("css/styles.css")
	fmt.Println(dir, file)

	var dir1, file1 string
	// path.Split function allow user to obmit return values
	// in this example, dir is obmited
	// dir1, file1 = path.Split("videos/mnteverest.css")
	// fmt.Println(dir1, file1)
	_, file1 = path.Split("videos/mnteverest.css")
	fmt.Println(dir1, file1)

}
