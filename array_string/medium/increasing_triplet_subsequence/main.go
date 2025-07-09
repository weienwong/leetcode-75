package main

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first, second := nums[0], int(^uint(0)>>1) // Initialize second to max int

	for _, num := range nums {
		if num <= first {
			first = num // Update first if current number is smaller or equal
		} else if num <= second {
			second = num // Update second if current number is smaller or equal to second
		} else { // If we find a number greater than both first and second, we have our triplet
			return true
		}
	}
	return false // If we finish the loop without finding a triplet, return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := increasingTriplet(nums)
	println(result) // Output: true

	nums = []int{5, 4, 3, 2, 1}
	result = increasingTriplet(nums)
	println(result) // Output: false

	nums = []int{2, 1, 5, 0, 4, 6}
	result = increasingTriplet(nums)
	println(result) // Output: true
}
