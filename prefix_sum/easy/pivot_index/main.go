package main

func pivotIndex(nums []int) int {
	index := -1
	leftSum := 0
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			index = i
			break
		}
		leftSum += num
	}
	return index
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	result := pivotIndex(nums)
	println(result) // Output: 3

	nums = []int{1, 2, 3}
	result = pivotIndex(nums)
	println(result) // Output: -1

	nums = []int{2, 1, -1}
	result = pivotIndex(nums)
	println(result) // Output: 0
}
