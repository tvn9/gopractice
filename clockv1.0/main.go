package main

import (
	"fmt"
	"time"

	"github.com/tvn9/learngo/clockv1.0/screen"
)

func main() {
	// Declare a holder array for holding the digital digits from 0 - 9
	// and the hour, minute, second separator
	type holder [5]string

	// digit zero, one, two, three and so on...
	zero := holder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := holder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := holder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := holder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := holder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := holder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := holder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := holder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := holder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := holder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	sep := holder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// digits array stores all digits in sub array type holder
	digits := [...]holder{zero, one, two, three, four, five, six, seven, eight, nine, sep}
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
