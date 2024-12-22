package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println("Original Array: ", citations)
	output := hIndex(citations)
	fmt.Println("Result: ", output)
	fmt.Println("")

	citations = []int{1, 3, 1}
	fmt.Println("Original Array: ", citations)
	output = hIndex(citations)
	fmt.Println("Result: ", output)
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	size := len(citations)
	hIndex := 0
	for i := 0; i < size; i = i + 1 {
		if citations[i] >= i+1 {
			hIndex++
		} else {
			break
		}
	}

	return hIndex
}
