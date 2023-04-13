package main

import "fmt"

func main() {

	// Veriables rules
	// In Go, variable must be declareed before being used
	// Can do: start with letter, _variable, or unicode like こんにちは, upcase letter use for Export
	// Can not do: start with number 3speed, !letter, let!ter, var
	// synctax: var name int

	// declare a variable without define a value
	// defaul value for this emply variable is 0
	var number int

	fmt.Println(number) // output => 0 for numeric type

}
