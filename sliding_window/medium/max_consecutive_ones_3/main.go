package main

func longestOnes(nums []int, k int) int {
	maxlength := 0
	length := 0
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] == 1 {
			length++
		} else {
			if k > 0 {
				k--
				length++
			} else {
				for nums[i] == 1 {
					i++
					length--
				}
				i++
			}
		}

		if length > maxlength {
			maxlength = length
		}

	}

	return maxlength

}

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	result := longestOnes(nums, k)
	println(result) // Output: 6

	nums = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k = 3
	result = longestOnes(nums, k)
	println(result) // Output: 10

}
