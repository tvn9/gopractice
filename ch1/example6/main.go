// Getting the current process PID
package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)

	
}
