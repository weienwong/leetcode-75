package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Calculate prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]

		fmt.Println("Prefix product at index", i, ":", result[i], "with prefix", prefix)
	}

	// fmt.Println("Prefix products:", result)

	// Calculate suffix products and multiply with prefix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]

		fmt.Println("Suffix product at index", i, ":", result[i], "with suffix", suffix)
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	for _, v := range result {
		println(v)
	}

}
