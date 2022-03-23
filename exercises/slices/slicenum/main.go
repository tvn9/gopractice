package main

import (
	"fmt"
	"strings"
)

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	// 1. converting string data into nums slice
	nums := strings.Split(data, " ")

	// 2. Print the slice
	fmt.Printf("%T, %-5d nums %v\n", nums, len(nums), nums)

	// 3. Slice for even numbers and print out
	evens := nums[:3]
	fmt.Printf("%T, %-5d evens %v\n", evens, len(evens), evens)

	// 4. Slice for the odd numbers and print it
	odds := nums[3:]
	fmt.Printf("%T, %-5d odds %v\n", odds, len(odds), odds)

	// 5. Slice for 2 numbers at the middle and print it
	middle := nums[2:4]
	fmt.Printf("%T, %-5d middle %v\n", middle, len(middle), middle)

	// 6. Slice for the first 2 numbers and print it
	first2 := nums[0:2]
	fmt.Printf("%T, %-5d first 2 %v\n", first2, len(first2), first2)

	// 7. Slice for the last two numbers
	last2 := nums[len(nums)-2:]
	fmt.Printf("%T, %-5d last 2 %v\n", last2, len(last2), last2)

	// 8. Slice the evens for the last one number
	evensLast1 := evens[len(evens)-1:]
	fmt.Printf("%T, %-5d even last 1 %v\n", evensLast1, len(evensLast1), evensLast1)

	// Slice fhe odds slice for the last two numbers
	oddsLast2 := odds[len(odds)-2:]
	fmt.Printf("%T, %-5d odd last 2 %v\n", oddsLast2, len(oddsLast2), oddsLast2)

}
