package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Println("Original Array: ", nums)
	fmt.Println("Value of K is: ", k)
	rotate(nums, k)
	fmt.Println("Modified Array", nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	fmt.Println("Original Array: ", nums)
	fmt.Println("Value of K is: ", k)
	rotate(nums, k)
	fmt.Println("Modified Array", nums)
}

// Space Complexity: O(1)
// Time Complexity: O(n)
func rotate(nums []int, k int) {
	length := len(nums)

	if k == 0 || length < 2 {
		return
	}

	k = k % length

	start, mid, end := 0, length-k-1, length-1

	temp := 0
	for i, j := start, mid; i < j; i, j = i+1, j-1 {
		temp = nums[j]
		nums[j] = nums[i]
		nums[i] = temp
	}

	for i, j := mid+1, end; i < j; i, j = i+1, j-1 {
		temp = nums[j]
		nums[j] = nums[i]
		nums[i] = temp
	}

	for i, j := 0, end; i < j; i, j = i+1, j-1 {
		temp = nums[j]
		nums[j] = nums[i]
		nums[i] = temp
	}
}
