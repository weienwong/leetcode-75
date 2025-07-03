package main

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for _, num := range nums {
		println(num) // Output: 1, 3, 12, 0, 0
	}

	nums = []int{0}
	moveZeroes(nums)
	for _, num := range nums {
		println(num) // Output: 0
	}

}
