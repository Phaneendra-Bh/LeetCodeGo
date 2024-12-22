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
	prices = []int{1, 2, 3, 4, 5}
	fmt.Println("Original Array: ", prices)
	result = maxProfit(prices)
	fmt.Println("Result: ", result)
	fmt.Println("")

	// Test case 3
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println("Original Array: ", prices)
	result = maxProfit(prices)
	fmt.Println("Result: ", result)
}

func maxProfit(prices []int) int {
	length := len(prices)

	if length < 2 {
		return 0
	}

	profit := 0
	for i, j := 0, 1; j < length; i, j = i+1, j+1 {
		if prices[j] > prices[i] {
			profit = profit + (prices[j] - prices[i])
		}
	}

	return profit
}
