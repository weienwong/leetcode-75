package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	// Initialize the first three tribonacci numbers
	a, b, c := 0, 1, 1

	// Calculate the tribonacci number iteratively
	for i := 3; i <= n; i++ {
		next := a + b + c
		a, b, c = b, c, next
	}

	return c
}

func main() {
	n := 4
	result := tribonacci(n)
	println(result) // Output: 4
}
