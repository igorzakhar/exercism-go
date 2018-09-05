package raindrops

import (
	"strconv"
)

func Convert(in int) string {
	var s string
	if in%3 == 0 {
		s += "Pling"
	}
	if in%5 == 0 {
		s += "Plang"
	}
	if in%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(in)
	}
	return s
}
