// Defer co example
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 0, 30)
	bytes, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		// file will close with deferred call
		return
	}
	fmt.Printf("%c\n", bytes)

}
