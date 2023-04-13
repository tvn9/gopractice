package main

import (
	"fmt"
	"path"
)

func main() {

	/*
		// Declare two variables
		var dir, file string

		// path.Split() package from standard library
		dir, file = path.Split("css/style.css")
	*/

	// Short hand definition
	dir, file := path.Split("css/styles.css")

	fmt.Printf("Directory: %s file: %s\n", dir, file)
}
