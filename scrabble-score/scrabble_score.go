package scrabble

import (
	"unicode"
)

func Score(word string) int {
	points := 0
	for _, letter := range word {
		switch unicode.ToLower(letter) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			points += 1
		case 'd', 'g':
			points += 2
		case 'b', 'c', 'm', 'p':
			points += 3
		case 'f', 'h', 'v', 'w', 'y':
			points += 4
		case 'k':
			points += 5
		case 'j', 'x':
			points += 8
		case 'q', 'z':
			points += 10
		}
	}
	return points
}
