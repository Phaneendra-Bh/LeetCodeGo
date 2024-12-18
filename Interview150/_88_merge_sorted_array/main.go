package main

import "fmt"

func main() {
	// Test case 1
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	fmt.Println("Test case 1")
	fmt.Println("Array 1: ", nums1)
	fmt.Println("Array 2: ", nums2)
	merge(nums1, m, nums2, n)
	fmt.Println("Merged array: ", nums1)
	fmt.Println("")

	// Test case 2
	nums1 = []int{1}
	nums2 = []int{}
	m, n = 1, 0
	fmt.Println("Test case 2")
	fmt.Println("Array 1: ", nums1)
	fmt.Println("Array 2: ", nums2)
	merge(nums1, m, nums2, n)
	fmt.Println("Merged array: ", nums1)
	fmt.Println("")

	// Test case 3
	nums1 = []int{0}
	nums2 = []int{1}
	m, n = 0, 1
	fmt.Println("Test case 3")
	fmt.Println("Array 1: ", nums1)
	fmt.Println("Array 2: ", nums2)
	merge(nums1, m, nums2, n)
	fmt.Println("Merged array: ", nums1)
	fmt.Println("")
}

// Space Complexity: O(m+n)
// Time Complexity: O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	merged := make([]int, m+n)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			merged[k] = nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			merged[k] = nums2[j]
			j++
		}
		k++
	}

	if i < m {
		for i < m {
			merged[k] = nums1[i]
			i++
			k++
		}
	}

	if j < n {
		for j < n {
			merged[k] = nums2[j]
			j++
			k++
		}
	}

	c := 0
	for _, val := range merged {
		nums1[c] = val
		c++
	}
}
