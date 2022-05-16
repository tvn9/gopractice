// Function Literals: 01, Anonymous Function | 02, Closures | 03, Aliases
// Closures Function
package main

import "fmt"

// calculatePrice func accept Closures func that return a float64
func calculatePrice(sub float64, discFn func(sub float64) float64) float64 {
	return sub - (sub * discFn(sub))
}

func main() {
	amntSpent := 200.0
	discAllow := 100.0
	discount := 0.3
	// discFun is a Closures
	discFn := func(sub float64) float64 {
		var amnt float64
		// Buy more save more!
		if sub > discAllow {
			amnt += discount
		}
		// Max discount
		if amnt > 0.3 {
			amnt = 0.3
		}
		fmt.Println("Discount amount", 100*amnt) // just to make sure the correct discount
		return amnt
	}
	disc := discFn(discount)               // => 0 (this is because Closures)
	fmt.Println("Discount amount =", disc) // print out the discount amount
	fmt.Println("Price: ", amntSpent)
	totalPrice := calculatePrice(amntSpent, discFn)
	fmt.Println("Total price after discount =", totalPrice)
}
