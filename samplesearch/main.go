package main

import (
	"log"
	"os"

	"github.com/tvn9/gopractice/samplesearch/search"
)

// init is call before main
func init() {
	// Set the log output to standard out
	log.SetOutput(os.Stdout)
}

// entry point of the program
func main() {

	// call search function
	search.Run("President")

}
