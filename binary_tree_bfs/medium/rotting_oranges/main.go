package main

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := [][]int{}
	rottenCount := 0
	freshCount := 0

	// Initialize the queue with rotten oranges and count fresh oranges
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
				rottenCount++
			} else if grid[r][c] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0 // No fresh oranges to rot
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	for len(queue) > 0 && freshCount > 0 {
		minutes++
		nextQueue := [][]int{}
		for _, pos := range queue {
			r, c := pos[0], pos[1]
			for _, dir := range directions {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < rows && newC >= 0 && newC < cols && grid[newR][newC] == 1 {
					grid[newR][newC] = 2 // Rot the fresh orange
					freshCount--         // Decrease fresh count
					nextQueue = append(nextQueue, []int{newR, newC})
				}
			}
		}
		queue = nextQueue // Move to the next level of rotten oranges
	}

	if freshCount > 0 {
		return -1 // Not all fresh oranges could rot
	}
	return minutes

}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	result := orangesRotting(grid)
	println(result) // Output: 4

	grid = [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}

	result = orangesRotting(grid)
	println(result) // Output: -1

	grid = [][]int{
		{0, 2},
	}

	result = orangesRotting(grid)
	println(result) // Output: 0
}
