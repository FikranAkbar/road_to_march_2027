package main

import "log"

func main() {
	intList := []int{2, 2, 3, 1, 3, 2, 1, 1}
	log.Println(`MAJORITY ELEMENT N DIV THREE OF`, intList, `IS`, majorityElementNDivThree(intList))
}

func majorityElementNDivThree(intList []int) []int {
	var (
		mapCounts       = make(map[int]int, len(intList))
		majoritiesArray []int
	)

	for _, element := range intList {
		mapCounts[element]++
	}

	for element, count := range mapCounts {
		if count > len(intList)/3 {
			majoritiesArray = append(majoritiesArray, element)
		}
	}

	return majoritiesArray
}
