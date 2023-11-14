// Error Interface from Go standard library
/*
type error interface {
	Error() string
}
*/

package main

import "fmt"

type DivError struct {
	a, b int
}

func (d *DivError) Error() string {
	return fmt.Sprintf("Cannot divide by zero: %d / %d", d.a, d.b)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivError{a, b}
	} else {
		return a / b, nil
	}
}

func main() {
	result, err := div(9, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
