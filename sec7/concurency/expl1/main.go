// Concurrnt Programming example
package main

import (
	"fmt"
	"time"
)

func count(cnt int) {
	for i := 1; i <= cnt; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	go count(10)
	fmt.Println("wait for goroutine")
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("end program")
}
