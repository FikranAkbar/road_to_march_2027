package main

import "log"

func main() {
	log.Println("Maximum Product Subarray:", maximumProductSubarrayOptimum([]int{-5, -4, -3, -2, -1}))
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

func maximumProductSubarrayOptimum(intList []int) int {
	var (
		currMin int
		currMax int
		maxProd int
	)

	if len(intList) <= 0 {
		return 0
	}

	currMin = intList[0]
	currMax = intList[0]
	maxProd = currMax

	for i := 1; i < len(intList); i++ {
		if intList[i] == 0 {
			currMax = 1
			currMin = 1
		}

		currMin = min(currMin, currMin*intList[i], intList[i])
		currMax = max(currMax, currMax*intList[i], intList[i])
		maxProd = max(currMax, maxProd)
	}

	return maxProd
}
