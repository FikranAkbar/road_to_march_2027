package main

import "log"

func main() {
	intList := []int{1, 1, 2, 3, 4}
	log.Println(`Find majority element of`, intList, `:`, findMajorityElement(intList))
}

func findMajorityElement(intList []int) int {
	var (
		elementMap = make(map[int]int)
	)

	for _, element := range intList {
		elementMap[element]++
		if elementMap[element] > len(intList)/2 {
			return element
		}
	}

	return -1
}
