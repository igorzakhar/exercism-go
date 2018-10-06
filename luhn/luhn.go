package luhn

import (
	"strconv"
	"strings"
)

func Valid(s string) bool {
	input := strings.Replace(s, " ", "", -1)
	_, err := strconv.ParseInt(input, 10, 64)
	if err != nil || len(input) == 1 {
		return false
	}

	sum := 0
	digits := []rune(input)
	parity := len(digits) % 2

	for i := 0; i < len(digits); i++ {
		digit := int(digits[i]) - '0'
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
