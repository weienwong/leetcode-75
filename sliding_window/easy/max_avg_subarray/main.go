package main

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	currentSum := 0

	// Calculate the sum of the first 'k' elements
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}
	maxSum = currentSum

	// Slide the window across the array
	for i := k; i < len(nums); i++ {
		currentSum += nums[i] - nums[i-k]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	result := findMaxAverage(nums, k)
	println(result) // Output: 12.75

}
