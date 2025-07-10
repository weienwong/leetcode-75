package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // Update the minimum price
		} else {
			priceDiff := price - minPrice
			if priceDiff > maxProfit {
				maxProfit = priceDiff // Update the maximum profit
			}
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	println(result) // Output: 5

	prices = []int{7, 6, 4, 3, 1}
	result = maxProfit(prices)
	println(result) // Output: 0

}
