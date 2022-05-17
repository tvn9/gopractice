// Channels example
package main

import "fmt"

func main() {
	// create a channel
	channel := make(chan int)

	// Send message to channel
	go func() {
		channel <- 1
	}()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	// Receive message from channel
	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}
