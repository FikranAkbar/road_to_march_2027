package main

import "log"

func main() {
	dailyStockPrices := []int{7, 10, 1, 3, 6, 9, 2}
	profit, boughtPrice, soldPrice := findMaximumProfit(dailyStockPrices)
	log.Printf(`Maximum Profit:%d, Bought At: %d, Sold At: %d`, profit, boughtPrice, soldPrice)
	log.Println(`Maximum Profit:`, findMaximumProfitV2(dailyStockPrices))
}

func findMaximumProfitV2(dailyStockPrices []int) int {
	var (
		boughtPrice = dailyStockPrices[0]
		profit      = 0
	)

	if len(dailyStockPrices) <= 0 {
		return 0
	}

	for i := 1; i < len(dailyStockPrices); i++ {
		boughtPrice = min(boughtPrice, dailyStockPrices[i])
		profit = max(profit, dailyStockPrices[i]-boughtPrice)
	}

	return profit
}

func findMaximumProfit(dailyStockPrices []int) (int, int, int) {
	var (
		boughtPrice int
		soldPrice   int
		profit      int
	)

	for i, price := range dailyStockPrices {
		if i == 0 {
			boughtPrice = price
			continue
		}

		currentProfit := price - boughtPrice
		if currentProfit > 0 && currentProfit > profit {
			soldPrice = price
			profit = currentProfit
			continue
		}

		if currentProfit <= 0 {
			boughtPrice = price
		}
	}

	return profit, boughtPrice, soldPrice
}
