// Retrieving the current working directory
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Path to execuable file
	fmt.Println(ex)

	// Resole the directory of the execuable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path:", exPath)

	// Use EvalSymlinks to get the real path.
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Symlink evaluated:", realPath)

}
