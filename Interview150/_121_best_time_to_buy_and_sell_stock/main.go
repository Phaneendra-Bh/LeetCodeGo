package main

import "fmt"

func main() {
	// Test case 1
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Original Array: ", prices)
	result := maxProfit(prices)
	fmt.Println("Result: ", result)
	fmt.Println("")

	// Test case 2
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println("Original Array: ", prices)
	result = maxProfit(prices)
	fmt.Println("Result: ", result)
}

// Space Complexity: O(1)
// Time Complexity: O(n)
func maxProfit(prices []int) int {
	length := len(prices)

	if length < 2 {
		return 0
	}

	min := prices[0]
	profit := 0

	for i := 1; i < length; i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}

	return profit
}
