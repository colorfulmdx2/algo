package __array

func maxProfit(prices []int) int {
	bestProfit := 0

	if len(prices) == 0 {
		return bestProfit
	}

	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > bestProfit {
			bestProfit = profit
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

	}

	return bestProfit

}
