package main

func moveZeroes(nums []int) {
	// Solution 1: Create a new slice and copy non-zero elements
	// This approach is simple and easy to understand, but it uses extra space.
	// It iterates through the original slice, collects non-zero elements,
	// and fills the rest of the slice with zeros.

	/*
		if len(nums) == 0 {
			return
		}
		res := make([]int, 0, len(nums))

		for _, num := range nums {
			if num != 0 {
				res = append(res, num)
			}
		}
		// Fill the rest of the slice with zeros
		for i := len(res); i < len(nums); i++ {
			res = append(res, 0)
		}

		// Copy the result back to the original slice
		copy(nums, res)
	*/

	// Solution 2: Two pointers approach
	// This approach is more space-efficient and runs in O(n) time.
	// It uses two pointers: one for the current position of non-zero elements
	// and another to iterate through the slice.

	if len(nums) == 0 {
		return
	}

	j := 0 // Pointer for the position of the next non-zero element

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	// Fill the rest of the slice with zeros
	for j < len(nums) {
		nums[j] = 0
		j++
	}
	// The original slice is now modified in place with all non-zero elements at the front
	// and zeros at the end.
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for _, num := range nums {
		println(num) // Output: 1, 3, 12, 0, 0
	}

	// nums = []int{0}
	// moveZeroes(nums)
	// for _, num := range nums {
	// 	println(num) // Output: 0
	// }

}
