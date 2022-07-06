package main

import "fmt"

func main() {
	c := make(chan bool)
	go sayWorld(c, "world!")
	sayHello("Hello")

	// send a signal to c in order to allow sayWold to continue
	c <- true

	// wait to receive another signal on c before exit
	<-c
}

func sayWorld(c chan bool, s string) {
	// wait for data on channel c
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}

func sayHello(s string) {
	fmt.Printf("%s ", "Hello")
}
