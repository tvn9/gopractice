package main

import "fmt"

func main() {
	// Declare a slice
	var nums []int

	nums = append(nums, 1, 3)
	fmt.Println(nums)

	nums = append(nums, 2)
	fmt.Println(nums)
	fmt.Printf("nums len: %d nums cap: %d\n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Println(nums)
	fmt.Printf("nums len: %d nums cap: %d\n", len(nums), cap(nums))

	nums = append(nums, 6, 7)
	fmt.Println(nums)
	fmt.Printf("nums len: %d nums cap: %d\n", len(nums), cap(nums))

	nums = append(nums, nums[2:]...)
	nums = append(nums, 2, 4)

	nums = nums[:]
	fmt.Println(nums)
	fmt.Printf("nums len: %d nums cap: %d prt: %p\n", len(nums), cap(nums), &nums)
}
