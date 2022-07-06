package main

import "fmt"

func sayHellos(c chan string, n int) {
	for i := 1; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}

func main() {
	c := make(chan string)
	go sayHellos(c, 5)
	for s := range c {
		fmt.Printf("%s\n", s)
	}
	fmt.Println()

	v, ok := <-c
	fmt.Println("Channel close?", !ok, "value", v)
}
