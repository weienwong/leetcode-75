package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	// count number of empty spots that are not next to a flower
	count := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			// Check if the current spot is empty and the previous and next spots are also empty or out of bounds
			if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				// Place a flower here
				flowerbed[i] = 1
				count++
			}
		}
		if count >= n {
			return true
		}
	}

	return false
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1

	result := canPlaceFlowers(flowerbed, n)
	println(result) // Output: true

	flowerbed = []int{1, 0, 0, 0, 1}
	n = 2

	result = canPlaceFlowers(flowerbed, n)
	println(result) // Output: false
}
