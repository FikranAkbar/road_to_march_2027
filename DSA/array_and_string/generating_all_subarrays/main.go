package main

import (
	"log"
	"slices"
)

func main() {
	arr := []int{9, 1, 4, 8}
	log.Println("MASTER ARRAY:", arr)
	log.Println("SUB ARRAYS GENERATED:", generateSubArrays(arr))
	arr[0] = 10
	log.Println("MASTER ARRAY:", arr)
}

func generateSubArrays(arr []int) [][]int {
	var subArrayList [][]int

	for i := 0; i < len(arr); i++ {
		for l := i; l < len(arr); l++ {
			subArrayList = append(subArrayList, slices.Clone(arr[i:l+1]))
		}
	}

	return subArrayList
}
