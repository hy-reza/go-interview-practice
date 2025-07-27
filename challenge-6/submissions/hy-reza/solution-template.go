// Package challenge6 contains the solution for Challenge 6.
package challenge6

import (
	"strings"
)

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func CountWordFrequency(text string) map[string]int {
	result := make(map[string]int)
	currentWord := ""

	lowerText := strings.ToLower(text)
	processedText := strings.ReplaceAll(lowerText, "'", "")

	for _, r := range processedText {
		if isAlphaNumeric(r) {
			currentWord += string(r)
		} else {
			if currentWord != "" {
				result[currentWord]++
			}
			currentWord = ""
		}
	}

	if currentWord != "" {
		result[currentWord]++
	}

	return result
}