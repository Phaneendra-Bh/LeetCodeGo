package main

import "fmt"

func main() {
	// Test case 1
	nums := []int{3, 2, 3}
	output := majorityElement(nums)
	fmt.Println("Result: ", output)

	// Test case 2
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	output = majorityElement(nums)
	fmt.Println("Result: ", output)
}

func majorityElement(nums []int) int {
	lookup := make(map[int]int, 0)

	for _, key := range nums {
		if val, found := lookup[key]; found {
			lookup[key] = val + 1
		} else {
			lookup[key] = 1
		}
	}

	for key, val := range lookup {
		if val > len(nums)/2 {
			return key
		}
	}

	return -1
}
