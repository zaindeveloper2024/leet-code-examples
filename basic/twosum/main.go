package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	return []int{}
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Printf("Example 1: %v\n", twoSum(nums1, target1))

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Printf("Example 2: %v\n", twoSum(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Printf("Example 3: %v\n", twoSum(nums3, target3))
}
