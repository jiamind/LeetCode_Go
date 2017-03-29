package main

func maxProfit(prices []int) int {

	if prices == nil || len(prices) < 2 {
		return 0
	}

	minToPurchase := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > minToPurchase {
			maxProfit = Max(maxProfit, prices[i] - minToPurchase)
		} else {
			minToPurchase = prices[i]
		}
	}

	return maxProfit
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}