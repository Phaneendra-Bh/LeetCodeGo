package main

import "fmt"

func main() {
	// Test case 1
	nums := []int{2, 4, 4, 4, 0}
	val := 4
	fmt.Println("Test case 1")
	fmt.Println("Original array: ", nums)
	result := removeElement(nums, val)
	fmt.Println("Result: ", result)
	fmt.Println("Modified array: ", nums)
	fmt.Println("")

	// Test case 2
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println("Test case 2")
	fmt.Println("Original array: ", nums)
	result = removeElement(nums, val)
	fmt.Println("Result: ", result)
	fmt.Println("Modified array: ", nums)
	fmt.Println("")

	// Test case 3
	nums = []int{4, 4, 4, 4}
	val = 4
	fmt.Println("Test case 3")
	fmt.Println("Original array: ", nums)
	result = removeElement(nums, val)
	fmt.Println("Result: ", result)
	fmt.Println("Modified array: ", nums)
	fmt.Println("")

	// Test case 4
	nums = []int{2, 2, 3}
	val = 2
	fmt.Println("Test case 4")
	fmt.Println("Original array: ", nums)
	result = removeElement(nums, val)
	fmt.Println("Result: ", result)
	fmt.Println("Modified array: ", nums)
	fmt.Println("")

	// Test case 5
	nums = []int{0, 4, 4, 0, 4, 4, 4, 0, 2}
	val = 4
	fmt.Println("Test case 5")
	fmt.Println("Original array: ", nums)
	result = removeElement(nums, val)
	fmt.Println("Result: ", result)
	fmt.Println("Modified array: ", nums)
}

// Space Complexity: O(1)
// Time Complexity: O(n)
func removeElement(nums []int, val int) int {
	result := len(nums)

	if result == 0 {
		return result
	}

	i, j := 0, result-1

	for ; i < j; i++ {
		for j >= 0 {
			if nums[j] != val {
				break
			}
			j--
			result--
		}
		if j <= 0 || i >= j {
			break
		}

		if nums[i] == val {
			nums[i] = nums[j]
			nums[j] = val
			j--
			result--
		}
	}

	if i == j && nums[i] == val {
		result--
	}

	return result
}
