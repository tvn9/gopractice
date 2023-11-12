package main

import (
	"fmt"
	"time"

	"github.com/tvn9/learngo/clockv1.3/screen"
)

func main() {

	// for loop run the clock until press clt-c to stop the clock
	for {
		// ClearScreen
		screen.ClearScreen()

		// Get the current time
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		// Declare a clock
		clock := [...]holder{ // {[0:],[0:],sep,[0:],[0:],sep,[0:],[0:]}
			digits[hour/10], digits[hour%10],
			sep,
			digits[min/10], digits[min%10],
			sep,
			digits[sec/10], digits[sec%10],
		}

		alramTime := sec%10 == 0

		if alramTime {
			clock = alarm
		}

		// Print the clock on the terminal screen
		fmt.Println()
		for i := range clock[0] {
			for j, d := range clock {
				nextClock := clock[j][i]
				// skip printing the separator every 2 second
				if d == sep && sec%2 == 0 {
					nextClock = "   "
				}
				fmt.Print(nextClock, "  ")
			}
			fmt.Println()
		}
		fmt.Println()

		// Allow this bock of code to run every second
		time.Sleep(time.Second)
	}
}
