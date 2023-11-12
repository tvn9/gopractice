package main

import (
	"fmt"
	"time"

	"github.com/tvn9/learngo/clockv1.1/screen"
)

func main() {

	// for loop run the clock until press clt-c to stop the clock
	for {
		// ClearScreen
		screen.ClearScreen()

		// Get the current time
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		// Declare a clock display
		clock := []holder{
			digits[hour/10], digits[hour%10],
			sep,
			digits[min/10], digits[min%10],
			sep,
			digits[sec/10], digits[sec%10],
		}

		// Print the clock on the terminal screen
		fmt.Println()
		for i := range clock[0] {
			for j, d := range clock {
				next := clock[j][i]
				// skip printing the separator every 2 second
				if d == sep && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
