// channel example
package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan int)
	two := make(chan int)

	for {
		select {
		case o := <-one:
			fmt.Println("one:", o)
		case t := <-two:
			fmt.Println("two:", t)
		case <-time.After(30 * time.Millisecond):
			third("Time out")
			return
		}
	}
}

func third(st string) {
	fmt.Println(st)
}
