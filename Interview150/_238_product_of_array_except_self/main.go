package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Original Array: ", nums)
	result := productExceptSelf(nums)
	fmt.Println("Result: ", result)

	nums = []int{-1, 1, 0, -3, 3}
	fmt.Println("Original Array: ", nums)
	result = productExceptSelf(nums)
	fmt.Println("Result: ", result)
}

func productExceptSelf(nums []int) []int {
	product := 1
	zeroCount := 0

	for _, val := range nums {
		if val == 0 {
			zeroCount++
			continue
		}
		product = product * val
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if zeroCount >= 2 {
			result = append(result, 0)
		} else if zeroCount == 1 {
			if nums[i] == 0 {
				result = append(result, product)
			} else {
				result = append(result, 0)
			}
		} else {
			result = append(result, divide(product, nums[i]))
		}
	}

	return result
}

func divide(dividend, divisor int) int {
	result := 0
	a, b := 1, 1
	if dividend < 0 {
		a = -1
	}
	if divisor < 0 {
		b = -1
	}

	for dividend > 0 {
		dividend = dividend - divisor
		result++
	}
	return result * a * b
}
