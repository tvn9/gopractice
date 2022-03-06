package main

import "fmt"

func main() {
	var names []string
	var distances []float64
	var data []byte
	var ratios []float64
	var alives []bool

	fmt.Printf("names     - type: %T, len: %d, is equal to nil? %t\n", names, len(names), names == nil)
	fmt.Printf("distances - type: %T, len: %d, is equal to nil? %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data      - type: %T, len: %d, is equal to nil? %t\n", data, len(data), data == nil)
	fmt.Printf("ratios    - type: %T, len: %d, is equal to nil? %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives    - type: %T, len: %d, is equal to nil? %t\n", alives, len(alives), alives == nil)

}
