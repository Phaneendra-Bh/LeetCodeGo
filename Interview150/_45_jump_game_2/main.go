package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 3, 1, 4}
	fmt.Println("Original Array: ", nums)
	output := jump(nums)
	fmt.Println("Result: ", output)

	nums = []int{2, 3, 0, 1, 4}
	fmt.Println("Original Array: ", nums)
	output = jump(nums)
	fmt.Println("Result: ", output)
}

// jump function to return the minimum number of jumps to reach the last index
func jump(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0 // No jumps needed if the array has only one element
	}

	jumps := 0      // Number of jumps taken
	currentEnd := 0 // The farthest index we can reach with the current jump
	farthest := 0   // The farthest index we can reach in the next jump

	for i := 0; i < n-1; i++ {
		// Update the farthest index we can reach from the current position
		farthest = max(farthest, i+nums[i])

		// If we've reached the end of the current range, we need to make another jump
		if i == currentEnd {
			jumps++
			currentEnd = farthest

			// If we've reached or passed the last index, no need to continue
			if currentEnd >= n-1 {
				break
			}
		}
	}
	return jumps
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
