package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// get max number of candies
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	// create a result slice with the same length as candies
	result := make([]bool, len(candies))
	// check if each kid can have the most candies
	for i, candy := range candies {
		if candy+extraCandies >= maxCandies {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	result := kidsWithCandies(candies, extraCandies)
	for _, v := range result {
		println(v) // Output: [true, true, true, false, true]
	}

	candies = []int{4, 2, 1, 1, 2}
	extraCandies = 1

	result = kidsWithCandies(candies, extraCandies)
	for _, v := range result {
		println(v) // Output: [true, false, false, false, false]
	}

	candies = []int{12, 1, 12}
	extraCandies = 10

	result = kidsWithCandies(candies, extraCandies)
	for _, v := range result {
		println(v) // Output: [true, false, true]
	}

}
