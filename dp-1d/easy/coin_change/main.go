package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// Initialize dp array with a value larger than any possible number of coins
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0 // Base case: 0 coins needed to make amount 0

	// Fill the dp array
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1 // Not possible to make the amount
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11

	result := coinChange(coins, amount)
	println(result) // Output: 3

	coins = []int{2}
	amount = 3

	result = coinChange(coins, amount)
	println(result) // Output: -1

	coins = []int{1}
	amount = 0

	result = coinChange(coins, amount)
	println(result) // Output: 0
}
