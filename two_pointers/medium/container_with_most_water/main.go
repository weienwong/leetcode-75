package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Calculate the width and height of the container
		width := right - left
		h := min(height[left], height[right])
		// Calculate the area and update maxArea if it's larger
		area := width * h
		if area > maxArea {
			maxArea = area
		}

		// Move the pointer pointing to the shorter line inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	println(result) // Output: 49
}
