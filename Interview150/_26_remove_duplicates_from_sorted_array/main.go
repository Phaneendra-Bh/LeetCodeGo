package main

import "fmt"

func main() {
	// Test case 1
	nums := []int{1, 1, 2}
	fmt.Println("Test case 1")
	fmt.Println("Original Array: ", nums)
	result := removeDuplicates(nums)
	fmt.Println("Result: ", result)
	fmt.Println("Modified Array:", nums)
	fmt.Println("")

	// Test case 2
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Test case 2")
	fmt.Println("Original Array: ", nums)
	result = removeDuplicates(nums)
	fmt.Println("Result: ", result)
	fmt.Println("Modified Array:", nums)
	fmt.Println("")
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	prevElement := nums[0]
	nextPosition := 1
	result := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == prevElement {
			continue
		}

		prevElement = nums[i]
		nums[nextPosition] = prevElement
		nextPosition++
		result++
	}

	return result
}
