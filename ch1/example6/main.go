// Getting the current process PID
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)

	prc := exec.Command("ps", "-ps", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
