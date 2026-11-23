package main

import (
	"log"
)

func main() {
	log.Println("FIND MAXIMUM SUB ARRAY SUM", maximumSubarraySum([]int{1, -1}))
	log.Println("EXPECT 1")
}

func maximumSubarraySum(listInt []int) int {
	var (
		result    = listInt[0]
		maxEnding = listInt[0]
	)

	for i := 1; i < len(listInt); i++ {
		maxEnding = max(maxEnding+listInt[i], listInt[i])
		result = max(result, maxEnding)
	}
	return result
}

func maximumSubarraySumBackup(listInt []int) int {
	var (
		globalBiggestSum  = listInt[0]
		currentBiggestSum = listInt[0]
	)

	for i := 1; i < len(listInt); i++ {
		if currentBiggestSum+listInt[i] >= listInt[i] {
			currentBiggestSum += listInt[i]
			globalBiggestSum = currentBiggestSum
			continue
		}

		log.Println("SKIP SUM")
		if currentBiggestSum < listInt[i] {
			log.Println("CurrentBiggestSum+listInt[i]", currentBiggestSum+listInt[i])
			log.Println("INT I Smaller", listInt[i])
			currentBiggestSum = listInt[i]
			log.Println("CurrentBiggestSum", currentBiggestSum)
		}
	}

	if globalBiggestSum > currentBiggestSum {
		return globalBiggestSum
	}
	return currentBiggestSum
}
