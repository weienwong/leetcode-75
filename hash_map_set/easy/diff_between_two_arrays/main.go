package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	answer := make([][]int, 2)
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// Fill set1 with elements from nums1
	for _, num := range nums1 {
		set1[num] = true
	}

	// Fill set2 with elements from nums2
	for _, num := range nums2 {
		set2[num] = true
	}

	// Find elements in nums1 that are not in nums2
	for num := range set1 {
		if !set2[num] {
			answer[0] = append(answer[0], num)
		}
	}

	// Find elements in nums2 that are not in nums1
	for num := range set2 {
		if !set1[num] {
			answer[1] = append(answer[1], num)
		}
	}

	// Optional: Sort the results if required
	// sort.Ints(answer[0])
	// sort.Ints(answer[1])

	return answer
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	result := findDifference(nums1, nums2)

	println(result) // Output: [[1, 3], [4, 6]]

	nums1 = []int{1, 2, 3, 3}
	nums2 = []int{1, 1, 2, 2}
	result = findDifference(nums1, nums2)

	println(result) // Output: [[3], []]

	// This is a placeholder for the main function.
}
