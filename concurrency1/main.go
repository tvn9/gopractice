package main

import (
	"fmt"
	"time"
)

func sayWorld(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func sayHello(s string) {
	fmt.Print(s, " ")
}

func main() {
	go sayWorld("World!")
	sayHello("Hello")
	time.Sleep(3 * time.Second)
}
