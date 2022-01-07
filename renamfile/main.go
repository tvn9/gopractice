package main

import (
	"fmt"
	"log"
	"os"
)

const usage = `
usage: renfile [oldname newname] [enter]

`

func main() {
	args := os.Args[1:]
	if len(os.Args[1:]) != 2 {
		fmt.Print(usage)
		return
	}

	oldName := args[0]
	newName := args[1]

	renameFiles(oldName, newName)
}

// renameFiles
func renameFiles(n1, n2 string) {
	err := os.Rename(n1, n2)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
