package main

import "log"

func main() {
	log.Println("Maximum Product Subarray:", maximumProductSubarray([]int{-5, -4, -3, -2, -1}))
}

func maximumProductSubarray(intList []int) int {
	var (
		maxProd int
	)

	if len(intList) <= 0 {
		return 0
	}

	maxProd = intList[0]

	for i := range intList {
		currentProd := 1
		for j := i; j < len(intList); j++ {
			currentProd *= intList[j]
			maxProd = max(maxProd, currentProd)
		}
	}

	return maxProd
}

