package main

import "log"

func main() {
	dailyStockPrices := []int{100, 180, 260, 310, 40, 535, 695}
	log.Println("MAXIMUM TOTAL PROFIT:", findMaximumProfitMultipleTxImproved(dailyStockPrices))
}

func findMaximumProfitMultipleTxImproved(dailyStockPrices []int) int {
	var (
		totalProfit = 0
	)

	for i := 1; i < len(dailyStockPrices); i++ {
		if dailyStockPrices[i] > dailyStockPrices[i-1] {
			totalProfit += dailyStockPrices[i] - dailyStockPrices[i-1]
		}
	}

	return totalProfit
}

func findMaximumProfitMultipleTx(dailyStockPrices []int) int {
	var (
		boughtPrice     = dailyStockPrices[0]
		profit          = 0
		collectedProfit = 0
	)

	if len(dailyStockPrices) <= 0 {
		return 0
	}

	for i := 1; i < len(dailyStockPrices); i++ {
		log.Println("STEP i", i)
		profit = max(profit, dailyStockPrices[i]-boughtPrice)
		log.Println("PROFIT", profit)
		boughtPrice = min(boughtPrice, dailyStockPrices[i])
		log.Println("BOUGHT PRICE", boughtPrice)
		if boughtPrice == dailyStockPrices[i] {
			collectedProfit += profit
			log.Println("COLLECTED PROFIT", profit)
			profit = 0
		}

		if i == len(dailyStockPrices)-1 {
			collectedProfit += profit
		}
	}

	return collectedProfit
}
