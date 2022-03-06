package main

import "fmt"

func main() {
	names := []string{"Dexter Watts", "Logan Bernes", "James Hunt"}
	distances := []int{1494, 1235, 2345, 5888, 3435}
	data := []byte{'A', 'B', 'C', 'D', 'E'}
	ratios := []float64{122.50, 22871.80}
	alives := []bool{true, false, true, false}

	fmt.Printf("names     - type: %T, len: %d, is equal to nil? %t\n", names, len(names), names == nil)
	fmt.Printf("distances - type: %T, len: %d, is equal to nil? %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data      - type: %T, len: %d, is equal to nil? %t\n", data, len(data), data == nil)
	fmt.Printf("ratios    - type: %T, len: %d, is equal to nil? %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives    - type: %T, len: %d, is equal to nil? %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Println("len of distances and len of data slices are equal")
	}
}
