package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if _, exists := m[num]; exists {
			return true // Duplicate found
		} else {
			m[num] = true // Add the number to the map
		}
	}

	return false // No duplicates found
}

func main() {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	println(result) // Output: true

	nums = []int{1, 2, 3, 4}
	result = containsDuplicate(nums)
	println(result) // Output: false

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result = containsDuplicate(nums)
	println(result) // Output: true
}
