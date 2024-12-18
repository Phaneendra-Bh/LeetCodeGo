package main

import "fmt"

func main() {
	// Test case 1
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println("Test case 1")
	fmt.Println("Original Array: ", nums)
	result := removeDuplicates(nums)
	fmt.Println("Result: ", result)
	fmt.Println("Modified Array: ", nums)
	fmt.Println("")

	// Test case 2
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println("Test case 2")
	fmt.Println("Original Array: ", nums)
	result = removeDuplicates(nums)
	fmt.Println("Result: ", result)
	fmt.Println("Modified Array: ", nums)
}

// Space Complexity: O(1)
// Time Complexity: O(n)
func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	lastElement := nums[0]
	nextPosition := 1
	count := 1
	result := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == lastElement {
			if count < 2 {
				nums[nextPosition] = lastElement
				nextPosition++
				count++
				result++
			}
			continue
		}

		lastElement = nums[i]
		nums[nextPosition] = lastElement
		count = 1
		nextPosition++
		result++
	}

	return result
}
