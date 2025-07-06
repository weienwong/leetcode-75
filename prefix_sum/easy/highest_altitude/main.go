package main

import "fmt"

func largestAltitude(gain []int) int {
	maxAltitude := 0
	currentAltitude := 0
	for _, g := range gain {
		currentAltitude += g
		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}

	return maxAltitude
}

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	largestAltitudeGained := largestAltitude(gain)
	fmt.Println(largestAltitudeGained) // Output: 1

	gain = []int{-4, -3, -2, -1, 4, 3, 2}
	largestAltitudeGained = largestAltitude(gain)
	fmt.Println(largestAltitudeGained) // Output: 0

	gain = []int{1, 2, 3, 4, 5}
	largestAltitudeGained = largestAltitude(gain)
	fmt.Println(largestAltitudeGained) // Output: 15
}
