package main

import (
	"log"
	"strings"
)

func main() {
	text := `i love programming`
	log.Println(`UrlifyString this text:`, text, `Result is:`, urlifyString(text))
}

func urlifyString(text string) string {
	var res []string

	for _, r := range text {
		char := string(r)
		if char == ` ` {
			res = append(res, "%20")
			continue
		}

		res = append(res, string(r))
	}

	return strings.Join(res, ``)
}
