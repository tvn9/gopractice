package main

import (
	"fmt"
	"learngo/clockv1.0/screen"
	"time"
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

		// display alarm every 10 seconds
		if sec%10 == 0 {
			alarm()
		}

		// Allow this bock of code to run every second
		time.Sleep(time.Second)
	}
}

// Display alarm
func alarm() {
	screen.ClearScreen()
	alarm := []lettersHolder{
		a, l, a, r, m, qm,
	}
	fmt.Println()
	for i := range alarm[0] {
		for j := range alarm {
			nextAl := alarm[j][i]
			fmt.Print(nextAl, "  ")
		}
		fmt.Println()
	}
	fmt.Println()
}
