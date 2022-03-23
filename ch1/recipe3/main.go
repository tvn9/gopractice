// Creating a program interface with the flag package
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Custom type need to implement flag.Value
type Value []string

func (s *Value) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *Value) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {
	// Extracting flag values with methods returning pointers
	retry := flag.Int("retry", -1, "Defines max retry count")

	// Read the flag using the XXXVar function.
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")

	var arr Value
	flag.Var(&arr, "array", "Input array to iterate through.")

	// Execute the flag.Parse function
	flag.Parse()

	// Sample logic not related to flags
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}
