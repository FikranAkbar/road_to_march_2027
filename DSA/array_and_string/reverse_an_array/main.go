package main

import "log"

func main() {
	arr := []int{0, 4, 1, 2, 5, 6}
	log.Println("STARTING ARRAY:", arr)
	arr = reverseArray(arr)
	log.Println("REVERSED ARRAY:", arr)
}

func reverseArray(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = temp
	}

	return arr
}
