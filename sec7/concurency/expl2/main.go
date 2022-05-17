// Goroutine with closure function
package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	wait := func(ms time.Duration) {
		time.Sleep(ms * time.Millisecond)
		counter += 1
	}

	fmt.Println("Launching goroutines")
	go wait(300)
	go wait(1200)
	go wait(1000)

	fmt.Println("Launched.		Counter =", counter)
	time.Sleep(1300 * time.Millisecond)
	fmt.Println("Waited 1300ms. Counter =", counter)
}
