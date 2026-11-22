package main

import (
	"log"
	"slices"
)

func main() {
	log.Println(findFirstNonRepeatingChar("gyeeksforgeekzs"))
}

func findFirstNonRepeatingChar(word string) string {
	var maxChars = slices.Repeat([]int32{-1}, 26)
	for index, element := range word {
		charIndex := element - 'a'
		if maxChars[charIndex] != -1 {
			maxChars[charIndex] = -2
		}

		if maxChars[charIndex] == -1 {
			maxChars[charIndex] = int32(index)
		}
	}

	var selectedIndex = rune(-1)
	for _, appearedAt := range maxChars {
		if appearedAt >= 0 && (selectedIndex == -1 || appearedAt < selectedIndex) {
			selectedIndex = rune(appearedAt)
		}
	}

	if selectedIndex == -1 {
		return string('$')
	}

	return string(word[selectedIndex])
}
