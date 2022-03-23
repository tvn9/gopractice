package main

import (
	"log"
	"runtime"
)

const info = `
Application %s starting.
The binary was build by Go: %s`

func main() {
	// Getting Go's runtime version
	log.Printf(info, "Example", runtime.Version())
}
