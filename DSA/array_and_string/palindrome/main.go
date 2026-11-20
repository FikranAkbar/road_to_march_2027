package main

import "log"

func main() {
	word := "été"
	log.Printf(`Is Word %s palindrome ? %v`, word, isPalindrome(word))
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}

	return true
}
