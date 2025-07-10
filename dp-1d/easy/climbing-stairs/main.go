package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// Initialize the first two steps
	first, second := 1, 2

	// Calculate the number of ways to reach each step from 3 to n
	for i := 3; i <= n; i++ {
		next := first + second
		first = second
		second = next
	}

	return second

}

func main() {
	n := 2
	result := climbStairs(n)
	println(result) // Output: 2

	n = 3
	result = climbStairs(n)
	println(result) // Output: 3
}
