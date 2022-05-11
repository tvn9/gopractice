// Concurrnt Programming example
package main

import (
	"fmt"
	"time"
)

func count(cnt int) {
	for i := 1; i <= cnt; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	go count(10)
	fmt.Println("wait for goroutine")
	time.Sleep(11 * time.Second)
	fmt.Println("end program")
}
