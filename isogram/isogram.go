package isogram

import "strings"

func IsIsogram(word string) bool {
	wordLower := strings.ToLower(word)
	for _, char := range wordLower {
		numberOfLetter := strings.Count(wordLower, string(char))
		if char != ' ' && char != '-' && numberOfLetter > 1 {
			return false
		}
	}
	return true
}
