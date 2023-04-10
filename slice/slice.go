package main

import "fmt"

func main() {
	// define a slice
	nums := []float64{10.0, 21.50, 25.00, 30, 35, 87.2, 98.3, 100}
	fmt.Println(nums, len(nums), cap(nums))

	// adding element to existing slice
	nums = append(nums, 5.99)
	nums = append(nums, 100.99, 123, 222, 333, 444, 555)
	nums = append(nums, 666, 777)
	fmt.Println(nums, len(nums), cap(nums))

	nums = nums[0:6]
	fmt.Println(nums, len(nums), cap(nums))

	newNums := []float64{88, 99, 111}

	nums = append(nums, newNums...)

	fmt.Println(nums)

	nums = nums[0:cap(nums)]

	fmt.Println(nums, len(nums), cap(nums))

}
