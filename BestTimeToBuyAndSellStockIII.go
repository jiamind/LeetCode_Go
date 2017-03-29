package main


func maxProfit(prices []int) int {

	if prices == nil || len(prices) < 2 {
		return 0
	}

	minPriceFirstPurchase := prices[0]
	profitAfterFirstSell := 0
	profitAfterSecondPurchase := 0 - int(^uint(0) >> 1) - 1
	profitAfterSecondSell := 0

	for i := 1; i < len(prices); i++ {
		minPriceFirstPurchase = Min(minPriceFirstPurchase, prices[i])
		profitAfterFirstSell = Max(profitAfterFirstSell, prices[i] - minPriceFirstPurchase)
		profitAfterSecondPurchase = Max(profitAfterSecondPurchase, profitAfterFirstSell - prices[i])
		profitAfterSecondSell = Max(profitAfterSecondSell, profitAfterSecondPurchase + prices[i])
	}

	return profitAfterSecondSell
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int  {
	if a < b {
		return a
	} else {
		return b
	}
}