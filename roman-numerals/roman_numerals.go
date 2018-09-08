package romannumerals

import (
	"fmt"
)

var decimal = [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var roman = [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func ToRomanNumeral(num int) (string, error) {
	var romanNum string
	if num <= 0 || num >= 3001 {
		return "", fmt.Errorf("%d. Values over 3001 and negative values are not allowed.", num)
	}
	for index, decimalValue := range decimal {
		for num%decimalValue < num {
			romanNum += roman[index]
			num -= decimalValue
		}
	}
	return romanNum, nil
}
