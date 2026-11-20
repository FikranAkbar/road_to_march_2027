package main

import (
	"log"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println("STARTING ARRAY:", arr)
	arr = rotateArray(arr, 3)
	log.Println("ARRAY AFTER ROTATION:", arr)
}

func rotateArray(arr []int, rotation int) []int {
	actualRotation := rotation % (len(arr) - 1)
	return append(arr[len(arr)-actualRotation-1:], arr[:len(arr)-actualRotation-1]...)
}
