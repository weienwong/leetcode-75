package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}

	occurrences := make(map[int]bool)
	for _, count := range counts {
		if occurrences[count] {
			return false
		}
		occurrences[count] = true
	}

	return true
}

func main() {
	// Example usage
	arr := []int{1, 2, 2, 1, 1, 3}
	result := uniqueOccurrences(arr)
	fmt.Println(result) // Output: true

	arr = []int{1, 2}
	result = uniqueOccurrences(arr)
	fmt.Println(result) // Output: false

	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	result = uniqueOccurrences(arr)
	fmt.Println(result) // Output: true
}
