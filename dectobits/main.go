package main

import (
	"fmt"
	"os"
)

func main() {
	num := 5

	// Decimal to Bits
	fmt.Printf("Converting Decimal %d to Bits\n", num)
	bits := DecToBits(num)
	fmt.Println(bits)

	// Converting number if bits to binary position decimal value
	fmt.Printf("Converting %d bits to decimal value\n", num)
	decs := BitToDecValue(num)
	fmt.Println(decs)
}

// DecToBits convert a decimal 0 to 256 number to bits
// The function is currently limited to 8 bits convertion only
func DecToBits(num int) []string {
	nums := BitToDecValue(num)

	var (
		bit []string
		b   string
	)
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] > num {
			b = "0"
		} else if nums[i] <= num {
			b = "1"
			num -= nums[i]
		} else {
			b = "0"
		}
		bit = append(bit, b)
	}
	return bit
}

// BitToDecValue converts the number of bits to binary position decimal value
// the function accept the number of bits
func BitToDecValue(num int) []int {
	// Set default number if bit to 8 or 16
	if num <= 0 {
		fmt.Printf("%d, is not a valid entry, try 1 - 16.\n", num)
		os.Exit(1)
	} else if num <= 8 {
		num = 8
	} else {
		num = 16
	}

	// Defind the return value []bitVal
	var bitVal []int

	// Defind a temp value to append to bitVal slice
	var dec int

	// Defind a binary value of 2 for binary calculation
	pow := 2

	for i := 0; i <= num-1; i++ {
		if i == 0 {
			dec = 1
		} else if i == 1 {
			dec = 2
		} else {
			dec *= pow
		}
		bitVal = append(bitVal, dec)
	}
	// Reverse sort bitVal for binany representation
	for i, j := 0, len(bitVal)-1; i < j; i, j = i+1, j-1 {
		bitVal[i], bitVal[j] = bitVal[j], bitVal[i]
	}

	return bitVal
}
