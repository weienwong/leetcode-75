package main

import "sort"

func maxOperations(nums []int, k int) int {
	count := 0
	left, right := 0, len(nums)-1

	sort.Ints(nums) // Sort the array to use two pointers effectively

	// Sort the array to use two pointers effectively
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			count++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}

	return count
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 5
	result := maxOperations(nums, k)
	println(result) // Output: 2

	nums = []int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}
	k = 2
	result = maxOperations(nums, k)
	println(result) // Output: 2
}
