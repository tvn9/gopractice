// Part of the day challenge
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	h := 18
	// h := t.Hour()
	m := t.Minute()
	s := t.Second()

	fmt.Println(h, ":", m, ":", s)

	/*
		if h >= 6 && h < 12 {
			fmt.Printf("Good morining, it is now %d:%d:%d\n", h, m, s)
		} else if h >= 12 && h < 18 {
			fmt.Printf("Good afternoon, it is now %d:%d:%d\n", h, m, s)
		} else if h >= 18 && h <= 24 {
			fmt.Printf("Good evening, it is now %d:%d:%d\n", h, m, s)
		} else {
			fmt.Printf("Good night, it is now %d:%d:%d\n", h, m, s)
		}
	*/

	switch {
	case h >= 6 && h < 12:
		fmt.Printf("Good morining, it is now %d:%d:%d\n", h, m, s)
	case h >= 12:
		fmt.Printf("Good afternoon, it is now %d:%d:%d\n", h, m, s)
	case h >= 18:
		fmt.Printf("Good evening, it is now %d:%d:%d\n", h, m, s)
	default:
		fmt.Printf("Good night, it is now %d:%d:%d\n", h, m, s)
	}
}
