// Function Literals: 01, Anonymous Function | 02, Closures | 03, Aliases
// Aliases Function
package main

import "fmt"

func calculatePrice(sub float64, discFn Disc) float64 {
	return sub - (sub * discFn(sub))
}

type Disc func(sub float64) float64

func main() {
	msrp := 20.0
	disc := 0.1
	discFn := func(sub float64) float64 {
		// Buy more save more!
		if sub > 100.0 {
			disc += 0.1
		}
		// Max discount
		if disc > 0.3 {
			disc = 0.3
		}
		return disc
	}
	fmt.Println("Price: ", msrp)
	totalPrice := calculatePrice(msrp, discFn)
	fmt.Println("Total price after discount =", totalPrice)
}
