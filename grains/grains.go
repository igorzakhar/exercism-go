package grains

import (
	"errors"
)

func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, errors.New("out of range 1...64")
	}
	var square uint64
	square = 1 << uint64(i-1)
	return square, nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		sum += sq
	}
	return sum
}
